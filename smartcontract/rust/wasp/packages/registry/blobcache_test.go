package registry

import (
	"github.com/iotaledger/wasp/packages/dbprovider"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBlobPutGet(t *testing.T) {
	log := testutil.NewLogger(t)
	db := dbprovider.NewInMemoryDBProvider(log)
	reg := NewRegistry(nil, log, db)

	data := []byte("data-data-data-data-data-data-data-data-data")
	h := hashing.HashData(data)

	hback, err := reg.PutBlob(data)
	require.NoError(t, err)
	require.EqualValues(t, h, hback)

	back, ok, err := reg.GetBlob(h)
	require.NoError(t, err)
	require.True(t, ok)
	require.EqualValues(t, data, back)
}
