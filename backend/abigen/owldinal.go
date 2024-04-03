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

// OwldinalMetaData contains all meta data concerning the Owldinal contract.
var OwldinalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"endedBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voyaTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"idOneOwnerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defenderAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boxId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"MintBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_VOYA_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"checkCanMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canMint\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintByAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainVoyaCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whiteListEndedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OwldinalABI is the input ABI used to generate the binding from.
// Deprecated: Use OwldinalMetaData.ABI instead.
var OwldinalABI = OwldinalMetaData.ABI

// Owldinal is an auto generated Go binding around an Ethereum contract.
type Owldinal struct {
	OwldinalCaller     // Read-only binding to the contract
	OwldinalTransactor // Write-only binding to the contract
	OwldinalFilterer   // Log filterer for contract events
}

// OwldinalCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwldinalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwldinalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwldinalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwldinalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwldinalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwldinalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwldinalSession struct {
	Contract     *Owldinal         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwldinalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwldinalCallerSession struct {
	Contract *OwldinalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OwldinalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwldinalTransactorSession struct {
	Contract     *OwldinalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OwldinalRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwldinalRaw struct {
	Contract *Owldinal // Generic contract binding to access the raw methods on
}

// OwldinalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwldinalCallerRaw struct {
	Contract *OwldinalCaller // Generic read-only contract binding to access the raw methods on
}

// OwldinalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwldinalTransactorRaw struct {
	Contract *OwldinalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwldinal creates a new instance of Owldinal, bound to a specific deployed contract.
func NewOwldinal(address common.Address, backend bind.ContractBackend) (*Owldinal, error) {
	contract, err := bindOwldinal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owldinal{OwldinalCaller: OwldinalCaller{contract: contract}, OwldinalTransactor: OwldinalTransactor{contract: contract}, OwldinalFilterer: OwldinalFilterer{contract: contract}}, nil
}

// NewOwldinalCaller creates a new read-only instance of Owldinal, bound to a specific deployed contract.
func NewOwldinalCaller(address common.Address, caller bind.ContractCaller) (*OwldinalCaller, error) {
	contract, err := bindOwldinal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwldinalCaller{contract: contract}, nil
}

// NewOwldinalTransactor creates a new write-only instance of Owldinal, bound to a specific deployed contract.
func NewOwldinalTransactor(address common.Address, transactor bind.ContractTransactor) (*OwldinalTransactor, error) {
	contract, err := bindOwldinal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwldinalTransactor{contract: contract}, nil
}

// NewOwldinalFilterer creates a new log filterer instance of Owldinal, bound to a specific deployed contract.
func NewOwldinalFilterer(address common.Address, filterer bind.ContractFilterer) (*OwldinalFilterer, error) {
	contract, err := bindOwldinal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwldinalFilterer{contract: contract}, nil
}

// bindOwldinal binds a generic wrapper to an already deployed contract.
func bindOwldinal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwldinalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owldinal *OwldinalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owldinal.Contract.OwldinalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owldinal *OwldinalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owldinal.Contract.OwldinalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owldinal *OwldinalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owldinal.Contract.OwldinalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owldinal *OwldinalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owldinal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owldinal *OwldinalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owldinal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owldinal *OwldinalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owldinal.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Owldinal *OwldinalCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Owldinal *OwldinalSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Owldinal.Contract.DEFAULTADMINROLE(&_Owldinal.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Owldinal *OwldinalCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Owldinal.Contract.DEFAULTADMINROLE(&_Owldinal.CallOpts)
}

// MINTPRICE is a free data retrieval call binding the contract method 0xc002d23d.
//
// Solidity: function MINT_PRICE() view returns(uint256)
func (_Owldinal *OwldinalCaller) MINTPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "MINT_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTPRICE is a free data retrieval call binding the contract method 0xc002d23d.
//
// Solidity: function MINT_PRICE() view returns(uint256)
func (_Owldinal *OwldinalSession) MINTPRICE() (*big.Int, error) {
	return _Owldinal.Contract.MINTPRICE(&_Owldinal.CallOpts)
}

// MINTPRICE is a free data retrieval call binding the contract method 0xc002d23d.
//
// Solidity: function MINT_PRICE() view returns(uint256)
func (_Owldinal *OwldinalCallerSession) MINTPRICE() (*big.Int, error) {
	return _Owldinal.Contract.MINTPRICE(&_Owldinal.CallOpts)
}

// MINTVOYATHRESHOLD is a free data retrieval call binding the contract method 0x40ba20bc.
//
// Solidity: function MINT_VOYA_THRESHOLD() view returns(uint256)
func (_Owldinal *OwldinalCaller) MINTVOYATHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "MINT_VOYA_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTVOYATHRESHOLD is a free data retrieval call binding the contract method 0x40ba20bc.
//
// Solidity: function MINT_VOYA_THRESHOLD() view returns(uint256)
func (_Owldinal *OwldinalSession) MINTVOYATHRESHOLD() (*big.Int, error) {
	return _Owldinal.Contract.MINTVOYATHRESHOLD(&_Owldinal.CallOpts)
}

// MINTVOYATHRESHOLD is a free data retrieval call binding the contract method 0x40ba20bc.
//
// Solidity: function MINT_VOYA_THRESHOLD() view returns(uint256)
func (_Owldinal *OwldinalCallerSession) MINTVOYATHRESHOLD() (*big.Int, error) {
	return _Owldinal.Contract.MINTVOYATHRESHOLD(&_Owldinal.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Owldinal *OwldinalCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Owldinal *OwldinalSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Owldinal.Contract.BalanceOf(&_Owldinal.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Owldinal *OwldinalCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Owldinal.Contract.BalanceOf(&_Owldinal.CallOpts, owner)
}

// CheckCanMint is a free data retrieval call binding the contract method 0xd339d8d0.
//
// Solidity: function checkCanMint(address addr) view returns(bool canMint, uint256 mintType)
func (_Owldinal *OwldinalCaller) CheckCanMint(opts *bind.CallOpts, addr common.Address) (struct {
	CanMint  bool
	MintType *big.Int
}, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "checkCanMint", addr)

	outstruct := new(struct {
		CanMint  bool
		MintType *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CanMint = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.MintType = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CheckCanMint is a free data retrieval call binding the contract method 0xd339d8d0.
//
// Solidity: function checkCanMint(address addr) view returns(bool canMint, uint256 mintType)
func (_Owldinal *OwldinalSession) CheckCanMint(addr common.Address) (struct {
	CanMint  bool
	MintType *big.Int
}, error) {
	return _Owldinal.Contract.CheckCanMint(&_Owldinal.CallOpts, addr)
}

// CheckCanMint is a free data retrieval call binding the contract method 0xd339d8d0.
//
// Solidity: function checkCanMint(address addr) view returns(bool canMint, uint256 mintType)
func (_Owldinal *OwldinalCallerSession) CheckCanMint(addr common.Address) (struct {
	CanMint  bool
	MintType *big.Int
}, error) {
	return _Owldinal.Contract.CheckCanMint(&_Owldinal.CallOpts, addr)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Owldinal *OwldinalCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Owldinal *OwldinalSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Owldinal.Contract.GetApproved(&_Owldinal.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Owldinal *OwldinalCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Owldinal.Contract.GetApproved(&_Owldinal.CallOpts, tokenId)
}

// GetCurrentBalance is a free data retrieval call binding the contract method 0xa5749710.
//
// Solidity: function getCurrentBalance() view returns(uint256)
func (_Owldinal *OwldinalCaller) GetCurrentBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "getCurrentBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBalance is a free data retrieval call binding the contract method 0xa5749710.
//
// Solidity: function getCurrentBalance() view returns(uint256)
func (_Owldinal *OwldinalSession) GetCurrentBalance() (*big.Int, error) {
	return _Owldinal.Contract.GetCurrentBalance(&_Owldinal.CallOpts)
}

// GetCurrentBalance is a free data retrieval call binding the contract method 0xa5749710.
//
// Solidity: function getCurrentBalance() view returns(uint256)
func (_Owldinal *OwldinalCallerSession) GetCurrentBalance() (*big.Int, error) {
	return _Owldinal.Contract.GetCurrentBalance(&_Owldinal.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Owldinal *OwldinalCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Owldinal *OwldinalSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Owldinal.Contract.GetRoleAdmin(&_Owldinal.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Owldinal *OwldinalCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Owldinal.Contract.GetRoleAdmin(&_Owldinal.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Owldinal *OwldinalCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Owldinal *OwldinalSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Owldinal.Contract.HasRole(&_Owldinal.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Owldinal *OwldinalCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Owldinal.Contract.HasRole(&_Owldinal.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Owldinal *OwldinalCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Owldinal *OwldinalSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Owldinal.Contract.IsApprovedForAll(&_Owldinal.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Owldinal *OwldinalCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Owldinal.Contract.IsApprovedForAll(&_Owldinal.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Owldinal *OwldinalCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Owldinal *OwldinalSession) Name() (string, error) {
	return _Owldinal.Contract.Name(&_Owldinal.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Owldinal *OwldinalCallerSession) Name() (string, error) {
	return _Owldinal.Contract.Name(&_Owldinal.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Owldinal *OwldinalCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Owldinal *OwldinalSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Owldinal.Contract.OwnerOf(&_Owldinal.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Owldinal *OwldinalCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Owldinal.Contract.OwnerOf(&_Owldinal.CallOpts, tokenId)
}

// RemainVoyaCount is a free data retrieval call binding the contract method 0x04d75faa.
//
// Solidity: function remainVoyaCount() view returns(uint256)
func (_Owldinal *OwldinalCaller) RemainVoyaCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "remainVoyaCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainVoyaCount is a free data retrieval call binding the contract method 0x04d75faa.
//
// Solidity: function remainVoyaCount() view returns(uint256)
func (_Owldinal *OwldinalSession) RemainVoyaCount() (*big.Int, error) {
	return _Owldinal.Contract.RemainVoyaCount(&_Owldinal.CallOpts)
}

// RemainVoyaCount is a free data retrieval call binding the contract method 0x04d75faa.
//
// Solidity: function remainVoyaCount() view returns(uint256)
func (_Owldinal *OwldinalCallerSession) RemainVoyaCount() (*big.Int, error) {
	return _Owldinal.Contract.RemainVoyaCount(&_Owldinal.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Owldinal *OwldinalCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Owldinal *OwldinalSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Owldinal.Contract.SupportsInterface(&_Owldinal.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Owldinal *OwldinalCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Owldinal.Contract.SupportsInterface(&_Owldinal.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Owldinal *OwldinalCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Owldinal *OwldinalSession) Symbol() (string, error) {
	return _Owldinal.Contract.Symbol(&_Owldinal.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Owldinal *OwldinalCallerSession) Symbol() (string, error) {
	return _Owldinal.Contract.Symbol(&_Owldinal.CallOpts)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_Owldinal *OwldinalCaller) TokenIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "tokenIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_Owldinal *OwldinalSession) TokenIdCounter() (*big.Int, error) {
	return _Owldinal.Contract.TokenIdCounter(&_Owldinal.CallOpts)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_Owldinal *OwldinalCallerSession) TokenIdCounter() (*big.Int, error) {
	return _Owldinal.Contract.TokenIdCounter(&_Owldinal.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Owldinal *OwldinalCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Owldinal *OwldinalSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Owldinal.Contract.TokenURI(&_Owldinal.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Owldinal *OwldinalCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Owldinal.Contract.TokenURI(&_Owldinal.CallOpts, tokenId)
}

// WhiteListEndedBlock is a free data retrieval call binding the contract method 0xd7379ba1.
//
// Solidity: function whiteListEndedBlock() view returns(uint256)
func (_Owldinal *OwldinalCaller) WhiteListEndedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Owldinal.contract.Call(opts, &out, "whiteListEndedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhiteListEndedBlock is a free data retrieval call binding the contract method 0xd7379ba1.
//
// Solidity: function whiteListEndedBlock() view returns(uint256)
func (_Owldinal *OwldinalSession) WhiteListEndedBlock() (*big.Int, error) {
	return _Owldinal.Contract.WhiteListEndedBlock(&_Owldinal.CallOpts)
}

// WhiteListEndedBlock is a free data retrieval call binding the contract method 0xd7379ba1.
//
// Solidity: function whiteListEndedBlock() view returns(uint256)
func (_Owldinal *OwldinalCallerSession) WhiteListEndedBlock() (*big.Int, error) {
	return _Owldinal.Contract.WhiteListEndedBlock(&_Owldinal.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Owldinal *OwldinalTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Owldinal *OwldinalSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Owldinal.Contract.Approve(&_Owldinal.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Owldinal *OwldinalTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Owldinal.Contract.Approve(&_Owldinal.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Owldinal *OwldinalTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Owldinal *OwldinalSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Owldinal.Contract.GrantRole(&_Owldinal.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Owldinal *OwldinalTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Owldinal.Contract.GrantRole(&_Owldinal.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x7ba0e2e7.
//
// Solidity: function mint(bytes signature) payable returns(uint256 tokenId, uint256 mintType)
func (_Owldinal *OwldinalTransactor) Mint(opts *bind.TransactOpts, signature []byte) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "mint", signature)
}

// Mint is a paid mutator transaction binding the contract method 0x7ba0e2e7.
//
// Solidity: function mint(bytes signature) payable returns(uint256 tokenId, uint256 mintType)
func (_Owldinal *OwldinalSession) Mint(signature []byte) (*types.Transaction, error) {
	return _Owldinal.Contract.Mint(&_Owldinal.TransactOpts, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x7ba0e2e7.
//
// Solidity: function mint(bytes signature) payable returns(uint256 tokenId, uint256 mintType)
func (_Owldinal *OwldinalTransactorSession) Mint(signature []byte) (*types.Transaction, error) {
	return _Owldinal.Contract.Mint(&_Owldinal.TransactOpts, signature)
}

// MintByAdmin is a paid mutator transaction binding the contract method 0x09d10d46.
//
// Solidity: function mintByAdmin() returns(uint256 tokenId)
func (_Owldinal *OwldinalTransactor) MintByAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "mintByAdmin")
}

// MintByAdmin is a paid mutator transaction binding the contract method 0x09d10d46.
//
// Solidity: function mintByAdmin() returns(uint256 tokenId)
func (_Owldinal *OwldinalSession) MintByAdmin() (*types.Transaction, error) {
	return _Owldinal.Contract.MintByAdmin(&_Owldinal.TransactOpts)
}

// MintByAdmin is a paid mutator transaction binding the contract method 0x09d10d46.
//
// Solidity: function mintByAdmin() returns(uint256 tokenId)
func (_Owldinal *OwldinalTransactorSession) MintByAdmin() (*types.Transaction, error) {
	return _Owldinal.Contract.MintByAdmin(&_Owldinal.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Owldinal *OwldinalTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Owldinal *OwldinalSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Owldinal.Contract.RenounceRole(&_Owldinal.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Owldinal *OwldinalTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Owldinal.Contract.RenounceRole(&_Owldinal.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Owldinal *OwldinalTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Owldinal *OwldinalSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Owldinal.Contract.RevokeRole(&_Owldinal.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Owldinal *OwldinalTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Owldinal.Contract.RevokeRole(&_Owldinal.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Owldinal *OwldinalTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Owldinal *OwldinalSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Owldinal.Contract.SafeTransferFrom(&_Owldinal.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Owldinal *OwldinalTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Owldinal.Contract.SafeTransferFrom(&_Owldinal.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Owldinal *OwldinalTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Owldinal *OwldinalSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Owldinal.Contract.SafeTransferFrom0(&_Owldinal.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Owldinal *OwldinalTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Owldinal.Contract.SafeTransferFrom0(&_Owldinal.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Owldinal *OwldinalTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Owldinal *OwldinalSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Owldinal.Contract.SetApprovalForAll(&_Owldinal.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Owldinal *OwldinalTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Owldinal.Contract.SetApprovalForAll(&_Owldinal.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Owldinal *OwldinalTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Owldinal *OwldinalSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Owldinal.Contract.TransferFrom(&_Owldinal.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Owldinal *OwldinalTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Owldinal.Contract.TransferFrom(&_Owldinal.TransactOpts, from, to, tokenId)
}

// ValidSignature is a paid mutator transaction binding the contract method 0xfec3f3f8.
//
// Solidity: function validSignature(bytes signature) returns(bool)
func (_Owldinal *OwldinalTransactor) ValidSignature(opts *bind.TransactOpts, signature []byte) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "validSignature", signature)
}

// ValidSignature is a paid mutator transaction binding the contract method 0xfec3f3f8.
//
// Solidity: function validSignature(bytes signature) returns(bool)
func (_Owldinal *OwldinalSession) ValidSignature(signature []byte) (*types.Transaction, error) {
	return _Owldinal.Contract.ValidSignature(&_Owldinal.TransactOpts, signature)
}

// ValidSignature is a paid mutator transaction binding the contract method 0xfec3f3f8.
//
// Solidity: function validSignature(bytes signature) returns(bool)
func (_Owldinal *OwldinalTransactorSession) ValidSignature(signature []byte) (*types.Transaction, error) {
	return _Owldinal.Contract.ValidSignature(&_Owldinal.TransactOpts, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns(uint256)
func (_Owldinal *OwldinalTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Owldinal.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns(uint256)
func (_Owldinal *OwldinalSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Owldinal.Contract.Withdraw(&_Owldinal.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns(uint256)
func (_Owldinal *OwldinalTransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Owldinal.Contract.Withdraw(&_Owldinal.TransactOpts, recipient)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Owldinal *OwldinalTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Owldinal.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Owldinal *OwldinalSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Owldinal.Contract.Fallback(&_Owldinal.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Owldinal *OwldinalTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Owldinal.Contract.Fallback(&_Owldinal.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Owldinal *OwldinalTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owldinal.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Owldinal *OwldinalSession) Receive() (*types.Transaction, error) {
	return _Owldinal.Contract.Receive(&_Owldinal.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Owldinal *OwldinalTransactorSession) Receive() (*types.Transaction, error) {
	return _Owldinal.Contract.Receive(&_Owldinal.TransactOpts)
}

// OwldinalApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Owldinal contract.
type OwldinalApprovalIterator struct {
	Event *OwldinalApproval // Event containing the contract specifics and raw log

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
func (it *OwldinalApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalApproval)
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
		it.Event = new(OwldinalApproval)
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
func (it *OwldinalApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalApproval represents a Approval event raised by the Owldinal contract.
type OwldinalApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Owldinal *OwldinalFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*OwldinalApprovalIterator, error) {

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

	logs, sub, err := _Owldinal.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalApprovalIterator{contract: _Owldinal.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Owldinal *OwldinalFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OwldinalApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Owldinal.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalApproval)
				if err := _Owldinal.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Owldinal *OwldinalFilterer) ParseApproval(log types.Log) (*OwldinalApproval, error) {
	event := new(OwldinalApproval)
	if err := _Owldinal.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Owldinal contract.
type OwldinalApprovalForAllIterator struct {
	Event *OwldinalApprovalForAll // Event containing the contract specifics and raw log

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
func (it *OwldinalApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalApprovalForAll)
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
		it.Event = new(OwldinalApprovalForAll)
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
func (it *OwldinalApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalApprovalForAll represents a ApprovalForAll event raised by the Owldinal contract.
type OwldinalApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Owldinal *OwldinalFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*OwldinalApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Owldinal.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalApprovalForAllIterator{contract: _Owldinal.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Owldinal *OwldinalFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *OwldinalApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Owldinal.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalApprovalForAll)
				if err := _Owldinal.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Owldinal *OwldinalFilterer) ParseApprovalForAll(log types.Log) (*OwldinalApprovalForAll, error) {
	event := new(OwldinalApprovalForAll)
	if err := _Owldinal.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Owldinal contract.
type OwldinalBatchMetadataUpdateIterator struct {
	Event *OwldinalBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *OwldinalBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalBatchMetadataUpdate)
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
		it.Event = new(OwldinalBatchMetadataUpdate)
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
func (it *OwldinalBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Owldinal contract.
type OwldinalBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Owldinal *OwldinalFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*OwldinalBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Owldinal.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &OwldinalBatchMetadataUpdateIterator{contract: _Owldinal.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Owldinal *OwldinalFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *OwldinalBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Owldinal.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalBatchMetadataUpdate)
				if err := _Owldinal.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Owldinal *OwldinalFilterer) ParseBatchMetadataUpdate(log types.Log) (*OwldinalBatchMetadataUpdate, error) {
	event := new(OwldinalBatchMetadataUpdate)
	if err := _Owldinal.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Owldinal contract.
type OwldinalMetadataUpdateIterator struct {
	Event *OwldinalMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *OwldinalMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalMetadataUpdate)
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
		it.Event = new(OwldinalMetadataUpdate)
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
func (it *OwldinalMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalMetadataUpdate represents a MetadataUpdate event raised by the Owldinal contract.
type OwldinalMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Owldinal *OwldinalFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*OwldinalMetadataUpdateIterator, error) {

	logs, sub, err := _Owldinal.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &OwldinalMetadataUpdateIterator{contract: _Owldinal.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Owldinal *OwldinalFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *OwldinalMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Owldinal.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalMetadataUpdate)
				if err := _Owldinal.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_Owldinal *OwldinalFilterer) ParseMetadataUpdate(log types.Log) (*OwldinalMetadataUpdate, error) {
	event := new(OwldinalMetadataUpdate)
	if err := _Owldinal.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalMintBoxIterator is returned from FilterMintBox and is used to iterate over the raw logs and unpacked data for MintBox events raised by the Owldinal contract.
type OwldinalMintBoxIterator struct {
	Event *OwldinalMintBox // Event containing the contract specifics and raw log

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
func (it *OwldinalMintBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalMintBox)
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
		it.Event = new(OwldinalMintBox)
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
func (it *OwldinalMintBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalMintBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalMintBox represents a MintBox event raised by the Owldinal contract.
type OwldinalMintBox struct {
	User     common.Address
	BoxId    *big.Int
	MintType *big.Int
	Url      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintBox is a free log retrieval operation binding the contract event 0xf5d3f864a50c2df29b92152f2936fc5520ee555438f668048785c1868cd34230.
//
// Solidity: event MintBox(address indexed user, uint256 boxId, uint256 mintType, string url)
func (_Owldinal *OwldinalFilterer) FilterMintBox(opts *bind.FilterOpts, user []common.Address) (*OwldinalMintBoxIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Owldinal.contract.FilterLogs(opts, "MintBox", userRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalMintBoxIterator{contract: _Owldinal.contract, event: "MintBox", logs: logs, sub: sub}, nil
}

// WatchMintBox is a free log subscription operation binding the contract event 0xf5d3f864a50c2df29b92152f2936fc5520ee555438f668048785c1868cd34230.
//
// Solidity: event MintBox(address indexed user, uint256 boxId, uint256 mintType, string url)
func (_Owldinal *OwldinalFilterer) WatchMintBox(opts *bind.WatchOpts, sink chan<- *OwldinalMintBox, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Owldinal.contract.WatchLogs(opts, "MintBox", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalMintBox)
				if err := _Owldinal.contract.UnpackLog(event, "MintBox", log); err != nil {
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

// ParseMintBox is a log parse operation binding the contract event 0xf5d3f864a50c2df29b92152f2936fc5520ee555438f668048785c1868cd34230.
//
// Solidity: event MintBox(address indexed user, uint256 boxId, uint256 mintType, string url)
func (_Owldinal *OwldinalFilterer) ParseMintBox(log types.Log) (*OwldinalMintBox, error) {
	event := new(OwldinalMintBox)
	if err := _Owldinal.contract.UnpackLog(event, "MintBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Owldinal contract.
type OwldinalRoleAdminChangedIterator struct {
	Event *OwldinalRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OwldinalRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalRoleAdminChanged)
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
		it.Event = new(OwldinalRoleAdminChanged)
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
func (it *OwldinalRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalRoleAdminChanged represents a RoleAdminChanged event raised by the Owldinal contract.
type OwldinalRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Owldinal *OwldinalFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OwldinalRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Owldinal.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalRoleAdminChangedIterator{contract: _Owldinal.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Owldinal *OwldinalFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OwldinalRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Owldinal.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalRoleAdminChanged)
				if err := _Owldinal.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Owldinal *OwldinalFilterer) ParseRoleAdminChanged(log types.Log) (*OwldinalRoleAdminChanged, error) {
	event := new(OwldinalRoleAdminChanged)
	if err := _Owldinal.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Owldinal contract.
type OwldinalRoleGrantedIterator struct {
	Event *OwldinalRoleGranted // Event containing the contract specifics and raw log

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
func (it *OwldinalRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalRoleGranted)
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
		it.Event = new(OwldinalRoleGranted)
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
func (it *OwldinalRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalRoleGranted represents a RoleGranted event raised by the Owldinal contract.
type OwldinalRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Owldinal *OwldinalFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OwldinalRoleGrantedIterator, error) {

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

	logs, sub, err := _Owldinal.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalRoleGrantedIterator{contract: _Owldinal.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Owldinal *OwldinalFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OwldinalRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Owldinal.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalRoleGranted)
				if err := _Owldinal.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Owldinal *OwldinalFilterer) ParseRoleGranted(log types.Log) (*OwldinalRoleGranted, error) {
	event := new(OwldinalRoleGranted)
	if err := _Owldinal.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Owldinal contract.
type OwldinalRoleRevokedIterator struct {
	Event *OwldinalRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OwldinalRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalRoleRevoked)
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
		it.Event = new(OwldinalRoleRevoked)
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
func (it *OwldinalRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalRoleRevoked represents a RoleRevoked event raised by the Owldinal contract.
type OwldinalRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Owldinal *OwldinalFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OwldinalRoleRevokedIterator, error) {

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

	logs, sub, err := _Owldinal.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalRoleRevokedIterator{contract: _Owldinal.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Owldinal *OwldinalFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OwldinalRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Owldinal.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalRoleRevoked)
				if err := _Owldinal.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Owldinal *OwldinalFilterer) ParseRoleRevoked(log types.Log) (*OwldinalRoleRevoked, error) {
	event := new(OwldinalRoleRevoked)
	if err := _Owldinal.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwldinalTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Owldinal contract.
type OwldinalTransferIterator struct {
	Event *OwldinalTransfer // Event containing the contract specifics and raw log

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
func (it *OwldinalTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwldinalTransfer)
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
		it.Event = new(OwldinalTransfer)
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
func (it *OwldinalTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwldinalTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwldinalTransfer represents a Transfer event raised by the Owldinal contract.
type OwldinalTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Owldinal *OwldinalFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*OwldinalTransferIterator, error) {

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

	logs, sub, err := _Owldinal.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OwldinalTransferIterator{contract: _Owldinal.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Owldinal *OwldinalFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OwldinalTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Owldinal.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwldinalTransfer)
				if err := _Owldinal.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Owldinal *OwldinalFilterer) ParseTransfer(log types.Log) (*OwldinalTransfer, error) {
	event := new(OwldinalTransfer)
	if err := _Owldinal.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
