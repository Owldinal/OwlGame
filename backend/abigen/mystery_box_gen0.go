// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

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
)

// MysteryBoxGen0MetaData contains all meta data concerning the MysteryBoxGen0 contract.
var MysteryBoxGen0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"endedBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voyaTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"idOneOwnerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defenderAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boxId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"}],\"name\":\"MintBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boxId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"OpenBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_VOYA_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkCanMintByVoya\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkCanMintByWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkCanMintTokenOne\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintByVoya\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintByWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintTokenOne\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"openBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whiteListEndedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MysteryBoxGen0ABI is the input ABI used to generate the binding from.
// Deprecated: Use MysteryBoxGen0MetaData.ABI instead.
var MysteryBoxGen0ABI = MysteryBoxGen0MetaData.ABI

// MysteryBoxGen0 is an auto generated Go binding around an Ethereum contract.
type MysteryBoxGen0 struct {
	MysteryBoxGen0Caller     // Read-only binding to the contract
	MysteryBoxGen0Transactor // Write-only binding to the contract
	MysteryBoxGen0Filterer   // Log filterer for contract events
}

// MysteryBoxGen0Caller is an auto generated read-only Go binding around an Ethereum contract.
type MysteryBoxGen0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MysteryBoxGen0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MysteryBoxGen0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MysteryBoxGen0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MysteryBoxGen0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MysteryBoxGen0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MysteryBoxGen0Session struct {
	Contract     *MysteryBoxGen0   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MysteryBoxGen0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MysteryBoxGen0CallerSession struct {
	Contract *MysteryBoxGen0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MysteryBoxGen0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MysteryBoxGen0TransactorSession struct {
	Contract     *MysteryBoxGen0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MysteryBoxGen0Raw is an auto generated low-level Go binding around an Ethereum contract.
type MysteryBoxGen0Raw struct {
	Contract *MysteryBoxGen0 // Generic contract binding to access the raw methods on
}

// MysteryBoxGen0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MysteryBoxGen0CallerRaw struct {
	Contract *MysteryBoxGen0Caller // Generic read-only contract binding to access the raw methods on
}

// MysteryBoxGen0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MysteryBoxGen0TransactorRaw struct {
	Contract *MysteryBoxGen0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMysteryBoxGen0 creates a new instance of MysteryBoxGen0, bound to a specific deployed contract.
func NewMysteryBoxGen0(address common.Address, backend bind.ContractBackend) (*MysteryBoxGen0, error) {
	contract, err := bindMysteryBoxGen0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0{MysteryBoxGen0Caller: MysteryBoxGen0Caller{contract: contract}, MysteryBoxGen0Transactor: MysteryBoxGen0Transactor{contract: contract}, MysteryBoxGen0Filterer: MysteryBoxGen0Filterer{contract: contract}}, nil
}

// NewMysteryBoxGen0Caller creates a new read-only instance of MysteryBoxGen0, bound to a specific deployed contract.
func NewMysteryBoxGen0Caller(address common.Address, caller bind.ContractCaller) (*MysteryBoxGen0Caller, error) {
	contract, err := bindMysteryBoxGen0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0Caller{contract: contract}, nil
}

// NewMysteryBoxGen0Transactor creates a new write-only instance of MysteryBoxGen0, bound to a specific deployed contract.
func NewMysteryBoxGen0Transactor(address common.Address, transactor bind.ContractTransactor) (*MysteryBoxGen0Transactor, error) {
	contract, err := bindMysteryBoxGen0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0Transactor{contract: contract}, nil
}

// NewMysteryBoxGen0Filterer creates a new log filterer instance of MysteryBoxGen0, bound to a specific deployed contract.
func NewMysteryBoxGen0Filterer(address common.Address, filterer bind.ContractFilterer) (*MysteryBoxGen0Filterer, error) {
	contract, err := bindMysteryBoxGen0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0Filterer{contract: contract}, nil
}

// bindMysteryBoxGen0 binds a generic wrapper to an already deployed contract.
func bindMysteryBoxGen0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MysteryBoxGen0ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MysteryBoxGen0 *MysteryBoxGen0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MysteryBoxGen0.Contract.MysteryBoxGen0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MysteryBoxGen0 *MysteryBoxGen0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.MysteryBoxGen0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MysteryBoxGen0 *MysteryBoxGen0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.MysteryBoxGen0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MysteryBoxGen0 *MysteryBoxGen0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MysteryBoxGen0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _MysteryBoxGen0.Contract.DEFAULTADMINROLE(&_MysteryBoxGen0.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MysteryBoxGen0.Contract.DEFAULTADMINROLE(&_MysteryBoxGen0.CallOpts)
}

// MINTVOYATHRESHOLD is a free data retrieval call binding the contract method 0x40ba20bc.
//
// Solidity: function MINT_VOYA_THRESHOLD() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) MINTVOYATHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "MINT_VOYA_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTVOYATHRESHOLD is a free data retrieval call binding the contract method 0x40ba20bc.
//
// Solidity: function MINT_VOYA_THRESHOLD() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) MINTVOYATHRESHOLD() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.MINTVOYATHRESHOLD(&_MysteryBoxGen0.CallOpts)
}

// MINTVOYATHRESHOLD is a free data retrieval call binding the contract method 0x40ba20bc.
//
// Solidity: function MINT_VOYA_THRESHOLD() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) MINTVOYATHRESHOLD() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.MINTVOYATHRESHOLD(&_MysteryBoxGen0.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MysteryBoxGen0.Contract.BalanceOf(&_MysteryBoxGen0.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MysteryBoxGen0.Contract.BalanceOf(&_MysteryBoxGen0.CallOpts, owner)
}

