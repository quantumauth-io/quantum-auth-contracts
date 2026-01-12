// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quantumauthaccount

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

// BaseAccountCall is an auto generated low-level Go binding around an user-defined struct.
type BaseAccountCall struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// QuantumAuthAccountMetaData contains all meta data concerning the QuantumAuthAccount contract.
var QuantumAuthAccountMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"entryPoint_\",\"type\":\"address\",\"internalType\":\"contractIEntryPoint\"},{\"name\":\"eoa1_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"eoa2_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tpmVerifier_\",\"type\":\"address\",\"internalType\":\"contractITPMVerifier\"},{\"name\":\"tpmKeyId_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"EOA1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EOA2\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TPM_KEY_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TPM_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITPMVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"entryPoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEntryPoint\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeBatch\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structBaseAccount.Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"missingAccountFunds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ExecuteError\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEOA\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTPMVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotFromEntryPoint\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entity\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entryPoint\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyEntryPoint\",\"inputs\":[]}]",
	Bin: "0x6101203461016c57601f6110da38819003918201601f19168301916001600160401b038311848410176101705780849260a09460405283398101031261016c5780516001600160a01b038116810361016c5761005d60208301610184565b61006960408401610184565b9060608401519260018060a01b0384169485850361016c5760800151946001600160a01b038316801590811561015a575b8115610147575b5061013857156101295760805260a05260c05260e05261010052604051610f419081610199823960805181818161031f015281816104d6015281816108d901526109a7015260a0518181816104350152610b28015260c0518181816102b10152610b51015260e0518181816102430152610c5501526101005181818161057e0152610c140152f35b633c4ac85360e21b5f5260045ffd5b6303602a5960e51b5f5260045ffd5b6001600160a01b0386161490505f6100a1565b6001600160a01b03861615915061009a565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361016c5756fe6080604052600436101561001a575b3615610018575f80fd5b005b5f3560e01c806319822f7c146100b957806334fcd5be146100b45780637e8ee814146100af578063a6b18bc0146100aa578063b0d691fe146100a5578063b61d27f6146100a0578063bfbbda751461009b578063d087d28814610096578063d9260ecf146100915763f3fef3a30361000e576105a1565b610549565b610459565b6103eb565b610361565b6102d5565b610267565b6101f9565b61017e565b3461017a5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a5760043567ffffffffffffffff811161017a576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261017a5761015e906101466044359161013a610990565b60243590600401610acb565b9080610162575b506040519081529081906020820190565b0390f35b5f80808093335af150610173610956565b505f61014d565b5f80fd5b3461017a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a5760043567ffffffffffffffff811161017a573660238201121561017a57806004013567ffffffffffffffff811161017a573660248260051b8401011161017a576024610018920161080b565b3461017a575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461017a575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461017a575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b73ffffffffffffffffffffffffffffffffffffffff81160361017a57565b3461017a5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a5760043561039c81610343565b6024356044359167ffffffffffffffff831161017a573660238401121561017a5782600401359167ffffffffffffffff831161017a57366024848601011161017a5760246100189401916108c0565b3461017a575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461017a575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a576040517f35567e1a0000000000000000000000000000000000000000000000000000000081523060048201525f602482015260208160448173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015610544575f90610510575b604051908152602090f35b506020813d60201161053c575b8161052a602093836106fb565b8101031261017a5761015e9051610505565b3d915061051d565b610985565b3461017a575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b3461017a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017a576105db600435610343565b7fea8e4eb5000000000000000000000000000000000000000000000000000000005f5260045ffd5b91908110156106435760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18136030182121561017a570190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b3561067a81610343565b90565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561017a570180359067ffffffffffffffff821161017a5760200191813603831361017a57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761073c57604052565b6106ce565b67ffffffffffffffff811161073c57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b92919261078782610741565b9161079560405193846106fb565b82948184528183011161017a578281602093845f960137010152565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b60409061067a9392815281602082015201906107b1565b90610814610990565b5f5b81811061082257505050565b61086861086461085e610836848688610603565b61083f81610670565b90610853602082013591604081019061067d565b93905a94369161077b565b91610ccc565b1590565b61087457600101610816565b600182145f03610cf857610886610cde565b906108bc6040519283927f5a154675000000000000000000000000000000000000000000000000000000008452600484016107f4565b0390fd5b909273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361092e575f938493826040519384928337810185815203925af161091d610956565b90156109265750565b602081519101fd5b7fbd07c551000000000000000000000000000000000000000000000000000000005f5260045ffd5b3d15610980573d9061096782610741565b9161097560405193846106fb565b82523d5f602084013e565b606090565b6040513d5f823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168033036109d15750565b7ffe34a6d3000000000000000000000000000000000000000000000000000000005f52336004523060245260445260645ffd5b9080601f8301121561017a5781602061067a9335910161077b565b9060808282031261017a57813560ff8116810361017a5792602083013567ffffffffffffffff811161017a5782610a57918501610a04565b92604081013567ffffffffffffffff811161017a5783610a78918301610a04565b92606082013567ffffffffffffffff811161017a5761067a9201610a04565b9081602091031261017a5751801515810361017a5790565b61067a93926060928252602082015281604082015201906107b1565b610b7691610aeb610ae360ff9361010081019061067d565b810190610a1f565b929491610b4e610b25869893987f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f52601c52603c5f2090565b917f00000000000000000000000000000000000000000000000000000000000000009083610d00565b967f000000000000000000000000000000000000000000000000000000000000000091610d00565b9316918215610bcd575050600114610bb0577fa0042b17000000000000000000000000000000000000000000000000000000005f5260045ffd5b81610bc5575b5015610bc0575f90565b600190565b90505f610bb6565b9290938092509115610cc4575b5015610cbd57610c3c9160209160405193849283927fde12c6400000000000000000000000000000000000000000000000000000000084527f000000000000000000000000000000000000000000000000000000000000000060048501610aaf565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610544575f91610c8e575015610bc0575f90565b610cb0915060203d602011610cb6575b610ca881836106fb565b810190610a97565b5f610bb6565b503d610c9e565b5050600190565b90505f610bda565b925f939184939260208451940192f190565b3d604051906020818301016040528082525f602083013e90565b610926610cde565b90805115610d3d57610d3692610d2d73ffffffffffffffffffffffffffffffffffffffff93928493610d44565b90959195610db5565b1691161490565b5050505f90565b8151919060418303610d7457610d6d9250602082015190606060408401519301515f1a90610e7c565b9192909190565b50505f9160029190565b60041115610d8857565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b610dbe81610d7e565b80610dc7575050565b610dd081610d7e565b60018103610e00577ff645eedf000000000000000000000000000000000000000000000000000000005f5260045ffd5b610e0981610d7e565b60028103610e3d57507ffce698f7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b80610e49600392610d7e565b14610e515750565b7fd78bce0c000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411610f00579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15610544575f5173ffffffffffffffffffffffffffffffffffffffff811615610ef657905f905f90565b505f906001905f90565b5050505f916003919056fea2646970667358221220c3c4a75527f66bf53b92908637068e561474c0dc46f192bababc4924c4b0f46264736f6c634300081c0033",
}

// QuantumAuthAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use QuantumAuthAccountMetaData.ABI instead.
var QuantumAuthAccountABI = QuantumAuthAccountMetaData.ABI

// QuantumAuthAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QuantumAuthAccountMetaData.Bin instead.
var QuantumAuthAccountBin = QuantumAuthAccountMetaData.Bin

// DeployQuantumAuthAccount deploys a new Ethereum contract, binding an instance of QuantumAuthAccount to it.
func DeployQuantumAuthAccount(auth *bind.TransactOpts, backend bind.ContractBackend, entryPoint_ common.Address, eoa1_ common.Address, eoa2_ common.Address, tpmVerifier_ common.Address, tpmKeyId_ [32]byte) (common.Address, *types.Transaction, *QuantumAuthAccount, error) {
	parsed, err := QuantumAuthAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuantumAuthAccountBin), backend, entryPoint_, eoa1_, eoa2_, tpmVerifier_, tpmKeyId_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QuantumAuthAccount{QuantumAuthAccountCaller: QuantumAuthAccountCaller{contract: contract}, QuantumAuthAccountTransactor: QuantumAuthAccountTransactor{contract: contract}, QuantumAuthAccountFilterer: QuantumAuthAccountFilterer{contract: contract}}, nil
}

// QuantumAuthAccount is an auto generated Go binding around an Ethereum contract.
type QuantumAuthAccount struct {
	QuantumAuthAccountCaller     // Read-only binding to the contract
	QuantumAuthAccountTransactor // Write-only binding to the contract
	QuantumAuthAccountFilterer   // Log filterer for contract events
}

// QuantumAuthAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuantumAuthAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuantumAuthAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuantumAuthAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuantumAuthAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuantumAuthAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuantumAuthAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuantumAuthAccountSession struct {
	Contract     *QuantumAuthAccount // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QuantumAuthAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuantumAuthAccountCallerSession struct {
	Contract *QuantumAuthAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// QuantumAuthAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuantumAuthAccountTransactorSession struct {
	Contract     *QuantumAuthAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// QuantumAuthAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuantumAuthAccountRaw struct {
	Contract *QuantumAuthAccount // Generic contract binding to access the raw methods on
}

// QuantumAuthAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuantumAuthAccountCallerRaw struct {
	Contract *QuantumAuthAccountCaller // Generic read-only contract binding to access the raw methods on
}

// QuantumAuthAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuantumAuthAccountTransactorRaw struct {
	Contract *QuantumAuthAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuantumAuthAccount creates a new instance of QuantumAuthAccount, bound to a specific deployed contract.
func NewQuantumAuthAccount(address common.Address, backend bind.ContractBackend) (*QuantumAuthAccount, error) {
	contract, err := bindQuantumAuthAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuantumAuthAccount{QuantumAuthAccountCaller: QuantumAuthAccountCaller{contract: contract}, QuantumAuthAccountTransactor: QuantumAuthAccountTransactor{contract: contract}, QuantumAuthAccountFilterer: QuantumAuthAccountFilterer{contract: contract}}, nil
}

// NewQuantumAuthAccountCaller creates a new read-only instance of QuantumAuthAccount, bound to a specific deployed contract.
func NewQuantumAuthAccountCaller(address common.Address, caller bind.ContractCaller) (*QuantumAuthAccountCaller, error) {
	contract, err := bindQuantumAuthAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuantumAuthAccountCaller{contract: contract}, nil
}

// NewQuantumAuthAccountTransactor creates a new write-only instance of QuantumAuthAccount, bound to a specific deployed contract.
func NewQuantumAuthAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*QuantumAuthAccountTransactor, error) {
	contract, err := bindQuantumAuthAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuantumAuthAccountTransactor{contract: contract}, nil
}

// NewQuantumAuthAccountFilterer creates a new log filterer instance of QuantumAuthAccount, bound to a specific deployed contract.
func NewQuantumAuthAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*QuantumAuthAccountFilterer, error) {
	contract, err := bindQuantumAuthAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuantumAuthAccountFilterer{contract: contract}, nil
}

// bindQuantumAuthAccount binds a generic wrapper to an already deployed contract.
func bindQuantumAuthAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuantumAuthAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuantumAuthAccount *QuantumAuthAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuantumAuthAccount.Contract.QuantumAuthAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuantumAuthAccount *QuantumAuthAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.QuantumAuthAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuantumAuthAccount *QuantumAuthAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.QuantumAuthAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuantumAuthAccount *QuantumAuthAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuantumAuthAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuantumAuthAccount *QuantumAuthAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuantumAuthAccount *QuantumAuthAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.contract.Transact(opts, method, params...)
}

// EOA1 is a free data retrieval call binding the contract method 0xbfbbda75.
//
// Solidity: function EOA1() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountCaller) EOA1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuantumAuthAccount.contract.Call(opts, &out, "EOA1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EOA1 is a free data retrieval call binding the contract method 0xbfbbda75.
//
// Solidity: function EOA1() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountSession) EOA1() (common.Address, error) {
	return _QuantumAuthAccount.Contract.EOA1(&_QuantumAuthAccount.CallOpts)
}

// EOA1 is a free data retrieval call binding the contract method 0xbfbbda75.
//
// Solidity: function EOA1() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountCallerSession) EOA1() (common.Address, error) {
	return _QuantumAuthAccount.Contract.EOA1(&_QuantumAuthAccount.CallOpts)
}

// EOA2 is a free data retrieval call binding the contract method 0xa6b18bc0.
//
// Solidity: function EOA2() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountCaller) EOA2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuantumAuthAccount.contract.Call(opts, &out, "EOA2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EOA2 is a free data retrieval call binding the contract method 0xa6b18bc0.
//
// Solidity: function EOA2() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountSession) EOA2() (common.Address, error) {
	return _QuantumAuthAccount.Contract.EOA2(&_QuantumAuthAccount.CallOpts)
}

// EOA2 is a free data retrieval call binding the contract method 0xa6b18bc0.
//
// Solidity: function EOA2() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountCallerSession) EOA2() (common.Address, error) {
	return _QuantumAuthAccount.Contract.EOA2(&_QuantumAuthAccount.CallOpts)
}

