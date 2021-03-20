// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package chain

import (
	"fmt"
	"io"

	"github.com/iotaledger/goshimmer/dapps/waspconn/packages/waspconn"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/util"
)

func (msg *StateIndexPingPongMsg) Write(w io.Writer) error {
	if err := util.WriteUint32(w, msg.BlockIndex); err != nil {
		return err
	}
	return util.WriteBoolByte(w, msg.RSVP)
}

func (msg *StateIndexPingPongMsg) Read(r io.Reader) error {
	if err := util.ReadUint32(r, &msg.BlockIndex); err != nil {
		return err
	}
	return util.ReadBoolByte(r, &msg.RSVP)
}

func (msg *NotifyReqMsg) Write(w io.Writer) error {
	if err := util.WriteUint32(w, msg.BlockIndex); err != nil {
		return err
	}
	if err := util.WriteUint16(w, uint16(len(msg.RequestIDs))); err != nil {
		return err
	}
	for _, reqid := range msg.RequestIDs {
		if _, err := w.Write(reqid[:]); err != nil {
			return err
		}
	}
	return nil
}

func (msg *NotifyReqMsg) Read(r io.Reader) error {
	err := util.ReadUint32(r, &msg.BlockIndex)
	if err != nil {
		return err
	}
	var arrLen uint16
	err = util.ReadUint16(r, &arrLen)
	if err != nil {
		return err
	}
	if arrLen == 0 {
		return nil
	}
	msg.RequestIDs = make([]coretypes.RequestID, arrLen)
	for i := range msg.RequestIDs {
		_, err = r.Read(msg.RequestIDs[i][:])
		if err != nil {
			return err
		}
	}
	return nil
}

func (msg *NotifyFinalResultPostedMsg) Write(w io.Writer) error {
	if err := util.WriteUint32(w, msg.BlockIndex); err != nil {
		return err
	}
	if _, err := w.Write(msg.TxId.Bytes()); err != nil {
		return err
	}
	return nil
}

func (msg *NotifyFinalResultPostedMsg) Read(r io.Reader) error {
	err := util.ReadUint32(r, &msg.BlockIndex)
	if err != nil {
		return err
	}
	if err := util.ReadTransactionId(r, &msg.TxId); err != nil {
		return err
	}
	return nil
}

func (msg *StartProcessingBatchMsg) Write(w io.Writer) error {
	if err := util.WriteUint32(w, msg.BlockIndex); err != nil {
		return err
	}
	if err := util.WriteUint16(w, uint16(len(msg.RequestIds))); err != nil {
		return err
	}
	for i := range msg.RequestIds {
		if _, err := w.Write(msg.RequestIds[i][:]); err != nil {
			return err
		}
	}
	if _, err := w.Write(msg.FeeDestination[:]); err != nil {
		return err
	}
	if err := waspconn.WriteBalances(w, msg.Balances); err != nil {
		return err
	}
	return nil
}

func (msg *StartProcessingBatchMsg) Read(r io.Reader) error {
	if err := util.ReadUint32(r, &msg.BlockIndex); err != nil {
		return err
	}
	var size uint16
	if err := util.ReadUint16(r, &size); err != nil {
		return err
	}
	msg.RequestIds = make([]coretypes.RequestID, size)
	for i := range msg.RequestIds {
		if err := msg.RequestIds[i].Read(r); err != nil {
			return err
		}
	}
	if err := coretypes.ReadAgentID(r, &msg.FeeDestination); err != nil {
		return err
	}
	var err error
	if msg.Balances, err = waspconn.ReadBalances(r); err != nil {
		return err
	}
	return nil
}

func (msg *SignedHashMsg) Write(w io.Writer) error {
	if err := util.WriteUint32(w, msg.BlockIndex); err != nil {
		return err
	}
	if err := util.WriteUint64(w, uint64(msg.OrigTimestamp)); err != nil {
		return err
	}
	if _, err := w.Write(msg.BatchHash[:]); err != nil {
		return err
	}
	if _, err := w.Write(msg.EssenceHash[:]); err != nil {
		return err
	}
	if err := util.WriteBytes16(w, msg.SigShare); err != nil {
		return err
	}
	return nil
}

