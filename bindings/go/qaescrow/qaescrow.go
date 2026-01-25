// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package qaescrow

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

// QAEscrowEscrowTerms is an auto generated low-level Go binding around an user-defined struct.
type QAEscrowEscrowTerms struct {
	AppId        [32]byte
	OrderId      [32]byte
	Buyer        common.Address
	Amount       *big.Int
	DelaySeconds uint64
	ExpiresAt    uint64
}

// QAEscrowMetaData contains all meta data concerning the QAEscrow contract.
var QAEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"usdc\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"APP_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ARBITRATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ESCROW_TERMS_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SIGNER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"appOwner\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"availableCapacity\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositUsdc\",\"inputs\":[{\"name\":\"t\",\"type\":\"tuple\",\"internalType\":\"structQAEscrow.EscrowTerms\",\"components\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delaySeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expiresAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emergencyWithdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"locked\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"openDispute\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"orderKey\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"orders\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"merchant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"releaseAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumQAEscrow.OrderStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"buyerWins\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"buyerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAppOwner\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTermsSigner\",\"inputs\":[{\"name\":\"newSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"staked\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"termsSigner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferAppOwner\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AppOwnerUpdated\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"orderKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"merchant\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"orderKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"merchant\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"releaseAt\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeOpened\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"orderKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeResolved\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"orderKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"buyerWins\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"buyerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"merchantAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"merchant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stakedTotal\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TermsSignerUpdated\",\"inputs\":[{\"name\":\"oldSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unstaked\",\"inputs\":[{\"name\":\"appId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"merchant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stakedTotal\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x61018080604052346102e757604081613b03803803809161002082856102eb565b8339810103126102e75761003f602061003883610322565b9201610322565b9060405161004e6040826102eb565b600881526020810190675141457363726f7760c01b8252604051916100746040846102eb565b600183526020830191603160f81b835233156102d4575f8054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a36100d781610565565b610120526100e484610700565b61014052519020918260e05251902080610100524660a0526040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261014d60c0826102eb565b5190206080523060c05260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00556001600160a01b031680156102a6576001600160a01b03821615610276576101605260048054610100600160a81b03191660089290921b610100600160a81b03169190911790556101cb33610336565b506101d5336103ac565b506101df3361043f565b506101e9336104d2565b5060405161322a9081610839823960805181612f98015260a0518161304f015260c05181612f69015260e05181612fe70152610100518161300d015261012051816118bf015261014051816118e801526101605181818161042c01528181610478015281816108d0015281816111140152818161133801528181611481015281816115bf01526120e20152f35b60405162461bcd60e51b815260206004820152600860248201526707369676e65723d360c41b6044820152606490fd5b60405162461bcd60e51b81526020600482015260066024820152650555344433d360d41b6044820152606490fd5b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761030e57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036102e757565b6001600160a01b0381165f9081525f516020613a835f395f51905f52602052604090205460ff166103a7576001600160a01b03165f8181525f516020613a835f395f51905f5260205260408120805460ff191660011790553391905f516020613a635f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f516020613ac35f395f51905f52602052604090205460ff166103a7576001600160a01b03165f8181525f516020613ac35f395f51905f5260205260408120805460ff191660011790553391907f371a0078bf8859908953848339bea5f1d5775487f6c2f50fd279fcc2cafd8c60905f516020613a635f395f51905f529080a4600190565b6001600160a01b0381165f9081525f516020613ae35f395f51905f52602052604090205460ff166103a7576001600160a01b03165f8181525f516020613ae35f395f51905f5260205260408120805460ff191660011790553391907f16ceee8289685dd2a02b9c8ae81d2df373176ce53519e6284e2a2950d6546ffa905f516020613a635f395f51905f529080a4600190565b6001600160a01b0381165f9081525f516020613aa35f395f51905f52602052604090205460ff166103a7576001600160a01b03165f8181525f516020613aa35f395f51905f5260205260408120805460ff191660011790553391907f6f82f19233637990cc5fee2a338f11530b99de438ca6f1e3226d1410814c055e905f516020613a635f395f51905f529080a4600190565b908151602081105f146105df575090601f81511161059f576020815191015160208210610590571790565b5f198260200360031b1b161790565b604460209160405192839163305a27a960e01b83528160048401528051918291826024860152018484015e5f828201840152601f01601f19168101030190fd5b6001600160401b03811161030e57600254600181811c911680156106f6575b60208210146106e257601f81116106af575b50602092601f821160011461064e57928192935f92610643575b50508160011b915f199060031b1c19161760025560ff90565b015190505f8061062a565b601f1982169360025f52805f20915f5b868110610697575083600195961061067f575b505050811b0160025560ff90565b01515f1960f88460031b161c191690555f8080610671565b9192602060018192868501518155019401920161065e565b60025f52601f60205f20910160051c810190601f830160051c015b8181106106d75750610610565b5f81556001016106ca565b634e487b7160e01b5f52602260045260245ffd5b90607f16906105fe565b908151602081105f1461072b575090601f81511161059f576020815191015160208210610590571790565b6001600160401b03811161030e57600354600181811c9116801561082e575b60208210146106e257601f81116107fb575b50602092601f821160011461079a57928192935f9261078f575b50508160011b915f199060031b1c19161760035560ff90565b015190505f80610776565b601f1982169360035f52805f20915f5b8681106107e357508360019596106107cb575b505050811b0160035560ff90565b01515f1960f88460031b161c191690555f80806107bd565b919260206001819286850151815501940192016107aa565b60035f52601f60205f20910160051c810190601f830160051c015b818110610823575061075c565b5f8155600101610816565b90607f169061074a56fe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a714612468575080630581c36b1461242c5780630e117e731461226e578063120c857c1461222657806316ae261d14612030578063248a9ca314611fe75780632c4bc09314611dc75780632f2ff15d14611d6b57806336568abe14611ce35780633823eefc14611cc25780633f4ba83a14611c075780635c39e04314611bb35780635c975abb14611b7357806364f025b114611aea578063715018a614611a505780638456cb59146119bb57806384b0196e1461188957806384b41e331461173257806384cc9dfb146114a557806389a30271146114375780638bf04c45146113d95780638caa5230146112d95780638da5cb5b1461128957806391d148541461121457806392fb9db2146111bc57806395ccea67146110a757806395ce9b441461104f5780639c3f1e9014610f8e5780639cea078714610f36578063a217fddf14610efe578063a95bdf7f14610ea6578063cbe9e76414610e5e578063d547741f14610dfb578063e2e8df6f14610600578063ea35e258146102985763f2fde38b146101a8575f80fd5b346102945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945773ffffffffffffffffffffffffffffffffffffffff6101f4612524565b6101fc612c35565b1680156102685773ffffffffffffffffffffffffffffffffffffffff5f54827fffffffffffffffffffffffff00000000000000000000000000000000000000008216175f55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b7f1e4fbdf7000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b5f80fd5b346102945760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457600435602435604435918215159283810361029457606435906102e8612928565b6102f061299f565b335f9081527f4ea974b431a0b4c2e2044c8e7ab9ed5c93a2a8529b05175b6aea7d6522678c2b602052604090205460ff16156105b0576103308484612770565b94855f52600560205260405f20600481019360ff855460401c16600581101561058357600203610525576003820154918282116104c7577f8e9c1dd6f57a92254350fa9e6a6da6e64dc9ff0838bc0d7b82970f37e691d0d69560609561039684866125e1565b94895f5260076020526103ae60405f209182546125e1565b9055156104a9576103c690846104a1576004906126ec565b81610457575b828061040c575b505060405192835260208301526040820152a460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b73ffffffffffffffffffffffffffffffffffffffff6002610450930154167f00000000000000000000000000000000000000000000000000000000000000006129d3565b88826103d3565b61049c8273ffffffffffffffffffffffffffffffffffffffff6001840154167f00000000000000000000000000000000000000000000000000000000000000006129d3565b6103cc565b6003906126ec565b6104ba90836104bf576003906126ec565b6103c6565b6004906126ec565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f6275796572416d6f756e7420746f6f20626967000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f6e6f7420646973707574656400000000000000000000000000000000000000006044820152fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004527f16ceee8289685dd2a02b9c8ae81d2df373176ce53519e6284e2a2950d6546ffa60245260445ffd5b34610294577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360160e081126102945760c0136102945760c43567ffffffffffffffff811161029457366023820112156102945780600401359067ffffffffffffffff82116102945736602483830101116102945761067d612928565b61068561299f565b73ffffffffffffffffffffffffffffffffffffffff6106a2612808565b163303610d9d576064356106b7811515612622565b67ffffffffffffffff6106c861282b565b1615610d3f5767ffffffffffffffff6106df612842565b164211610ce15760043590815f52600860205273ffffffffffffffffffffffffffffffffffffffff60405f205416908115610c8357825f52600760205261072a8160405f20546127fb565b835f52600660205260405f205410610c25576024359361074a8585612770565b95865f52600560205260ff600460405f20015460401c16600581101561058357610bc7576108a49161089b915f60206042898b67ffffffffffffffff61078e612808565b8161079761282b565b73ffffffffffffffffffffffffffffffffffffffff6107b4612842565b93604051968a8801987ffe0d8c3be206e144ac41d1a18ce931cd63a5a8dd8b96834ad3b0ebc4cf9cc17f8a52604089015260608801521660808601528d60a08601521660c08401521660e082015260e081526108126101008261272f565b51902061081d612f52565b90604051917f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201522092806024837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116019661088a604051988961272f565b828852018387013784010152613075565b909291926130af565b73ffffffffffffffffffffffffffffffffffffffff8060045460081c16911603610b69576108f48130337f0000000000000000000000000000000000000000000000000000000000000000612c81565b67ffffffffffffffff61090561282b565b1667ffffffffffffffff4216019167ffffffffffffffff8311610b3c5761092a612808565b60405160c081019080821067ffffffffffffffff831117610b0f57600467ffffffffffffffff918a9360405288815273ffffffffffffffffffffffffffffffffffffffff60208201951685526040810186815273ffffffffffffffffffffffffffffffffffffffff60608301918983528186608086019c1698898d5260a086019860018a525f52600560205260405f20955186555116826001860191167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055511673ffffffffffffffffffffffffffffffffffffffff6002840191167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055516003820155019551167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008654161785555192600584101561058357610a9b6080947fe76a1d7f1e6df89045e72eaa27e22664ea43d53955a6cfe5e47ddb600b0b2239966126ec565b855f52600760205260405f20610ab28282546127fb565b9055610abc612808565b9273ffffffffffffffffffffffffffffffffffffffff60405194168452602084015260408301526060820152a460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f626164207369676e6174757265000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f6f726465722065786973747300000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6f766572206361706163697479000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f756e6b6e6f776e206170700000000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f7465726d732065787069726564000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f64656c61793d30000000000000000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6275796572206d69736d617463680000000000000000000000000000000000006044820152fd5b346102945760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457610e5c600435610e38612547565b90610e57610e52825f526001602052600160405f20015490565b6128c1565b612b6b565b005b346102945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610294576004355f526007602052602060405f2054604051908152f35b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945760206040517f6f82f19233637990cc5fee2a338f11530b99de438ca6f1e3226d1410814c055e8152f35b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945760206040515f8152f35b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945760206040517f16ceee8289685dd2a02b9c8ae81d2df373176ce53519e6284e2a2950d6546ffa8152f35b346102945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610294576004355f52600560205260405f20805467ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff6001840154169273ffffffffffffffffffffffffffffffffffffffff60028201541690600460038201549101549160ff8360401c169560405195865260208601526040850152606084015216608082015260058210156105835760c09160a0820152f35b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945760206040517ffe0d8c3be206e144ac41d1a18ce931cd63a5a8dd8b96834ad3b0ebc4cf9cc17f8152f35b346102945760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610294576110de612524565b6110e6612859565b6110ee612928565b73ffffffffffffffffffffffffffffffffffffffff81161561115e5761113890602435907f00000000000000000000000000000000000000000000000000000000000000006129d3565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600460248201527f746f3d30000000000000000000000000000000000000000000000000000000006044820152fd5b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945760206040517f371a0078bf8859908953848339bea5f1d5775487f6c2f50fd279fcc2cafd8c608152f35b346102945760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945761124b612547565b6004355f52600160205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457602073ffffffffffffffffffffffffffffffffffffffff5f5416604051908152f35b34610294576112e73661256a565b906112f0612928565b6112f861299f565b611303821515612622565b805f52600860205261133073ffffffffffffffffffffffffffffffffffffffff60405f2054163314612687565b61135c8230337f0000000000000000000000000000000000000000000000000000000000000000612c81565b805f52600660205260405f206113738382546127fb565b9055805f52600660205260405f205460405192835260208301527fa1fdccfe567643a44425efdd141171e8d992854a81e5c819c1432b0de47c9a1160403393a360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b346102945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610294576004355f526008602052602073ffffffffffffffffffffffffffffffffffffffff60405f205416604051908152f35b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610294576114b33661256a565b6114bb612928565b6114c361299f565b6114cd8183612770565b91825f52600560205260405f2060048101805460ff8160401c166005811015610583576001036116d457600283019073ffffffffffffffffffffffffffffffffffffffff82541633036116765767ffffffffffffffff1642106116185773ffffffffffffffffffffffffffffffffffffffff60037f08d0ecb8e19e562308f200cd130d9dba4e5e86592a6470ee028732e44edbe22994604094680300000000000000007fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff82541617905501918254865f5260076020526115b1855f209182546125e1565b90556115e3828254168454907f00000000000000000000000000000000000000000000000000000000000000006129d3565b5416905482519182526020820152a460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f746f6f206561726c7900000000000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f6e6f74206d65726368616e7400000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6e6f7420636c61696d61626c65000000000000000000000000000000000000006044820152fd5b346102945760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945760043561176c612547565b335f9081527f44b31f1bc25958e4acb8433a7d0a722ee5ff13180d0cc5d5ced58fccd53eb77c602052604090205460ff16156118395773ffffffffffffffffffffffffffffffffffffffff16906117c4821515612796565b5f81815260086020526040812080547fffffffffffffffffffffffff00000000000000000000000000000000000000008116851790915573ffffffffffffffffffffffffffffffffffffffff1691907f9a5c7e2e887f2ba1250b313ad37205003e92ec9bcd2bc57b593646cd6c31bfdc9080a4005b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004527f371a0078bf8859908953848339bea5f1d5775487f6c2f50fd279fcc2cafd8c6060245260445ffd5b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945761195f6118e37f0000000000000000000000000000000000000000000000000000000000000000612d0c565b61190c7f0000000000000000000000000000000000000000000000000000000000000000612e82565b602061196d6040519261191f838561272f565b5f84525f3681376040519586957f0f00000000000000000000000000000000000000000000000000000000000000875260e08588015260e087019061259e565b90858203604087015261259e565b4660608501523060808501525f60a085015283810360c08501528180845192838152019301915f5b8281106119a457505050500390f35b835185528695509381019392810192600101611995565b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610294576119f1612859565b6119f961299f565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060045416176004557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457611a86612c35565b5f73ffffffffffffffffffffffffffffffffffffffff81547fffffffffffffffffffffffff000000000000000000000000000000000000000081168355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346102945760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102945760043573ffffffffffffffffffffffffffffffffffffffff611b39612547565b1690611b46821515612796565b805f5260086020526117c473ffffffffffffffffffffffffffffffffffffffff60405f2054163314612687565b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457602060ff600454166040519015158152f35b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457602073ffffffffffffffffffffffffffffffffffffffff60045460081c16604051908152f35b34610294575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457611c3d612859565b60045460ff811615611c9a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166004557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610294576020611cdb611cd53661256a565b90612770565b604051908152f35b346102945760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457611d1a612547565b3373ffffffffffffffffffffffffffffffffffffffff821603611d4357610e5c90600435612b6b565b7f6697b232000000000000000000000000000000000000000000000000000000005f5260045ffd5b346102945760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457610e5c600435611da8612547565b90611dc2610e52825f526001602052600160405f20015490565b612a9d565b3461029457611dd53661256a565b611ddd612928565b611de561299f565b611def8183612770565b91825f52600560205260405f20600481019081549060ff8260401c16600581101561058357600103611f89576001015473ffffffffffffffffffffffffffffffffffffffff163303611f2b5767ffffffffffffffff16421015611ecd57680200000000000000007fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff8254161790557f7ebc677916fc5b3dc8915102e25ae9a58a5b0f21d7f29f1def498d234d5be92a6020604051338152a460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f616c72656164792072656c65617361626c6500000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f7420627579657200000000000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6e6f742064697370757461626c650000000000000000000000000000000000006044820152fd5b346102945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610294576020611cdb6004355f526001602052600160405f20015490565b346102945761203e3661256a565b90612047612928565b61204f61299f565b61205a821515612622565b805f52600860205261208773ffffffffffffffffffffffffffffffffffffffff60405f2054163314612687565b805f5260066020528160405f2054106121c857805f5260066020526120b08260405f20546125e1565b815f52600760205260405f20541161216a57805f52600660205260405f206120d98382546125e1565b905561210682337f00000000000000000000000000000000000000000000000000000000000000006129d3565b805f52600660205260405f205460405192835260208301527f213b7fab17342c2519d094fa4cd369b94b037e91f6d88a17b008bc3735607b1060403393a360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f756e646572636f6c6c61746572616c697a6500000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f696e73756666696369656e74207374616b6564000000000000000000000000006044820152fd5b346102945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610294576004355f526006602052602060405f2054604051908152f35b346102945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610294576122a5612524565b335f9081527f13a2c628591c5dedeb901e513a407216c0cb502840c34ca4c2e8f5159663bb5c602052604090205460ff16156123dc5773ffffffffffffffffffffffffffffffffffffffff811690811561237e5773ffffffffffffffffffffffffffffffffffffffff9074ffffffffffffffffffffffffffffffffffffffff006004549160081b167fffffffffffffffffffffff0000000000000000000000000000000000000000ff82161760045560081c167fc9a86b1ec2fde1df5e7074d929e13023beef2041df6abdcf8d669d94c427b72d5f80a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f7369676e65723d300000000000000000000000000000000000000000000000006044820152fd5b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004527f6f82f19233637990cc5fee2a338f11530b99de438ca6f1e3226d1410814c055e60245260445ffd5b346102945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610294576020611cdb6004356125ee565b346102945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261029457600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361029457817f7965db0b00000000000000000000000000000000000000000000000000000000602093149081156124fa575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836124f3565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361029457565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361029457565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112610294576004359060243590565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b91908203918211610b3c57565b805f52600660205260405f2054905f52600760205260405f20548082111561261c57612619916125e1565b90565b50505f90565b1561262957565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f616d6f756e743d300000000000000000000000000000000000000000000000006044820152fd5b1561268e57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f74206f776e657200000000000000000000000000000000000000000000006044820152fd5b906005811015610583577fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff68ff000000000000000083549260401b169116179055565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610b0f57604052565b90604051906020820192835260408201526040815261279060608261272f565b51902090565b1561279d57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f6f776e65723d30000000000000000000000000000000000000000000000000006044820152fd5b91908201809211610b3c57565b60443573ffffffffffffffffffffffffffffffffffffffff811681036102945790565b60843567ffffffffffffffff811681036102945790565b60a43567ffffffffffffffff811681036102945790565b335f9081527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49602052604090205460ff161561289157565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004525f60245260445ffd5b805f52600160205260405f2073ffffffffffffffffffffffffffffffffffffffff33165f5260205260ff60405f205416156128f95750565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f523360045260245260445ffd5b60027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0054146129775760027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f3ee5aeb5000000000000000000000000000000000000000000000000000000005f5260045ffd5b60ff600454166129ab57565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b9173ffffffffffffffffffffffffffffffffffffffff604051927fa9059cbb000000000000000000000000000000000000000000000000000000005f521660045260245260205f60448180865af19060015f5114821615612a7c575b60405215612a3a5750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b906001811516612a9457823b15153d15161690612a2f565b503d5f823e3d90fd5b805f52600160205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f205416155f1461261c57805f52600160205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905573ffffffffffffffffffffffffffffffffffffffff339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b805f52600160205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f2054165f1461261c57805f52600160205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815416905573ffffffffffffffffffffffffffffffffffffffff339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b73ffffffffffffffffffffffffffffffffffffffff5f54163303612c5557565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b92909173ffffffffffffffffffffffffffffffffffffffff9081604051947f23b872dd000000000000000000000000000000000000000000000000000000005f52166004521660245260445260205f60648180865af19060015f5114821615612cf4575b6040525f60605215612a3a5750565b906001811516612a9457823b15153d15161690612ce5565b60ff8114612d6b5760ff811690601f8211612d435760405191612d3060408461272f565b6020808452838101919036833783525290565b7fb3512b0c000000000000000000000000000000000000000000000000000000005f5260045ffd5b506040515f6002548060011c9160018216918215612e78575b602084108314612e4b578385528492908115612e0e5750600114612daf575b6126199250038261272f565b5060025f90815290917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace5b818310612df257505090602061261992820101612da3565b6020919350806001915483858801015201910190918392612dda565b602092506126199491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101612da3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b92607f1692612d84565b60ff8114612ea65760ff811690601f8211612d435760405191612d3060408461272f565b506040515f6003548060011c9160018216918215612f48575b602084108314612e4b578385528492908115612e0e5750600114612ee9576126199250038261272f565b5060035f90815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b818310612f2c57505090602061261992820101612da3565b6020919350806001915483858801015201910190918392612f14565b92607f1692612ebf565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001630148061304c575b15612fba577f000000000000000000000000000000000000000000000000000000000000000090565b60405160208101907f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f82527f000000000000000000000000000000000000000000000000000000000000000060408201527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260a0815261279060c08261272f565b507f00000000000000000000000000000000000000000000000000000000000000004614612f91565b81519190604183036130a55761309e9250602082015190606060408401519301515f1a9061315a565b9192909190565b50505f9160029190565b600481101561058357806130c1575050565b600181036130f1577ff645eedf000000000000000000000000000000000000000000000000000000005f5260045ffd5b6002810361312557507ffce698f7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b60031461312f5750565b7fd78bce0c000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116131e9579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa156131de575f5173ffffffffffffffffffffffffffffffffffffffff8116156131d457905f905f90565b505f906001905f90565b6040513d5f823e3d90fd5b5050505f916003919056fea2646970667358221220c900b11b0128e704e2da09c67243d23eeaac605b54f2c9664bbfa8bfe902c23964736f6c634300081c00332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0da6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4913a2c628591c5dedeb901e513a407216c0cb502840c34ca4c2e8f5159663bb5c44b31f1bc25958e4acb8433a7d0a722ee5ff13180d0cc5d5ced58fccd53eb77c4ea974b431a0b4c2e2044c8e7ab9ed5c93a2a8529b05175b6aea7d6522678c2b",
}

// QAEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use QAEscrowMetaData.ABI instead.
var QAEscrowABI = QAEscrowMetaData.ABI

// QAEscrowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QAEscrowMetaData.Bin instead.
var QAEscrowBin = QAEscrowMetaData.Bin

// DeployQAEscrow deploys a new Ethereum contract, binding an instance of QAEscrow to it.
func DeployQAEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, usdc common.Address, initialSigner common.Address) (common.Address, *types.Transaction, *QAEscrow, error) {
	parsed, err := QAEscrowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QAEscrowBin), backend, usdc, initialSigner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QAEscrow{QAEscrowCaller: QAEscrowCaller{contract: contract}, QAEscrowTransactor: QAEscrowTransactor{contract: contract}, QAEscrowFilterer: QAEscrowFilterer{contract: contract}}, nil
}

// QAEscrow is an auto generated Go binding around an Ethereum contract.
type QAEscrow struct {
	QAEscrowCaller     // Read-only binding to the contract
	QAEscrowTransactor // Write-only binding to the contract
	QAEscrowFilterer   // Log filterer for contract events
}

// QAEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type QAEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QAEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QAEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QAEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QAEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QAEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QAEscrowSession struct {
	Contract     *QAEscrow         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QAEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QAEscrowCallerSession struct {
	Contract *QAEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QAEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QAEscrowTransactorSession struct {
	Contract     *QAEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QAEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type QAEscrowRaw struct {
	Contract *QAEscrow // Generic contract binding to access the raw methods on
}

// QAEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QAEscrowCallerRaw struct {
	Contract *QAEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// QAEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QAEscrowTransactorRaw struct {
	Contract *QAEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQAEscrow creates a new instance of QAEscrow, bound to a specific deployed contract.
func NewQAEscrow(address common.Address, backend bind.ContractBackend) (*QAEscrow, error) {
	contract, err := bindQAEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QAEscrow{QAEscrowCaller: QAEscrowCaller{contract: contract}, QAEscrowTransactor: QAEscrowTransactor{contract: contract}, QAEscrowFilterer: QAEscrowFilterer{contract: contract}}, nil
}

// NewQAEscrowCaller creates a new read-only instance of QAEscrow, bound to a specific deployed contract.
func NewQAEscrowCaller(address common.Address, caller bind.ContractCaller) (*QAEscrowCaller, error) {
	contract, err := bindQAEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QAEscrowCaller{contract: contract}, nil
}

// NewQAEscrowTransactor creates a new write-only instance of QAEscrow, bound to a specific deployed contract.
func NewQAEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*QAEscrowTransactor, error) {
	contract, err := bindQAEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QAEscrowTransactor{contract: contract}, nil
}

// NewQAEscrowFilterer creates a new log filterer instance of QAEscrow, bound to a specific deployed contract.
func NewQAEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*QAEscrowFilterer, error) {
	contract, err := bindQAEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QAEscrowFilterer{contract: contract}, nil
}

