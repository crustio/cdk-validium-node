package ethda

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostSequence(t *testing.T) {
	rpcUrl := "https://rpc-devnet2.ethda.io"

	// test key
	key, err := crypto.HexToECDSA("f26d6aca18e0c75ac948a262d4b9435a8173515f84c258d1f90d171143039024")
	require.NoError(t, err)

	da, err := New(rpcUrl, key)
	require.NoError(t, err)

	err = da.Init()
	require.NoError(t, err)

	batchData := [][]byte{[]byte("ethda")}
	msg, err := da.PostSequence(context.Background(), batchData)
	require.NoError(t, err)
	fmt.Println(common.Bytes2Hex(msg))
}

func TestGetSequence(t *testing.T) {
	rpcUrl := "https://rpc-devnet2.ethda.io"

	// test key
	key, err := crypto.HexToECDSA("f26d6aca18e0c75ac948a262d4b9435a8173515f84c258d1f90d171143039024")
	require.NoError(t, err)
	da, err := New(rpcUrl, key)
	require.NoError(t, err)

	err = da.Init()
	require.NoError(t, err)

	data, err := da.GetSequence(context.Background(), nil, common.Hex2Bytes("76b56f65beab1cfe55242eb97eebfe2f32aacd957f202122835026bffb3ba282"))
	require.NoError(t, err)
	require.Equal(t, "ethda", string(data[0]))
}
