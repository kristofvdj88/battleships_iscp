// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// Package group implements a generic peering.GroupProvider.
package group

import (
	"errors"
	"fmt"
	"time"

	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/peering"
)

const NotInGroup uint16 = 0xFFFF

//
// groupImpl implements peering.GroupProvider
//
type groupImpl struct {
	netProvider peering.NetworkProvider
	nodes       []peering.PeerSender
	other       map[uint16]peering.PeerSender
	log         *logger.Logger
}

// NewPeeringGroupProvider creates a generic peering group.
// That should be used as a helper for peering implementations.
func NewPeeringGroupProvider(netProvider peering.NetworkProvider, nodes []peering.PeerSender, log *logger.Logger) peering.GroupProvider {
	other := make(map[uint16]peering.PeerSender)
	for i := range nodes {
		if nodes[i].NetID() != netProvider.Self().NetID() {
			other[uint16(i)] = nodes[i]
		}
	}
	return &groupImpl{
		netProvider: netProvider,
		nodes:       nodes,
		other:       other,
		log:         log,
	}
}

// PeerIndex implements peering.GroupProvider.
func (g *groupImpl) PeerIndex(peer peering.PeerSender) (uint16, error) {
	return g.PeerIndexByNetID(peer.NetID())
}

// PeerIndexByNetID implements peering.GroupProvider.
func (g *groupImpl) PeerIndexByNetID(peerNetID string) (uint16, error) {
	for i := range g.nodes {
		if g.nodes[i].NetID() == peerNetID {
			return uint16(i), nil
		}
	}
	return uint16(NotInGroup), errors.New("peer_not_found_by_net_id")
}

// SendMsgByIndex implements peering.GroupProvider.
func (g *groupImpl) SendMsgByIndex(peerIdx uint16, msg *peering.PeerMessage) {
	g.nodes[peerIdx].SendMsg(msg)
}

// Broadcast implements peering.GroupProvider.
func (g *groupImpl) Broadcast(msg *peering.PeerMessage, includingSelf bool) {
	var peers map[uint16]peering.PeerSender
	if includingSelf {
		peers = g.AllNodes()
	} else {
		peers = g.OtherNodes()
	}
	for i := range peers {
		peers[i].SendMsg(msg)
	}
}

// ExchangeRound sends a message to the specified set of peers and waits for acks.
// Resends the messages if acks are not received for some time.
func (g *groupImpl) ExchangeRound(
	peers map[uint16]peering.PeerSender,
	recvCh chan *peering.RecvEvent,
	retryTimeout time.Duration,
	giveUpTimeout time.Duration,
	sendCB func(peerIdx uint16, peer peering.PeerSender),
	recvCB func(recv *peering.RecvEvent) (bool, error),
) error {
	var err error
	acks := make(map[uint16]bool)
	errs := make(map[uint16]error)
	retryCh := time.After(retryTimeout)
	giveUpCh := time.After(giveUpTimeout)
	for i := range peers {
		acks[i] = false
		sendCB(i, peers[i])
	}
	haveAllAcks := func() bool {
		for i := range acks {
			if !acks[i] {
				return false
			}
		}
		return true
	}
	for !haveAllAcks() {
		select {
		case recvMsg, ok := <-recvCh:
			if !ok {
				return errors.New("recv_channel_closed")
			}
			if recvMsg.Msg.SenderIndex, err = g.PeerIndex(recvMsg.From); err != nil {
				g.log.Warnf(
					"Dropping message %v -> %v, MsgType=%v because of %v",
					recvMsg.From.NetID(), g.netProvider.Self().NetID(),
					recvMsg.Msg.MsgType, err,
				)
				continue
			}
			if acks[recvMsg.Msg.SenderIndex] { // Only consider first successful message.
				g.log.Warnf(
					"Dropping duplicate message %v -> %v, MsgType=%v",
					recvMsg.From.NetID(), g.netProvider.Self().NetID(),
					recvMsg.Msg.MsgType,
				)
				continue
			}
			if acks[recvMsg.Msg.SenderIndex], err = recvCB(recvMsg); err != nil {
				errs[recvMsg.Msg.SenderIndex] = err
				continue
			}
			if acks[recvMsg.Msg.SenderIndex] {
				// Clear previous errors on success.
				delete(errs, recvMsg.Msg.SenderIndex)
			}
		case <-retryCh:
			for i := range peers {
				if !acks[i] {
					sendCB(i, peers[i])
				}
			}
			retryCh = time.After(retryTimeout)
		case <-giveUpCh:
			var errMsg string
			for i := range peers {
				if acks[i] {
					continue
				}
				if errs[i] != nil {
					errMsg = errMsg + fmt.Sprintf("[%v:%v]", i, errs[i].Error())
				} else {
					errMsg = errMsg + fmt.Sprintf("[%v:%v]", i, "round_timeout")
				}
			}
			return errors.New(errMsg)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	var errMsg string
	for i := range errs {
		errMsg = errMsg + fmt.Sprintf("[%v:%v]", i, errs[i].Error())
	}
	return errors.New(errMsg)
}

// AllNodes returns a map of all nodes in the group.
func (g *groupImpl) AllNodes() map[uint16]peering.PeerSender {
	all := make(map[uint16]peering.PeerSender)
	for i := range g.nodes {
		all[uint16(i)] = g.nodes[i]
	}
	return all
}

// OtherNodes returns a map of all nodes in the group, excluding the self node.
func (g *groupImpl) OtherNodes() map[uint16]peering.PeerSender {
	return g.other
}

// Attach starts listening for messages. Messages in this case will be filtered
// to those received from nodes in the group only. SenderIndex will be filled
// for the messages according to the message source.
func (g *groupImpl) Attach(chainID *coretypes.ChainID, callback func(recv *peering.RecvEvent)) interface{} {
	return g.netProvider.Attach(chainID, func(recv *peering.RecvEvent) {
		if idx, err := g.PeerIndexByNetID(recv.From.NetID()); err == nil && idx != NotInGroup {
			recv.Msg.SenderIndex = idx
			callback(recv)
			return
		}
		g.log.Warnf("Dropping message MsgType=%v from %v, it does not belong to the group.", recv.Msg.MsgType, recv.From.NetID())
	})
}

// Detach terminates listening for messages.
func (g *groupImpl) Detach(attachID interface{}) {
	g.netProvider.Detach(attachID)
}

// Close implements peering.GroupProvider.
func (g *groupImpl) Close() {
	for i := range g.nodes {
		g.nodes[i].Close()
	}
}
