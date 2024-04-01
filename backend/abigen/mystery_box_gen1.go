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

// MysteryBoxGen1MetaData contains all meta data concerning the MysteryBoxGen1 contract.
var MysteryBoxGen1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC20Burnable\",\"name\":\"_owlToken\",\"type\":\"address\"},{\"internalType\":\"contractOwlGame\",\"name\":\"_gameContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boxId\",\"type\":\"uint256\"}],\"name\":\"MintBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumBoxType\",\"name\":\"boxType\",\"type\":\"uint8\"}],\"name\":\"OpenBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAME_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boxTypes\",\"outputs\":[{\"internalType\":\"enumBoxType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameContract\",\"outputs\":[{\"internalType\":\"contractOwlGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBoxType\",\"outputs\":[{\"internalType\":\"enumBoxType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isBoxOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"mintBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasBuff\",\"type\":\"bool\"}],\"name\":\"openBox\",\"outputs\":[{\"internalType\":\"enumBoxType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owlToken\",\"outputs\":[{\"internalType\":\"contractERC20Burnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MysteryBoxGen1ABI is the input ABI used to generate the binding from.
// Deprecated: Use MysteryBoxGen1MetaData.ABI instead.
var MysteryBoxGen1ABI = MysteryBoxGen1MetaData.ABI

// MysteryBoxGen1 is an auto generated Go binding around an Ethereum contract.
type MysteryBoxGen1 struct {
	MysteryBoxGen1Caller     // Read-only binding to the contract
	MysteryBoxGen1Transactor // Write-only binding to the contract
	MysteryBoxGen1Filterer   // Log filterer for contract events
}

