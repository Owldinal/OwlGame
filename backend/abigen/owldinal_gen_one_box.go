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

// OwldinalGenOneBoxMetaData contains all meta data concerning the OwldinalGenOneBox contract.
var OwldinalGenOneBoxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC20Burnable\",\"name\":\"_owlToken\",\"type\":\"address\"},{\"internalType\":\"contractOwlGame\",\"name\":\"_gameContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumBoxType\",\"name\":\"boxType\",\"type\":\"uint8\"}],\"name\":\"MintBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAME_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameContract\",\"outputs\":[{\"internalType\":\"contractOwlGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBoxType\",\"outputs\":[{\"internalType\":\"enumBoxType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isBoxOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasBuff\",\"type\":\"bool\"}],\"name\":\"mintAndOpenBoxes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owlToken\",\"outputs\":[{\"internalType\":\"contractERC20Burnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenBoxTypes\",\"outputs\":[{\"internalType\":\"enumBoxType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwldinalGenOneBoxABI is the input ABI used to generate the binding from.
// Deprecated: Use OwldinalGenOneBoxMetaData.ABI instead.
var OwldinalGenOneBoxABI = OwldinalGenOneBoxMetaData.ABI

// OwldinalGenOneBox is an auto generated Go binding around an Ethereum contract.
type OwldinalGenOneBox struct {
	OwldinalGenOneBoxCaller     // Read-only binding to the contract
	OwldinalGenOneBoxTransactor // Write-only binding to the contract
	OwldinalGenOneBoxFilterer   // Log filterer for contract events
}

// OwldinalGenOneBoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwldinalGenOneBoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwldinalGenOneBoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwldinalGenOneBoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwldinalGenOneBoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwldinalGenOneBoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwldinalGenOneBoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwldinalGenOneBoxSession struct {
	Contract     *OwldinalGenOneBox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwldinalGenOneBoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwldinalGenOneBoxCallerSession struct {
	Contract *OwldinalGenOneBoxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OwldinalGenOneBoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwldinalGenOneBoxTransactorSession struct {
	Contract     *OwldinalGenOneBoxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OwldinalGenOneBoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwldinalGenOneBoxRaw struct {
	Contract *OwldinalGenOneBox // Generic contract binding to access the raw methods on
}

// OwldinalGenOneBoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwldinalGenOneBoxCallerRaw struct {
	Contract *OwldinalGenOneBoxCaller // Generic read-only contract binding to access the raw methods on
}

// OwldinalGenOneBoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwldinalGenOneBoxTransactorRaw struct {
	Contract *OwldinalGenOneBoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwldinalGenOneBox creates a new instance of OwldinalGenOneBox, bound to a specific deployed contract.
func NewOwldinalGenOneBox(address common.Address, backend bind.ContractBackend) (*OwldinalGenOneBox, error) {
	contract, err := bindOwldinalGenOneBox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBox{OwldinalGenOneBoxCaller: OwldinalGenOneBoxCaller{contract: contract}, OwldinalGenOneBoxTransactor: OwldinalGenOneBoxTransactor{contract: contract}, OwldinalGenOneBoxFilterer: OwldinalGenOneBoxFilterer{contract: contract}}, nil
}

// NewOwldinalGenOneBoxCaller creates a new read-only instance of OwldinalGenOneBox, bound to a specific deployed contract.
func NewOwldinalGenOneBoxCaller(address common.Address, caller bind.ContractCaller) (*OwldinalGenOneBoxCaller, error) {
	contract, err := bindOwldinalGenOneBox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxCaller{contract: contract}, nil
}

// NewOwldinalGenOneBoxTransactor creates a new write-only instance of OwldinalGenOneBox, bound to a specific deployed contract.
func NewOwldinalGenOneBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*OwldinalGenOneBoxTransactor, error) {
	contract, err := bindOwldinalGenOneBox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxTransactor{contract: contract}, nil
}

// NewOwldinalGenOneBoxFilterer creates a new log filterer instance of OwldinalGenOneBox, bound to a specific deployed contract.
func NewOwldinalGenOneBoxFilterer(address common.Address, filterer bind.ContractFilterer) (*OwldinalGenOneBoxFilterer, error) {
	contract, err := bindOwldinalGenOneBox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxFilterer{contract: contract}, nil
}