// TPMKEYID is a free data retrieval call binding the contract method 0xd9260ecf.
//
// Solidity: function TPM_KEY_ID() view returns(bytes32)
func (_QuantumAuthAccount *QuantumAuthAccountCaller) TPMKEYID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QuantumAuthAccount.contract.Call(opts, &out, "TPM_KEY_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TPMKEYID is a free data retrieval call binding the contract method 0xd9260ecf.
//
// Solidity: function TPM_KEY_ID() view returns(bytes32)
func (_QuantumAuthAccount *QuantumAuthAccountSession) TPMKEYID() ([32]byte, error) {
	return _QuantumAuthAccount.Contract.TPMKEYID(&_QuantumAuthAccount.CallOpts)
}

// TPMKEYID is a free data retrieval call binding the contract method 0xd9260ecf.
//
// Solidity: function TPM_KEY_ID() view returns(bytes32)
func (_QuantumAuthAccount *QuantumAuthAccountCallerSession) TPMKEYID() ([32]byte, error) {
	return _QuantumAuthAccount.Contract.TPMKEYID(&_QuantumAuthAccount.CallOpts)
}

// TPMVERIFIER is a free data retrieval call binding the contract method 0x7e8ee814.
//
// Solidity: function TPM_VERIFIER() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountCaller) TPMVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuantumAuthAccount.contract.Call(opts, &out, "TPM_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TPMVERIFIER is a free data retrieval call binding the contract method 0x7e8ee814.
//
// Solidity: function TPM_VERIFIER() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountSession) TPMVERIFIER() (common.Address, error) {
	return _QuantumAuthAccount.Contract.TPMVERIFIER(&_QuantumAuthAccount.CallOpts)
}

// TPMVERIFIER is a free data retrieval call binding the contract method 0x7e8ee814.
//
// Solidity: function TPM_VERIFIER() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountCallerSession) TPMVERIFIER() (common.Address, error) {
	return _QuantumAuthAccount.Contract.TPMVERIFIER(&_QuantumAuthAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuantumAuthAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountSession) EntryPoint() (common.Address, error) {
	return _QuantumAuthAccount.Contract.EntryPoint(&_QuantumAuthAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_QuantumAuthAccount *QuantumAuthAccountCallerSession) EntryPoint() (common.Address, error) {
	return _QuantumAuthAccount.Contract.EntryPoint(&_QuantumAuthAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_QuantumAuthAccount *QuantumAuthAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuantumAuthAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_QuantumAuthAccount *QuantumAuthAccountSession) GetNonce() (*big.Int, error) {
	return _QuantumAuthAccount.Contract.GetNonce(&_QuantumAuthAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_QuantumAuthAccount *QuantumAuthAccountCallerSession) GetNonce() (*big.Int, error) {
	return _QuantumAuthAccount.Contract.GetNonce(&_QuantumAuthAccount.CallOpts)
}

// Withdraw is a free data retrieval call binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address , uint256 ) pure returns()
func (_QuantumAuthAccount *QuantumAuthAccountCaller) Withdraw(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) error {
	var out []interface{}
	err := _QuantumAuthAccount.contract.Call(opts, &out, "withdraw", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// Withdraw is a free data retrieval call binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address , uint256 ) pure returns()
func (_QuantumAuthAccount *QuantumAuthAccountSession) Withdraw(arg0 common.Address, arg1 *big.Int) error {
	return _QuantumAuthAccount.Contract.Withdraw(&_QuantumAuthAccount.CallOpts, arg0, arg1)
}

// Withdraw is a free data retrieval call binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address , uint256 ) pure returns()
func (_QuantumAuthAccount *QuantumAuthAccountCallerSession) Withdraw(arg0 common.Address, arg1 *big.Int) error {
	return _QuantumAuthAccount.Contract.Withdraw(&_QuantumAuthAccount.CallOpts, arg0, arg1)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_QuantumAuthAccount *QuantumAuthAccountTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _QuantumAuthAccount.contract.Transact(opts, "execute", target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_QuantumAuthAccount *QuantumAuthAccountSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.Execute(&_QuantumAuthAccount.TransactOpts, target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_QuantumAuthAccount *QuantumAuthAccountTransactorSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.Execute(&_QuantumAuthAccount.TransactOpts, target, value, data)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_QuantumAuthAccount *QuantumAuthAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, calls []BaseAccountCall) (*types.Transaction, error) {
	return _QuantumAuthAccount.contract.Transact(opts, "executeBatch", calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_QuantumAuthAccount *QuantumAuthAccountSession) ExecuteBatch(calls []BaseAccountCall) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.ExecuteBatch(&_QuantumAuthAccount.TransactOpts, calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_QuantumAuthAccount *QuantumAuthAccountTransactorSession) ExecuteBatch(calls []BaseAccountCall) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.ExecuteBatch(&_QuantumAuthAccount.TransactOpts, calls)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_QuantumAuthAccount *QuantumAuthAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _QuantumAuthAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_QuantumAuthAccount *QuantumAuthAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.ValidateUserOp(&_QuantumAuthAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_QuantumAuthAccount *QuantumAuthAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.ValidateUserOp(&_QuantumAuthAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_QuantumAuthAccount *QuantumAuthAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuantumAuthAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_QuantumAuthAccount *QuantumAuthAccountSession) Receive() (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.Receive(&_QuantumAuthAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_QuantumAuthAccount *QuantumAuthAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _QuantumAuthAccount.Contract.Receive(&_QuantumAuthAccount.TransactOpts)
}
