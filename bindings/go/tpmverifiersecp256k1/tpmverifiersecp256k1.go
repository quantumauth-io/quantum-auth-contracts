// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tpmverifiersecp256k1

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

// TPMVerifierSecp256k1MetaData contains all meta data concerning the TPMVerifierSecp256k1 contract.
var TPMVerifierSecp256k1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BadSigLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102638061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063de12c6401461002d575b5f5ffd5b61004061003b366004610186565b610054565b604051901515815260200160405180910390f35b5f846001600160a01b03811661006d575f915050610132565b6041831461007e575f915050610132565b5f5f5f61008b878761013a565b9250925092505f6001898386866040515f81526020016040526040516100cd949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156100ed573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610115575f95505050505050610132565b846001600160a01b0316816001600160a01b031614955050505050505b949350505050565b8135602083013560408401355f1a601b81101561015f5761015c601b82610202565b90505b8060ff16601b1415801561017757508060ff16601c14155b1561017f57505f5b9250925092565b5f5f5f5f60608587031215610199575f5ffd5b8435935060208501359250604085013567ffffffffffffffff8111156101bd575f5ffd5b8501601f810187136101cd575f5ffd5b803567ffffffffffffffff8111156101e3575f5ffd5b8760208284010111156101f4575f5ffd5b949793965060200194505050565b60ff818116838216019081111561022757634e487b7160e01b5f52601160045260245ffd5b9291505056fea26469706673582212205ed1d85925916de2e12d43ae2ef5b0dc27576a9e537681641fd292e3969f518464736f6c634300081c0033",
}

// TPMVerifierSecp256k1ABI is the input ABI used to generate the binding from.
// Deprecated: Use TPMVerifierSecp256k1MetaData.ABI instead.
var TPMVerifierSecp256k1ABI = TPMVerifierSecp256k1MetaData.ABI

// TPMVerifierSecp256k1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TPMVerifierSecp256k1MetaData.Bin instead.
var TPMVerifierSecp256k1Bin = TPMVerifierSecp256k1MetaData.Bin

// DeployTPMVerifierSecp256k1 deploys a new Ethereum contract, binding an instance of TPMVerifierSecp256k1 to it.
func DeployTPMVerifierSecp256k1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TPMVerifierSecp256k1, error) {
	parsed, err := TPMVerifierSecp256k1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TPMVerifierSecp256k1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TPMVerifierSecp256k1{TPMVerifierSecp256k1Caller: TPMVerifierSecp256k1Caller{contract: contract}, TPMVerifierSecp256k1Transactor: TPMVerifierSecp256k1Transactor{contract: contract}, TPMVerifierSecp256k1Filterer: TPMVerifierSecp256k1Filterer{contract: contract}}, nil
}

// TPMVerifierSecp256k1 is an auto generated Go binding around an Ethereum contract.
type TPMVerifierSecp256k1 struct {
	TPMVerifierSecp256k1Caller     // Read-only binding to the contract
	TPMVerifierSecp256k1Transactor // Write-only binding to the contract
	TPMVerifierSecp256k1Filterer   // Log filterer for contract events
}