// bindOwldinalGenOneBox binds a generic wrapper to an already deployed contract.
func bindOwldinalGenOneBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwldinalGenOneBoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwldinalGenOneBox *OwldinalGenOneBoxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwldinalGenOneBox.Contract.OwldinalGenOneBoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwldinalGenOneBox *OwldinalGenOneBoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.OwldinalGenOneBoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwldinalGenOneBox *OwldinalGenOneBoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.OwldinalGenOneBoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwldinalGenOneBox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OwldinalGenOneBox.Contract.DEFAULTADMINROLE(&_OwldinalGenOneBox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OwldinalGenOneBox.Contract.DEFAULTADMINROLE(&_OwldinalGenOneBox.CallOpts)
}

// GAMECONTRACTROLE is a free data retrieval call binding the contract method 0xb5e352da.
//
// Solidity: function GAME_CONTRACT_ROLE() view returns(bytes32)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) GAMECONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "GAME_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GAMECONTRACTROLE is a free data retrieval call binding the contract method 0xb5e352da.
//
// Solidity: function GAME_CONTRACT_ROLE() view returns(bytes32)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) GAMECONTRACTROLE() ([32]byte, error) {
	return _OwldinalGenOneBox.Contract.GAMECONTRACTROLE(&_OwldinalGenOneBox.CallOpts)
}

// GAMECONTRACTROLE is a free data retrieval call binding the contract method 0xb5e352da.
//
// Solidity: function GAME_CONTRACT_ROLE() view returns(bytes32)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) GAMECONTRACTROLE() ([32]byte, error) {
	return _OwldinalGenOneBox.Contract.GAMECONTRACTROLE(&_OwldinalGenOneBox.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _OwldinalGenOneBox.Contract.BalanceOf(&_OwldinalGenOneBox.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _OwldinalGenOneBox.Contract.BalanceOf(&_OwldinalGenOneBox.CallOpts, owner)
}

// GameContract is a free data retrieval call binding the contract method 0xd3f33009.
//
// Solidity: function gameContract() view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) GameContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "gameContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameContract is a free data retrieval call binding the contract method 0xd3f33009.
//
// Solidity: function gameContract() view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) GameContract() (common.Address, error) {
	return _OwldinalGenOneBox.Contract.GameContract(&_OwldinalGenOneBox.CallOpts)
}

