package main

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/ethda"
	etrog "github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/etrogpolygonzkevm"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/0xPolygonHermez/zkevm-node/test/operations"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Deployments struct {
	ethdaContract *ethda.Ethda
	ethdaAddr     common.Address
}

func main() {
	ctx := context.Background()
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(operations.DefaultSequencerPrivateKey, "0x"))
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	deployments := DeployContractsAndSetDAP(client, auth)

	log.Infof("Deploy contract: %v", deployments.ethdaAddr.Hex())
}

func DeployContractsAndSetDAP(client *ethclient.Client, auth *bind.TransactOpts) Deployments {
	addr, tx, ethdaContract, err := ethda.DeployEthda(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	err = operations.WaitTxToBeMined(context.Background(), client, tx, 20*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = nil
	auth.Value = nil
	log.Debugf("Deploy tx: %v", tx.Hash().Hex())
	log.Debugf("Addr: %v", addr.Hex())
	// Set DAP
	log.Debugf("Setting DAP", operations.DefaultL1ZkEVMSmartContract, addr)
	etrog, err := etrog.NewEtrogpolygonzkevm(common.HexToAddress(operations.DefaultL1ZkEVMSmartContract), client)
	if err != nil {
		log.Fatal(err)
	}
	_, err = etrog.SetDataAvailabilityProtocol(auth, addr)
	if err != nil {
		log.Fatal(err)
	}

	return Deployments{
		ethdaContract: ethdaContract,
		ethdaAddr:     addr,
	}
}