// CheckCanMintByVoya is a free data retrieval call binding the contract method 0x5c05b3c2.
//
// Solidity: function checkCanMintByVoya() view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) CheckCanMintByVoya(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "checkCanMintByVoya")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCanMintByVoya is a free data retrieval call binding the contract method 0x5c05b3c2.
//
// Solidity: function checkCanMintByVoya() view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) CheckCanMintByVoya() (bool, error) {
	return _MysteryBoxGen0.Contract.CheckCanMintByVoya(&_MysteryBoxGen0.CallOpts)
}

// CheckCanMintByVoya is a free data retrieval call binding the contract method 0x5c05b3c2.
//
// Solidity: function checkCanMintByVoya() view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) CheckCanMintByVoya() (bool, error) {
	return _MysteryBoxGen0.Contract.CheckCanMintByVoya(&_MysteryBoxGen0.CallOpts)
}

// CheckCanMintByWhiteList is a free data retrieval call binding the contract method 0x2af2f1e6.
//
// Solidity: function checkCanMintByWhiteList() view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) CheckCanMintByWhiteList(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "checkCanMintByWhiteList")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCanMintByWhiteList is a free data retrieval call binding the contract method 0x2af2f1e6.
//
// Solidity: function checkCanMintByWhiteList() view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) CheckCanMintByWhiteList() (bool, error) {
	return _MysteryBoxGen0.Contract.CheckCanMintByWhiteList(&_MysteryBoxGen0.CallOpts)
}

// CheckCanMintByWhiteList is a free data retrieval call binding the contract method 0x2af2f1e6.
//
// Solidity: function checkCanMintByWhiteList() view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) CheckCanMintByWhiteList() (bool, error) {
	return _MysteryBoxGen0.Contract.CheckCanMintByWhiteList(&_MysteryBoxGen0.CallOpts)
}

// CheckCanMintTokenOne is a free data retrieval call binding the contract method 0x7522cf9d.
//
// Solidity: function checkCanMintTokenOne() view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) CheckCanMintTokenOne(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "checkCanMintTokenOne")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCanMintTokenOne is a free data retrieval call binding the contract method 0x7522cf9d.
//
// Solidity: function checkCanMintTokenOne() view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) CheckCanMintTokenOne() (bool, error) {
	return _MysteryBoxGen0.Contract.CheckCanMintTokenOne(&_MysteryBoxGen0.CallOpts)
}