// bindQAEscrow binds a generic wrapper to an already deployed contract.
func bindQAEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QAEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QAEscrow *QAEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QAEscrow.Contract.QAEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QAEscrow *QAEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QAEscrow.Contract.QAEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QAEscrow *QAEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QAEscrow.Contract.QAEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QAEscrow *QAEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QAEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QAEscrow *QAEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QAEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QAEscrow *QAEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QAEscrow.Contract.contract.Transact(opts, method, params...)
}

// APPADMINROLE is a free data retrieval call binding the contract method 0x92fb9db2.
//
// Solidity: function APP_ADMIN_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowCaller) APPADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "APP_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// APPADMINROLE is a free data retrieval call binding the contract method 0x92fb9db2.
//
// Solidity: function APP_ADMIN_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowSession) APPADMINROLE() ([32]byte, error) {
	return _QAEscrow.Contract.APPADMINROLE(&_QAEscrow.CallOpts)
}

// APPADMINROLE is a free data retrieval call binding the contract method 0x92fb9db2.
//
// Solidity: function APP_ADMIN_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowCallerSession) APPADMINROLE() ([32]byte, error) {
	return _QAEscrow.Contract.APPADMINROLE(&_QAEscrow.CallOpts)
}

// ARBITRATORROLE is a free data retrieval call binding the contract method 0x9cea0787.
//
// Solidity: function ARBITRATOR_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowCaller) ARBITRATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "ARBITRATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ARBITRATORROLE is a free data retrieval call binding the contract method 0x9cea0787.
//
// Solidity: function ARBITRATOR_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowSession) ARBITRATORROLE() ([32]byte, error) {
	return _QAEscrow.Contract.ARBITRATORROLE(&_QAEscrow.CallOpts)
}

