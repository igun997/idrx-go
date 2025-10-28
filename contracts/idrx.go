// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IDRXMetaData contains all meta data concerning the IDRX contract.
var IDRXMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountAfterCut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bridgeNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"BurnBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hashedAccountNumber\",\"type\":\"string\"}],\"name\":\"BurnWithAccountNumber\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountAfterCut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromBridgeNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"MintBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_burnBridgeFee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_mintBridgeFee\",\"type\":\"uint64\"}],\"name\":\"PlatformFeeInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLACKLIST_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PLATFORM_FEE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"}],\"name\":\"burnBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"accountNumber\",\"type\":\"string\"}],\"name\":\"burnWithAccountNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fromChainNonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPlatformFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainBridgeNonce\",\"type\":\"uint256\"}],\"name\":\"mintBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_burnBridgeFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_mintBridgeFee\",\"type\":\"uint64\"}],\"name\":\"setPlatformFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IDRXABI is the input ABI used to generate the binding from.
// Deprecated: Use IDRXMetaData.ABI instead.
var IDRXABI = IDRXMetaData.ABI

// IDRX is an auto generated Go binding around an Ethereum contract.
type IDRX struct {
	IDRXCaller     // Read-only binding to the contract
	IDRXTransactor // Write-only binding to the contract
	IDRXFilterer   // Log filterer for contract events
}

// IDRXCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDRXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDRXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDRXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDRXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDRXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDRXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDRXSession struct {
	Contract     *IDRX             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDRXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDRXCallerSession struct {
	Contract *IDRXCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IDRXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDRXTransactorSession struct {
	Contract     *IDRXTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDRXRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDRXRaw struct {
	Contract *IDRX // Generic contract binding to access the raw methods on
}

// IDRXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDRXCallerRaw struct {
	Contract *IDRXCaller // Generic read-only contract binding to access the raw methods on
}

// IDRXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDRXTransactorRaw struct {
	Contract *IDRXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDRX creates a new instance of IDRX, bound to a specific deployed contract.
func NewIDRX(address common.Address, backend bind.ContractBackend) (*IDRX, error) {
	contract, err := bindIDRX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDRX{IDRXCaller: IDRXCaller{contract: contract}, IDRXTransactor: IDRXTransactor{contract: contract}, IDRXFilterer: IDRXFilterer{contract: contract}}, nil
}

// NewIDRXCaller creates a new read-only instance of IDRX, bound to a specific deployed contract.
func NewIDRXCaller(address common.Address, caller bind.ContractCaller) (*IDRXCaller, error) {
	contract, err := bindIDRX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDRXCaller{contract: contract}, nil
}

// NewIDRXTransactor creates a new write-only instance of IDRX, bound to a specific deployed contract.
func NewIDRXTransactor(address common.Address, transactor bind.ContractTransactor) (*IDRXTransactor, error) {
	contract, err := bindIDRX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDRXTransactor{contract: contract}, nil
}

// NewIDRXFilterer creates a new log filterer instance of IDRX, bound to a specific deployed contract.
func NewIDRXFilterer(address common.Address, filterer bind.ContractFilterer) (*IDRXFilterer, error) {
	contract, err := bindIDRX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDRXFilterer{contract: contract}, nil
}

// bindIDRX binds a generic wrapper to an already deployed contract.
func bindIDRX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDRXMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDRX *IDRXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDRX.Contract.IDRXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDRX *IDRXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDRX.Contract.IDRXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDRX *IDRXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDRX.Contract.IDRXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDRX *IDRXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDRX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDRX *IDRXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDRX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDRX *IDRXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDRX.Contract.contract.Transact(opts, method, params...)
}

// BLACKLISTROLE is a free data retrieval call binding the contract method 0x19afe463.
//
// Solidity: function BLACKLIST_ROLE() view returns(bytes32)
func (_IDRX *IDRXCaller) BLACKLISTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "BLACKLIST_ROLE")
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// BLACKLISTROLE is a free data retrieval call binding the contract method 0x19afe463.
//
// Solidity: function BLACKLIST_ROLE() view returns(bytes32)
func (_IDRX *IDRXSession) BLACKLISTROLE() ([32]byte, error) {
	return _IDRX.Contract.BLACKLISTROLE(&_IDRX.CallOpts)
}

