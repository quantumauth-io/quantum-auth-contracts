// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package qaerc20

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

// QAERC20MetaData contains all meta data concerning the QAERC20 contract.
var QAERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051610cca380380610cca83398101604081905261002e91610181565b808484600361003d838261029f565b50600461004a828261029f565b5050506001600160a01b03811661007a57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61008381610093565b505060ff16608052506103599050565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610107575f5ffd5b81516001600160401b03811115610120576101206100e4565b604051601f8201601f19908116603f011681016001600160401b038111828210171561014e5761014e6100e4565b604052818152838201602001851015610165575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f60808587031215610194575f5ffd5b84516001600160401b038111156101a9575f5ffd5b6101b5878288016100f8565b602087015190955090506001600160401b038111156101d2575f5ffd5b6101de878288016100f8565b935050604085015160ff811681146101f4575f5ffd5b60608601519092506001600160a01b0381168114610210575f5ffd5b939692955090935050565b600181811c9082168061022f57607f821691505b60208210810361024d57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561029a57805f5260205f20601f840160051c810160208510156102785750805b601f840160051c820191505b81811015610297575f8155600101610284565b50505b505050565b81516001600160401b038111156102b8576102b86100e4565b6102cc816102c6845461021b565b84610253565b6020601f8211600181146102fe575f83156102e75750848201515b5f19600385901b1c1916600184901b178455610297565b5f84815260208120601f198516915b8281101561032d578785015182556020948501946001909201910161030d565b508482101561034a57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6080516109596103715f395f61016101526109595ff3fe608060405234801561000f575f5ffd5b50600436106100f0575f3560e01c806370a082311161009357806395d89b411161006357806395d89b4114610211578063a9059cbb14610219578063dd62ed3e1461022c578063f2fde38b14610264575f5ffd5b806370a08231146101b3578063715018a6146101db57806379cc6790146101e35780638da5cb5b146101f6575f5ffd5b806323b872dd116100ce57806323b872dd14610147578063313ce5671461015a57806340c10f191461018b57806342966c68146101a0575f5ffd5b806306fdde03146100f4578063095ea7b31461011257806318160ddd14610135575b5f5ffd5b6100fc610277565b60405161010991906107b2565b60405180910390f35b610125610120366004610802565b610307565b6040519015158152602001610109565b6002545b604051908152602001610109565b61012561015536600461082a565b610320565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610109565b61019e610199366004610802565b610343565b005b61019e6101ae366004610864565b610359565b6101396101c136600461087b565b6001600160a01b03165f9081526020819052604090205490565b61019e610366565b61019e6101f1366004610802565b610379565b6005546040516001600160a01b039091168152602001610109565b6100fc61038e565b610125610227366004610802565b61039d565b61013961023a36600461089b565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b61019e61027236600461087b565b6103aa565b606060038054610286906108cc565b80601f01602080910402602001604051908101604052809291908181526020018280546102b2906108cc565b80156102fd5780601f106102d4576101008083540402835291602001916102fd565b820191905f5260205f20905b8154815290600101906020018083116102e057829003601f168201915b5050505050905090565b5f336103148185856103e9565b60019150505b92915050565b5f3361032d8582856103fb565b610338858585610477565b506001949350505050565b61034b6104d4565b6103558282610501565b5050565b6103633382610535565b50565b61036e6104d4565b6103775f610569565b565b6103848233836103fb565b6103558282610535565b606060048054610286906108cc565b5f33610314818585610477565b6103b26104d4565b6001600160a01b0381166103e057604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61036381610569565b6103f683838360016105ba565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811015610471578181101561046357604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103d7565b61047184848484035f6105ba565b50505050565b6001600160a01b0383166104a057604051634b637e8f60e11b81525f60048201526024016103d7565b6001600160a01b0382166104c95760405163ec442f0560e01b81525f60048201526024016103d7565b6103f683838361068c565b6005546001600160a01b031633146103775760405163118cdaa760e01b81523360048201526024016103d7565b6001600160a01b03821661052a5760405163ec442f0560e01b81525f60048201526024016103d7565b6103555f838361068c565b6001600160a01b03821661055e57604051634b637e8f60e11b81525f60048201526024016103d7565b610355825f8361068c565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6001600160a01b0384166105e35760405163e602df0560e01b81525f60048201526024016103d7565b6001600160a01b03831661060c57604051634a1406b160e11b81525f60048201526024016103d7565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561047157826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161067e91815260200190565b60405180910390a350505050565b6001600160a01b0383166106b6578060025f8282546106ab9190610904565b909155506107269050565b6001600160a01b0383165f90815260208190526040902054818110156107085760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103d7565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661074257600280548290039055610760565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107a591815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b03811681146107fd575f5ffd5b919050565b5f5f60408385031215610813575f5ffd5b61081c836107e7565b946020939093013593505050565b5f5f5f6060848603121561083c575f5ffd5b610845846107e7565b9250610853602085016107e7565b929592945050506040919091013590565b5f60208284031215610874575f5ffd5b5035919050565b5f6020828403121561088b575f5ffd5b610894826107e7565b9392505050565b5f5f604083850312156108ac575f5ffd5b6108b5836107e7565b91506108c3602084016107e7565b90509250929050565b600181811c908216806108e057607f821691505b6020821081036108fe57634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561031a57634e487b7160e01b5f52601160045260245ffdfea264697066735822122080ed19c8dfe288f5fff44844379297d6f5dcaedb36fb8b9e3c00dee16786e1b364736f6c634300081c0033",
}

// QAERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use QAERC20MetaData.ABI instead.
var QAERC20ABI = QAERC20MetaData.ABI

// QAERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QAERC20MetaData.Bin instead.
var QAERC20Bin = QAERC20MetaData.Bin

// DeployQAERC20 deploys a new Ethereum contract, binding an instance of QAERC20 to it.
func DeployQAERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, owner_ common.Address) (common.Address, *types.Transaction, *QAERC20, error) {
	parsed, err := QAERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QAERC20Bin), backend, name_, symbol_, decimals_, owner_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QAERC20{QAERC20Caller: QAERC20Caller{contract: contract}, QAERC20Transactor: QAERC20Transactor{contract: contract}, QAERC20Filterer: QAERC20Filterer{contract: contract}}, nil
}

// QAERC20 is an auto generated Go binding around an Ethereum contract.
type QAERC20 struct {
	QAERC20Caller     // Read-only binding to the contract
	QAERC20Transactor // Write-only binding to the contract
	QAERC20Filterer   // Log filterer for contract events
}

// QAERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type QAERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QAERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type QAERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QAERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QAERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QAERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QAERC20Session struct {
	Contract     *QAERC20          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QAERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QAERC20CallerSession struct {
	Contract *QAERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// QAERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QAERC20TransactorSession struct {
	Contract     *QAERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// QAERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type QAERC20Raw struct {
	Contract *QAERC20 // Generic contract binding to access the raw methods on
}

// QAERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QAERC20CallerRaw struct {
	Contract *QAERC20Caller // Generic read-only contract binding to access the raw methods on
}

// QAERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QAERC20TransactorRaw struct {
	Contract *QAERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewQAERC20 creates a new instance of QAERC20, bound to a specific deployed contract.
func NewQAERC20(address common.Address, backend bind.ContractBackend) (*QAERC20, error) {
	contract, err := bindQAERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QAERC20{QAERC20Caller: QAERC20Caller{contract: contract}, QAERC20Transactor: QAERC20Transactor{contract: contract}, QAERC20Filterer: QAERC20Filterer{contract: contract}}, nil
}

// NewQAERC20Caller creates a new read-only instance of QAERC20, bound to a specific deployed contract.
func NewQAERC20Caller(address common.Address, caller bind.ContractCaller) (*QAERC20Caller, error) {
	contract, err := bindQAERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QAERC20Caller{contract: contract}, nil
}

// NewQAERC20Transactor creates a new write-only instance of QAERC20, bound to a specific deployed contract.
func NewQAERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*QAERC20Transactor, error) {
	contract, err := bindQAERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QAERC20Transactor{contract: contract}, nil
}

// NewQAERC20Filterer creates a new log filterer instance of QAERC20, bound to a specific deployed contract.
func NewQAERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*QAERC20Filterer, error) {
	contract, err := bindQAERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QAERC20Filterer{contract: contract}, nil
}

