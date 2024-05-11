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
	daMessage := "b1e79a75baadbd688f97213920a01980e37e4fbf464cf9ffc94e6d10be2ee69253f1acb077c8355d30f4d8046ccbe4eb5f4eb05643ef31c4fe6e531c0ca7dc8e1bdb919f0ef6ffe754f3dfdfdfc4d1789b0bf2aee713567b550d330527eab00522"
	rpcUrl := "https://rpc-devnet2.ethda.io"

	// test key
	key, err := crypto.HexToECDSA("f26d6aca18e0c75ac948a262d4b9435a8173515f84c258d1f90d171143039024")
	require.NoError(t, err)
	da, err := New(rpcUrl, key)
	require.NoError(t, err)

	err = da.Init()
	require.NoError(t, err)

	data, err := da.GetSequence(context.Background(), nil, common.Hex2Bytes(daMessage))
	require.NoError(t, err)
	require.Equal(t, "ethda", string(data[0]))
}