// CheckCanMintTokenOne is a free data retrieval call binding the contract method 0x7522cf9d.
//
// Solidity: function checkCanMintTokenOne() view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) CheckCanMintTokenOne() (bool, error) {
	return _MysteryBoxGen0.Contract.CheckCanMintTokenOne(&_MysteryBoxGen0.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) CurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "currentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) CurrentTokenId() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.CurrentTokenId(&_MysteryBoxGen0.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) CurrentTokenId() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.CurrentTokenId(&_MysteryBoxGen0.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MysteryBoxGen0.Contract.GetApproved(&_MysteryBoxGen0.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MysteryBoxGen0.Contract.GetApproved(&_MysteryBoxGen0.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MysteryBoxGen0.Contract.GetRoleAdmin(&_MysteryBoxGen0.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MysteryBoxGen0.Contract.GetRoleAdmin(&_MysteryBoxGen0.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MysteryBoxGen0.Contract.HasRole(&_MysteryBoxGen0.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MysteryBoxGen0.Contract.HasRole(&_MysteryBoxGen0.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MysteryBoxGen0.Contract.IsApprovedForAll(&_MysteryBoxGen0.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MysteryBoxGen0.Contract.IsApprovedForAll(&_MysteryBoxGen0.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) Name() (string, error) {
	return _MysteryBoxGen0.Contract.Name(&_MysteryBoxGen0.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) Name() (string, error) {
	return _MysteryBoxGen0.Contract.Name(&_MysteryBoxGen0.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MysteryBoxGen0.Contract.OwnerOf(&_MysteryBoxGen0.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MysteryBoxGen0.Contract.OwnerOf(&_MysteryBoxGen0.CallOpts, tokenId)
}

// RemainCount is a free data retrieval call binding the contract method 0xe06f0806.
//
// Solidity: function remainCount() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) RemainCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "remainCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainCount is a free data retrieval call binding the contract method 0xe06f0806.
//
// Solidity: function remainCount() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) RemainCount() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.RemainCount(&_MysteryBoxGen0.CallOpts)
}

// RemainCount is a free data retrieval call binding the contract method 0xe06f0806.
//
// Solidity: function remainCount() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) RemainCount() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.RemainCount(&_MysteryBoxGen0.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MysteryBoxGen0.Contract.SupportsInterface(&_MysteryBoxGen0.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MysteryBoxGen0.Contract.SupportsInterface(&_MysteryBoxGen0.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) Symbol() (string, error) {
	return _MysteryBoxGen0.Contract.Symbol(&_MysteryBoxGen0.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) Symbol() (string, error) {
	return _MysteryBoxGen0.Contract.Symbol(&_MysteryBoxGen0.CallOpts)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) TokenIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "tokenIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) TokenIdCounter() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.TokenIdCounter(&_MysteryBoxGen0.CallOpts)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) TokenIdCounter() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.TokenIdCounter(&_MysteryBoxGen0.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) TokenURI(tokenId *big.Int) (string, error) {
	return _MysteryBoxGen0.Contract.TokenURI(&_MysteryBoxGen0.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MysteryBoxGen0.Contract.TokenURI(&_MysteryBoxGen0.CallOpts, tokenId)
}

// WhiteListEndedBlock is a free data retrieval call binding the contract method 0xd7379ba1.
//
// Solidity: function whiteListEndedBlock() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Caller) WhiteListEndedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MysteryBoxGen0.contract.Call(opts, &out, "whiteListEndedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhiteListEndedBlock is a free data retrieval call binding the contract method 0xd7379ba1.
//
// Solidity: function whiteListEndedBlock() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0Session) WhiteListEndedBlock() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.WhiteListEndedBlock(&_MysteryBoxGen0.CallOpts)
}