// BLACKLISTROLE is a free data retrieval call binding the contract method 0x19afe463.
//
// Solidity: function BLACKLIST_ROLE() view returns(bytes32)
func (_IDRX *IDRXCallerSession) BLACKLISTROLE() ([32]byte, error) {
	return _IDRX.Contract.BLACKLISTROLE(&_IDRX.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IDRX *IDRXCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IDRX *IDRXSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IDRX.Contract.DEFAULTADMINROLE(&_IDRX.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IDRX *IDRXCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IDRX.Contract.DEFAULTADMINROLE(&_IDRX.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_IDRX *IDRXCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "MINTER_ROLE")
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_IDRX *IDRXSession) MINTERROLE() ([32]byte, error) {
	return _IDRX.Contract.MINTERROLE(&_IDRX.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_IDRX *IDRXCallerSession) MINTERROLE() ([32]byte, error) {
	return _IDRX.Contract.MINTERROLE(&_IDRX.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_IDRX *IDRXCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "PAUSER_ROLE")
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_IDRX *IDRXSession) PAUSERROLE() ([32]byte, error) {
	return _IDRX.Contract.PAUSERROLE(&_IDRX.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_IDRX *IDRXCallerSession) PAUSERROLE() ([32]byte, error) {
	return _IDRX.Contract.PAUSERROLE(&_IDRX.CallOpts)
}

// PLATFORMFEESETTERROLE is a free data retrieval call binding the contract method 0x40532003.
//
// Solidity: function PLATFORM_FEE_SETTER_ROLE() view returns(bytes32)
func (_IDRX *IDRXCaller) PLATFORMFEESETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "PLATFORM_FEE_SETTER_ROLE")
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// PLATFORMFEESETTERROLE is a free data retrieval call binding the contract method 0x40532003.
//
// Solidity: function PLATFORM_FEE_SETTER_ROLE() view returns(bytes32)
func (_IDRX *IDRXSession) PLATFORMFEESETTERROLE() ([32]byte, error) {
	return _IDRX.Contract.PLATFORMFEESETTERROLE(&_IDRX.CallOpts)
}

// PLATFORMFEESETTERROLE is a free data retrieval call binding the contract method 0x40532003.
//
// Solidity: function PLATFORM_FEE_SETTER_ROLE() view returns(bytes32)
func (_IDRX *IDRXCallerSession) PLATFORMFEESETTERROLE() ([32]byte, error) {
	return _IDRX.Contract.PLATFORMFEESETTERROLE(&_IDRX.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_IDRX *IDRXCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "UPGRADER_ROLE")
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_IDRX *IDRXSession) UPGRADERROLE() ([32]byte, error) {
	return _IDRX.Contract.UPGRADERROLE(&_IDRX.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_IDRX *IDRXCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _IDRX.Contract.UPGRADERROLE(&_IDRX.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_IDRX *IDRXCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "_balances", arg0)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_IDRX *IDRXSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IDRX.Contract.Balances(&_IDRX.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_IDRX *IDRXCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IDRX.Contract.Balances(&_IDRX.CallOpts, arg0)
}

// BridgeNonce is a free data retrieval call binding the contract method 0x1a91b6ce.
//
// Solidity: function _bridgeNonce() view returns(uint256)
func (_IDRX *IDRXCaller) BridgeNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "_bridgeNonce")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BridgeNonce is a free data retrieval call binding the contract method 0x1a91b6ce.
//
// Solidity: function _bridgeNonce() view returns(uint256)
func (_IDRX *IDRXSession) BridgeNonce() (*big.Int, error) {
	return _IDRX.Contract.BridgeNonce(&_IDRX.CallOpts)
}

// BridgeNonce is a free data retrieval call binding the contract method 0x1a91b6ce.
//
// Solidity: function _bridgeNonce() view returns(uint256)
func (_IDRX *IDRXCallerSession) BridgeNonce() (*big.Int, error) {
	return _IDRX.Contract.BridgeNonce(&_IDRX.CallOpts)
}

// InternalTotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() view returns(uint256)
func (_IDRX *IDRXCaller) InternalTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "_totalSupply")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// InternalTotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() view returns(uint256)
func (_IDRX *IDRXSession) InternalTotalSupply() (*big.Int, error) {
	return _IDRX.Contract.InternalTotalSupply(&_IDRX.CallOpts)
}

// InternalTotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() view returns(uint256)
func (_IDRX *IDRXCallerSession) InternalTotalSupply() (*big.Int, error) {
	return _IDRX.Contract.InternalTotalSupply(&_IDRX.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IDRX *IDRXCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "allowance", owner, spender)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IDRX *IDRXSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IDRX.Contract.Allowance(&_IDRX.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IDRX *IDRXCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IDRX.Contract.Allowance(&_IDRX.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IDRX *IDRXCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "balanceOf", account)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IDRX *IDRXSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IDRX.Contract.BalanceOf(&_IDRX.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IDRX *IDRXCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IDRX.Contract.BalanceOf(&_IDRX.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IDRX *IDRXCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "decimals")
	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IDRX *IDRXSession) Decimals() (uint8, error) {
	return _IDRX.Contract.Decimals(&_IDRX.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IDRX *IDRXCallerSession) Decimals() (uint8, error) {
	return _IDRX.Contract.Decimals(&_IDRX.CallOpts)
}

// FromChainNonceUsed is a free data retrieval call binding the contract method 0xf3e95776.
//
// Solidity: function fromChainNonceUsed(uint256 , uint256 ) view returns(bool)
func (_IDRX *IDRXCaller) FromChainNonceUsed(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "fromChainNonceUsed", arg0, arg1)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// FromChainNonceUsed is a free data retrieval call binding the contract method 0xf3e95776.
//
// Solidity: function fromChainNonceUsed(uint256 , uint256 ) view returns(bool)
func (_IDRX *IDRXSession) FromChainNonceUsed(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _IDRX.Contract.FromChainNonceUsed(&_IDRX.CallOpts, arg0, arg1)
}

// FromChainNonceUsed is a free data retrieval call binding the contract method 0xf3e95776.
//
// Solidity: function fromChainNonceUsed(uint256 , uint256 ) view returns(bool)
func (_IDRX *IDRXCallerSession) FromChainNonceUsed(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _IDRX.Contract.FromChainNonceUsed(&_IDRX.CallOpts, arg0, arg1)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_IDRX *IDRXCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "getBlackListStatus", _maker)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_IDRX *IDRXSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _IDRX.Contract.GetBlackListStatus(&_IDRX.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_IDRX *IDRXCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _IDRX.Contract.GetBlackListStatus(&_IDRX.CallOpts, _maker)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint64, uint64)
func (_IDRX *IDRXCaller) GetPlatformFeeInfo(opts *bind.CallOpts) (common.Address, uint64, uint64, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "getPlatformFeeInfo")
	if err != nil {
		return *new(common.Address), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint64, uint64)
func (_IDRX *IDRXSession) GetPlatformFeeInfo() (common.Address, uint64, uint64, error) {
	return _IDRX.Contract.GetPlatformFeeInfo(&_IDRX.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint64, uint64)
func (_IDRX *IDRXCallerSession) GetPlatformFeeInfo() (common.Address, uint64, uint64, error) {
	return _IDRX.Contract.GetPlatformFeeInfo(&_IDRX.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IDRX *IDRXCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "getRoleAdmin", role)
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IDRX *IDRXSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IDRX.Contract.GetRoleAdmin(&_IDRX.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IDRX *IDRXCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IDRX.Contract.GetRoleAdmin(&_IDRX.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IDRX *IDRXCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "hasRole", role, account)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IDRX *IDRXSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IDRX.Contract.HasRole(&_IDRX.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IDRX *IDRXCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IDRX.Contract.HasRole(&_IDRX.CallOpts, role, account)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_IDRX *IDRXCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "isBlackListed", arg0)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_IDRX *IDRXSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _IDRX.Contract.IsBlackListed(&_IDRX.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_IDRX *IDRXCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _IDRX.Contract.IsBlackListed(&_IDRX.CallOpts, arg0)
}

// MaxPlatformFee is a free data retrieval call binding the contract method 0x62a6f6a6.
//
// Solidity: function maxPlatformFee() view returns(uint64)
func (_IDRX *IDRXCaller) MaxPlatformFee(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "maxPlatformFee")
	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err
}

// MaxPlatformFee is a free data retrieval call binding the contract method 0x62a6f6a6.
//
// Solidity: function maxPlatformFee() view returns(uint64)
func (_IDRX *IDRXSession) MaxPlatformFee() (uint64, error) {
	return _IDRX.Contract.MaxPlatformFee(&_IDRX.CallOpts)
}

// MaxPlatformFee is a free data retrieval call binding the contract method 0x62a6f6a6.
//
// Solidity: function maxPlatformFee() view returns(uint64)
func (_IDRX *IDRXCallerSession) MaxPlatformFee() (uint64, error) {
	return _IDRX.Contract.MaxPlatformFee(&_IDRX.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IDRX *IDRXCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "name")
	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IDRX *IDRXSession) Name() (string, error) {
	return _IDRX.Contract.Name(&_IDRX.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IDRX *IDRXCallerSession) Name() (string, error) {
	return _IDRX.Contract.Name(&_IDRX.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IDRX *IDRXCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "paused")
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IDRX *IDRXSession) Paused() (bool, error) {
	return _IDRX.Contract.Paused(&_IDRX.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IDRX *IDRXCallerSession) Paused() (bool, error) {
	return _IDRX.Contract.Paused(&_IDRX.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IDRX *IDRXCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "proxiableUUID")
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IDRX *IDRXSession) ProxiableUUID() ([32]byte, error) {
	return _IDRX.Contract.ProxiableUUID(&_IDRX.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IDRX *IDRXCallerSession) ProxiableUUID() ([32]byte, error) {
	return _IDRX.Contract.ProxiableUUID(&_IDRX.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IDRX *IDRXCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "supportsInterface", interfaceId)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IDRX *IDRXSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IDRX.Contract.SupportsInterface(&_IDRX.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IDRX *IDRXCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IDRX.Contract.SupportsInterface(&_IDRX.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IDRX *IDRXCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "symbol")
	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IDRX *IDRXSession) Symbol() (string, error) {
	return _IDRX.Contract.Symbol(&_IDRX.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IDRX *IDRXCallerSession) Symbol() (string, error) {
	return _IDRX.Contract.Symbol(&_IDRX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IDRX *IDRXCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDRX.contract.Call(opts, &out, "totalSupply")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IDRX *IDRXSession) TotalSupply() (*big.Int, error) {
	return _IDRX.Contract.TotalSupply(&_IDRX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IDRX *IDRXCallerSession) TotalSupply() (*big.Int, error) {
	return _IDRX.Contract.TotalSupply(&_IDRX.CallOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_IDRX *IDRXTransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_IDRX *IDRXSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.AddBlackList(&_IDRX.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_IDRX *IDRXTransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.AddBlackList(&_IDRX.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IDRX *IDRXTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IDRX *IDRXSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.Approve(&_IDRX.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IDRX *IDRXTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.Approve(&_IDRX.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IDRX *IDRXTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IDRX *IDRXSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.Burn(&_IDRX.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IDRX *IDRXTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.Burn(&_IDRX.TransactOpts, amount)
}

// BurnBridge is a paid mutator transaction binding the contract method 0x875858c9.
//
// Solidity: function burnBridge(uint256 amount, uint256 toChain) returns()
func (_IDRX *IDRXTransactor) BurnBridge(opts *bind.TransactOpts, amount *big.Int, toChain *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "burnBridge", amount, toChain)
}

// BurnBridge is a paid mutator transaction binding the contract method 0x875858c9.
//
// Solidity: function burnBridge(uint256 amount, uint256 toChain) returns()
func (_IDRX *IDRXSession) BurnBridge(amount *big.Int, toChain *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.BurnBridge(&_IDRX.TransactOpts, amount, toChain)
}

// BurnBridge is a paid mutator transaction binding the contract method 0x875858c9.
//
// Solidity: function burnBridge(uint256 amount, uint256 toChain) returns()
func (_IDRX *IDRXTransactorSession) BurnBridge(amount *big.Int, toChain *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.BurnBridge(&_IDRX.TransactOpts, amount, toChain)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IDRX *IDRXTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IDRX *IDRXSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.BurnFrom(&_IDRX.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IDRX *IDRXTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.BurnFrom(&_IDRX.TransactOpts, account, amount)
}

// BurnWithAccountNumber is a paid mutator transaction binding the contract method 0x880939a1.
//
// Solidity: function burnWithAccountNumber(uint256 amount, string accountNumber) returns()
func (_IDRX *IDRXTransactor) BurnWithAccountNumber(opts *bind.TransactOpts, amount *big.Int, accountNumber string) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "burnWithAccountNumber", amount, accountNumber)
}

// BurnWithAccountNumber is a paid mutator transaction binding the contract method 0x880939a1.
//
// Solidity: function burnWithAccountNumber(uint256 amount, string accountNumber) returns()
func (_IDRX *IDRXSession) BurnWithAccountNumber(amount *big.Int, accountNumber string) (*types.Transaction, error) {
	return _IDRX.Contract.BurnWithAccountNumber(&_IDRX.TransactOpts, amount, accountNumber)
}

// BurnWithAccountNumber is a paid mutator transaction binding the contract method 0x880939a1.
//
// Solidity: function burnWithAccountNumber(uint256 amount, string accountNumber) returns()
func (_IDRX *IDRXTransactorSession) BurnWithAccountNumber(amount *big.Int, accountNumber string) (*types.Transaction, error) {
	return _IDRX.Contract.BurnWithAccountNumber(&_IDRX.TransactOpts, amount, accountNumber)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IDRX *IDRXTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IDRX *IDRXSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.DecreaseAllowance(&_IDRX.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IDRX *IDRXTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.DecreaseAllowance(&_IDRX.TransactOpts, spender, subtractedValue)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_IDRX *IDRXTransactor) DestroyBlackFunds(opts *bind.TransactOpts, _blackListedUser common.Address) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "destroyBlackFunds", _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_IDRX *IDRXSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.DestroyBlackFunds(&_IDRX.TransactOpts, _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_IDRX *IDRXTransactorSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.DestroyBlackFunds(&_IDRX.TransactOpts, _blackListedUser)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IDRX *IDRXTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IDRX *IDRXSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.GrantRole(&_IDRX.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IDRX *IDRXTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.GrantRole(&_IDRX.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IDRX *IDRXTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IDRX *IDRXSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.IncreaseAllowance(&_IDRX.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IDRX *IDRXTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.IncreaseAllowance(&_IDRX.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IDRX *IDRXTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IDRX *IDRXSession) Initialize() (*types.Transaction, error) {
	return _IDRX.Contract.Initialize(&_IDRX.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IDRX *IDRXTransactorSession) Initialize() (*types.Transaction, error) {
	return _IDRX.Contract.Initialize(&_IDRX.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IDRX *IDRXTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IDRX *IDRXSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.Mint(&_IDRX.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IDRX *IDRXTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.Mint(&_IDRX.TransactOpts, to, amount)
}

// MintBridge is a paid mutator transaction binding the contract method 0x83104170.
//
// Solidity: function mintBridge(address to, uint256 amount, uint256 fromChain, uint256 fromChainBridgeNonce) returns()
func (_IDRX *IDRXTransactor) MintBridge(opts *bind.TransactOpts, to common.Address, amount *big.Int, fromChain *big.Int, fromChainBridgeNonce *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "mintBridge", to, amount, fromChain, fromChainBridgeNonce)
}

// MintBridge is a paid mutator transaction binding the contract method 0x83104170.
//
// Solidity: function mintBridge(address to, uint256 amount, uint256 fromChain, uint256 fromChainBridgeNonce) returns()
func (_IDRX *IDRXSession) MintBridge(to common.Address, amount *big.Int, fromChain *big.Int, fromChainBridgeNonce *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.MintBridge(&_IDRX.TransactOpts, to, amount, fromChain, fromChainBridgeNonce)
}

// MintBridge is a paid mutator transaction binding the contract method 0x83104170.
//
// Solidity: function mintBridge(address to, uint256 amount, uint256 fromChain, uint256 fromChainBridgeNonce) returns()
func (_IDRX *IDRXTransactorSession) MintBridge(to common.Address, amount *big.Int, fromChain *big.Int, fromChainBridgeNonce *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.MintBridge(&_IDRX.TransactOpts, to, amount, fromChain, fromChainBridgeNonce)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IDRX *IDRXTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IDRX *IDRXSession) Pause() (*types.Transaction, error) {
	return _IDRX.Contract.Pause(&_IDRX.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IDRX *IDRXTransactorSession) Pause() (*types.Transaction, error) {
	return _IDRX.Contract.Pause(&_IDRX.TransactOpts)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_IDRX *IDRXTransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_IDRX *IDRXSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.RemoveBlackList(&_IDRX.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_IDRX *IDRXTransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.RemoveBlackList(&_IDRX.TransactOpts, _clearedUser)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IDRX *IDRXTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IDRX *IDRXSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.RenounceRole(&_IDRX.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IDRX *IDRXTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.RenounceRole(&_IDRX.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IDRX *IDRXTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IDRX *IDRXSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.RevokeRole(&_IDRX.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IDRX *IDRXTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.RevokeRole(&_IDRX.TransactOpts, role, account)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x635eaf34.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint64 _burnBridgeFee, uint64 _mintBridgeFee) returns()
func (_IDRX *IDRXTransactor) SetPlatformFeeInfo(opts *bind.TransactOpts, _platformFeeRecipient common.Address, _burnBridgeFee uint64, _mintBridgeFee uint64) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "setPlatformFeeInfo", _platformFeeRecipient, _burnBridgeFee, _mintBridgeFee)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x635eaf34.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint64 _burnBridgeFee, uint64 _mintBridgeFee) returns()
func (_IDRX *IDRXSession) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _burnBridgeFee uint64, _mintBridgeFee uint64) (*types.Transaction, error) {
	return _IDRX.Contract.SetPlatformFeeInfo(&_IDRX.TransactOpts, _platformFeeRecipient, _burnBridgeFee, _mintBridgeFee)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x635eaf34.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint64 _burnBridgeFee, uint64 _mintBridgeFee) returns()
func (_IDRX *IDRXTransactorSession) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _burnBridgeFee uint64, _mintBridgeFee uint64) (*types.Transaction, error) {
	return _IDRX.Contract.SetPlatformFeeInfo(&_IDRX.TransactOpts, _platformFeeRecipient, _burnBridgeFee, _mintBridgeFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IDRX *IDRXTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IDRX *IDRXSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.Transfer(&_IDRX.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IDRX *IDRXTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.Transfer(&_IDRX.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IDRX *IDRXTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IDRX *IDRXSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.TransferFrom(&_IDRX.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IDRX *IDRXTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDRX.Contract.TransferFrom(&_IDRX.TransactOpts, from, to, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IDRX *IDRXTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IDRX *IDRXSession) Unpause() (*types.Transaction, error) {
	return _IDRX.Contract.Unpause(&_IDRX.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IDRX *IDRXTransactorSession) Unpause() (*types.Transaction, error) {
	return _IDRX.Contract.Unpause(&_IDRX.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_IDRX *IDRXTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_IDRX *IDRXSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.UpgradeTo(&_IDRX.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_IDRX *IDRXTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _IDRX.Contract.UpgradeTo(&_IDRX.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IDRX *IDRXTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IDRX.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IDRX *IDRXSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IDRX.Contract.UpgradeToAndCall(&_IDRX.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IDRX *IDRXTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IDRX.Contract.UpgradeToAndCall(&_IDRX.TransactOpts, newImplementation, data)
}

// IDRXAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the IDRX contract.
type IDRXAddedBlackListIterator struct {
	Event *IDRXAddedBlackList // Event containing the contract specifics and raw log

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
func (it *IDRXAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXAddedBlackList)
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
		it.Event = new(IDRXAddedBlackList)
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
func (it *IDRXAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXAddedBlackList represents a AddedBlackList event raised by the IDRX contract.
type IDRXAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_IDRX *IDRXFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*IDRXAddedBlackListIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &IDRXAddedBlackListIterator{contract: _IDRX.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_IDRX *IDRXFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *IDRXAddedBlackList) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXAddedBlackList)
				if err := _IDRX.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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

// ParseAddedBlackList is a log parse operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_IDRX *IDRXFilterer) ParseAddedBlackList(log types.Log) (*IDRXAddedBlackList, error) {
	event := new(IDRXAddedBlackList)
	if err := _IDRX.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the IDRX contract.
type IDRXAdminChangedIterator struct {
	Event *IDRXAdminChanged // Event containing the contract specifics and raw log

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
func (it *IDRXAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXAdminChanged)
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
		it.Event = new(IDRXAdminChanged)
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
func (it *IDRXAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXAdminChanged represents a AdminChanged event raised by the IDRX contract.
type IDRXAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IDRX *IDRXFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*IDRXAdminChangedIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &IDRXAdminChangedIterator{contract: _IDRX.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IDRX *IDRXFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *IDRXAdminChanged) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXAdminChanged)
				if err := _IDRX.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IDRX *IDRXFilterer) ParseAdminChanged(log types.Log) (*IDRXAdminChanged, error) {
	event := new(IDRXAdminChanged)
	if err := _IDRX.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IDRX contract.
type IDRXApprovalIterator struct {
	Event *IDRXApproval // Event containing the contract specifics and raw log

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
func (it *IDRXApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXApproval)
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
		it.Event = new(IDRXApproval)
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
func (it *IDRXApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXApproval represents a Approval event raised by the IDRX contract.
type IDRXApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IDRX *IDRXFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IDRXApprovalIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IDRX.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IDRXApprovalIterator{contract: _IDRX.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IDRX *IDRXFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IDRXApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IDRX.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXApproval)
				if err := _IDRX.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IDRX *IDRXFilterer) ParseApproval(log types.Log) (*IDRXApproval, error) {
	event := new(IDRXApproval)
	if err := _IDRX.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the IDRX contract.
type IDRXBeaconUpgradedIterator struct {
	Event *IDRXBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *IDRXBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXBeaconUpgraded)
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
		it.Event = new(IDRXBeaconUpgraded)
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
func (it *IDRXBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXBeaconUpgraded represents a BeaconUpgraded event raised by the IDRX contract.
type IDRXBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IDRX *IDRXFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*IDRXBeaconUpgradedIterator, error) {
	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _IDRX.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &IDRXBeaconUpgradedIterator{contract: _IDRX.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IDRX *IDRXFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *IDRXBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {
	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _IDRX.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXBeaconUpgraded)
				if err := _IDRX.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IDRX *IDRXFilterer) ParseBeaconUpgraded(log types.Log) (*IDRXBeaconUpgraded, error) {
	event := new(IDRXBeaconUpgraded)
	if err := _IDRX.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXBurnBridgeIterator is returned from FilterBurnBridge and is used to iterate over the raw logs and unpacked data for BurnBridge events raised by the IDRX contract.
type IDRXBurnBridgeIterator struct {
	Event *IDRXBurnBridge // Event containing the contract specifics and raw log

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
func (it *IDRXBurnBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXBurnBridge)
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
		it.Event = new(IDRXBurnBridge)
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
func (it *IDRXBurnBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXBurnBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXBurnBridge represents a BurnBridge event raised by the IDRX contract.
type IDRXBurnBridge struct {
	User           common.Address
	Amount         *big.Int
	AmountAfterCut *big.Int
	ToChain        *big.Int
	BridgeNonce    *big.Int
	PlatformFee    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBurnBridge is a free log retrieval operation binding the contract event 0xa1cde98fa32e68f296db8cafbf1be704c08ff3aee6f94cc85e64c7d8615f58c3.
//
// Solidity: event BurnBridge(address _user, uint256 _amount, uint256 amountAfterCut, uint256 toChain, uint256 _bridgeNonce, uint256 platformFee)
func (_IDRX *IDRXFilterer) FilterBurnBridge(opts *bind.FilterOpts) (*IDRXBurnBridgeIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "BurnBridge")
	if err != nil {
		return nil, err
	}
	return &IDRXBurnBridgeIterator{contract: _IDRX.contract, event: "BurnBridge", logs: logs, sub: sub}, nil
}

// WatchBurnBridge is a free log subscription operation binding the contract event 0xa1cde98fa32e68f296db8cafbf1be704c08ff3aee6f94cc85e64c7d8615f58c3.
//
// Solidity: event BurnBridge(address _user, uint256 _amount, uint256 amountAfterCut, uint256 toChain, uint256 _bridgeNonce, uint256 platformFee)
func (_IDRX *IDRXFilterer) WatchBurnBridge(opts *bind.WatchOpts, sink chan<- *IDRXBurnBridge) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "BurnBridge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXBurnBridge)
				if err := _IDRX.contract.UnpackLog(event, "BurnBridge", log); err != nil {
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

// ParseBurnBridge is a log parse operation binding the contract event 0xa1cde98fa32e68f296db8cafbf1be704c08ff3aee6f94cc85e64c7d8615f58c3.
//
// Solidity: event BurnBridge(address _user, uint256 _amount, uint256 amountAfterCut, uint256 toChain, uint256 _bridgeNonce, uint256 platformFee)
func (_IDRX *IDRXFilterer) ParseBurnBridge(log types.Log) (*IDRXBurnBridge, error) {
	event := new(IDRXBurnBridge)
	if err := _IDRX.contract.UnpackLog(event, "BurnBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXBurnWithAccountNumberIterator is returned from FilterBurnWithAccountNumber and is used to iterate over the raw logs and unpacked data for BurnWithAccountNumber events raised by the IDRX contract.
type IDRXBurnWithAccountNumberIterator struct {
	Event *IDRXBurnWithAccountNumber // Event containing the contract specifics and raw log

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
func (it *IDRXBurnWithAccountNumberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXBurnWithAccountNumber)
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
		it.Event = new(IDRXBurnWithAccountNumber)
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
func (it *IDRXBurnWithAccountNumberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXBurnWithAccountNumberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXBurnWithAccountNumber represents a BurnWithAccountNumber event raised by the IDRX contract.
type IDRXBurnWithAccountNumber struct {
	User                common.Address
	Amount              *big.Int
	HashedAccountNumber string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBurnWithAccountNumber is a free log retrieval operation binding the contract event 0x35626062961ef23c4eea38fafb3055f7c3b7430fa14c95a8da518d2574f8cbb3.
//
// Solidity: event BurnWithAccountNumber(address _user, uint256 amount, string hashedAccountNumber)
func (_IDRX *IDRXFilterer) FilterBurnWithAccountNumber(opts *bind.FilterOpts) (*IDRXBurnWithAccountNumberIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "BurnWithAccountNumber")
	if err != nil {
		return nil, err
	}
	return &IDRXBurnWithAccountNumberIterator{contract: _IDRX.contract, event: "BurnWithAccountNumber", logs: logs, sub: sub}, nil
}

// WatchBurnWithAccountNumber is a free log subscription operation binding the contract event 0x35626062961ef23c4eea38fafb3055f7c3b7430fa14c95a8da518d2574f8cbb3.
//
// Solidity: event BurnWithAccountNumber(address _user, uint256 amount, string hashedAccountNumber)
func (_IDRX *IDRXFilterer) WatchBurnWithAccountNumber(opts *bind.WatchOpts, sink chan<- *IDRXBurnWithAccountNumber) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "BurnWithAccountNumber")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXBurnWithAccountNumber)
				if err := _IDRX.contract.UnpackLog(event, "BurnWithAccountNumber", log); err != nil {
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

// ParseBurnWithAccountNumber is a log parse operation binding the contract event 0x35626062961ef23c4eea38fafb3055f7c3b7430fa14c95a8da518d2574f8cbb3.
//
// Solidity: event BurnWithAccountNumber(address _user, uint256 amount, string hashedAccountNumber)
func (_IDRX *IDRXFilterer) ParseBurnWithAccountNumber(log types.Log) (*IDRXBurnWithAccountNumber, error) {
	event := new(IDRXBurnWithAccountNumber)
	if err := _IDRX.contract.UnpackLog(event, "BurnWithAccountNumber", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXDestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the IDRX contract.
type IDRXDestroyedBlackFundsIterator struct {
	Event *IDRXDestroyedBlackFunds // Event containing the contract specifics and raw log

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
func (it *IDRXDestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXDestroyedBlackFunds)
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
		it.Event = new(IDRXDestroyedBlackFunds)
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
func (it *IDRXDestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXDestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXDestroyedBlackFunds represents a DestroyedBlackFunds event raised by the IDRX contract.
type IDRXDestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_IDRX *IDRXFilterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*IDRXDestroyedBlackFundsIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &IDRXDestroyedBlackFundsIterator{contract: _IDRX.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_IDRX *IDRXFilterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *IDRXDestroyedBlackFunds) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXDestroyedBlackFunds)
				if err := _IDRX.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
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

// ParseDestroyedBlackFunds is a log parse operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_IDRX *IDRXFilterer) ParseDestroyedBlackFunds(log types.Log) (*IDRXDestroyedBlackFunds, error) {
	event := new(IDRXDestroyedBlackFunds)
	if err := _IDRX.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IDRX contract.
type IDRXInitializedIterator struct {
	Event *IDRXInitialized // Event containing the contract specifics and raw log

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
func (it *IDRXInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXInitialized)
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
		it.Event = new(IDRXInitialized)
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
func (it *IDRXInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXInitialized represents a Initialized event raised by the IDRX contract.
type IDRXInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IDRX *IDRXFilterer) FilterInitialized(opts *bind.FilterOpts) (*IDRXInitializedIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IDRXInitializedIterator{contract: _IDRX.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IDRX *IDRXFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IDRXInitialized) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXInitialized)
				if err := _IDRX.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IDRX *IDRXFilterer) ParseInitialized(log types.Log) (*IDRXInitialized, error) {
	event := new(IDRXInitialized)
	if err := _IDRX.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXMintBridgeIterator is returned from FilterMintBridge and is used to iterate over the raw logs and unpacked data for MintBridge events raised by the IDRX contract.
type IDRXMintBridgeIterator struct {
	Event *IDRXMintBridge // Event containing the contract specifics and raw log

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
func (it *IDRXMintBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXMintBridge)
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
		it.Event = new(IDRXMintBridge)
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
func (it *IDRXMintBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXMintBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXMintBridge represents a MintBridge event raised by the IDRX contract.
type IDRXMintBridge struct {
	User            common.Address
	Amount          *big.Int
	AmountAfterCut  *big.Int
	FromChain       *big.Int
	FromBridgeNonce *big.Int
	PlatformFee     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintBridge is a free log retrieval operation binding the contract event 0x410fe54a5fbff8c754e6adce20d0bd47194e3e36632e53e9c80aaa29d35184e8.
//
// Solidity: event MintBridge(address _user, uint256 _amount, uint256 amountAfterCut, uint256 fromChain, uint256 fromBridgeNonce, uint256 platformFee)
func (_IDRX *IDRXFilterer) FilterMintBridge(opts *bind.FilterOpts) (*IDRXMintBridgeIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "MintBridge")
	if err != nil {
		return nil, err
	}
	return &IDRXMintBridgeIterator{contract: _IDRX.contract, event: "MintBridge", logs: logs, sub: sub}, nil
}

// WatchMintBridge is a free log subscription operation binding the contract event 0x410fe54a5fbff8c754e6adce20d0bd47194e3e36632e53e9c80aaa29d35184e8.
//
// Solidity: event MintBridge(address _user, uint256 _amount, uint256 amountAfterCut, uint256 fromChain, uint256 fromBridgeNonce, uint256 platformFee)
func (_IDRX *IDRXFilterer) WatchMintBridge(opts *bind.WatchOpts, sink chan<- *IDRXMintBridge) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "MintBridge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXMintBridge)
				if err := _IDRX.contract.UnpackLog(event, "MintBridge", log); err != nil {
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

// ParseMintBridge is a log parse operation binding the contract event 0x410fe54a5fbff8c754e6adce20d0bd47194e3e36632e53e9c80aaa29d35184e8.
//
// Solidity: event MintBridge(address _user, uint256 _amount, uint256 amountAfterCut, uint256 fromChain, uint256 fromBridgeNonce, uint256 platformFee)
func (_IDRX *IDRXFilterer) ParseMintBridge(log types.Log) (*IDRXMintBridge, error) {
	event := new(IDRXMintBridge)
	if err := _IDRX.contract.UnpackLog(event, "MintBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IDRX contract.
type IDRXPausedIterator struct {
	Event *IDRXPaused // Event containing the contract specifics and raw log

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
func (it *IDRXPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXPaused)
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
		it.Event = new(IDRXPaused)
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
func (it *IDRXPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXPaused represents a Paused event raised by the IDRX contract.
type IDRXPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IDRX *IDRXFilterer) FilterPaused(opts *bind.FilterOpts) (*IDRXPausedIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IDRXPausedIterator{contract: _IDRX.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IDRX *IDRXFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IDRXPaused) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXPaused)
				if err := _IDRX.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_IDRX *IDRXFilterer) ParsePaused(log types.Log) (*IDRXPaused, error) {
	event := new(IDRXPaused)
	if err := _IDRX.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXPlatformFeeInfoUpdatedIterator is returned from FilterPlatformFeeInfoUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeInfoUpdated events raised by the IDRX contract.
type IDRXPlatformFeeInfoUpdatedIterator struct {
	Event *IDRXPlatformFeeInfoUpdated // Event containing the contract specifics and raw log

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
func (it *IDRXPlatformFeeInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXPlatformFeeInfoUpdated)
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
		it.Event = new(IDRXPlatformFeeInfoUpdated)
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
func (it *IDRXPlatformFeeInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXPlatformFeeInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXPlatformFeeInfoUpdated represents a PlatformFeeInfoUpdated event raised by the IDRX contract.
type IDRXPlatformFeeInfoUpdated struct {
	PlatformFeeRecipient common.Address
	BurnBridgeFee        uint64
	MintBridgeFee        uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeInfoUpdated is a free log retrieval operation binding the contract event 0x7803f8d90330912bc9f2d8d0ef0a2cef51897cc7cf4bdaef92ddd80dd36dc691.
//
// Solidity: event PlatformFeeInfoUpdated(address _platformFeeRecipient, uint64 _burnBridgeFee, uint64 _mintBridgeFee)
func (_IDRX *IDRXFilterer) FilterPlatformFeeInfoUpdated(opts *bind.FilterOpts) (*IDRXPlatformFeeInfoUpdatedIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "PlatformFeeInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &IDRXPlatformFeeInfoUpdatedIterator{contract: _IDRX.contract, event: "PlatformFeeInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeInfoUpdated is a free log subscription operation binding the contract event 0x7803f8d90330912bc9f2d8d0ef0a2cef51897cc7cf4bdaef92ddd80dd36dc691.
//
// Solidity: event PlatformFeeInfoUpdated(address _platformFeeRecipient, uint64 _burnBridgeFee, uint64 _mintBridgeFee)
func (_IDRX *IDRXFilterer) WatchPlatformFeeInfoUpdated(opts *bind.WatchOpts, sink chan<- *IDRXPlatformFeeInfoUpdated) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "PlatformFeeInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXPlatformFeeInfoUpdated)
				if err := _IDRX.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
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

// ParsePlatformFeeInfoUpdated is a log parse operation binding the contract event 0x7803f8d90330912bc9f2d8d0ef0a2cef51897cc7cf4bdaef92ddd80dd36dc691.
//
// Solidity: event PlatformFeeInfoUpdated(address _platformFeeRecipient, uint64 _burnBridgeFee, uint64 _mintBridgeFee)
func (_IDRX *IDRXFilterer) ParsePlatformFeeInfoUpdated(log types.Log) (*IDRXPlatformFeeInfoUpdated, error) {
	event := new(IDRXPlatformFeeInfoUpdated)
	if err := _IDRX.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the IDRX contract.
type IDRXRemovedBlackListIterator struct {
	Event *IDRXRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *IDRXRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXRemovedBlackList)
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
		it.Event = new(IDRXRemovedBlackList)
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
func (it *IDRXRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXRemovedBlackList represents a RemovedBlackList event raised by the IDRX contract.
type IDRXRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_IDRX *IDRXFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*IDRXRemovedBlackListIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &IDRXRemovedBlackListIterator{contract: _IDRX.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_IDRX *IDRXFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *IDRXRemovedBlackList) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXRemovedBlackList)
				if err := _IDRX.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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

// ParseRemovedBlackList is a log parse operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_IDRX *IDRXFilterer) ParseRemovedBlackList(log types.Log) (*IDRXRemovedBlackList, error) {
	event := new(IDRXRemovedBlackList)
	if err := _IDRX.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IDRX contract.
type IDRXRoleAdminChangedIterator struct {
	Event *IDRXRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IDRXRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXRoleAdminChanged)
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
		it.Event = new(IDRXRoleAdminChanged)
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
func (it *IDRXRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXRoleAdminChanged represents a RoleAdminChanged event raised by the IDRX contract.
type IDRXRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IDRX *IDRXFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IDRXRoleAdminChangedIterator, error) {
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

	logs, sub, err := _IDRX.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IDRXRoleAdminChangedIterator{contract: _IDRX.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IDRX *IDRXFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IDRXRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {
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

	logs, sub, err := _IDRX.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXRoleAdminChanged)
				if err := _IDRX.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_IDRX *IDRXFilterer) ParseRoleAdminChanged(log types.Log) (*IDRXRoleAdminChanged, error) {
	event := new(IDRXRoleAdminChanged)
	if err := _IDRX.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IDRX contract.
type IDRXRoleGrantedIterator struct {
	Event *IDRXRoleGranted // Event containing the contract specifics and raw log

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
func (it *IDRXRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXRoleGranted)
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
		it.Event = new(IDRXRoleGranted)
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
func (it *IDRXRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXRoleGranted represents a RoleGranted event raised by the IDRX contract.
type IDRXRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IDRX *IDRXFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IDRXRoleGrantedIterator, error) {
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

	logs, sub, err := _IDRX.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IDRXRoleGrantedIterator{contract: _IDRX.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IDRX *IDRXFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IDRXRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {
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

	logs, sub, err := _IDRX.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXRoleGranted)
				if err := _IDRX.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_IDRX *IDRXFilterer) ParseRoleGranted(log types.Log) (*IDRXRoleGranted, error) {
	event := new(IDRXRoleGranted)
	if err := _IDRX.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IDRX contract.
type IDRXRoleRevokedIterator struct {
	Event *IDRXRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IDRXRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXRoleRevoked)
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
		it.Event = new(IDRXRoleRevoked)
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
func (it *IDRXRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXRoleRevoked represents a RoleRevoked event raised by the IDRX contract.
type IDRXRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IDRX *IDRXFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IDRXRoleRevokedIterator, error) {
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

	logs, sub, err := _IDRX.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IDRXRoleRevokedIterator{contract: _IDRX.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IDRX *IDRXFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IDRXRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {
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

	logs, sub, err := _IDRX.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXRoleRevoked)
				if err := _IDRX.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_IDRX *IDRXFilterer) ParseRoleRevoked(log types.Log) (*IDRXRoleRevoked, error) {
	event := new(IDRXRoleRevoked)
	if err := _IDRX.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IDRX contract.
type IDRXTransferIterator struct {
	Event *IDRXTransfer // Event containing the contract specifics and raw log

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
func (it *IDRXTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXTransfer)
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
		it.Event = new(IDRXTransfer)
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
func (it *IDRXTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXTransfer represents a Transfer event raised by the IDRX contract.
type IDRXTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IDRX *IDRXFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IDRXTransferIterator, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDRX.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IDRXTransferIterator{contract: _IDRX.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IDRX *IDRXFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IDRXTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDRX.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXTransfer)
				if err := _IDRX.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IDRX *IDRXFilterer) ParseTransfer(log types.Log) (*IDRXTransfer, error) {
	event := new(IDRXTransfer)
	if err := _IDRX.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IDRX contract.
type IDRXUnpausedIterator struct {
	Event *IDRXUnpaused // Event containing the contract specifics and raw log

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
func (it *IDRXUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXUnpaused)
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
		it.Event = new(IDRXUnpaused)
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
func (it *IDRXUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXUnpaused represents a Unpaused event raised by the IDRX contract.
type IDRXUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IDRX *IDRXFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IDRXUnpausedIterator, error) {
	logs, sub, err := _IDRX.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IDRXUnpausedIterator{contract: _IDRX.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IDRX *IDRXFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IDRXUnpaused) (event.Subscription, error) {
	logs, sub, err := _IDRX.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXUnpaused)
				if err := _IDRX.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_IDRX *IDRXFilterer) ParseUnpaused(log types.Log) (*IDRXUnpaused, error) {
	event := new(IDRXUnpaused)
	if err := _IDRX.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDRXUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IDRX contract.
type IDRXUpgradedIterator struct {
	Event *IDRXUpgraded // Event containing the contract specifics and raw log

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
func (it *IDRXUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDRXUpgraded)
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
		it.Event = new(IDRXUpgraded)
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
func (it *IDRXUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDRXUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDRXUpgraded represents a Upgraded event raised by the IDRX contract.
type IDRXUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IDRX *IDRXFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*IDRXUpgradedIterator, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IDRX.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &IDRXUpgradedIterator{contract: _IDRX.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IDRX *IDRXFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IDRXUpgraded, implementation []common.Address) (event.Subscription, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IDRX.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDRXUpgraded)
				if err := _IDRX.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IDRX *IDRXFilterer) ParseUpgraded(log types.Log) (*IDRXUpgraded, error) {
	event := new(IDRXUpgraded)
	if err := _IDRX.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