// ARBITRATORROLE is a free data retrieval call binding the contract method 0x9cea0787.
//
// Solidity: function ARBITRATOR_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowCallerSession) ARBITRATORROLE() ([32]byte, error) {
	return _QAEscrow.Contract.ARBITRATORROLE(&_QAEscrow.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _QAEscrow.Contract.DEFAULTADMINROLE(&_QAEscrow.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _QAEscrow.Contract.DEFAULTADMINROLE(&_QAEscrow.CallOpts)
}

// ESCROWTERMSTYPEHASH is a free data retrieval call binding the contract method 0x95ce9b44.
//
// Solidity: function ESCROW_TERMS_TYPEHASH() view returns(bytes32)
func (_QAEscrow *QAEscrowCaller) ESCROWTERMSTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "ESCROW_TERMS_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ESCROWTERMSTYPEHASH is a free data retrieval call binding the contract method 0x95ce9b44.
//
// Solidity: function ESCROW_TERMS_TYPEHASH() view returns(bytes32)
func (_QAEscrow *QAEscrowSession) ESCROWTERMSTYPEHASH() ([32]byte, error) {
	return _QAEscrow.Contract.ESCROWTERMSTYPEHASH(&_QAEscrow.CallOpts)
}

// ESCROWTERMSTYPEHASH is a free data retrieval call binding the contract method 0x95ce9b44.
//
// Solidity: function ESCROW_TERMS_TYPEHASH() view returns(bytes32)
func (_QAEscrow *QAEscrowCallerSession) ESCROWTERMSTYPEHASH() ([32]byte, error) {
	return _QAEscrow.Contract.ESCROWTERMSTYPEHASH(&_QAEscrow.CallOpts)
}

// SIGNERADMINROLE is a free data retrieval call binding the contract method 0xa95bdf7f.
//
// Solidity: function SIGNER_ADMIN_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowCaller) SIGNERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "SIGNER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNERADMINROLE is a free data retrieval call binding the contract method 0xa95bdf7f.
//
// Solidity: function SIGNER_ADMIN_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowSession) SIGNERADMINROLE() ([32]byte, error) {
	return _QAEscrow.Contract.SIGNERADMINROLE(&_QAEscrow.CallOpts)
}