// WhiteListEndedBlock is a free data retrieval call binding the contract method 0xd7379ba1.
//
// Solidity: function whiteListEndedBlock() view returns(uint256)
func (_MysteryBoxGen0 *MysteryBoxGen0CallerSession) WhiteListEndedBlock() (*big.Int, error) {
	return _MysteryBoxGen0.Contract.WhiteListEndedBlock(&_MysteryBoxGen0.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.Approve(&_MysteryBoxGen0.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.Approve(&_MysteryBoxGen0.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.GrantRole(&_MysteryBoxGen0.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.GrantRole(&_MysteryBoxGen0.TransactOpts, role, account)
}

// MintByVoya is a paid mutator transaction binding the contract method 0xd290dd6c.
//
// Solidity: function mintByVoya(bytes32 hash, bytes signature) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) MintByVoya(opts *bind.TransactOpts, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "mintByVoya", hash, signature)
}

// MintByVoya is a paid mutator transaction binding the contract method 0xd290dd6c.
//
// Solidity: function mintByVoya(bytes32 hash, bytes signature) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) MintByVoya(hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.MintByVoya(&_MysteryBoxGen0.TransactOpts, hash, signature)
}

// MintByVoya is a paid mutator transaction binding the contract method 0xd290dd6c.
//
// Solidity: function mintByVoya(bytes32 hash, bytes signature) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) MintByVoya(hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.MintByVoya(&_MysteryBoxGen0.TransactOpts, hash, signature)
}

// MintByWhiteList is a paid mutator transaction binding the contract method 0x16744e1c.
//
// Solidity: function mintByWhiteList(bytes32 hash, bytes signature) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) MintByWhiteList(opts *bind.TransactOpts, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "mintByWhiteList", hash, signature)
}

// MintByWhiteList is a paid mutator transaction binding the contract method 0x16744e1c.
//
// Solidity: function mintByWhiteList(bytes32 hash, bytes signature) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) MintByWhiteList(hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.MintByWhiteList(&_MysteryBoxGen0.TransactOpts, hash, signature)
}

// MintByWhiteList is a paid mutator transaction binding the contract method 0x16744e1c.
//
// Solidity: function mintByWhiteList(bytes32 hash, bytes signature) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) MintByWhiteList(hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.MintByWhiteList(&_MysteryBoxGen0.TransactOpts, hash, signature)
}

// MintTokenOne is a paid mutator transaction binding the contract method 0xd4862493.
//
// Solidity: function mintTokenOne(bytes32 hash, bytes signature) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) MintTokenOne(opts *bind.TransactOpts, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "mintTokenOne", hash, signature)
}

// MintTokenOne is a paid mutator transaction binding the contract method 0xd4862493.
//
// Solidity: function mintTokenOne(bytes32 hash, bytes signature) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) MintTokenOne(hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.MintTokenOne(&_MysteryBoxGen0.TransactOpts, hash, signature)
}

// MintTokenOne is a paid mutator transaction binding the contract method 0xd4862493.
//
// Solidity: function mintTokenOne(bytes32 hash, bytes signature) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) MintTokenOne(hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.MintTokenOne(&_MysteryBoxGen0.TransactOpts, hash, signature)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) OpenBox(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "openBox", tokenId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) OpenBox(tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.OpenBox(&_MysteryBoxGen0.TransactOpts, tokenId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) OpenBox(tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.OpenBox(&_MysteryBoxGen0.TransactOpts, tokenId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.RenounceRole(&_MysteryBoxGen0.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.RenounceRole(&_MysteryBoxGen0.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.RevokeRole(&_MysteryBoxGen0.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.RevokeRole(&_MysteryBoxGen0.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.SafeTransferFrom(&_MysteryBoxGen0.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.SafeTransferFrom(&_MysteryBoxGen0.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.SafeTransferFrom0(&_MysteryBoxGen0.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.SafeTransferFrom0(&_MysteryBoxGen0.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.SetApprovalForAll(&_MysteryBoxGen0.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.SetApprovalForAll(&_MysteryBoxGen0.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.TransferFrom(&_MysteryBoxGen0.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen0 *MysteryBoxGen0TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen0.Contract.TransferFrom(&_MysteryBoxGen0.TransactOpts, from, to, tokenId)
}

