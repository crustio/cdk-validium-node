package ethda

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	blobutils "github.com/crustio/blob-utils"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EthdaBackend struct {
	ethdaClient *blobutils.Client
	chainId     *big.Int
	ethdaRPCURL string
	privKey     *ecdsa.PrivateKey
}

func New(
	ethdaRPCURL string,
	privKey *ecdsa.PrivateKey,
) (*EthdaBackend, error) {
	if ethdaRPCURL == "" {
		return nil, fmt.Errorf("empty ethda rpc url")
	}

	if privKey == nil {
		return nil, fmt.Errorf("empty private key")
	}

	return &EthdaBackend{
		privKey:     privKey,
		ethdaRPCURL: ethdaRPCURL,
	}, nil
}

func (d *EthdaBackend) Init() error {
	ethdaClient, err := blobutils.New(d.ethdaRPCURL, d.privKey)
	if err != nil {
		return fmt.Errorf("create ethda client: %w", err)
	}

	d.ethdaClient = ethdaClient

	return nil
}

func (d *EthdaBackend) PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, error) {
	var hashes []byte
	for _, batch := range batchesData {
		hash, err := d.ethdaClient.PostBlob(ctx, batch)
		if err != nil {
			return nil, fmt.Errorf("post batch to ethda: %w", err)
		}
		hashes = append(hashes, hash.Bytes()...)
	}

	return hashes, nil
}

func (d *EthdaBackend) GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	if len(dataAvailabilityMessage)%common.HashLength != 0 {
		return nil, fmt.Errorf("wrong da message length: %d", len(dataAvailabilityMessage))
	}

	var data [][]byte
	for i := 0; i < len(dataAvailabilityMessage)/common.HashLength; i++ {
		start := common.HashLength * i
		hash := common.BytesToHash(dataAvailabilityMessage[start : start+common.HashLength])

		r, err := d.ethdaClient.GetBlob(hash)
		if err != nil {
			return nil, fmt.Errorf("get blob from ethda: %w", err)
		}
		data = append(data, r)
	}

	return data, nil
}