// SIGNERADMINROLE is a free data retrieval call binding the contract method 0xa95bdf7f.
//
// Solidity: function SIGNER_ADMIN_ROLE() view returns(bytes32)
func (_QAEscrow *QAEscrowCallerSession) SIGNERADMINROLE() ([32]byte, error) {
	return _QAEscrow.Contract.SIGNERADMINROLE(&_QAEscrow.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_QAEscrow *QAEscrowCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_QAEscrow *QAEscrowSession) USDC() (common.Address, error) {
	return _QAEscrow.Contract.USDC(&_QAEscrow.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_QAEscrow *QAEscrowCallerSession) USDC() (common.Address, error) {
	return _QAEscrow.Contract.USDC(&_QAEscrow.CallOpts)
}

// AppOwner is a free data retrieval call binding the contract method 0x8bf04c45.
//
// Solidity: function appOwner(bytes32 ) view returns(address)
func (_QAEscrow *QAEscrowCaller) AppOwner(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "appOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppOwner is a free data retrieval call binding the contract method 0x8bf04c45.
//
// Solidity: function appOwner(bytes32 ) view returns(address)
func (_QAEscrow *QAEscrowSession) AppOwner(arg0 [32]byte) (common.Address, error) {
	return _QAEscrow.Contract.AppOwner(&_QAEscrow.CallOpts, arg0)
}

// AppOwner is a free data retrieval call binding the contract method 0x8bf04c45.
//
// Solidity: function appOwner(bytes32 ) view returns(address)
func (_QAEscrow *QAEscrowCallerSession) AppOwner(arg0 [32]byte) (common.Address, error) {
	return _QAEscrow.Contract.AppOwner(&_QAEscrow.CallOpts, arg0)
}

// AvailableCapacity is a free data retrieval call binding the contract method 0x0581c36b.
//
// Solidity: function availableCapacity(bytes32 appId) view returns(uint256)
func (_QAEscrow *QAEscrowCaller) AvailableCapacity(opts *bind.CallOpts, appId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "availableCapacity", appId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableCapacity is a free data retrieval call binding the contract method 0x0581c36b.
//
// Solidity: function availableCapacity(bytes32 appId) view returns(uint256)
func (_QAEscrow *QAEscrowSession) AvailableCapacity(appId [32]byte) (*big.Int, error) {
	return _QAEscrow.Contract.AvailableCapacity(&_QAEscrow.CallOpts, appId)
}

// AvailableCapacity is a free data retrieval call binding the contract method 0x0581c36b.
//
// Solidity: function availableCapacity(bytes32 appId) view returns(uint256)
func (_QAEscrow *QAEscrowCallerSession) AvailableCapacity(appId [32]byte) (*big.Int, error) {
	return _QAEscrow.Contract.AvailableCapacity(&_QAEscrow.CallOpts, appId)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_QAEscrow *QAEscrowCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_QAEscrow *QAEscrowSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _QAEscrow.Contract.Eip712Domain(&_QAEscrow.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_QAEscrow *QAEscrowCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _QAEscrow.Contract.Eip712Domain(&_QAEscrow.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_QAEscrow *QAEscrowCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_QAEscrow *QAEscrowSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _QAEscrow.Contract.GetRoleAdmin(&_QAEscrow.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_QAEscrow *QAEscrowCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _QAEscrow.Contract.GetRoleAdmin(&_QAEscrow.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_QAEscrow *QAEscrowCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_QAEscrow *QAEscrowSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _QAEscrow.Contract.HasRole(&_QAEscrow.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_QAEscrow *QAEscrowCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _QAEscrow.Contract.HasRole(&_QAEscrow.CallOpts, role, account)
}

// Locked is a free data retrieval call binding the contract method 0xcbe9e764.
//
// Solidity: function locked(bytes32 ) view returns(uint256)
func (_QAEscrow *QAEscrowCaller) Locked(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "locked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcbe9e764.
//
// Solidity: function locked(bytes32 ) view returns(uint256)
func (_QAEscrow *QAEscrowSession) Locked(arg0 [32]byte) (*big.Int, error) {
	return _QAEscrow.Contract.Locked(&_QAEscrow.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbe9e764.
//
// Solidity: function locked(bytes32 ) view returns(uint256)
func (_QAEscrow *QAEscrowCallerSession) Locked(arg0 [32]byte) (*big.Int, error) {
	return _QAEscrow.Contract.Locked(&_QAEscrow.CallOpts, arg0)
}

// OrderKey is a free data retrieval call binding the contract method 0x3823eefc.
//
// Solidity: function orderKey(bytes32 appId, bytes32 orderId) pure returns(bytes32)
func (_QAEscrow *QAEscrowCaller) OrderKey(opts *bind.CallOpts, appId [32]byte, orderId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "orderKey", appId, orderId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrderKey is a free data retrieval call binding the contract method 0x3823eefc.
//
// Solidity: function orderKey(bytes32 appId, bytes32 orderId) pure returns(bytes32)
func (_QAEscrow *QAEscrowSession) OrderKey(appId [32]byte, orderId [32]byte) ([32]byte, error) {
	return _QAEscrow.Contract.OrderKey(&_QAEscrow.CallOpts, appId, orderId)
}

// OrderKey is a free data retrieval call binding the contract method 0x3823eefc.
//
// Solidity: function orderKey(bytes32 appId, bytes32 orderId) pure returns(bytes32)
func (_QAEscrow *QAEscrowCallerSession) OrderKey(appId [32]byte, orderId [32]byte) ([32]byte, error) {
	return _QAEscrow.Contract.OrderKey(&_QAEscrow.CallOpts, appId, orderId)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(bytes32 appId, address buyer, address merchant, uint256 amount, uint64 releaseAt, uint8 status)
func (_QAEscrow *QAEscrowCaller) Orders(opts *bind.CallOpts, arg0 [32]byte) (struct {
	AppId     [32]byte
	Buyer     common.Address
	Merchant  common.Address
	Amount    *big.Int
	ReleaseAt uint64
	Status    uint8
}, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		AppId     [32]byte
		Buyer     common.Address
		Merchant  common.Address
		Amount    *big.Int
		ReleaseAt uint64
		Status    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Buyer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Merchant = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReleaseAt = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(bytes32 appId, address buyer, address merchant, uint256 amount, uint64 releaseAt, uint8 status)
func (_QAEscrow *QAEscrowSession) Orders(arg0 [32]byte) (struct {
	AppId     [32]byte
	Buyer     common.Address
	Merchant  common.Address
	Amount    *big.Int
	ReleaseAt uint64
	Status    uint8
}, error) {
	return _QAEscrow.Contract.Orders(&_QAEscrow.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(bytes32 appId, address buyer, address merchant, uint256 amount, uint64 releaseAt, uint8 status)
func (_QAEscrow *QAEscrowCallerSession) Orders(arg0 [32]byte) (struct {
	AppId     [32]byte
	Buyer     common.Address
	Merchant  common.Address
	Amount    *big.Int
	ReleaseAt uint64
	Status    uint8
}, error) {
	return _QAEscrow.Contract.Orders(&_QAEscrow.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QAEscrow *QAEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QAEscrow *QAEscrowSession) Owner() (common.Address, error) {
	return _QAEscrow.Contract.Owner(&_QAEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QAEscrow *QAEscrowCallerSession) Owner() (common.Address, error) {
	return _QAEscrow.Contract.Owner(&_QAEscrow.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_QAEscrow *QAEscrowCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_QAEscrow *QAEscrowSession) Paused() (bool, error) {
	return _QAEscrow.Contract.Paused(&_QAEscrow.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_QAEscrow *QAEscrowCallerSession) Paused() (bool, error) {
	return _QAEscrow.Contract.Paused(&_QAEscrow.CallOpts)
}

// Staked is a free data retrieval call binding the contract method 0x120c857c.
//
// Solidity: function staked(bytes32 ) view returns(uint256)
func (_QAEscrow *QAEscrowCaller) Staked(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "staked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Staked is a free data retrieval call binding the contract method 0x120c857c.
//
// Solidity: function staked(bytes32 ) view returns(uint256)
func (_QAEscrow *QAEscrowSession) Staked(arg0 [32]byte) (*big.Int, error) {
	return _QAEscrow.Contract.Staked(&_QAEscrow.CallOpts, arg0)
}

// Staked is a free data retrieval call binding the contract method 0x120c857c.
//
// Solidity: function staked(bytes32 ) view returns(uint256)
func (_QAEscrow *QAEscrowCallerSession) Staked(arg0 [32]byte) (*big.Int, error) {
	return _QAEscrow.Contract.Staked(&_QAEscrow.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_QAEscrow *QAEscrowCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_QAEscrow *QAEscrowSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _QAEscrow.Contract.SupportsInterface(&_QAEscrow.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_QAEscrow *QAEscrowCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _QAEscrow.Contract.SupportsInterface(&_QAEscrow.CallOpts, interfaceId)
}

// TermsSigner is a free data retrieval call binding the contract method 0x5c39e043.
//
// Solidity: function termsSigner() view returns(address)
func (_QAEscrow *QAEscrowCaller) TermsSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QAEscrow.contract.Call(opts, &out, "termsSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TermsSigner is a free data retrieval call binding the contract method 0x5c39e043.
//
// Solidity: function termsSigner() view returns(address)
func (_QAEscrow *QAEscrowSession) TermsSigner() (common.Address, error) {
	return _QAEscrow.Contract.TermsSigner(&_QAEscrow.CallOpts)
}

// TermsSigner is a free data retrieval call binding the contract method 0x5c39e043.
//
// Solidity: function termsSigner() view returns(address)
func (_QAEscrow *QAEscrowCallerSession) TermsSigner() (common.Address, error) {
	return _QAEscrow.Contract.TermsSigner(&_QAEscrow.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x84cc9dfb.
//
// Solidity: function claim(bytes32 appId, bytes32 orderId) returns()
func (_QAEscrow *QAEscrowTransactor) Claim(opts *bind.TransactOpts, appId [32]byte, orderId [32]byte) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "claim", appId, orderId)
}

// Claim is a paid mutator transaction binding the contract method 0x84cc9dfb.
//
// Solidity: function claim(bytes32 appId, bytes32 orderId) returns()
func (_QAEscrow *QAEscrowSession) Claim(appId [32]byte, orderId [32]byte) (*types.Transaction, error) {
	return _QAEscrow.Contract.Claim(&_QAEscrow.TransactOpts, appId, orderId)
}

// Claim is a paid mutator transaction binding the contract method 0x84cc9dfb.
//
// Solidity: function claim(bytes32 appId, bytes32 orderId) returns()
func (_QAEscrow *QAEscrowTransactorSession) Claim(appId [32]byte, orderId [32]byte) (*types.Transaction, error) {
	return _QAEscrow.Contract.Claim(&_QAEscrow.TransactOpts, appId, orderId)
}

// DepositUsdc is a paid mutator transaction binding the contract method 0xe2e8df6f.
//
// Solidity: function depositUsdc((bytes32,bytes32,address,uint256,uint64,uint64) t, bytes sig) returns()
func (_QAEscrow *QAEscrowTransactor) DepositUsdc(opts *bind.TransactOpts, t QAEscrowEscrowTerms, sig []byte) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "depositUsdc", t, sig)
}

// DepositUsdc is a paid mutator transaction binding the contract method 0xe2e8df6f.
//
// Solidity: function depositUsdc((bytes32,bytes32,address,uint256,uint64,uint64) t, bytes sig) returns()
func (_QAEscrow *QAEscrowSession) DepositUsdc(t QAEscrowEscrowTerms, sig []byte) (*types.Transaction, error) {
	return _QAEscrow.Contract.DepositUsdc(&_QAEscrow.TransactOpts, t, sig)
}

// DepositUsdc is a paid mutator transaction binding the contract method 0xe2e8df6f.
//
// Solidity: function depositUsdc((bytes32,bytes32,address,uint256,uint64,uint64) t, bytes sig) returns()
func (_QAEscrow *QAEscrowTransactorSession) DepositUsdc(t QAEscrowEscrowTerms, sig []byte) (*types.Transaction, error) {
	return _QAEscrow.Contract.DepositUsdc(&_QAEscrow.TransactOpts, t, sig)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address to, uint256 amount) returns()
func (_QAEscrow *QAEscrowTransactor) EmergencyWithdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "emergencyWithdraw", to, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address to, uint256 amount) returns()
func (_QAEscrow *QAEscrowSession) EmergencyWithdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.Contract.EmergencyWithdraw(&_QAEscrow.TransactOpts, to, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address to, uint256 amount) returns()
func (_QAEscrow *QAEscrowTransactorSession) EmergencyWithdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.Contract.EmergencyWithdraw(&_QAEscrow.TransactOpts, to, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_QAEscrow *QAEscrowTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_QAEscrow *QAEscrowSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.GrantRole(&_QAEscrow.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_QAEscrow *QAEscrowTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.GrantRole(&_QAEscrow.TransactOpts, role, account)
}

// OpenDispute is a paid mutator transaction binding the contract method 0x2c4bc093.
//
// Solidity: function openDispute(bytes32 appId, bytes32 orderId) returns()
func (_QAEscrow *QAEscrowTransactor) OpenDispute(opts *bind.TransactOpts, appId [32]byte, orderId [32]byte) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "openDispute", appId, orderId)
}

// OpenDispute is a paid mutator transaction binding the contract method 0x2c4bc093.
//
// Solidity: function openDispute(bytes32 appId, bytes32 orderId) returns()
func (_QAEscrow *QAEscrowSession) OpenDispute(appId [32]byte, orderId [32]byte) (*types.Transaction, error) {
	return _QAEscrow.Contract.OpenDispute(&_QAEscrow.TransactOpts, appId, orderId)
}

// OpenDispute is a paid mutator transaction binding the contract method 0x2c4bc093.
//
// Solidity: function openDispute(bytes32 appId, bytes32 orderId) returns()
func (_QAEscrow *QAEscrowTransactorSession) OpenDispute(appId [32]byte, orderId [32]byte) (*types.Transaction, error) {
	return _QAEscrow.Contract.OpenDispute(&_QAEscrow.TransactOpts, appId, orderId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_QAEscrow *QAEscrowTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_QAEscrow *QAEscrowSession) Pause() (*types.Transaction, error) {
	return _QAEscrow.Contract.Pause(&_QAEscrow.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_QAEscrow *QAEscrowTransactorSession) Pause() (*types.Transaction, error) {
	return _QAEscrow.Contract.Pause(&_QAEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QAEscrow *QAEscrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QAEscrow *QAEscrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _QAEscrow.Contract.RenounceOwnership(&_QAEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QAEscrow *QAEscrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QAEscrow.Contract.RenounceOwnership(&_QAEscrow.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_QAEscrow *QAEscrowTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_QAEscrow *QAEscrowSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.RenounceRole(&_QAEscrow.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_QAEscrow *QAEscrowTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.RenounceRole(&_QAEscrow.TransactOpts, role, callerConfirmation)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xea35e258.
//
// Solidity: function resolveDispute(bytes32 appId, bytes32 orderId, bool buyerWins, uint256 buyerAmount) returns()
func (_QAEscrow *QAEscrowTransactor) ResolveDispute(opts *bind.TransactOpts, appId [32]byte, orderId [32]byte, buyerWins bool, buyerAmount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "resolveDispute", appId, orderId, buyerWins, buyerAmount)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xea35e258.
//
// Solidity: function resolveDispute(bytes32 appId, bytes32 orderId, bool buyerWins, uint256 buyerAmount) returns()
func (_QAEscrow *QAEscrowSession) ResolveDispute(appId [32]byte, orderId [32]byte, buyerWins bool, buyerAmount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.Contract.ResolveDispute(&_QAEscrow.TransactOpts, appId, orderId, buyerWins, buyerAmount)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xea35e258.
//
// Solidity: function resolveDispute(bytes32 appId, bytes32 orderId, bool buyerWins, uint256 buyerAmount) returns()
func (_QAEscrow *QAEscrowTransactorSession) ResolveDispute(appId [32]byte, orderId [32]byte, buyerWins bool, buyerAmount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.Contract.ResolveDispute(&_QAEscrow.TransactOpts, appId, orderId, buyerWins, buyerAmount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_QAEscrow *QAEscrowTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_QAEscrow *QAEscrowSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.RevokeRole(&_QAEscrow.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_QAEscrow *QAEscrowTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.RevokeRole(&_QAEscrow.TransactOpts, role, account)
}

// SetAppOwner is a paid mutator transaction binding the contract method 0x84b41e33.
//
// Solidity: function setAppOwner(bytes32 appId, address owner) returns()
func (_QAEscrow *QAEscrowTransactor) SetAppOwner(opts *bind.TransactOpts, appId [32]byte, owner common.Address) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "setAppOwner", appId, owner)
}

// SetAppOwner is a paid mutator transaction binding the contract method 0x84b41e33.
//
// Solidity: function setAppOwner(bytes32 appId, address owner) returns()
func (_QAEscrow *QAEscrowSession) SetAppOwner(appId [32]byte, owner common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.SetAppOwner(&_QAEscrow.TransactOpts, appId, owner)
}

// SetAppOwner is a paid mutator transaction binding the contract method 0x84b41e33.
//
// Solidity: function setAppOwner(bytes32 appId, address owner) returns()
func (_QAEscrow *QAEscrowTransactorSession) SetAppOwner(appId [32]byte, owner common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.SetAppOwner(&_QAEscrow.TransactOpts, appId, owner)
}

// SetTermsSigner is a paid mutator transaction binding the contract method 0x0e117e73.
//
// Solidity: function setTermsSigner(address newSigner) returns()
func (_QAEscrow *QAEscrowTransactor) SetTermsSigner(opts *bind.TransactOpts, newSigner common.Address) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "setTermsSigner", newSigner)
}

// SetTermsSigner is a paid mutator transaction binding the contract method 0x0e117e73.
//
// Solidity: function setTermsSigner(address newSigner) returns()
func (_QAEscrow *QAEscrowSession) SetTermsSigner(newSigner common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.SetTermsSigner(&_QAEscrow.TransactOpts, newSigner)
}

// SetTermsSigner is a paid mutator transaction binding the contract method 0x0e117e73.
//
// Solidity: function setTermsSigner(address newSigner) returns()
func (_QAEscrow *QAEscrowTransactorSession) SetTermsSigner(newSigner common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.SetTermsSigner(&_QAEscrow.TransactOpts, newSigner)
}

// Stake is a paid mutator transaction binding the contract method 0x8caa5230.
//
// Solidity: function stake(bytes32 appId, uint256 amount) returns()
func (_QAEscrow *QAEscrowTransactor) Stake(opts *bind.TransactOpts, appId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "stake", appId, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x8caa5230.
//
// Solidity: function stake(bytes32 appId, uint256 amount) returns()
func (_QAEscrow *QAEscrowSession) Stake(appId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.Contract.Stake(&_QAEscrow.TransactOpts, appId, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x8caa5230.
//
// Solidity: function stake(bytes32 appId, uint256 amount) returns()
func (_QAEscrow *QAEscrowTransactorSession) Stake(appId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.Contract.Stake(&_QAEscrow.TransactOpts, appId, amount)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x64f025b1.
//
// Solidity: function transferAppOwner(bytes32 appId, address newOwner) returns()
func (_QAEscrow *QAEscrowTransactor) TransferAppOwner(opts *bind.TransactOpts, appId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "transferAppOwner", appId, newOwner)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x64f025b1.
//
// Solidity: function transferAppOwner(bytes32 appId, address newOwner) returns()
func (_QAEscrow *QAEscrowSession) TransferAppOwner(appId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.TransferAppOwner(&_QAEscrow.TransactOpts, appId, newOwner)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x64f025b1.
//
// Solidity: function transferAppOwner(bytes32 appId, address newOwner) returns()
func (_QAEscrow *QAEscrowTransactorSession) TransferAppOwner(appId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.TransferAppOwner(&_QAEscrow.TransactOpts, appId, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QAEscrow *QAEscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QAEscrow *QAEscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.TransferOwnership(&_QAEscrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QAEscrow *QAEscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QAEscrow.Contract.TransferOwnership(&_QAEscrow.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_QAEscrow *QAEscrowTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_QAEscrow *QAEscrowSession) Unpause() (*types.Transaction, error) {
	return _QAEscrow.Contract.Unpause(&_QAEscrow.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_QAEscrow *QAEscrowTransactorSession) Unpause() (*types.Transaction, error) {
	return _QAEscrow.Contract.Unpause(&_QAEscrow.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x16ae261d.
//
// Solidity: function unstake(bytes32 appId, uint256 amount) returns()
func (_QAEscrow *QAEscrowTransactor) Unstake(opts *bind.TransactOpts, appId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.contract.Transact(opts, "unstake", appId, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x16ae261d.
//
// Solidity: function unstake(bytes32 appId, uint256 amount) returns()
func (_QAEscrow *QAEscrowSession) Unstake(appId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.Contract.Unstake(&_QAEscrow.TransactOpts, appId, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x16ae261d.
//
// Solidity: function unstake(bytes32 appId, uint256 amount) returns()
func (_QAEscrow *QAEscrowTransactorSession) Unstake(appId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _QAEscrow.Contract.Unstake(&_QAEscrow.TransactOpts, appId, amount)
}

// QAEscrowAppOwnerUpdatedIterator is returned from FilterAppOwnerUpdated and is used to iterate over the raw logs and unpacked data for AppOwnerUpdated events raised by the QAEscrow contract.
type QAEscrowAppOwnerUpdatedIterator struct {
	Event *QAEscrowAppOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *QAEscrowAppOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowAppOwnerUpdated)
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
		it.Event = new(QAEscrowAppOwnerUpdated)
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
func (it *QAEscrowAppOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowAppOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowAppOwnerUpdated represents a AppOwnerUpdated event raised by the QAEscrow contract.
type QAEscrowAppOwnerUpdated struct {
	AppId    [32]byte
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAppOwnerUpdated is a free log retrieval operation binding the contract event 0x9a5c7e2e887f2ba1250b313ad37205003e92ec9bcd2bc57b593646cd6c31bfdc.
//
// Solidity: event AppOwnerUpdated(bytes32 indexed appId, address indexed oldOwner, address indexed newOwner)
func (_QAEscrow *QAEscrowFilterer) FilterAppOwnerUpdated(opts *bind.FilterOpts, appId [][32]byte, oldOwner []common.Address, newOwner []common.Address) (*QAEscrowAppOwnerUpdatedIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "AppOwnerUpdated", appIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowAppOwnerUpdatedIterator{contract: _QAEscrow.contract, event: "AppOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchAppOwnerUpdated is a free log subscription operation binding the contract event 0x9a5c7e2e887f2ba1250b313ad37205003e92ec9bcd2bc57b593646cd6c31bfdc.
//
// Solidity: event AppOwnerUpdated(bytes32 indexed appId, address indexed oldOwner, address indexed newOwner)
func (_QAEscrow *QAEscrowFilterer) WatchAppOwnerUpdated(opts *bind.WatchOpts, sink chan<- *QAEscrowAppOwnerUpdated, appId [][32]byte, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "AppOwnerUpdated", appIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowAppOwnerUpdated)
				if err := _QAEscrow.contract.UnpackLog(event, "AppOwnerUpdated", log); err != nil {
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

// ParseAppOwnerUpdated is a log parse operation binding the contract event 0x9a5c7e2e887f2ba1250b313ad37205003e92ec9bcd2bc57b593646cd6c31bfdc.
//
// Solidity: event AppOwnerUpdated(bytes32 indexed appId, address indexed oldOwner, address indexed newOwner)
func (_QAEscrow *QAEscrowFilterer) ParseAppOwnerUpdated(log types.Log) (*QAEscrowAppOwnerUpdated, error) {
	event := new(QAEscrowAppOwnerUpdated)
	if err := _QAEscrow.contract.UnpackLog(event, "AppOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the QAEscrow contract.
type QAEscrowClaimedIterator struct {
	Event *QAEscrowClaimed // Event containing the contract specifics and raw log

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
func (it *QAEscrowClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowClaimed)
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
		it.Event = new(QAEscrowClaimed)
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
func (it *QAEscrowClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowClaimed represents a Claimed event raised by the QAEscrow contract.
type QAEscrowClaimed struct {
	AppId    [32]byte
	OrderId  [32]byte
	OrderKey [32]byte
	Merchant common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x08d0ecb8e19e562308f200cd130d9dba4e5e86592a6470ee028732e44edbe229.
//
// Solidity: event Claimed(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, address merchant, uint256 amount)
func (_QAEscrow *QAEscrowFilterer) FilterClaimed(opts *bind.FilterOpts, appId [][32]byte, orderId [][32]byte, orderKey [][32]byte) (*QAEscrowClaimedIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "Claimed", appIdRule, orderIdRule, orderKeyRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowClaimedIterator{contract: _QAEscrow.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x08d0ecb8e19e562308f200cd130d9dba4e5e86592a6470ee028732e44edbe229.
//
// Solidity: event Claimed(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, address merchant, uint256 amount)
func (_QAEscrow *QAEscrowFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *QAEscrowClaimed, appId [][32]byte, orderId [][32]byte, orderKey [][32]byte) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "Claimed", appIdRule, orderIdRule, orderKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowClaimed)
				if err := _QAEscrow.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x08d0ecb8e19e562308f200cd130d9dba4e5e86592a6470ee028732e44edbe229.
//
// Solidity: event Claimed(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, address merchant, uint256 amount)
func (_QAEscrow *QAEscrowFilterer) ParseClaimed(log types.Log) (*QAEscrowClaimed, error) {
	event := new(QAEscrowClaimed)
	if err := _QAEscrow.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the QAEscrow contract.
type QAEscrowDepositedIterator struct {
	Event *QAEscrowDeposited // Event containing the contract specifics and raw log

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
func (it *QAEscrowDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowDeposited)
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
		it.Event = new(QAEscrowDeposited)
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
func (it *QAEscrowDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowDeposited represents a Deposited event raised by the QAEscrow contract.
type QAEscrowDeposited struct {
	AppId     [32]byte
	OrderId   [32]byte
	OrderKey  [32]byte
	Buyer     common.Address
	Merchant  common.Address
	Amount    *big.Int
	ReleaseAt uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xe76a1d7f1e6df89045e72eaa27e22664ea43d53955a6cfe5e47ddb600b0b2239.
//
// Solidity: event Deposited(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, address buyer, address merchant, uint256 amount, uint64 releaseAt)
func (_QAEscrow *QAEscrowFilterer) FilterDeposited(opts *bind.FilterOpts, appId [][32]byte, orderId [][32]byte, orderKey [][32]byte) (*QAEscrowDepositedIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "Deposited", appIdRule, orderIdRule, orderKeyRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowDepositedIterator{contract: _QAEscrow.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xe76a1d7f1e6df89045e72eaa27e22664ea43d53955a6cfe5e47ddb600b0b2239.
//
// Solidity: event Deposited(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, address buyer, address merchant, uint256 amount, uint64 releaseAt)
func (_QAEscrow *QAEscrowFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *QAEscrowDeposited, appId [][32]byte, orderId [][32]byte, orderKey [][32]byte) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "Deposited", appIdRule, orderIdRule, orderKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowDeposited)
				if err := _QAEscrow.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xe76a1d7f1e6df89045e72eaa27e22664ea43d53955a6cfe5e47ddb600b0b2239.
//
// Solidity: event Deposited(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, address buyer, address merchant, uint256 amount, uint64 releaseAt)
func (_QAEscrow *QAEscrowFilterer) ParseDeposited(log types.Log) (*QAEscrowDeposited, error) {
	event := new(QAEscrowDeposited)
	if err := _QAEscrow.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowDisputeOpenedIterator is returned from FilterDisputeOpened and is used to iterate over the raw logs and unpacked data for DisputeOpened events raised by the QAEscrow contract.
type QAEscrowDisputeOpenedIterator struct {
	Event *QAEscrowDisputeOpened // Event containing the contract specifics and raw log

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
func (it *QAEscrowDisputeOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowDisputeOpened)
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
		it.Event = new(QAEscrowDisputeOpened)
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
func (it *QAEscrowDisputeOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowDisputeOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowDisputeOpened represents a DisputeOpened event raised by the QAEscrow contract.
type QAEscrowDisputeOpened struct {
	AppId    [32]byte
	OrderId  [32]byte
	OrderKey [32]byte
	Buyer    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDisputeOpened is a free log retrieval operation binding the contract event 0x7ebc677916fc5b3dc8915102e25ae9a58a5b0f21d7f29f1def498d234d5be92a.
//
// Solidity: event DisputeOpened(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, address buyer)
func (_QAEscrow *QAEscrowFilterer) FilterDisputeOpened(opts *bind.FilterOpts, appId [][32]byte, orderId [][32]byte, orderKey [][32]byte) (*QAEscrowDisputeOpenedIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "DisputeOpened", appIdRule, orderIdRule, orderKeyRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowDisputeOpenedIterator{contract: _QAEscrow.contract, event: "DisputeOpened", logs: logs, sub: sub}, nil
}

// WatchDisputeOpened is a free log subscription operation binding the contract event 0x7ebc677916fc5b3dc8915102e25ae9a58a5b0f21d7f29f1def498d234d5be92a.
//
// Solidity: event DisputeOpened(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, address buyer)
func (_QAEscrow *QAEscrowFilterer) WatchDisputeOpened(opts *bind.WatchOpts, sink chan<- *QAEscrowDisputeOpened, appId [][32]byte, orderId [][32]byte, orderKey [][32]byte) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "DisputeOpened", appIdRule, orderIdRule, orderKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowDisputeOpened)
				if err := _QAEscrow.contract.UnpackLog(event, "DisputeOpened", log); err != nil {
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

// ParseDisputeOpened is a log parse operation binding the contract event 0x7ebc677916fc5b3dc8915102e25ae9a58a5b0f21d7f29f1def498d234d5be92a.
//
// Solidity: event DisputeOpened(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, address buyer)
func (_QAEscrow *QAEscrowFilterer) ParseDisputeOpened(log types.Log) (*QAEscrowDisputeOpened, error) {
	event := new(QAEscrowDisputeOpened)
	if err := _QAEscrow.contract.UnpackLog(event, "DisputeOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the QAEscrow contract.
type QAEscrowDisputeResolvedIterator struct {
	Event *QAEscrowDisputeResolved // Event containing the contract specifics and raw log

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
func (it *QAEscrowDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowDisputeResolved)
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
		it.Event = new(QAEscrowDisputeResolved)
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
func (it *QAEscrowDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowDisputeResolved represents a DisputeResolved event raised by the QAEscrow contract.
type QAEscrowDisputeResolved struct {
	AppId          [32]byte
	OrderId        [32]byte
	OrderKey       [32]byte
	BuyerWins      bool
	BuyerAmount    *big.Int
	MerchantAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x8e9c1dd6f57a92254350fa9e6a6da6e64dc9ff0838bc0d7b82970f37e691d0d6.
//
// Solidity: event DisputeResolved(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, bool buyerWins, uint256 buyerAmount, uint256 merchantAmount)
func (_QAEscrow *QAEscrowFilterer) FilterDisputeResolved(opts *bind.FilterOpts, appId [][32]byte, orderId [][32]byte, orderKey [][32]byte) (*QAEscrowDisputeResolvedIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "DisputeResolved", appIdRule, orderIdRule, orderKeyRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowDisputeResolvedIterator{contract: _QAEscrow.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x8e9c1dd6f57a92254350fa9e6a6da6e64dc9ff0838bc0d7b82970f37e691d0d6.
//
// Solidity: event DisputeResolved(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, bool buyerWins, uint256 buyerAmount, uint256 merchantAmount)
func (_QAEscrow *QAEscrowFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *QAEscrowDisputeResolved, appId [][32]byte, orderId [][32]byte, orderKey [][32]byte) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "DisputeResolved", appIdRule, orderIdRule, orderKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowDisputeResolved)
				if err := _QAEscrow.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
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

// ParseDisputeResolved is a log parse operation binding the contract event 0x8e9c1dd6f57a92254350fa9e6a6da6e64dc9ff0838bc0d7b82970f37e691d0d6.
//
// Solidity: event DisputeResolved(bytes32 indexed appId, bytes32 indexed orderId, bytes32 indexed orderKey, bool buyerWins, uint256 buyerAmount, uint256 merchantAmount)
func (_QAEscrow *QAEscrowFilterer) ParseDisputeResolved(log types.Log) (*QAEscrowDisputeResolved, error) {
	event := new(QAEscrowDisputeResolved)
	if err := _QAEscrow.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the QAEscrow contract.
type QAEscrowEIP712DomainChangedIterator struct {
	Event *QAEscrowEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *QAEscrowEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowEIP712DomainChanged)
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
		it.Event = new(QAEscrowEIP712DomainChanged)
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
func (it *QAEscrowEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowEIP712DomainChanged represents a EIP712DomainChanged event raised by the QAEscrow contract.
type QAEscrowEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_QAEscrow *QAEscrowFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*QAEscrowEIP712DomainChangedIterator, error) {

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &QAEscrowEIP712DomainChangedIterator{contract: _QAEscrow.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_QAEscrow *QAEscrowFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *QAEscrowEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowEIP712DomainChanged)
				if err := _QAEscrow.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_QAEscrow *QAEscrowFilterer) ParseEIP712DomainChanged(log types.Log) (*QAEscrowEIP712DomainChanged, error) {
	event := new(QAEscrowEIP712DomainChanged)
	if err := _QAEscrow.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QAEscrow contract.
type QAEscrowOwnershipTransferredIterator struct {
	Event *QAEscrowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QAEscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowOwnershipTransferred)
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
		it.Event = new(QAEscrowOwnershipTransferred)
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
func (it *QAEscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowOwnershipTransferred represents a OwnershipTransferred event raised by the QAEscrow contract.
type QAEscrowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QAEscrow *QAEscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QAEscrowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowOwnershipTransferredIterator{contract: _QAEscrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QAEscrow *QAEscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QAEscrowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowOwnershipTransferred)
				if err := _QAEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QAEscrow *QAEscrowFilterer) ParseOwnershipTransferred(log types.Log) (*QAEscrowOwnershipTransferred, error) {
	event := new(QAEscrowOwnershipTransferred)
	if err := _QAEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the QAEscrow contract.
type QAEscrowPausedIterator struct {
	Event *QAEscrowPaused // Event containing the contract specifics and raw log

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
func (it *QAEscrowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowPaused)
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
		it.Event = new(QAEscrowPaused)
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
func (it *QAEscrowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowPaused represents a Paused event raised by the QAEscrow contract.
type QAEscrowPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_QAEscrow *QAEscrowFilterer) FilterPaused(opts *bind.FilterOpts) (*QAEscrowPausedIterator, error) {

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &QAEscrowPausedIterator{contract: _QAEscrow.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_QAEscrow *QAEscrowFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *QAEscrowPaused) (event.Subscription, error) {

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowPaused)
				if err := _QAEscrow.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_QAEscrow *QAEscrowFilterer) ParsePaused(log types.Log) (*QAEscrowPaused, error) {
	event := new(QAEscrowPaused)
	if err := _QAEscrow.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the QAEscrow contract.
type QAEscrowRoleAdminChangedIterator struct {
	Event *QAEscrowRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *QAEscrowRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowRoleAdminChanged)
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
		it.Event = new(QAEscrowRoleAdminChanged)
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
func (it *QAEscrowRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowRoleAdminChanged represents a RoleAdminChanged event raised by the QAEscrow contract.
type QAEscrowRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_QAEscrow *QAEscrowFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*QAEscrowRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowRoleAdminChangedIterator{contract: _QAEscrow.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_QAEscrow *QAEscrowFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *QAEscrowRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowRoleAdminChanged)
				if err := _QAEscrow.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_QAEscrow *QAEscrowFilterer) ParseRoleAdminChanged(log types.Log) (*QAEscrowRoleAdminChanged, error) {
	event := new(QAEscrowRoleAdminChanged)
	if err := _QAEscrow.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the QAEscrow contract.
type QAEscrowRoleGrantedIterator struct {
	Event *QAEscrowRoleGranted // Event containing the contract specifics and raw log

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
func (it *QAEscrowRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowRoleGranted)
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
		it.Event = new(QAEscrowRoleGranted)
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
func (it *QAEscrowRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowRoleGranted represents a RoleGranted event raised by the QAEscrow contract.
type QAEscrowRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_QAEscrow *QAEscrowFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*QAEscrowRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowRoleGrantedIterator{contract: _QAEscrow.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_QAEscrow *QAEscrowFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *QAEscrowRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowRoleGranted)
				if err := _QAEscrow.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_QAEscrow *QAEscrowFilterer) ParseRoleGranted(log types.Log) (*QAEscrowRoleGranted, error) {
	event := new(QAEscrowRoleGranted)
	if err := _QAEscrow.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the QAEscrow contract.
type QAEscrowRoleRevokedIterator struct {
	Event *QAEscrowRoleRevoked // Event containing the contract specifics and raw log

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
func (it *QAEscrowRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowRoleRevoked)
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
		it.Event = new(QAEscrowRoleRevoked)
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
func (it *QAEscrowRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowRoleRevoked represents a RoleRevoked event raised by the QAEscrow contract.
type QAEscrowRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_QAEscrow *QAEscrowFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*QAEscrowRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowRoleRevokedIterator{contract: _QAEscrow.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_QAEscrow *QAEscrowFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *QAEscrowRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowRoleRevoked)
				if err := _QAEscrow.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_QAEscrow *QAEscrowFilterer) ParseRoleRevoked(log types.Log) (*QAEscrowRoleRevoked, error) {
	event := new(QAEscrowRoleRevoked)
	if err := _QAEscrow.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the QAEscrow contract.
type QAEscrowStakedIterator struct {
	Event *QAEscrowStaked // Event containing the contract specifics and raw log

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
func (it *QAEscrowStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowStaked)
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
		it.Event = new(QAEscrowStaked)
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
func (it *QAEscrowStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowStaked represents a Staked event raised by the QAEscrow contract.
type QAEscrowStaked struct {
	AppId       [32]byte
	Merchant    common.Address
	Amount      *big.Int
	StakedTotal *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xa1fdccfe567643a44425efdd141171e8d992854a81e5c819c1432b0de47c9a11.
//
// Solidity: event Staked(bytes32 indexed appId, address indexed merchant, uint256 amount, uint256 stakedTotal)
func (_QAEscrow *QAEscrowFilterer) FilterStaked(opts *bind.FilterOpts, appId [][32]byte, merchant []common.Address) (*QAEscrowStakedIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "Staked", appIdRule, merchantRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowStakedIterator{contract: _QAEscrow.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xa1fdccfe567643a44425efdd141171e8d992854a81e5c819c1432b0de47c9a11.
//
// Solidity: event Staked(bytes32 indexed appId, address indexed merchant, uint256 amount, uint256 stakedTotal)
func (_QAEscrow *QAEscrowFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *QAEscrowStaked, appId [][32]byte, merchant []common.Address) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "Staked", appIdRule, merchantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowStaked)
				if err := _QAEscrow.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xa1fdccfe567643a44425efdd141171e8d992854a81e5c819c1432b0de47c9a11.
//
// Solidity: event Staked(bytes32 indexed appId, address indexed merchant, uint256 amount, uint256 stakedTotal)
func (_QAEscrow *QAEscrowFilterer) ParseStaked(log types.Log) (*QAEscrowStaked, error) {
	event := new(QAEscrowStaked)
	if err := _QAEscrow.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowTermsSignerUpdatedIterator is returned from FilterTermsSignerUpdated and is used to iterate over the raw logs and unpacked data for TermsSignerUpdated events raised by the QAEscrow contract.
type QAEscrowTermsSignerUpdatedIterator struct {
	Event *QAEscrowTermsSignerUpdated // Event containing the contract specifics and raw log

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
func (it *QAEscrowTermsSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowTermsSignerUpdated)
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
		it.Event = new(QAEscrowTermsSignerUpdated)
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
func (it *QAEscrowTermsSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowTermsSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowTermsSignerUpdated represents a TermsSignerUpdated event raised by the QAEscrow contract.
type QAEscrowTermsSignerUpdated struct {
	OldSigner common.Address
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTermsSignerUpdated is a free log retrieval operation binding the contract event 0xc9a86b1ec2fde1df5e7074d929e13023beef2041df6abdcf8d669d94c427b72d.
//
// Solidity: event TermsSignerUpdated(address indexed oldSigner, address indexed newSigner)
func (_QAEscrow *QAEscrowFilterer) FilterTermsSignerUpdated(opts *bind.FilterOpts, oldSigner []common.Address, newSigner []common.Address) (*QAEscrowTermsSignerUpdatedIterator, error) {

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "TermsSignerUpdated", oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowTermsSignerUpdatedIterator{contract: _QAEscrow.contract, event: "TermsSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchTermsSignerUpdated is a free log subscription operation binding the contract event 0xc9a86b1ec2fde1df5e7074d929e13023beef2041df6abdcf8d669d94c427b72d.
//
// Solidity: event TermsSignerUpdated(address indexed oldSigner, address indexed newSigner)
func (_QAEscrow *QAEscrowFilterer) WatchTermsSignerUpdated(opts *bind.WatchOpts, sink chan<- *QAEscrowTermsSignerUpdated, oldSigner []common.Address, newSigner []common.Address) (event.Subscription, error) {

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "TermsSignerUpdated", oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowTermsSignerUpdated)
				if err := _QAEscrow.contract.UnpackLog(event, "TermsSignerUpdated", log); err != nil {
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

// ParseTermsSignerUpdated is a log parse operation binding the contract event 0xc9a86b1ec2fde1df5e7074d929e13023beef2041df6abdcf8d669d94c427b72d.
//
// Solidity: event TermsSignerUpdated(address indexed oldSigner, address indexed newSigner)
func (_QAEscrow *QAEscrowFilterer) ParseTermsSignerUpdated(log types.Log) (*QAEscrowTermsSignerUpdated, error) {
	event := new(QAEscrowTermsSignerUpdated)
	if err := _QAEscrow.contract.UnpackLog(event, "TermsSignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the QAEscrow contract.
type QAEscrowUnpausedIterator struct {
	Event *QAEscrowUnpaused // Event containing the contract specifics and raw log

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
func (it *QAEscrowUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowUnpaused)
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
		it.Event = new(QAEscrowUnpaused)
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
func (it *QAEscrowUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowUnpaused represents a Unpaused event raised by the QAEscrow contract.
type QAEscrowUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_QAEscrow *QAEscrowFilterer) FilterUnpaused(opts *bind.FilterOpts) (*QAEscrowUnpausedIterator, error) {

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &QAEscrowUnpausedIterator{contract: _QAEscrow.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_QAEscrow *QAEscrowFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *QAEscrowUnpaused) (event.Subscription, error) {

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowUnpaused)
				if err := _QAEscrow.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_QAEscrow *QAEscrowFilterer) ParseUnpaused(log types.Log) (*QAEscrowUnpaused, error) {
	event := new(QAEscrowUnpaused)
	if err := _QAEscrow.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QAEscrowUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the QAEscrow contract.
type QAEscrowUnstakedIterator struct {
	Event *QAEscrowUnstaked // Event containing the contract specifics and raw log

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
func (it *QAEscrowUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QAEscrowUnstaked)
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
		it.Event = new(QAEscrowUnstaked)
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
func (it *QAEscrowUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QAEscrowUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QAEscrowUnstaked represents a Unstaked event raised by the QAEscrow contract.
type QAEscrowUnstaked struct {
	AppId       [32]byte
	Merchant    common.Address
	Amount      *big.Int
	StakedTotal *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x213b7fab17342c2519d094fa4cd369b94b037e91f6d88a17b008bc3735607b10.
//
// Solidity: event Unstaked(bytes32 indexed appId, address indexed merchant, uint256 amount, uint256 stakedTotal)
func (_QAEscrow *QAEscrowFilterer) FilterUnstaked(opts *bind.FilterOpts, appId [][32]byte, merchant []common.Address) (*QAEscrowUnstakedIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}

	logs, sub, err := _QAEscrow.contract.FilterLogs(opts, "Unstaked", appIdRule, merchantRule)
	if err != nil {
		return nil, err
	}
	return &QAEscrowUnstakedIterator{contract: _QAEscrow.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x213b7fab17342c2519d094fa4cd369b94b037e91f6d88a17b008bc3735607b10.
//
// Solidity: event Unstaked(bytes32 indexed appId, address indexed merchant, uint256 amount, uint256 stakedTotal)
func (_QAEscrow *QAEscrowFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *QAEscrowUnstaked, appId [][32]byte, merchant []common.Address) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}
	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}

	logs, sub, err := _QAEscrow.contract.WatchLogs(opts, "Unstaked", appIdRule, merchantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QAEscrowUnstaked)
				if err := _QAEscrow.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x213b7fab17342c2519d094fa4cd369b94b037e91f6d88a17b008bc3735607b10.
//
// Solidity: event Unstaked(bytes32 indexed appId, address indexed merchant, uint256 amount, uint256 stakedTotal)
func (_QAEscrow *QAEscrowFilterer) ParseUnstaked(log types.Log) (*QAEscrowUnstaked, error) {
	event := new(QAEscrowUnstaked)
	if err := _QAEscrow.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
