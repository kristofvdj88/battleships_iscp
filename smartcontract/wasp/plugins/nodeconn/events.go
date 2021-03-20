package nodeconn

import (
	"time"

	"github.com/iotaledger/goshimmer/dapps/waspconn/packages/waspconn"
	"github.com/iotaledger/goshimmer/packages/tangle"
	"github.com/iotaledger/hive.go/events"
)

var EventMessageReceived = events.NewEvent(param1Caller)

func param1Caller(handler interface{}, params ...interface{}) {
	handler.(func(interface{}))(params[0])
}

func msgDataToEvent(data []byte) {
	msg, err := waspconn.DecodeMsg(data, true)
	if err != nil {
		log.Errorf("wrong message from node: %v", err)
		return
	}

	//log.Debugf("received msg type %T data len = %d", msg, len(data))

	switch msgt := msg.(type) {
	case *waspconn.WaspMsgChunk:
		finalData, err := msgChopper.IncomingChunk(msgt.Data, tangle.MaxMessageSize, waspconn.ChunkMessageHeaderSize)
		if err != nil {
			log.Errorf("receiving message chunk: %v", err)
			return
		}
		if finalData != nil {
			msgDataToEvent(finalData)
		}

	case *waspconn.WaspPingMsg:
		roundtrip := time.Since(time.Unix(0, msgt.Timestamp))
		log.Infof("PING %d response from node. Roundtrip %v", msgt.Id, roundtrip)

	default:
		EventMessageReceived.Trigger(msgt)
	}
}
