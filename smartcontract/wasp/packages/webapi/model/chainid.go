package model

import (
	"encoding/json"

	"github.com/iotaledger/wasp/packages/coretypes"
)

// ChainID is the base58 representation of coretypes.ChainID
type ChainID string

func NewChainID(chainID *coretypes.ChainID) ChainID {
	return ChainID(chainID.String())
}

func (ch ChainID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ch))
}

func (ch *ChainID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	_, err := coretypes.NewChainIDFromBase58(s)
	*ch = ChainID(s)
	return err
}

func (ch ChainID) ChainID() coretypes.ChainID {
	chainID, err := coretypes.NewChainIDFromBase58(string(ch))
	if err != nil {
		panic(err)
	}
	return chainID
}