// GameContract is a free data retrieval call binding the contract method 0xd3f33009.
//
// Solidity: function gameContract() view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) GameContract() (common.Address, error) {
	return _OwldinalGenOneBox.Contract.GameContract(&_OwldinalGenOneBox.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _OwldinalGenOneBox.Contract.GetApproved(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _OwldinalGenOneBox.Contract.GetApproved(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// GetBoxType is a free data retrieval call binding the contract method 0xb546a3bc.
//
// Solidity: function getBoxType(uint256 tokenId) view returns(uint8)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) GetBoxType(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "getBoxType", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetBoxType is a free data retrieval call binding the contract method 0xb546a3bc.
//
// Solidity: function getBoxType(uint256 tokenId) view returns(uint8)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) GetBoxType(tokenId *big.Int) (uint8, error) {
	return _OwldinalGenOneBox.Contract.GetBoxType(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// GetBoxType is a free data retrieval call binding the contract method 0xb546a3bc.
//
// Solidity: function getBoxType(uint256 tokenId) view returns(uint8)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) GetBoxType(tokenId *big.Int) (uint8, error) {
	return _OwldinalGenOneBox.Contract.GetBoxType(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OwldinalGenOneBox.Contract.GetRoleAdmin(&_OwldinalGenOneBox.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OwldinalGenOneBox.Contract.GetRoleAdmin(&_OwldinalGenOneBox.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OwldinalGenOneBox.Contract.HasRole(&_OwldinalGenOneBox.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OwldinalGenOneBox.Contract.HasRole(&_OwldinalGenOneBox.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _OwldinalGenOneBox.Contract.IsApprovedForAll(&_OwldinalGenOneBox.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _OwldinalGenOneBox.Contract.IsApprovedForAll(&_OwldinalGenOneBox.CallOpts, owner, operator)
}

// IsBoxOpened is a free data retrieval call binding the contract method 0xa0520689.
//
// Solidity: function isBoxOpened(uint256 tokenId) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) IsBoxOpened(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "isBoxOpened", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBoxOpened is a free data retrieval call binding the contract method 0xa0520689.
//
// Solidity: function isBoxOpened(uint256 tokenId) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) IsBoxOpened(tokenId *big.Int) (bool, error) {
	return _OwldinalGenOneBox.Contract.IsBoxOpened(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// IsBoxOpened is a free data retrieval call binding the contract method 0xa0520689.
//
// Solidity: function isBoxOpened(uint256 tokenId) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) IsBoxOpened(tokenId *big.Int) (bool, error) {
	return _OwldinalGenOneBox.Contract.IsBoxOpened(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) Name() (string, error) {
	return _OwldinalGenOneBox.Contract.Name(&_OwldinalGenOneBox.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) Name() (string, error) {
	return _OwldinalGenOneBox.Contract.Name(&_OwldinalGenOneBox.CallOpts)
}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) OwlToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "owlToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) OwlToken() (common.Address, error) {
	return _OwldinalGenOneBox.Contract.OwlToken(&_OwldinalGenOneBox.CallOpts)
}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) OwlToken() (common.Address, error) {
	return _OwldinalGenOneBox.Contract.OwlToken(&_OwldinalGenOneBox.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _OwldinalGenOneBox.Contract.OwnerOf(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _OwldinalGenOneBox.Contract.OwnerOf(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OwldinalGenOneBox.Contract.SupportsInterface(&_OwldinalGenOneBox.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OwldinalGenOneBox.Contract.SupportsInterface(&_OwldinalGenOneBox.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) Symbol() (string, error) {
	return _OwldinalGenOneBox.Contract.Symbol(&_OwldinalGenOneBox.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) Symbol() (string, error) {
	return _OwldinalGenOneBox.Contract.Symbol(&_OwldinalGenOneBox.CallOpts)
}

// TokenBoxTypes is a free data retrieval call binding the contract method 0xcdb13d20.
//
// Solidity: function tokenBoxTypes(uint256 ) view returns(uint8)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) TokenBoxTypes(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "tokenBoxTypes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenBoxTypes is a free data retrieval call binding the contract method 0xcdb13d20.
//
// Solidity: function tokenBoxTypes(uint256 ) view returns(uint8)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) TokenBoxTypes(arg0 *big.Int) (uint8, error) {
	return _OwldinalGenOneBox.Contract.TokenBoxTypes(&_OwldinalGenOneBox.CallOpts, arg0)
}

// TokenBoxTypes is a free data retrieval call binding the contract method 0xcdb13d20.
//
// Solidity: function tokenBoxTypes(uint256 ) view returns(uint8)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) TokenBoxTypes(arg0 *big.Int) (uint8, error) {
	return _OwldinalGenOneBox.Contract.TokenBoxTypes(&_OwldinalGenOneBox.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _OwldinalGenOneBox.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) TokenURI(tokenId *big.Int) (string, error) {
	return _OwldinalGenOneBox.Contract.TokenURI(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_OwldinalGenOneBox *OwldinalGenOneBoxCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _OwldinalGenOneBox.Contract.TokenURI(&_OwldinalGenOneBox.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OwldinalGenOneBox.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.Approve(&_OwldinalGenOneBox.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.Approve(&_OwldinalGenOneBox.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwldinalGenOneBox.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.GrantRole(&_OwldinalGenOneBox.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.GrantRole(&_OwldinalGenOneBox.TransactOpts, role, account)
}

// MintAndOpenBoxes is a paid mutator transaction binding the contract method 0x96cd5900.
//
// Solidity: function mintAndOpenBoxes(address owner, uint256 count, bool hasBuff) returns(uint256[] tokenIdList)
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactor) MintAndOpenBoxes(opts *bind.TransactOpts, owner common.Address, count *big.Int, hasBuff bool) (*types.Transaction, error) {
	return _OwldinalGenOneBox.contract.Transact(opts, "mintAndOpenBoxes", owner, count, hasBuff)
}

// MintAndOpenBoxes is a paid mutator transaction binding the contract method 0x96cd5900.
//
// Solidity: function mintAndOpenBoxes(address owner, uint256 count, bool hasBuff) returns(uint256[] tokenIdList)
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) MintAndOpenBoxes(owner common.Address, count *big.Int, hasBuff bool) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.MintAndOpenBoxes(&_OwldinalGenOneBox.TransactOpts, owner, count, hasBuff)
}

// MintAndOpenBoxes is a paid mutator transaction binding the contract method 0x96cd5900.
//
// Solidity: function mintAndOpenBoxes(address owner, uint256 count, bool hasBuff) returns(uint256[] tokenIdList)
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorSession) MintAndOpenBoxes(owner common.Address, count *big.Int, hasBuff bool) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.MintAndOpenBoxes(&_OwldinalGenOneBox.TransactOpts, owner, count, hasBuff)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OwldinalGenOneBox.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.RenounceRole(&_OwldinalGenOneBox.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.RenounceRole(&_OwldinalGenOneBox.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwldinalGenOneBox.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.RevokeRole(&_OwldinalGenOneBox.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.RevokeRole(&_OwldinalGenOneBox.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OwldinalGenOneBox.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.SafeTransferFrom(&_OwldinalGenOneBox.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.SafeTransferFrom(&_OwldinalGenOneBox.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _OwldinalGenOneBox.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.SafeTransferFrom0(&_OwldinalGenOneBox.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.SafeTransferFrom0(&_OwldinalGenOneBox.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _OwldinalGenOneBox.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.SetApprovalForAll(&_OwldinalGenOneBox.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.SetApprovalForAll(&_OwldinalGenOneBox.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OwldinalGenOneBox.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.TransferFrom(&_OwldinalGenOneBox.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_OwldinalGenOneBox *OwldinalGenOneBoxTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OwldinalGenOneBox.Contract.TransferFrom(&_OwldinalGenOneBox.TransactOpts, from, to, tokenId)
}

// OwldinalGenOneBoxApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxApprovalIterator struct {
	Event *OwldinalGenOneBoxApproval // Event containing the contract specifics and raw log

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
func (it *OwldinalGenOneBoxApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalGenOneBoxApproval)
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
		it.Event = new(OwldinalGenOneBoxApproval)
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
func (it *OwldinalGenOneBoxApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalGenOneBoxApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalGenOneBoxApproval represents a Approval event raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*OwldinalGenOneBoxApprovalIterator, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxApprovalIterator{contract: _OwldinalGenOneBox.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OwldinalGenOneBoxApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalGenOneBoxApproval)
				if err := _OwldinalGenOneBox.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) ParseApproval(log types.Log) (*OwldinalGenOneBoxApproval, error) {
	event := new(OwldinalGenOneBoxApproval)
	if err := _OwldinalGenOneBox.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalGenOneBoxApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxApprovalForAllIterator struct {
	Event *OwldinalGenOneBoxApprovalForAll // Event containing the contract specifics and raw log

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
func (it *OwldinalGenOneBoxApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalGenOneBoxApprovalForAll)
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
		it.Event = new(OwldinalGenOneBoxApprovalForAll)
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
func (it *OwldinalGenOneBoxApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalGenOneBoxApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalGenOneBoxApprovalForAll represents a ApprovalForAll event raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*OwldinalGenOneBoxApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OwldinalGenOneBox.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxApprovalForAllIterator{contract: _OwldinalGenOneBox.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *OwldinalGenOneBoxApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OwldinalGenOneBox.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalGenOneBoxApprovalForAll)
				if err := _OwldinalGenOneBox.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) ParseApprovalForAll(log types.Log) (*OwldinalGenOneBoxApprovalForAll, error) {
	event := new(OwldinalGenOneBoxApprovalForAll)
	if err := _OwldinalGenOneBox.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalGenOneBoxMintBoxIterator is returned from FilterMintBox and is used to iterate over the raw logs and unpacked data for MintBox events raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxMintBoxIterator struct {
	Event *OwldinalGenOneBoxMintBox // Event containing the contract specifics and raw log

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
func (it *OwldinalGenOneBoxMintBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalGenOneBoxMintBox)
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
		it.Event = new(OwldinalGenOneBoxMintBox)
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
func (it *OwldinalGenOneBoxMintBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalGenOneBoxMintBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalGenOneBoxMintBox represents a MintBox event raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxMintBox struct {
	User    common.Address
	TokenId *big.Int
	BoxType uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintBox is a free log retrieval operation binding the contract event 0x4cce2d7ca388465a90e71f76235d389abe1ede028b09c07d4f86519e5adb078c.
//
// Solidity: event MintBox(address indexed user, uint256 tokenId, uint8 boxType)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) FilterMintBox(opts *bind.FilterOpts, user []common.Address) (*OwldinalGenOneBoxMintBoxIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwldinalGenOneBox.contract.FilterLogs(opts, "MintBox", userRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxMintBoxIterator{contract: _OwldinalGenOneBox.contract, event: "MintBox", logs: logs, sub: sub}, nil
}

// WatchMintBox is a free log subscription operation binding the contract event 0x4cce2d7ca388465a90e71f76235d389abe1ede028b09c07d4f86519e5adb078c.
//
// Solidity: event MintBox(address indexed user, uint256 tokenId, uint8 boxType)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) WatchMintBox(opts *bind.WatchOpts, sink chan<- *OwldinalGenOneBoxMintBox, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwldinalGenOneBox.contract.WatchLogs(opts, "MintBox", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalGenOneBoxMintBox)
				if err := _OwldinalGenOneBox.contract.UnpackLog(event, "MintBox", log); err != nil {
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

// ParseMintBox is a log parse operation binding the contract event 0x4cce2d7ca388465a90e71f76235d389abe1ede028b09c07d4f86519e5adb078c.
//
// Solidity: event MintBox(address indexed user, uint256 tokenId, uint8 boxType)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) ParseMintBox(log types.Log) (*OwldinalGenOneBoxMintBox, error) {
	event := new(OwldinalGenOneBoxMintBox)
	if err := _OwldinalGenOneBox.contract.UnpackLog(event, "MintBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalGenOneBoxRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxRoleAdminChangedIterator struct {
	Event *OwldinalGenOneBoxRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OwldinalGenOneBoxRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalGenOneBoxRoleAdminChanged)
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
		it.Event = new(OwldinalGenOneBoxRoleAdminChanged)
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
func (it *OwldinalGenOneBoxRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalGenOneBoxRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalGenOneBoxRoleAdminChanged represents a RoleAdminChanged event raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OwldinalGenOneBoxRoleAdminChangedIterator, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxRoleAdminChangedIterator{contract: _OwldinalGenOneBox.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OwldinalGenOneBoxRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalGenOneBoxRoleAdminChanged)
				if err := _OwldinalGenOneBox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) ParseRoleAdminChanged(log types.Log) (*OwldinalGenOneBoxRoleAdminChanged, error) {
	event := new(OwldinalGenOneBoxRoleAdminChanged)
	if err := _OwldinalGenOneBox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalGenOneBoxRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxRoleGrantedIterator struct {
	Event *OwldinalGenOneBoxRoleGranted // Event containing the contract specifics and raw log

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
func (it *OwldinalGenOneBoxRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalGenOneBoxRoleGranted)
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
		it.Event = new(OwldinalGenOneBoxRoleGranted)
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
func (it *OwldinalGenOneBoxRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalGenOneBoxRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalGenOneBoxRoleGranted represents a RoleGranted event raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OwldinalGenOneBoxRoleGrantedIterator, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxRoleGrantedIterator{contract: _OwldinalGenOneBox.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OwldinalGenOneBoxRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalGenOneBoxRoleGranted)
				if err := _OwldinalGenOneBox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) ParseRoleGranted(log types.Log) (*OwldinalGenOneBoxRoleGranted, error) {
	event := new(OwldinalGenOneBoxRoleGranted)
	if err := _OwldinalGenOneBox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalGenOneBoxRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxRoleRevokedIterator struct {
	Event *OwldinalGenOneBoxRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OwldinalGenOneBoxRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalGenOneBoxRoleRevoked)
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
		it.Event = new(OwldinalGenOneBoxRoleRevoked)
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
func (it *OwldinalGenOneBoxRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalGenOneBoxRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalGenOneBoxRoleRevoked represents a RoleRevoked event raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OwldinalGenOneBoxRoleRevokedIterator, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxRoleRevokedIterator{contract: _OwldinalGenOneBox.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OwldinalGenOneBoxRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalGenOneBoxRoleRevoked)
				if err := _OwldinalGenOneBox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) ParseRoleRevoked(log types.Log) (*OwldinalGenOneBoxRoleRevoked, error) {
	event := new(OwldinalGenOneBoxRoleRevoked)
	if err := _OwldinalGenOneBox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalGenOneBoxTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxTransferIterator struct {
	Event *OwldinalGenOneBoxTransfer // Event containing the contract specifics and raw log

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
func (it *OwldinalGenOneBoxTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalGenOneBoxTransfer)
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
		it.Event = new(OwldinalGenOneBoxTransfer)
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
func (it *OwldinalGenOneBoxTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalGenOneBoxTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalGenOneBoxTransfer represents a Transfer event raised by the OwldinalGenOneBox contract.
type OwldinalGenOneBoxTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*OwldinalGenOneBoxTransferIterator, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalGenOneBoxTransferIterator{contract: _OwldinalGenOneBox.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OwldinalGenOneBoxTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _OwldinalGenOneBox.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalGenOneBoxTransfer)
				if err := _OwldinalGenOneBox.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_OwldinalGenOneBox *OwldinalGenOneBoxFilterer) ParseTransfer(log types.Log) (*OwldinalGenOneBoxTransfer, error) {
	event := new(OwldinalGenOneBoxTransfer)
	if err := _OwldinalGenOneBox.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
