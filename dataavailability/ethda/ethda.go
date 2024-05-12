package ethda

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/0xPolygonHermez/zkevm-node/log"
	blobutils "github.com/crustio/blob-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
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

	return &EthdaBackend{
		privKey:     privKey,
		ethdaRPCURL: ethdaRPCURL,
	}, nil
}

func (d *EthdaBackend) Init() error {
	if d.privKey == nil {
		return nil
	}

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
		log.Infof("Post da to ethda: %s", hash.Hex())
	}

	currentHash := common.Hash{}.Bytes()
	for _, batchData := range batchesData {
		types := []string{
			"bytes32",
			"bytes32",
		}
		values := []interface{}{
			currentHash,
			crypto.Keccak256(batchData),
		}
		currentHash = solsha3.SoliditySHA3(types, values)
	}

	sig, err := d.ethdaClient.SignBatchHash(common.BytesToHash(currentHash))
	if err != nil {
		return nil, err
	}

	return append(sig, hashes...), nil
}

func (d *EthdaBackend) GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	msgLen := len(dataAvailabilityMessage)

	if msgLen < crypto.SignatureLength || (msgLen-crypto.SignatureLength)%common.HashLength != 0 {
		return nil, fmt.Errorf("wrong da message length: %d", msgLen)
	}

	var data [][]byte
	for i := 0; i < (msgLen-crypto.SignatureLength)/common.HashLength; i++ {
		start := common.HashLength*i + crypto.SignatureLength
		hash := common.BytesToHash(dataAvailabilityMessage[start : start+common.HashLength])

		r, err := d.ethdaClient.GetBlob(hash)
		if err != nil {
			return nil, fmt.Errorf("get blob from ethda: %w", err)
		}
		data = append(data, r)
	}

	return data, nil
}
