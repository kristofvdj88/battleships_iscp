// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package consensus

import (
	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/util"
)

// sendRequestNotificationsToLeader sends current leader the backlog of requests
// it is only possible in the `consensusStageLeaderStarting` stage for non-leader
func (op *operator) sendRequestNotificationsToLeader() {
	if len(op.requests) == 0 {
		return
	}
	if op.iAmCurrentLeader() {
		return
	}
	if op.consensusStage != consensusStageSubStarting {
		return
	}
	if !op.chain.HasQuorum() {
		op.log.Debugf("sendRequestNotificationsToLeader: postponed due to no quorum. Peer status: %s",
			op.chain.PeerStatus())
		return
	}
	currentLeaderPeerIndex, _ := op.currentLeader()
	reqs := op.requestCandidateList()
	//reqs = op.filterOutRequestsWithoutTokens(reqs)

	// get not time-locked requests with the message known
	if len(reqs) == 0 {
		// nothing to notify about
		return
	}
	op.log.Debugf("sending notifications to #%d, backlog: %d, candidates (with tokens): %d",
		currentLeaderPeerIndex, len(op.requests), len(reqs))

	reqIds := takeIds(reqs)
	msgData := util.MustBytes(&chain.NotifyReqMsg{
		PeerMsgHeader: chain.PeerMsgHeader{
			BlockIndex: op.mustStateIndex(),
		},
		RequestIDs: reqIds,
	})

	// send until first success, but no more than number of nodes in the committee
	op.log.Infow("sendRequestNotificationsToLeader",
		"leader", currentLeaderPeerIndex,
		"state index", op.mustStateIndex(),
		"reqs", idsShortStr(reqIds),
	)
	if err := op.chain.SendMsg(currentLeaderPeerIndex, chain.MsgNotifyRequests, msgData); err != nil {
		op.log.Errorf("sending notifications to %d: %v", currentLeaderPeerIndex, err)
	}
	op.setNextConsensusStage(consensusStageSubNotificationsSent)
}

func (op *operator) storeNotification(msg *chain.NotifyReqMsg) {
	stateIndex, stateDefined := op.blockIndex()
	if stateDefined && msg.BlockIndex < stateIndex {
		// don't save from earlier. The current currentState saved only for tracking
		return
	}
	op.notificationsBacklog = append(op.notificationsBacklog, msg)
}

// markRequestsNotified stores information about notification in the current currentState
func (op *operator) markRequestsNotified(msgs []*chain.NotifyReqMsg) {
	stateIndex, stateDefined := op.blockIndex()
	if !stateDefined {
		return
	}
	for _, msg := range msgs {
		if msg.BlockIndex != stateIndex {
			continue
		}
		for _, reqid := range msg.RequestIDs {
			req, ok := op.requestFromId(reqid)
			if !ok {
				continue
			}
			// mark request was seen by sender
			req.notifications[msg.SenderIndex] = true
		}
	}
}

// adjust all notification information to the current state index
func (op *operator) adjustNotifications() {
	stateIndex, stateDefined := op.blockIndex()
	if !stateDefined {
		return
	}
	// clear all the notification markers
	for _, req := range op.requests {
		setAllFalse(req.notifications)
		req.notifications[op.peerIndex()] = req.reqTx != nil
	}
	// put markers of the current state
	op.markRequestsNotified(op.notificationsBacklog)

	// clean notification backlog from messages from current and and past stages
	newBacklog := op.notificationsBacklog[:0] // new slice, same underlying array!
	for _, msg := range op.notificationsBacklog {
		if msg.BlockIndex < stateIndex {
			continue
		}
		newBacklog = append(newBacklog, msg)
	}
	op.notificationsBacklog = newBacklog
}

func setAllFalse(bs []bool) {
	for i := range bs {
		bs[i] = false
	}
}