// MysteryBoxGen1Caller is an auto generated read-only Go binding around an Ethereum contract.
type MysteryBoxGen1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MysteryBoxGen1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MysteryBoxGen1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MysteryBoxGen1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MysteryBoxGen1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MysteryBoxGen1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MysteryBoxGen1Session struct {
	Contract     *MysteryBoxGen1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MysteryBoxGen1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MysteryBoxGen1CallerSession struct {
	Contract *MysteryBoxGen1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MysteryBoxGen1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MysteryBoxGen1TransactorSession struct {
	Contract     *MysteryBoxGen1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MysteryBoxGen1Raw is an auto generated low-level Go binding around an Ethereum contract.
type MysteryBoxGen1Raw struct {
	Contract *MysteryBoxGen1 // Generic contract binding to access the raw methods on
}

// MysteryBoxGen1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MysteryBoxGen1CallerRaw struct {
	Contract *MysteryBoxGen1Caller // Generic read-only contract binding to access the raw methods on
}

// MysteryBoxGen1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MysteryBoxGen1TransactorRaw struct {
	Contract *MysteryBoxGen1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMysteryBoxGen1 creates a new instance of MysteryBoxGen1, bound to a specific deployed contract.
func NewMysteryBoxGen1(address common.Address, backend bind.ContractBackend) (*MysteryBoxGen1, error) {
	contract, err := bindMysteryBoxGen1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1{MysteryBoxGen1Caller: MysteryBoxGen1Caller{contract: contract}, MysteryBoxGen1Transactor: MysteryBoxGen1Transactor{contract: contract}, MysteryBoxGen1Filterer: MysteryBoxGen1Filterer{contract: contract}}, nil
}

// NewMysteryBoxGen1Caller creates a new read-only instance of MysteryBoxGen1, bound to a specific deployed contract.
func NewMysteryBoxGen1Caller(address common.Address, caller bind.ContractCaller) (*MysteryBoxGen1Caller, error) {
	contract, err := bindMysteryBoxGen1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1Caller{contract: contract}, nil
}

// NewMysteryBoxGen1Transactor creates a new write-only instance of MysteryBoxGen1, bound to a specific deployed contract.
func NewMysteryBoxGen1Transactor(address common.Address, transactor bind.ContractTransactor) (*MysteryBoxGen1Transactor, error) {
	contract, err := bindMysteryBoxGen1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1Transactor{contract: contract}, nil
}

// NewMysteryBoxGen1Filterer creates a new log filterer instance of MysteryBoxGen1, bound to a specific deployed contract.
func NewMysteryBoxGen1Filterer(address common.Address, filterer bind.ContractFilterer) (*MysteryBoxGen1Filterer, error) {
	contract, err := bindMysteryBoxGen1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1Filterer{contract: contract}, nil
}

// bindMysteryBoxGen1 binds a generic wrapper to an already deployed contract.
func bindMysteryBoxGen1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MysteryBoxGen1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MysteryBoxGen1 *MysteryBoxGen1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MysteryBoxGen1.Contract.MysteryBoxGen1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MysteryBoxGen1 *MysteryBoxGen1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.MysteryBoxGen1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MysteryBoxGen1 *MysteryBoxGen1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.MysteryBoxGen1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MysteryBoxGen1 *MysteryBoxGen1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MysteryBoxGen1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _MysteryBoxGen1.Contract.DEFAULTADMINROLE(&_MysteryBoxGen1.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MysteryBoxGen1.Contract.DEFAULTADMINROLE(&_MysteryBoxGen1.CallOpts)
}

// GAMECONTRACTROLE is a free data retrieval call binding the contract method 0xb5e352da.
//
// Solidity: function GAME_CONTRACT_ROLE() view returns(bytes32)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) GAMECONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "GAME_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GAMECONTRACTROLE is a free data retrieval call binding the contract method 0xb5e352da.
//
// Solidity: function GAME_CONTRACT_ROLE() view returns(bytes32)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) GAMECONTRACTROLE() ([32]byte, error) {
	return _MysteryBoxGen1.Contract.GAMECONTRACTROLE(&_MysteryBoxGen1.CallOpts)
}

// GAMECONTRACTROLE is a free data retrieval call binding the contract method 0xb5e352da.
//
// Solidity: function GAME_CONTRACT_ROLE() view returns(bytes32)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) GAMECONTRACTROLE() ([32]byte, error) {
	return _MysteryBoxGen1.Contract.GAMECONTRACTROLE(&_MysteryBoxGen1.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MysteryBoxGen1.Contract.BalanceOf(&_MysteryBoxGen1.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MysteryBoxGen1.Contract.BalanceOf(&_MysteryBoxGen1.CallOpts, owner)
}

// BoxTypes is a free data retrieval call binding the contract method 0x3a931a1d.
//
// Solidity: function boxTypes(uint256 ) view returns(uint8)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) BoxTypes(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "boxTypes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BoxTypes is a free data retrieval call binding the contract method 0x3a931a1d.
//
// Solidity: function boxTypes(uint256 ) view returns(uint8)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) BoxTypes(arg0 *big.Int) (uint8, error) {
	return _MysteryBoxGen1.Contract.BoxTypes(&_MysteryBoxGen1.CallOpts, arg0)
}

// BoxTypes is a free data retrieval call binding the contract method 0x3a931a1d.
//
// Solidity: function boxTypes(uint256 ) view returns(uint8)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) BoxTypes(arg0 *big.Int) (uint8, error) {
	return _MysteryBoxGen1.Contract.BoxTypes(&_MysteryBoxGen1.CallOpts, arg0)
}

// GameContract is a free data retrieval call binding the contract method 0xd3f33009.
//
// Solidity: function gameContract() view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) GameContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "gameContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameContract is a free data retrieval call binding the contract method 0xd3f33009.
//
// Solidity: function gameContract() view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) GameContract() (common.Address, error) {
	return _MysteryBoxGen1.Contract.GameContract(&_MysteryBoxGen1.CallOpts)
}