// MysteryBoxGen0ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0ApprovalIterator struct {
	Event *MysteryBoxGen0Approval // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0Approval)
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
		it.Event = new(MysteryBoxGen0Approval)
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
func (it *MysteryBoxGen0ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0Approval represents a Approval event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MysteryBoxGen0ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0ApprovalIterator{contract: _MysteryBoxGen0.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0Approval)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseApproval(log types.Log) (*MysteryBoxGen0Approval, error) {
	event := new(MysteryBoxGen0Approval)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen0ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0ApprovalForAllIterator struct {
	Event *MysteryBoxGen0ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0ApprovalForAll)
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
		it.Event = new(MysteryBoxGen0ApprovalForAll)
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
func (it *MysteryBoxGen0ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0ApprovalForAll represents a ApprovalForAll event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MysteryBoxGen0ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0ApprovalForAllIterator{contract: _MysteryBoxGen0.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0ApprovalForAll)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseApprovalForAll(log types.Log) (*MysteryBoxGen0ApprovalForAll, error) {
	event := new(MysteryBoxGen0ApprovalForAll)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen0BatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0BatchMetadataUpdateIterator struct {
	Event *MysteryBoxGen0BatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0BatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0BatchMetadataUpdate)
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
		it.Event = new(MysteryBoxGen0BatchMetadataUpdate)
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
func (it *MysteryBoxGen0BatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0BatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0BatchMetadataUpdate represents a BatchMetadataUpdate event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0BatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*MysteryBoxGen0BatchMetadataUpdateIterator, error) {

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0BatchMetadataUpdateIterator{contract: _MysteryBoxGen0.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0BatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0BatchMetadataUpdate)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseBatchMetadataUpdate(log types.Log) (*MysteryBoxGen0BatchMetadataUpdate, error) {
	event := new(MysteryBoxGen0BatchMetadataUpdate)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen0MetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0MetadataUpdateIterator struct {
	Event *MysteryBoxGen0MetadataUpdate // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0MetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0MetadataUpdate)
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
		it.Event = new(MysteryBoxGen0MetadataUpdate)
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
func (it *MysteryBoxGen0MetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0MetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0MetadataUpdate represents a MetadataUpdate event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0MetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*MysteryBoxGen0MetadataUpdateIterator, error) {

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0MetadataUpdateIterator{contract: _MysteryBoxGen0.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0MetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0MetadataUpdate)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseMetadataUpdate(log types.Log) (*MysteryBoxGen0MetadataUpdate, error) {
	event := new(MysteryBoxGen0MetadataUpdate)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen0MintBoxIterator is returned from FilterMintBox and is used to iterate over the raw logs and unpacked data for MintBox events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0MintBoxIterator struct {
	Event *MysteryBoxGen0MintBox // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0MintBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0MintBox)
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
		it.Event = new(MysteryBoxGen0MintBox)
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
func (it *MysteryBoxGen0MintBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0MintBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0MintBox represents a MintBox event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0MintBox struct {
	User     common.Address
	BoxId    *big.Int
	MintType *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintBox is a free log retrieval operation binding the contract event 0xd87153f724a1d6cd9dffd85366cd07948f87bc696b485edd446ab72c68c9a66d.
//
// Solidity: event MintBox(address indexed user, uint256 boxId, uint256 mintType)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterMintBox(opts *bind.FilterOpts, user []common.Address) (*MysteryBoxGen0MintBoxIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "MintBox", userRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0MintBoxIterator{contract: _MysteryBoxGen0.contract, event: "MintBox", logs: logs, sub: sub}, nil
}

// WatchMintBox is a free log subscription operation binding the contract event 0xd87153f724a1d6cd9dffd85366cd07948f87bc696b485edd446ab72c68c9a66d.
//
// Solidity: event MintBox(address indexed user, uint256 boxId, uint256 mintType)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchMintBox(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0MintBox, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "MintBox", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0MintBox)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "MintBox", log); err != nil {
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

// ParseMintBox is a log parse operation binding the contract event 0xd87153f724a1d6cd9dffd85366cd07948f87bc696b485edd446ab72c68c9a66d.
//
// Solidity: event MintBox(address indexed user, uint256 boxId, uint256 mintType)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseMintBox(log types.Log) (*MysteryBoxGen0MintBox, error) {
	event := new(MysteryBoxGen0MintBox)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "MintBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen0OpenBoxIterator is returned from FilterOpenBox and is used to iterate over the raw logs and unpacked data for OpenBox events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0OpenBoxIterator struct {
	Event *MysteryBoxGen0OpenBox // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0OpenBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0OpenBox)
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
		it.Event = new(MysteryBoxGen0OpenBox)
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
func (it *MysteryBoxGen0OpenBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0OpenBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0OpenBox represents a OpenBox event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0OpenBox struct {
	User  common.Address
	BoxId *big.Int
	Url   string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOpenBox is a free log retrieval operation binding the contract event 0xadcb2434e8f45d0b9bc8812cfc56fea4f6c1d4e2a7bc0526a9f1ae337901c3a3.
//
// Solidity: event OpenBox(address indexed user, uint256 boxId, string url)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterOpenBox(opts *bind.FilterOpts, user []common.Address) (*MysteryBoxGen0OpenBoxIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "OpenBox", userRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0OpenBoxIterator{contract: _MysteryBoxGen0.contract, event: "OpenBox", logs: logs, sub: sub}, nil
}

// WatchOpenBox is a free log subscription operation binding the contract event 0xadcb2434e8f45d0b9bc8812cfc56fea4f6c1d4e2a7bc0526a9f1ae337901c3a3.
//
// Solidity: event OpenBox(address indexed user, uint256 boxId, string url)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchOpenBox(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0OpenBox, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "OpenBox", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0OpenBox)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "OpenBox", log); err != nil {
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

// ParseOpenBox is a log parse operation binding the contract event 0xadcb2434e8f45d0b9bc8812cfc56fea4f6c1d4e2a7bc0526a9f1ae337901c3a3.
//
// Solidity: event OpenBox(address indexed user, uint256 boxId, string url)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseOpenBox(log types.Log) (*MysteryBoxGen0OpenBox, error) {
	event := new(MysteryBoxGen0OpenBox)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "OpenBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen0RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0RoleAdminChangedIterator struct {
	Event *MysteryBoxGen0RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0RoleAdminChanged)
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
		it.Event = new(MysteryBoxGen0RoleAdminChanged)
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
func (it *MysteryBoxGen0RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0RoleAdminChanged represents a RoleAdminChanged event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MysteryBoxGen0RoleAdminChangedIterator, error) {

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

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0RoleAdminChangedIterator{contract: _MysteryBoxGen0.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0RoleAdminChanged)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseRoleAdminChanged(log types.Log) (*MysteryBoxGen0RoleAdminChanged, error) {
	event := new(MysteryBoxGen0RoleAdminChanged)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen0RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0RoleGrantedIterator struct {
	Event *MysteryBoxGen0RoleGranted // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0RoleGranted)
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
		it.Event = new(MysteryBoxGen0RoleGranted)
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
func (it *MysteryBoxGen0RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0RoleGranted represents a RoleGranted event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MysteryBoxGen0RoleGrantedIterator, error) {

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

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0RoleGrantedIterator{contract: _MysteryBoxGen0.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0RoleGranted)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseRoleGranted(log types.Log) (*MysteryBoxGen0RoleGranted, error) {
	event := new(MysteryBoxGen0RoleGranted)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen0RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0RoleRevokedIterator struct {
	Event *MysteryBoxGen0RoleRevoked // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0RoleRevoked)
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
		it.Event = new(MysteryBoxGen0RoleRevoked)
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
func (it *MysteryBoxGen0RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0RoleRevoked represents a RoleRevoked event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MysteryBoxGen0RoleRevokedIterator, error) {

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

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0RoleRevokedIterator{contract: _MysteryBoxGen0.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0RoleRevoked)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseRoleRevoked(log types.Log) (*MysteryBoxGen0RoleRevoked, error) {
	event := new(MysteryBoxGen0RoleRevoked)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen0TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0TransferIterator struct {
	Event *MysteryBoxGen0Transfer // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen0TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen0Transfer)
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
		it.Event = new(MysteryBoxGen0Transfer)
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
func (it *MysteryBoxGen0TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen0TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen0Transfer represents a Transfer event raised by the MysteryBoxGen0 contract.
type MysteryBoxGen0Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MysteryBoxGen0TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen0TransferIterator{contract: _MysteryBoxGen0.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen0Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MysteryBoxGen0.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen0Transfer)
				if err := _MysteryBoxGen0.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MysteryBoxGen0 *MysteryBoxGen0Filterer) ParseTransfer(log types.Log) (*MysteryBoxGen0Transfer, error) {
	event := new(MysteryBoxGen0Transfer)
	if err := _MysteryBoxGen0.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
