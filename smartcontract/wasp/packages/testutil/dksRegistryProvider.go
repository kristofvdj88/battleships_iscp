// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"fmt"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/wasp/packages/tcrypto"
)

// DkgRegistryProvider stands for a mock for dkg.RegistryProvider.
type DkgRegistryProvider struct {
	DB    map[string][]byte
	Suite tcrypto.Suite
}

// NewDkgRegistryProvider creates new mocked DKG registry provider.
func NewDkgRegistryProvider(suite tcrypto.Suite) *DkgRegistryProvider {
	return &DkgRegistryProvider{
		DB:    map[string][]byte{},
		Suite: suite,
	}
}

// SaveDKShare implements dkg.RegistryProvider.
func (p *DkgRegistryProvider) SaveDKShare(dkShare *tcrypto.DKShare) error {
	var err error
	var dkShareBytes []byte
	if dkShareBytes, err = dkShare.Bytes(); err != nil {
		return err
	}
	p.DB[dkShare.Address.String()] = dkShareBytes
	return nil
}

// LoadDKShare implements dkg.RegistryProvider.
func (p *DkgRegistryProvider) LoadDKShare(sharedAddress *address.Address) (*tcrypto.DKShare, error) {
	var dkShareBytes = p.DB[sharedAddress.String()]
	if dkShareBytes == nil {
		return nil, fmt.Errorf("DKShare not found for %v", sharedAddress)
	}
	return tcrypto.DKShareFromBytes(dkShareBytes, p.Suite)
}
