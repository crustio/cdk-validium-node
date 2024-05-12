// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethda

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EthdaMetaData contains all meta data concerning the Ethda contract.
var EthdaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"UnexpectedAddrsAndSignaturesSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongSignature\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ethdaSequencerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProcotolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethdaSequencerAddress\",\"type\":\"address\"}],\"name\":\"setEthdaSequencerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signaturesAndAddrs\",\"type\":\"bytes\"}],\"name\":\"verifyMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100da565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811610156100d8575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610bce806100e75f395ff3fe608060405234801561000f575f80fd5b5060043610610085575f3560e01c80638da5cb5b116100585780638da5cb5b146100f8578063d5deff3014610116578063e4f1712014610129578063f2fde38b14610168575f80fd5b80633b51be4b14610089578063715018a61461009e5780637e585d3e146100a65780638129fc1c146100f0575b5f80fd5b61009c6100973660046109bc565b61017b565b005b61009c61027f565b6065546100c69073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61009c610292565b60335473ffffffffffffffffffffffffffffffffffffffff166100c6565b61009c610124366004610a31565b610423565b604080518082018252600581527f4574686461000000000000000000000000000000000000000000000000000000602082015290516100e79190610a6b565b61009c610176366004610a31565b610472565b60418082108061019f575060206101928284610ad4565b61019c9190610b0c565b15155b156101d6576040517f6b8eec4600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610221856101e86041848789610b44565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061052692505050565b60655490915073ffffffffffffffffffffffffffffffffffffffff808316911614610278576040517f356a441800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b61028761054a565b6102905f6105cb565b565b5f54610100900460ff16158080156102b057505f54600160ff909116105b806102c95750303b1580156102c957505f5460ff166001145b61035a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156103b6575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6103be610641565b8015610420575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b61042b61054a565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b61047a61054a565b73ffffffffffffffffffffffffffffffffffffffff811661051d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610351565b610420816105cb565b5f805f61053385856106e0565b9150915061054081610722565b5090505b92915050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610290576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610351565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff166106d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610351565b610290336105cb565b5f808251604103610714576020830151604084015160608501515f1a610708878285856108d4565b9450945050505061071b565b505f905060025b9250929050565b5f81600481111561073557610735610b6b565b0361073d5750565b600181600481111561075157610751610b6b565b036107b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610351565b60028160048111156107cc576107cc610b6b565b03610833576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610351565b600381600481111561084757610847610b6b565b03610420576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610351565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561090957505f905060036109b3565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561095a573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166109ad575f600192509250506109b3565b91505f90505b94509492505050565b5f805f604084860312156109ce575f80fd5b83359250602084013567ffffffffffffffff808211156109ec575f80fd5b818601915086601f8301126109ff575f80fd5b813581811115610a0d575f80fd5b876020828501011115610a1e575f80fd5b6020830194508093505050509250925092565b5f60208284031215610a41575f80fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610a64575f80fd5b9392505050565b5f6020808352835180828501525f5b81811015610a9657858101830151858201604001528201610a7a565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b81810381811115610544577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f82610b3f577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500690565b5f8085851115610b52575f80fd5b83861115610b5e575f80fd5b5050820193919092039150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea264697066735822122010fdd197f73d092becd782054513afc2e54a220206a28c0f0f39e31fe955dcca64736f6c63430008140033",
}

// EthdaABI is the input ABI used to generate the binding from.
// Deprecated: Use EthdaMetaData.ABI instead.
var EthdaABI = EthdaMetaData.ABI

// EthdaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthdaMetaData.Bin instead.
var EthdaBin = EthdaMetaData.Bin

// DeployEthda deploys a new Ethereum contract, binding an instance of Ethda to it.
func DeployEthda(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ethda, error) {
	parsed, err := EthdaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthdaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethda{EthdaCaller: EthdaCaller{contract: contract}, EthdaTransactor: EthdaTransactor{contract: contract}, EthdaFilterer: EthdaFilterer{contract: contract}}, nil
}

// Ethda is an auto generated Go binding around an Ethereum contract.
type Ethda struct {
	EthdaCaller     // Read-only binding to the contract
	EthdaTransactor // Write-only binding to the contract
	EthdaFilterer   // Log filterer for contract events
}

// EthdaCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthdaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthdaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthdaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthdaSession struct {
	Contract     *Ethda            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthdaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthdaCallerSession struct {
	Contract *EthdaCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthdaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthdaTransactorSession struct {
	Contract     *EthdaTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthdaRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthdaRaw struct {
	Contract *Ethda // Generic contract binding to access the raw methods on
}

// EthdaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthdaCallerRaw struct {
	Contract *EthdaCaller // Generic read-only contract binding to access the raw methods on
}

// EthdaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthdaTransactorRaw struct {
	Contract *EthdaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthda creates a new instance of Ethda, bound to a specific deployed contract.
func NewEthda(address common.Address, backend bind.ContractBackend) (*Ethda, error) {
	contract, err := bindEthda(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethda{EthdaCaller: EthdaCaller{contract: contract}, EthdaTransactor: EthdaTransactor{contract: contract}, EthdaFilterer: EthdaFilterer{contract: contract}}, nil
}

// NewEthdaCaller creates a new read-only instance of Ethda, bound to a specific deployed contract.
func NewEthdaCaller(address common.Address, caller bind.ContractCaller) (*EthdaCaller, error) {
	contract, err := bindEthda(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthdaCaller{contract: contract}, nil
}

// NewEthdaTransactor creates a new write-only instance of Ethda, bound to a specific deployed contract.
func NewEthdaTransactor(address common.Address, transactor bind.ContractTransactor) (*EthdaTransactor, error) {
	contract, err := bindEthda(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthdaTransactor{contract: contract}, nil
}

// NewEthdaFilterer creates a new log filterer instance of Ethda, bound to a specific deployed contract.
func NewEthdaFilterer(address common.Address, filterer bind.ContractFilterer) (*EthdaFilterer, error) {
	contract, err := bindEthda(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthdaFilterer{contract: contract}, nil
}

// bindEthda binds a generic wrapper to an already deployed contract.
func bindEthda(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthdaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethda *EthdaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethda.Contract.EthdaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethda *EthdaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethda.Contract.EthdaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethda *EthdaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethda.Contract.EthdaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethda *EthdaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethda.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethda *EthdaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethda.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethda *EthdaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethda.Contract.contract.Transact(opts, method, params...)
}

// EthdaSequencerAddress is a free data retrieval call binding the contract method 0x7e585d3e.
//
// Solidity: function ethdaSequencerAddress() view returns(address)
func (_Ethda *EthdaCaller) EthdaSequencerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethda.contract.Call(opts, &out, "ethdaSequencerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthdaSequencerAddress is a free data retrieval call binding the contract method 0x7e585d3e.
//
// Solidity: function ethdaSequencerAddress() view returns(address)
func (_Ethda *EthdaSession) EthdaSequencerAddress() (common.Address, error) {
	return _Ethda.Contract.EthdaSequencerAddress(&_Ethda.CallOpts)
}

// EthdaSequencerAddress is a free data retrieval call binding the contract method 0x7e585d3e.
//
// Solidity: function ethdaSequencerAddress() view returns(address)
func (_Ethda *EthdaCallerSession) EthdaSequencerAddress() (common.Address, error) {
	return _Ethda.Contract.EthdaSequencerAddress(&_Ethda.CallOpts)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Ethda *EthdaCaller) GetProcotolName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ethda.contract.Call(opts, &out, "getProcotolName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Ethda *EthdaSession) GetProcotolName() (string, error) {
	return _Ethda.Contract.GetProcotolName(&_Ethda.CallOpts)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Ethda *EthdaCallerSession) GetProcotolName() (string, error) {
	return _Ethda.Contract.GetProcotolName(&_Ethda.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethda *EthdaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethda.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethda *EthdaSession) Owner() (common.Address, error) {
	return _Ethda.Contract.Owner(&_Ethda.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethda *EthdaCallerSession) Owner() (common.Address, error) {
	return _Ethda.Contract.Owner(&_Ethda.CallOpts)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 signedHash, bytes signaturesAndAddrs) view returns()
func (_Ethda *EthdaCaller) VerifyMessage(opts *bind.CallOpts, signedHash [32]byte, signaturesAndAddrs []byte) error {
	var out []interface{}
	err := _Ethda.contract.Call(opts, &out, "verifyMessage", signedHash, signaturesAndAddrs)

	if err != nil {
		return err
	}

	return err

}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 signedHash, bytes signaturesAndAddrs) view returns()
func (_Ethda *EthdaSession) VerifyMessage(signedHash [32]byte, signaturesAndAddrs []byte) error {
	return _Ethda.Contract.VerifyMessage(&_Ethda.CallOpts, signedHash, signaturesAndAddrs)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 signedHash, bytes signaturesAndAddrs) view returns()
func (_Ethda *EthdaCallerSession) VerifyMessage(signedHash [32]byte, signaturesAndAddrs []byte) error {
	return _Ethda.Contract.VerifyMessage(&_Ethda.CallOpts, signedHash, signaturesAndAddrs)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Ethda *EthdaTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethda.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Ethda *EthdaSession) Initialize() (*types.Transaction, error) {
	return _Ethda.Contract.Initialize(&_Ethda.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Ethda *EthdaTransactorSession) Initialize() (*types.Transaction, error) {
	return _Ethda.Contract.Initialize(&_Ethda.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethda *EthdaTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethda.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethda *EthdaSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethda.Contract.RenounceOwnership(&_Ethda.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethda *EthdaTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethda.Contract.RenounceOwnership(&_Ethda.TransactOpts)
}

// SetEthdaSequencerAddress is a paid mutator transaction binding the contract method 0xd5deff30.
//
// Solidity: function setEthdaSequencerAddress(address _ethdaSequencerAddress) returns()
func (_Ethda *EthdaTransactor) SetEthdaSequencerAddress(opts *bind.TransactOpts, _ethdaSequencerAddress common.Address) (*types.Transaction, error) {
	return _Ethda.contract.Transact(opts, "setEthdaSequencerAddress", _ethdaSequencerAddress)
}

// SetEthdaSequencerAddress is a paid mutator transaction binding the contract method 0xd5deff30.
//
// Solidity: function setEthdaSequencerAddress(address _ethdaSequencerAddress) returns()
func (_Ethda *EthdaSession) SetEthdaSequencerAddress(_ethdaSequencerAddress common.Address) (*types.Transaction, error) {
	return _Ethda.Contract.SetEthdaSequencerAddress(&_Ethda.TransactOpts, _ethdaSequencerAddress)
}

// SetEthdaSequencerAddress is a paid mutator transaction binding the contract method 0xd5deff30.
//
// Solidity: function setEthdaSequencerAddress(address _ethdaSequencerAddress) returns()
func (_Ethda *EthdaTransactorSession) SetEthdaSequencerAddress(_ethdaSequencerAddress common.Address) (*types.Transaction, error) {
	return _Ethda.Contract.SetEthdaSequencerAddress(&_Ethda.TransactOpts, _ethdaSequencerAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethda *EthdaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ethda.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethda *EthdaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethda.Contract.TransferOwnership(&_Ethda.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethda *EthdaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethda.Contract.TransferOwnership(&_Ethda.TransactOpts, newOwner)
}

// EthdaInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ethda contract.
type EthdaInitializedIterator struct {
	Event *EthdaInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthdaInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthdaInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthdaInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthdaInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthdaInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthdaInitialized represents a Initialized event raised by the Ethda contract.
type EthdaInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ethda *EthdaFilterer) FilterInitialized(opts *bind.FilterOpts) (*EthdaInitializedIterator, error) {

	logs, sub, err := _Ethda.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EthdaInitializedIterator{contract: _Ethda.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ethda *EthdaFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EthdaInitialized) (event.Subscription, error) {

	logs, sub, err := _Ethda.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthdaInitialized)
				if err := _Ethda.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ethda *EthdaFilterer) ParseInitialized(log types.Log) (*EthdaInitialized, error) {
	event := new(EthdaInitialized)
	if err := _Ethda.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthdaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ethda contract.
type EthdaOwnershipTransferredIterator struct {
	Event *EthdaOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthdaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthdaOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthdaOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthdaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthdaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthdaOwnershipTransferred represents a OwnershipTransferred event raised by the Ethda contract.
type EthdaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethda *EthdaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthdaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethda.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthdaOwnershipTransferredIterator{contract: _Ethda.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethda *EthdaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthdaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethda.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthdaOwnershipTransferred)
				if err := _Ethda.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethda *EthdaFilterer) ParseOwnershipTransferred(log types.Log) (*EthdaOwnershipTransferred, error) {
	event := new(EthdaOwnershipTransferred)
	if err := _Ethda.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