// bindQAERC20 binds a generic wrapper to an already deployed contract.
func bindQAERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QAERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QAERC20 *QAERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QAERC20.Contract.QAERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QAERC20 *QAERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QAERC20.Contract.QAERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QAERC20 *QAERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QAERC20.Contract.QAERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QAERC20 *QAERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QAERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QAERC20 *QAERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QAERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QAERC20 *QAERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QAERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_QAERC20 *QAERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QAERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_QAERC20 *QAERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _QAERC20.Contract.Allowance(&_QAERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_QAERC20 *QAERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _QAERC20.Contract.Allowance(&_QAERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_QAERC20 *QAERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QAERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_QAERC20 *QAERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _QAERC20.Contract.BalanceOf(&_QAERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_QAERC20 *QAERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _QAERC20.Contract.BalanceOf(&_QAERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_QAERC20 *QAERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _QAERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_QAERC20 *QAERC20Session) Decimals() (uint8, error) {
	return _QAERC20.Contract.Decimals(&_QAERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_QAERC20 *QAERC20CallerSession) Decimals() (uint8, error) {
	return _QAERC20.Contract.Decimals(&_QAERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QAERC20 *QAERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QAERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QAERC20 *QAERC20Session) Name() (string, error) {
	return _QAERC20.Contract.Name(&_QAERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QAERC20 *QAERC20CallerSession) Name() (string, error) {
	return _QAERC20.Contract.Name(&_QAERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QAERC20 *QAERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QAERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QAERC20 *QAERC20Session) Owner() (common.Address, error) {
	return _QAERC20.Contract.Owner(&_QAERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QAERC20 *QAERC20CallerSession) Owner() (common.Address, error) {
	return _QAERC20.Contract.Owner(&_QAERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_QAERC20 *QAERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QAERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_QAERC20 *QAERC20Session) Symbol() (string, error) {
	return _QAERC20.Contract.Symbol(&_QAERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_QAERC20 *QAERC20CallerSession) Symbol() (string, error) {
	return _QAERC20.Contract.Symbol(&_QAERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_QAERC20 *QAERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QAERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_QAERC20 *QAERC20Session) TotalSupply() (*big.Int, error) {
	return _QAERC20.Contract.TotalSupply(&_QAERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_QAERC20 *QAERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _QAERC20.Contract.TotalSupply(&_QAERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_QAERC20 *QAERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _QAERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_QAERC20 *QAERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.Approve(&_QAERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_QAERC20 *QAERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.Approve(&_QAERC20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_QAERC20 *QAERC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _QAERC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_QAERC20 *QAERC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.Burn(&_QAERC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_QAERC20 *QAERC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.Burn(&_QAERC20.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_QAERC20 *QAERC20Transactor) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QAERC20.contract.Transact(opts, "burnFrom", from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_QAERC20 *QAERC20Session) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.BurnFrom(&_QAERC20.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_QAERC20 *QAERC20TransactorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.BurnFrom(&_QAERC20.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_QAERC20 *QAERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QAERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_QAERC20 *QAERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.Mint(&_QAERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_QAERC20 *QAERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.Mint(&_QAERC20.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QAERC20 *QAERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QAERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QAERC20 *QAERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _QAERC20.Contract.RenounceOwnership(&_QAERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QAERC20 *QAERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QAERC20.Contract.RenounceOwnership(&_QAERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_QAERC20 *QAERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QAERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_QAERC20 *QAERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.Transfer(&_QAERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_QAERC20 *QAERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.Transfer(&_QAERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_QAERC20 *QAERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QAERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_QAERC20 *QAERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.TransferFrom(&_QAERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_QAERC20 *QAERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QAERC20.Contract.TransferFrom(&_QAERC20.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QAERC20 *QAERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QAERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QAERC20 *QAERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QAERC20.Contract.TransferOwnership(&_QAERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QAERC20 *QAERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QAERC20.Contract.TransferOwnership(&_QAERC20.TransactOpts, newOwner)
}

// QAERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the QAERC20 contract.
type QAERC20ApprovalIterator struct {
	Event *QAERC20Approval // Event containing the contract specifics and raw log

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
func (it *QAERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAERC20Approval)
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
		it.Event = new(QAERC20Approval)
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
func (it *QAERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAERC20Approval represents a Approval event raised by the QAERC20 contract.
type QAERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_QAERC20 *QAERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*QAERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _QAERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &QAERC20ApprovalIterator{contract: _QAERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_QAERC20 *QAERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *QAERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _QAERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAERC20Approval)
				if err := _QAERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_QAERC20 *QAERC20Filterer) ParseApproval(log types.Log) (*QAERC20Approval, error) {
	event := new(QAERC20Approval)
	if err := _QAERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QAERC20 contract.
type QAERC20OwnershipTransferredIterator struct {
	Event *QAERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QAERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAERC20OwnershipTransferred)
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
		it.Event = new(QAERC20OwnershipTransferred)
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
func (it *QAERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAERC20OwnershipTransferred represents a OwnershipTransferred event raised by the QAERC20 contract.
type QAERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QAERC20 *QAERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QAERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QAERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QAERC20OwnershipTransferredIterator{contract: _QAERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QAERC20 *QAERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QAERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QAERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAERC20OwnershipTransferred)
				if err := _QAERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QAERC20 *QAERC20Filterer) ParseOwnershipTransferred(log types.Log) (*QAERC20OwnershipTransferred, error) {
	event := new(QAERC20OwnershipTransferred)
	if err := _QAERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the QAERC20 contract.
type QAERC20TransferIterator struct {
	Event *QAERC20Transfer // Event containing the contract specifics and raw log

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
func (it *QAERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAERC20Transfer)
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
		it.Event = new(QAERC20Transfer)
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
func (it *QAERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAERC20Transfer represents a Transfer event raised by the QAERC20 contract.
type QAERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_QAERC20 *QAERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*QAERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QAERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &QAERC20TransferIterator{contract: _QAERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_QAERC20 *QAERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *QAERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QAERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAERC20Transfer)
				if err := _QAERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_QAERC20 *QAERC20Filterer) ParseTransfer(log types.Log) (*QAERC20Transfer, error) {
	event := new(QAERC20Transfer)
	if err := _QAERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