// TPMVerifierSecp256k1Caller is an auto generated read-only Go binding around an Ethereum contract.
type TPMVerifierSecp256k1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TPMVerifierSecp256k1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TPMVerifierSecp256k1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TPMVerifierSecp256k1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TPMVerifierSecp256k1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TPMVerifierSecp256k1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TPMVerifierSecp256k1Session struct {
	Contract     *TPMVerifierSecp256k1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TPMVerifierSecp256k1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TPMVerifierSecp256k1CallerSession struct {
	Contract *TPMVerifierSecp256k1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TPMVerifierSecp256k1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TPMVerifierSecp256k1TransactorSession struct {
	Contract     *TPMVerifierSecp256k1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TPMVerifierSecp256k1Raw is an auto generated low-level Go binding around an Ethereum contract.
type TPMVerifierSecp256k1Raw struct {
	Contract *TPMVerifierSecp256k1 // Generic contract binding to access the raw methods on
}

// TPMVerifierSecp256k1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TPMVerifierSecp256k1CallerRaw struct {
	Contract *TPMVerifierSecp256k1Caller // Generic read-only contract binding to access the raw methods on
}

// TPMVerifierSecp256k1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TPMVerifierSecp256k1TransactorRaw struct {
	Contract *TPMVerifierSecp256k1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTPMVerifierSecp256k1 creates a new instance of TPMVerifierSecp256k1, bound to a specific deployed contract.
func NewTPMVerifierSecp256k1(address common.Address, backend bind.ContractBackend) (*TPMVerifierSecp256k1, error) {
	contract, err := bindTPMVerifierSecp256k1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TPMVerifierSecp256k1{TPMVerifierSecp256k1Caller: TPMVerifierSecp256k1Caller{contract: contract}, TPMVerifierSecp256k1Transactor: TPMVerifierSecp256k1Transactor{contract: contract}, TPMVerifierSecp256k1Filterer: TPMVerifierSecp256k1Filterer{contract: contract}}, nil
}

// NewTPMVerifierSecp256k1Caller creates a new read-only instance of TPMVerifierSecp256k1, bound to a specific deployed contract.
func NewTPMVerifierSecp256k1Caller(address common.Address, caller bind.ContractCaller) (*TPMVerifierSecp256k1Caller, error) {
	contract, err := bindTPMVerifierSecp256k1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TPMVerifierSecp256k1Caller{contract: contract}, nil
}

// NewTPMVerifierSecp256k1Transactor creates a new write-only instance of TPMVerifierSecp256k1, bound to a specific deployed contract.
func NewTPMVerifierSecp256k1Transactor(address common.Address, transactor bind.ContractTransactor) (*TPMVerifierSecp256k1Transactor, error) {
	contract, err := bindTPMVerifierSecp256k1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TPMVerifierSecp256k1Transactor{contract: contract}, nil
}

// NewTPMVerifierSecp256k1Filterer creates a new log filterer instance of TPMVerifierSecp256k1, bound to a specific deployed contract.
func NewTPMVerifierSecp256k1Filterer(address common.Address, filterer bind.ContractFilterer) (*TPMVerifierSecp256k1Filterer, error) {
	contract, err := bindTPMVerifierSecp256k1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TPMVerifierSecp256k1Filterer{contract: contract}, nil
}

// bindTPMVerifierSecp256k1 binds a generic wrapper to an already deployed contract.
func bindTPMVerifierSecp256k1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TPMVerifierSecp256k1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TPMVerifierSecp256k1 *TPMVerifierSecp256k1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TPMVerifierSecp256k1.Contract.TPMVerifierSecp256k1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TPMVerifierSecp256k1 *TPMVerifierSecp256k1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TPMVerifierSecp256k1.Contract.TPMVerifierSecp256k1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TPMVerifierSecp256k1 *TPMVerifierSecp256k1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TPMVerifierSecp256k1.Contract.TPMVerifierSecp256k1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TPMVerifierSecp256k1 *TPMVerifierSecp256k1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TPMVerifierSecp256k1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TPMVerifierSecp256k1 *TPMVerifierSecp256k1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TPMVerifierSecp256k1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TPMVerifierSecp256k1 *TPMVerifierSecp256k1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TPMVerifierSecp256k1.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0xde12c640.
//
// Solidity: function verify(bytes32 keyId, bytes32 messageHash, bytes signature) view returns(bool ok)
func (_TPMVerifierSecp256k1 *TPMVerifierSecp256k1Caller) Verify(opts *bind.CallOpts, keyId [32]byte, messageHash [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _TPMVerifierSecp256k1.contract.Call(opts, &out, "verify", keyId, messageHash, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xde12c640.
//
// Solidity: function verify(bytes32 keyId, bytes32 messageHash, bytes signature) view returns(bool ok)
func (_TPMVerifierSecp256k1 *TPMVerifierSecp256k1Session) Verify(keyId [32]byte, messageHash [32]byte, signature []byte) (bool, error) {
	return _TPMVerifierSecp256k1.Contract.Verify(&_TPMVerifierSecp256k1.CallOpts, keyId, messageHash, signature)
}

// Verify is a free data retrieval call binding the contract method 0xde12c640.
//
// Solidity: function verify(bytes32 keyId, bytes32 messageHash, bytes signature) view returns(bool ok)
func (_TPMVerifierSecp256k1 *TPMVerifierSecp256k1CallerSession) Verify(keyId [32]byte, messageHash [32]byte, signature []byte) (bool, error) {
	return _TPMVerifierSecp256k1.Contract.Verify(&_TPMVerifierSecp256k1.CallOpts, keyId, messageHash, signature)
}