// GameContract is a free data retrieval call binding the contract method 0xd3f33009.
//
// Solidity: function gameContract() view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) GameContract() (common.Address, error) {
	return _MysteryBoxGen1.Contract.GameContract(&_MysteryBoxGen1.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MysteryBoxGen1.Contract.GetApproved(&_MysteryBoxGen1.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MysteryBoxGen1.Contract.GetApproved(&_MysteryBoxGen1.CallOpts, tokenId)
}

// GetBoxType is a free data retrieval call binding the contract method 0xb546a3bc.
//
// Solidity: function getBoxType(uint256 tokenId) view returns(uint8)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) GetBoxType(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "getBoxType", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetBoxType is a free data retrieval call binding the contract method 0xb546a3bc.
//
// Solidity: function getBoxType(uint256 tokenId) view returns(uint8)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) GetBoxType(tokenId *big.Int) (uint8, error) {
	return _MysteryBoxGen1.Contract.GetBoxType(&_MysteryBoxGen1.CallOpts, tokenId)
}

// GetBoxType is a free data retrieval call binding the contract method 0xb546a3bc.
//
// Solidity: function getBoxType(uint256 tokenId) view returns(uint8)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) GetBoxType(tokenId *big.Int) (uint8, error) {
	return _MysteryBoxGen1.Contract.GetBoxType(&_MysteryBoxGen1.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MysteryBoxGen1.Contract.GetRoleAdmin(&_MysteryBoxGen1.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MysteryBoxGen1.Contract.GetRoleAdmin(&_MysteryBoxGen1.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MysteryBoxGen1.Contract.HasRole(&_MysteryBoxGen1.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MysteryBoxGen1.Contract.HasRole(&_MysteryBoxGen1.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MysteryBoxGen1.Contract.IsApprovedForAll(&_MysteryBoxGen1.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MysteryBoxGen1.Contract.IsApprovedForAll(&_MysteryBoxGen1.CallOpts, owner, operator)
}

// IsBoxOpened is a free data retrieval call binding the contract method 0xa0520689.
//
// Solidity: function isBoxOpened(uint256 tokenId) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) IsBoxOpened(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "isBoxOpened", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBoxOpened is a free data retrieval call binding the contract method 0xa0520689.
//
// Solidity: function isBoxOpened(uint256 tokenId) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) IsBoxOpened(tokenId *big.Int) (bool, error) {
	return _MysteryBoxGen1.Contract.IsBoxOpened(&_MysteryBoxGen1.CallOpts, tokenId)
}

// IsBoxOpened is a free data retrieval call binding the contract method 0xa0520689.
//
// Solidity: function isBoxOpened(uint256 tokenId) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) IsBoxOpened(tokenId *big.Int) (bool, error) {
	return _MysteryBoxGen1.Contract.IsBoxOpened(&_MysteryBoxGen1.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) Name() (string, error) {
	return _MysteryBoxGen1.Contract.Name(&_MysteryBoxGen1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) Name() (string, error) {
	return _MysteryBoxGen1.Contract.Name(&_MysteryBoxGen1.CallOpts)
}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) OwlToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "owlToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) OwlToken() (common.Address, error) {
	return _MysteryBoxGen1.Contract.OwlToken(&_MysteryBoxGen1.CallOpts)
}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) OwlToken() (common.Address, error) {
	return _MysteryBoxGen1.Contract.OwlToken(&_MysteryBoxGen1.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MysteryBoxGen1.Contract.OwnerOf(&_MysteryBoxGen1.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MysteryBoxGen1.Contract.OwnerOf(&_MysteryBoxGen1.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MysteryBoxGen1.Contract.SupportsInterface(&_MysteryBoxGen1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MysteryBoxGen1.Contract.SupportsInterface(&_MysteryBoxGen1.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) Symbol() (string, error) {
	return _MysteryBoxGen1.Contract.Symbol(&_MysteryBoxGen1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) Symbol() (string, error) {
	return _MysteryBoxGen1.Contract.Symbol(&_MysteryBoxGen1.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MysteryBoxGen1 *MysteryBoxGen1Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MysteryBoxGen1.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) TokenURI(tokenId *big.Int) (string, error) {
	return _MysteryBoxGen1.Contract.TokenURI(&_MysteryBoxGen1.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MysteryBoxGen1 *MysteryBoxGen1CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MysteryBoxGen1.Contract.TokenURI(&_MysteryBoxGen1.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.Approve(&_MysteryBoxGen1.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.Approve(&_MysteryBoxGen1.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.GrantRole(&_MysteryBoxGen1.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.GrantRole(&_MysteryBoxGen1.TransactOpts, role, account)
}

// MintBox is a paid mutator transaction binding the contract method 0x71d02c9c.
//
// Solidity: function mintBox(address owner) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) MintBox(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "mintBox", owner)
}

// MintBox is a paid mutator transaction binding the contract method 0x71d02c9c.
//
// Solidity: function mintBox(address owner) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Session) MintBox(owner common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.MintBox(&_MysteryBoxGen1.TransactOpts, owner)
}

// MintBox is a paid mutator transaction binding the contract method 0x71d02c9c.
//
// Solidity: function mintBox(address owner) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) MintBox(owner common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.MintBox(&_MysteryBoxGen1.TransactOpts, owner)
}

// OpenBox is a paid mutator transaction binding the contract method 0xcc1c6db1.
//
// Solidity: function openBox(uint256 tokenId, bool hasBuff) returns(uint8)
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) OpenBox(opts *bind.TransactOpts, tokenId *big.Int, hasBuff bool) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "openBox", tokenId, hasBuff)
}

// OpenBox is a paid mutator transaction binding the contract method 0xcc1c6db1.
//
// Solidity: function openBox(uint256 tokenId, bool hasBuff) returns(uint8)
func (_MysteryBoxGen1 *MysteryBoxGen1Session) OpenBox(tokenId *big.Int, hasBuff bool) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.OpenBox(&_MysteryBoxGen1.TransactOpts, tokenId, hasBuff)
}

// OpenBox is a paid mutator transaction binding the contract method 0xcc1c6db1.
//
// Solidity: function openBox(uint256 tokenId, bool hasBuff) returns(uint8)
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) OpenBox(tokenId *big.Int, hasBuff bool) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.OpenBox(&_MysteryBoxGen1.TransactOpts, tokenId, hasBuff)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.RenounceRole(&_MysteryBoxGen1.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.RenounceRole(&_MysteryBoxGen1.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.RevokeRole(&_MysteryBoxGen1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.RevokeRole(&_MysteryBoxGen1.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.SafeTransferFrom(&_MysteryBoxGen1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.SafeTransferFrom(&_MysteryBoxGen1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.SafeTransferFrom0(&_MysteryBoxGen1.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.SafeTransferFrom0(&_MysteryBoxGen1.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.SetApprovalForAll(&_MysteryBoxGen1.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.SetApprovalForAll(&_MysteryBoxGen1.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen1.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.TransferFrom(&_MysteryBoxGen1.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MysteryBoxGen1 *MysteryBoxGen1TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MysteryBoxGen1.Contract.TransferFrom(&_MysteryBoxGen1.TransactOpts, from, to, tokenId)
}

// MysteryBoxGen1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1ApprovalIterator struct {
	Event *MysteryBoxGen1Approval // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen1Approval)
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
		it.Event = new(MysteryBoxGen1Approval)
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
func (it *MysteryBoxGen1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen1Approval represents a Approval event raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MysteryBoxGen1ApprovalIterator, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1ApprovalIterator{contract: _MysteryBoxGen1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen1Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen1Approval)
				if err := _MysteryBoxGen1.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) ParseApproval(log types.Log) (*MysteryBoxGen1Approval, error) {
	event := new(MysteryBoxGen1Approval)
	if err := _MysteryBoxGen1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen1ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1ApprovalForAllIterator struct {
	Event *MysteryBoxGen1ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen1ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen1ApprovalForAll)
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
		it.Event = new(MysteryBoxGen1ApprovalForAll)
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
func (it *MysteryBoxGen1ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen1ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen1ApprovalForAll represents a ApprovalForAll event raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MysteryBoxGen1ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MysteryBoxGen1.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1ApprovalForAllIterator{contract: _MysteryBoxGen1.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen1ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MysteryBoxGen1.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen1ApprovalForAll)
				if err := _MysteryBoxGen1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) ParseApprovalForAll(log types.Log) (*MysteryBoxGen1ApprovalForAll, error) {
	event := new(MysteryBoxGen1ApprovalForAll)
	if err := _MysteryBoxGen1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen1MintBoxIterator is returned from FilterMintBox and is used to iterate over the raw logs and unpacked data for MintBox events raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1MintBoxIterator struct {
	Event *MysteryBoxGen1MintBox // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen1MintBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen1MintBox)
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
		it.Event = new(MysteryBoxGen1MintBox)
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
func (it *MysteryBoxGen1MintBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen1MintBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen1MintBox represents a MintBox event raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1MintBox struct {
	User  common.Address
	BoxId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMintBox is a free log retrieval operation binding the contract event 0x92b26e707b024019810eb9f81fb26ab92c318a8efdc77204f029ed774a4ca84b.
//
// Solidity: event MintBox(address indexed user, uint256 boxId)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) FilterMintBox(opts *bind.FilterOpts, user []common.Address) (*MysteryBoxGen1MintBoxIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MysteryBoxGen1.contract.FilterLogs(opts, "MintBox", userRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1MintBoxIterator{contract: _MysteryBoxGen1.contract, event: "MintBox", logs: logs, sub: sub}, nil
}

// WatchMintBox is a free log subscription operation binding the contract event 0x92b26e707b024019810eb9f81fb26ab92c318a8efdc77204f029ed774a4ca84b.
//
// Solidity: event MintBox(address indexed user, uint256 boxId)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) WatchMintBox(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen1MintBox, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MysteryBoxGen1.contract.WatchLogs(opts, "MintBox", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen1MintBox)
				if err := _MysteryBoxGen1.contract.UnpackLog(event, "MintBox", log); err != nil {
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

// ParseMintBox is a log parse operation binding the contract event 0x92b26e707b024019810eb9f81fb26ab92c318a8efdc77204f029ed774a4ca84b.
//
// Solidity: event MintBox(address indexed user, uint256 boxId)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) ParseMintBox(log types.Log) (*MysteryBoxGen1MintBox, error) {
	event := new(MysteryBoxGen1MintBox)
	if err := _MysteryBoxGen1.contract.UnpackLog(event, "MintBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen1OpenBoxIterator is returned from FilterOpenBox and is used to iterate over the raw logs and unpacked data for OpenBox events raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1OpenBoxIterator struct {
	Event *MysteryBoxGen1OpenBox // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen1OpenBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen1OpenBox)
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
		it.Event = new(MysteryBoxGen1OpenBox)
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
func (it *MysteryBoxGen1OpenBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen1OpenBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen1OpenBox represents a OpenBox event raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1OpenBox struct {
	User    common.Address
	TokenId *big.Int
	BoxType uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenBox is a free log retrieval operation binding the contract event 0x36f124809754b34f903ebc2995e02198ae2d4bc6c088078a48a6fade21444584.
//
// Solidity: event OpenBox(address indexed user, uint256 tokenId, uint8 boxType)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) FilterOpenBox(opts *bind.FilterOpts, user []common.Address) (*MysteryBoxGen1OpenBoxIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MysteryBoxGen1.contract.FilterLogs(opts, "OpenBox", userRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1OpenBoxIterator{contract: _MysteryBoxGen1.contract, event: "OpenBox", logs: logs, sub: sub}, nil
}

// WatchOpenBox is a free log subscription operation binding the contract event 0x36f124809754b34f903ebc2995e02198ae2d4bc6c088078a48a6fade21444584.
//
// Solidity: event OpenBox(address indexed user, uint256 tokenId, uint8 boxType)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) WatchOpenBox(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen1OpenBox, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MysteryBoxGen1.contract.WatchLogs(opts, "OpenBox", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen1OpenBox)
				if err := _MysteryBoxGen1.contract.UnpackLog(event, "OpenBox", log); err != nil {
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

// ParseOpenBox is a log parse operation binding the contract event 0x36f124809754b34f903ebc2995e02198ae2d4bc6c088078a48a6fade21444584.
//
// Solidity: event OpenBox(address indexed user, uint256 tokenId, uint8 boxType)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) ParseOpenBox(log types.Log) (*MysteryBoxGen1OpenBox, error) {
	event := new(MysteryBoxGen1OpenBox)
	if err := _MysteryBoxGen1.contract.UnpackLog(event, "OpenBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen1RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1RoleAdminChangedIterator struct {
	Event *MysteryBoxGen1RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen1RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen1RoleAdminChanged)
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
		it.Event = new(MysteryBoxGen1RoleAdminChanged)
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
func (it *MysteryBoxGen1RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen1RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen1RoleAdminChanged represents a RoleAdminChanged event raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MysteryBoxGen1RoleAdminChangedIterator, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1RoleAdminChangedIterator{contract: _MysteryBoxGen1.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen1RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen1RoleAdminChanged)
				if err := _MysteryBoxGen1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) ParseRoleAdminChanged(log types.Log) (*MysteryBoxGen1RoleAdminChanged, error) {
	event := new(MysteryBoxGen1RoleAdminChanged)
	if err := _MysteryBoxGen1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen1RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1RoleGrantedIterator struct {
	Event *MysteryBoxGen1RoleGranted // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen1RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen1RoleGranted)
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
		it.Event = new(MysteryBoxGen1RoleGranted)
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
func (it *MysteryBoxGen1RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen1RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen1RoleGranted represents a RoleGranted event raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MysteryBoxGen1RoleGrantedIterator, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1RoleGrantedIterator{contract: _MysteryBoxGen1.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen1RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen1RoleGranted)
				if err := _MysteryBoxGen1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) ParseRoleGranted(log types.Log) (*MysteryBoxGen1RoleGranted, error) {
	event := new(MysteryBoxGen1RoleGranted)
	if err := _MysteryBoxGen1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen1RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1RoleRevokedIterator struct {
	Event *MysteryBoxGen1RoleRevoked // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen1RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen1RoleRevoked)
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
		it.Event = new(MysteryBoxGen1RoleRevoked)
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
func (it *MysteryBoxGen1RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen1RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen1RoleRevoked represents a RoleRevoked event raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MysteryBoxGen1RoleRevokedIterator, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1RoleRevokedIterator{contract: _MysteryBoxGen1.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen1RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen1RoleRevoked)
				if err := _MysteryBoxGen1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) ParseRoleRevoked(log types.Log) (*MysteryBoxGen1RoleRevoked, error) {
	event := new(MysteryBoxGen1RoleRevoked)
	if err := _MysteryBoxGen1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MysteryBoxGen1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1TransferIterator struct {
	Event *MysteryBoxGen1Transfer // Event containing the contract specifics and raw log

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
func (it *MysteryBoxGen1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MysteryBoxGen1Transfer)
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
		it.Event = new(MysteryBoxGen1Transfer)
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
func (it *MysteryBoxGen1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MysteryBoxGen1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MysteryBoxGen1Transfer represents a Transfer event raised by the MysteryBoxGen1 contract.
type MysteryBoxGen1Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MysteryBoxGen1TransferIterator, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MysteryBoxGen1TransferIterator{contract: _MysteryBoxGen1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MysteryBoxGen1Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _MysteryBoxGen1.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MysteryBoxGen1Transfer)
				if err := _MysteryBoxGen1.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MysteryBoxGen1 *MysteryBoxGen1Filterer) ParseTransfer(log types.Log) (*MysteryBoxGen1Transfer, error) {
	event := new(MysteryBoxGen1Transfer)
	if err := _MysteryBoxGen1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