func (msg *SignedHashMsg) Read(r io.Reader) error {
	if err := util.ReadUint32(r, &msg.BlockIndex); err != nil {
		return err
	}
	var ts uint64
	if err := util.ReadUint64(r, &ts); err != nil {
		return err
	}
	msg.OrigTimestamp = int64(ts)

	if err := util.ReadHashValue(r, &msg.BatchHash); err != nil {
		return err
	}
	if err := util.ReadHashValue(r, &msg.EssenceHash); err != nil {
		return err
	}
	var err error
	if msg.SigShare, err = util.ReadBytes16(r); err != nil {
		return err
	}
	return nil
}

func (msg *GetBlockMsg) Write(w io.Writer) error {
	return util.WriteUint32(w, msg.BlockIndex)
}

func (msg *GetBlockMsg) Read(r io.Reader) error {
	return util.ReadUint32(r, &msg.BlockIndex)
}

func (msg *BlockHeaderMsg) Write(w io.Writer) error {
	if err := util.WriteUint32(w, msg.BlockIndex); err != nil {
		return err
	}
	if err := util.WriteUint16(w, msg.Size); err != nil {
		return err
	}
	if _, err := w.Write(msg.AnchorTransactionID.Bytes()); err != nil {
		return err
	}
	return nil
}

func (msg *BlockHeaderMsg) Read(r io.Reader) error {
	if err := util.ReadUint32(r, &msg.BlockIndex); err != nil {
		return err
	}
	if err := util.ReadUint16(r, &msg.Size); err != nil {
		return err
	}
	if _, err := r.Read(msg.AnchorTransactionID[:]); err != nil {
		return err
	}
	return nil
}

func (msg *StateUpdateMsg) Write(w io.Writer) error {
	if err := util.WriteUint32(w, msg.BlockIndex); err != nil {
		return err
	}
	if err := msg.StateUpdate.Write(w); err != nil {
		return err
	}
	if err := util.WriteUint16(w, msg.IndexInTheBlock); err != nil {
		return err
	}
	return nil
}

func (msg *StateUpdateMsg) Read(r io.Reader) error {
	if err := util.ReadUint32(r, &msg.BlockIndex); err != nil {
		return err
	}
	msg.StateUpdate = state.NewStateUpdate(nil)
	if err := msg.StateUpdate.Read(r); err != nil {
		return err
	}
	if err := util.ReadUint16(r, &msg.IndexInTheBlock); err != nil {
		return err
	}
	return nil
}

func (msg *TestTraceMsg) Write(w io.Writer) error {
	if !util.ValidPermutation(msg.Sequence) {
		panic(fmt.Sprintf("Write: wrong permutation %+v", msg.Sequence))
	}
	if err := util.WriteUint64(w, uint64(msg.InitTime)); err != nil {
		return err
	}
	if err := util.WriteUint16(w, msg.InitPeerIndex); err != nil {
		return err
	}
	if err := util.WriteUint16(w, uint16(len(msg.Sequence))); err != nil {
		return err
	}
	for _, idx := range msg.Sequence {
		if err := util.WriteUint16(w, idx); err != nil {
			return err
		}
	}
	if err := util.WriteUint16(w, msg.NumHops); err != nil {
		return err
	}
	return nil
}

func (msg *TestTraceMsg) Read(r io.Reader) error {
	var initTime uint64
	if err := util.ReadUint64(r, &initTime); err != nil {
		return err
	}
	msg.InitTime = int64(initTime)
	if err := util.ReadUint16(r, &msg.InitPeerIndex); err != nil {
		return err
	}
	var size uint16
	if err := util.ReadUint16(r, &size); err != nil {
		return err
	}
	msg.Sequence = make([]uint16, size)
	for i := range msg.Sequence {
		if err := util.ReadUint16(r, &msg.Sequence[i]); err != nil {
			return err
		}
	}
	if err := util.ReadUint16(r, &msg.NumHops); err != nil {
		return err
	}
	if !util.ValidPermutation(msg.Sequence) {
		panic(fmt.Sprintf("Read: wrong permutation %+v", msg.Sequence))
	}
	return nil
}
