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

// OwlGameMetaData contains all meta data concerning the OwlGame contract.
var OwlGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRUIT_REWARD_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_REBATE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"}],\"name\":\"addPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boxGen0Contract\",\"outputs\":[{\"internalType\":\"contractOwldinal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boxGen1Contract\",\"outputs\":[{\"internalType\":\"contractMysteryBoxGen1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimInviterReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owlTokenAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owldinalNftAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mysteryBoxAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"inviteCode\",\"type\":\"uint32\"}],\"name\":\"mintGen1Box\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owlToken\",\"outputs\":[{\"internalType\":\"contractERC20Burnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isGen0\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"inviteCode\",\"type\":\"uint32\"}],\"name\":\"stakeBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"unstakeOwl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"unstakeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAllFruitRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwlGameABI is the input ABI used to generate the binding from.
// Deprecated: Use OwlGameMetaData.ABI instead.
var OwlGameABI = OwlGameMetaData.ABI

// OwlGame is an auto generated Go binding around an Ethereum contract.
type OwlGame struct {
	OwlGameCaller     // Read-only binding to the contract
	OwlGameTransactor // Write-only binding to the contract
	OwlGameFilterer   // Log filterer for contract events
}

// OwlGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwlGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwlGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwlGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwlGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwlGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwlGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwlGameSession struct {
	Contract     *OwlGame          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwlGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwlGameCallerSession struct {
	Contract *OwlGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwlGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwlGameTransactorSession struct {
	Contract     *OwlGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwlGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwlGameRaw struct {
	Contract *OwlGame // Generic contract binding to access the raw methods on
}

// OwlGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwlGameCallerRaw struct {
	Contract *OwlGameCaller // Generic read-only contract binding to access the raw methods on
}

// OwlGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwlGameTransactorRaw struct {
	Contract *OwlGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwlGame creates a new instance of OwlGame, bound to a specific deployed contract.
func NewOwlGame(address common.Address, backend bind.ContractBackend) (*OwlGame, error) {
	contract, err := bindOwlGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwlGame{OwlGameCaller: OwlGameCaller{contract: contract}, OwlGameTransactor: OwlGameTransactor{contract: contract}, OwlGameFilterer: OwlGameFilterer{contract: contract}}, nil
}

// NewOwlGameCaller creates a new read-only instance of OwlGame, bound to a specific deployed contract.
func NewOwlGameCaller(address common.Address, caller bind.ContractCaller) (*OwlGameCaller, error) {
	contract, err := bindOwlGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwlGameCaller{contract: contract}, nil
}

// NewOwlGameTransactor creates a new write-only instance of OwlGame, bound to a specific deployed contract.
func NewOwlGameTransactor(address common.Address, transactor bind.ContractTransactor) (*OwlGameTransactor, error) {
	contract, err := bindOwlGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwlGameTransactor{contract: contract}, nil
}

// NewOwlGameFilterer creates a new log filterer instance of OwlGame, bound to a specific deployed contract.
func NewOwlGameFilterer(address common.Address, filterer bind.ContractFilterer) (*OwlGameFilterer, error) {
	contract, err := bindOwlGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwlGameFilterer{contract: contract}, nil
}

// bindOwlGame binds a generic wrapper to an already deployed contract.
func bindOwlGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwlGameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwlGame *OwlGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwlGame.Contract.OwlGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwlGame *OwlGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwlGame.Contract.OwlGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwlGame *OwlGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwlGame.Contract.OwlGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwlGame *OwlGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwlGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwlGame *OwlGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwlGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwlGame *OwlGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwlGame.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OwlGame *OwlGameCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OwlGame *OwlGameSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OwlGame.Contract.DEFAULTADMINROLE(&_OwlGame.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OwlGame *OwlGameCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OwlGame.Contract.DEFAULTADMINROLE(&_OwlGame.CallOpts)
}

// FRUITREWARDINTERVAL is a free data retrieval call binding the contract method 0xd0fdeef1.
//
// Solidity: function FRUIT_REWARD_INTERVAL() view returns(uint256)
func (_OwlGame *OwlGameCaller) FRUITREWARDINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "FRUIT_REWARD_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FRUITREWARDINTERVAL is a free data retrieval call binding the contract method 0xd0fdeef1.
//
// Solidity: function FRUIT_REWARD_INTERVAL() view returns(uint256)
func (_OwlGame *OwlGameSession) FRUITREWARDINTERVAL() (*big.Int, error) {
	return _OwlGame.Contract.FRUITREWARDINTERVAL(&_OwlGame.CallOpts)
}

// FRUITREWARDINTERVAL is a free data retrieval call binding the contract method 0xd0fdeef1.
//
// Solidity: function FRUIT_REWARD_INTERVAL() view returns(uint256)
func (_OwlGame *OwlGameCallerSession) FRUITREWARDINTERVAL() (*big.Int, error) {
	return _OwlGame.Contract.FRUITREWARDINTERVAL(&_OwlGame.CallOpts)
}

// MINTPRICE is a free data retrieval call binding the contract method 0xc002d23d.
//
// Solidity: function MINT_PRICE() view returns(uint256)
func (_OwlGame *OwlGameCaller) MINTPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "MINT_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTPRICE is a free data retrieval call binding the contract method 0xc002d23d.
//
// Solidity: function MINT_PRICE() view returns(uint256)
func (_OwlGame *OwlGameSession) MINTPRICE() (*big.Int, error) {
	return _OwlGame.Contract.MINTPRICE(&_OwlGame.CallOpts)
}

// MINTPRICE is a free data retrieval call binding the contract method 0xc002d23d.
//
// Solidity: function MINT_PRICE() view returns(uint256)
func (_OwlGame *OwlGameCallerSession) MINTPRICE() (*big.Int, error) {
	return _OwlGame.Contract.MINTPRICE(&_OwlGame.CallOpts)
}

// MINTREBATE is a free data retrieval call binding the contract method 0xd84b9ac2.
//
// Solidity: function MINT_REBATE() view returns(uint32)
func (_OwlGame *OwlGameCaller) MINTREBATE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "MINT_REBATE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MINTREBATE is a free data retrieval call binding the contract method 0xd84b9ac2.
//
// Solidity: function MINT_REBATE() view returns(uint32)
func (_OwlGame *OwlGameSession) MINTREBATE() (uint32, error) {
	return _OwlGame.Contract.MINTREBATE(&_OwlGame.CallOpts)
}

// MINTREBATE is a free data retrieval call binding the contract method 0xd84b9ac2.
//
// Solidity: function MINT_REBATE() view returns(uint32)
func (_OwlGame *OwlGameCallerSession) MINTREBATE() (uint32, error) {
	return _OwlGame.Contract.MINTREBATE(&_OwlGame.CallOpts)
}

// BoxGen0Contract is a free data retrieval call binding the contract method 0xca3f21d0.
//
// Solidity: function boxGen0Contract() view returns(address)
func (_OwlGame *OwlGameCaller) BoxGen0Contract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "boxGen0Contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoxGen0Contract is a free data retrieval call binding the contract method 0xca3f21d0.
//
// Solidity: function boxGen0Contract() view returns(address)
func (_OwlGame *OwlGameSession) BoxGen0Contract() (common.Address, error) {
	return _OwlGame.Contract.BoxGen0Contract(&_OwlGame.CallOpts)
}

// BoxGen0Contract is a free data retrieval call binding the contract method 0xca3f21d0.
//
// Solidity: function boxGen0Contract() view returns(address)
func (_OwlGame *OwlGameCallerSession) BoxGen0Contract() (common.Address, error) {
	return _OwlGame.Contract.BoxGen0Contract(&_OwlGame.CallOpts)
}

// BoxGen1Contract is a free data retrieval call binding the contract method 0xdda18896.
//
// Solidity: function boxGen1Contract() view returns(address)
func (_OwlGame *OwlGameCaller) BoxGen1Contract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "boxGen1Contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoxGen1Contract is a free data retrieval call binding the contract method 0xdda18896.
//
// Solidity: function boxGen1Contract() view returns(address)
func (_OwlGame *OwlGameSession) BoxGen1Contract() (common.Address, error) {
	return _OwlGame.Contract.BoxGen1Contract(&_OwlGame.CallOpts)
}

// BoxGen1Contract is a free data retrieval call binding the contract method 0xdda18896.
//
// Solidity: function boxGen1Contract() view returns(address)
func (_OwlGame *OwlGameCallerSession) BoxGen1Contract() (common.Address, error) {
	return _OwlGame.Contract.BoxGen1Contract(&_OwlGame.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OwlGame *OwlGameCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OwlGame *OwlGameSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OwlGame.Contract.GetRoleAdmin(&_OwlGame.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OwlGame *OwlGameCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OwlGame.Contract.GetRoleAdmin(&_OwlGame.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OwlGame *OwlGameCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OwlGame *OwlGameSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OwlGame.Contract.HasRole(&_OwlGame.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OwlGame *OwlGameCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OwlGame.Contract.HasRole(&_OwlGame.CallOpts, role, account)
}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() view returns(address)
func (_OwlGame *OwlGameCaller) OwlToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "owlToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() view returns(address)
func (_OwlGame *OwlGameSession) OwlToken() (common.Address, error) {
	return _OwlGame.Contract.OwlToken(&_OwlGame.CallOpts)
}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() view returns(address)
func (_OwlGame *OwlGameCallerSession) OwlToken() (common.Address, error) {
	return _OwlGame.Contract.OwlToken(&_OwlGame.CallOpts)
}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(uint256)
func (_OwlGame *OwlGameCaller) PrizePool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "prizePool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(uint256)
func (_OwlGame *OwlGameSession) PrizePool() (*big.Int, error) {
	return _OwlGame.Contract.PrizePool(&_OwlGame.CallOpts)
}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(uint256)
func (_OwlGame *OwlGameCallerSession) PrizePool() (*big.Int, error) {
	return _OwlGame.Contract.PrizePool(&_OwlGame.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OwlGame *OwlGameCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OwlGame *OwlGameSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OwlGame.Contract.SupportsInterface(&_OwlGame.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OwlGame *OwlGameCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OwlGame.Contract.SupportsInterface(&_OwlGame.CallOpts, interfaceId)
}

// AddPrize is a paid mutator transaction binding the contract method 0x2a050024.
//
// Solidity: function addPrize(uint256 prizeAmount) returns()
func (_OwlGame *OwlGameTransactor) AddPrize(opts *bind.TransactOpts, prizeAmount *big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "addPrize", prizeAmount)
}

// AddPrize is a paid mutator transaction binding the contract method 0x2a050024.
//
// Solidity: function addPrize(uint256 prizeAmount) returns()
func (_OwlGame *OwlGameSession) AddPrize(prizeAmount *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.AddPrize(&_OwlGame.TransactOpts, prizeAmount)
}

// AddPrize is a paid mutator transaction binding the contract method 0x2a050024.
//
// Solidity: function addPrize(uint256 prizeAmount) returns()
func (_OwlGame *OwlGameTransactorSession) AddPrize(prizeAmount *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.AddPrize(&_OwlGame.TransactOpts, prizeAmount)
}

// ClaimInviterReward is a paid mutator transaction binding the contract method 0xb0b03638.
//
// Solidity: function claimInviterReward() returns()
func (_OwlGame *OwlGameTransactor) ClaimInviterReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "claimInviterReward")
}

// ClaimInviterReward is a paid mutator transaction binding the contract method 0xb0b03638.
//
// Solidity: function claimInviterReward() returns()
func (_OwlGame *OwlGameSession) ClaimInviterReward() (*types.Transaction, error) {
	return _OwlGame.Contract.ClaimInviterReward(&_OwlGame.TransactOpts)
}

// ClaimInviterReward is a paid mutator transaction binding the contract method 0xb0b03638.
//
// Solidity: function claimInviterReward() returns()
func (_OwlGame *OwlGameTransactorSession) ClaimInviterReward() (*types.Transaction, error) {
	return _OwlGame.Contract.ClaimInviterReward(&_OwlGame.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OwlGame *OwlGameTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OwlGame *OwlGameSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.GrantRole(&_OwlGame.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OwlGame *OwlGameTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.GrantRole(&_OwlGame.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owlTokenAddr, address owldinalNftAddr, address mysteryBoxAddr) returns()
func (_OwlGame *OwlGameTransactor) Initialize(opts *bind.TransactOpts, owlTokenAddr common.Address, owldinalNftAddr common.Address, mysteryBoxAddr common.Address) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "initialize", owlTokenAddr, owldinalNftAddr, mysteryBoxAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owlTokenAddr, address owldinalNftAddr, address mysteryBoxAddr) returns()
func (_OwlGame *OwlGameSession) Initialize(owlTokenAddr common.Address, owldinalNftAddr common.Address, mysteryBoxAddr common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.Initialize(&_OwlGame.TransactOpts, owlTokenAddr, owldinalNftAddr, mysteryBoxAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owlTokenAddr, address owldinalNftAddr, address mysteryBoxAddr) returns()
func (_OwlGame *OwlGameTransactorSession) Initialize(owlTokenAddr common.Address, owldinalNftAddr common.Address, mysteryBoxAddr common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.Initialize(&_OwlGame.TransactOpts, owlTokenAddr, owldinalNftAddr, mysteryBoxAddr)
}

// MintGen1Box is a paid mutator transaction binding the contract method 0xfde6c8ff.
//
// Solidity: function mintGen1Box(uint32 inviteCode) returns()
func (_OwlGame *OwlGameTransactor) MintGen1Box(opts *bind.TransactOpts, inviteCode uint32) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "mintGen1Box", inviteCode)
}

// MintGen1Box is a paid mutator transaction binding the contract method 0xfde6c8ff.
//
// Solidity: function mintGen1Box(uint32 inviteCode) returns()
func (_OwlGame *OwlGameSession) MintGen1Box(inviteCode uint32) (*types.Transaction, error) {
	return _OwlGame.Contract.MintGen1Box(&_OwlGame.TransactOpts, inviteCode)
}

// MintGen1Box is a paid mutator transaction binding the contract method 0xfde6c8ff.
//
// Solidity: function mintGen1Box(uint32 inviteCode) returns()
func (_OwlGame *OwlGameTransactorSession) MintGen1Box(inviteCode uint32) (*types.Transaction, error) {
	return _OwlGame.Contract.MintGen1Box(&_OwlGame.TransactOpts, inviteCode)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OwlGame *OwlGameTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OwlGame *OwlGameSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.RenounceRole(&_OwlGame.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OwlGame *OwlGameTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.RenounceRole(&_OwlGame.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OwlGame *OwlGameTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OwlGame *OwlGameSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.RevokeRole(&_OwlGame.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OwlGame *OwlGameTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.RevokeRole(&_OwlGame.TransactOpts, role, account)
}

// StakeBox is a paid mutator transaction binding the contract method 0x64d24687.
//
// Solidity: function stakeBox(uint256 tokenId, bool isGen0, uint32 inviteCode) returns()
func (_OwlGame *OwlGameTransactor) StakeBox(opts *bind.TransactOpts, tokenId *big.Int, isGen0 bool, inviteCode uint32) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "stakeBox", tokenId, isGen0, inviteCode)
}

// StakeBox is a paid mutator transaction binding the contract method 0x64d24687.
//
// Solidity: function stakeBox(uint256 tokenId, bool isGen0, uint32 inviteCode) returns()
func (_OwlGame *OwlGameSession) StakeBox(tokenId *big.Int, isGen0 bool, inviteCode uint32) (*types.Transaction, error) {
	return _OwlGame.Contract.StakeBox(&_OwlGame.TransactOpts, tokenId, isGen0, inviteCode)
}

// StakeBox is a paid mutator transaction binding the contract method 0x64d24687.
//
// Solidity: function stakeBox(uint256 tokenId, bool isGen0, uint32 inviteCode) returns()
func (_OwlGame *OwlGameTransactorSession) StakeBox(tokenId *big.Int, isGen0 bool, inviteCode uint32) (*types.Transaction, error) {
	return _OwlGame.Contract.StakeBox(&_OwlGame.TransactOpts, tokenId, isGen0, inviteCode)
}

// UnstakeOwl is a paid mutator transaction binding the contract method 0x4a12183b.
//
// Solidity: function unstakeOwl(uint256 tokenId) returns()
func (_OwlGame *OwlGameTransactor) UnstakeOwl(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "unstakeOwl", tokenId)
}

// UnstakeOwl is a paid mutator transaction binding the contract method 0x4a12183b.
//
// Solidity: function unstakeOwl(uint256 tokenId) returns()
func (_OwlGame *OwlGameSession) UnstakeOwl(tokenId *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.UnstakeOwl(&_OwlGame.TransactOpts, tokenId)
}

// UnstakeOwl is a paid mutator transaction binding the contract method 0x4a12183b.
//
// Solidity: function unstakeOwl(uint256 tokenId) returns()
func (_OwlGame *OwlGameTransactorSession) UnstakeOwl(tokenId *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.UnstakeOwl(&_OwlGame.TransactOpts, tokenId)
}

// UnstakeToken is a paid mutator transaction binding the contract method 0x2cfb6688.
//
// Solidity: function unstakeToken(uint256 tokenId) returns()
func (_OwlGame *OwlGameTransactor) UnstakeToken(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "unstakeToken", tokenId)
}

// UnstakeToken is a paid mutator transaction binding the contract method 0x2cfb6688.
//
// Solidity: function unstakeToken(uint256 tokenId) returns()
func (_OwlGame *OwlGameSession) UnstakeToken(tokenId *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.UnstakeToken(&_OwlGame.TransactOpts, tokenId)
}

// UnstakeToken is a paid mutator transaction binding the contract method 0x2cfb6688.
//
// Solidity: function unstakeToken(uint256 tokenId) returns()
func (_OwlGame *OwlGameTransactorSession) UnstakeToken(tokenId *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.UnstakeToken(&_OwlGame.TransactOpts, tokenId)
}

// UpdateAllFruitRewards is a paid mutator transaction binding the contract method 0xee2b3d54.
//
// Solidity: function updateAllFruitRewards() returns()
func (_OwlGame *OwlGameTransactor) UpdateAllFruitRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "updateAllFruitRewards")
}

// UpdateAllFruitRewards is a paid mutator transaction binding the contract method 0xee2b3d54.
//
// Solidity: function updateAllFruitRewards() returns()
func (_OwlGame *OwlGameSession) UpdateAllFruitRewards() (*types.Transaction, error) {
	return _OwlGame.Contract.UpdateAllFruitRewards(&_OwlGame.TransactOpts)
}

// UpdateAllFruitRewards is a paid mutator transaction binding the contract method 0xee2b3d54.
//
// Solidity: function updateAllFruitRewards() returns()
func (_OwlGame *OwlGameTransactorSession) UpdateAllFruitRewards() (*types.Transaction, error) {
	return _OwlGame.Contract.UpdateAllFruitRewards(&_OwlGame.TransactOpts)
}

// OwlGameRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the OwlGame contract.
type OwlGameRoleAdminChangedIterator struct {
	Event *OwlGameRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OwlGameRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameRoleAdminChanged)
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
		it.Event = new(OwlGameRoleAdminChanged)
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
func (it *OwlGameRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameRoleAdminChanged represents a RoleAdminChanged event raised by the OwlGame contract.
type OwlGameRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OwlGame *OwlGameFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OwlGameRoleAdminChangedIterator, error) {

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

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameRoleAdminChangedIterator{contract: _OwlGame.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OwlGame *OwlGameFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OwlGameRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameRoleAdminChanged)
				if err := _OwlGame.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_OwlGame *OwlGameFilterer) ParseRoleAdminChanged(log types.Log) (*OwlGameRoleAdminChanged, error) {
	event := new(OwlGameRoleAdminChanged)
	if err := _OwlGame.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the OwlGame contract.
type OwlGameRoleGrantedIterator struct {
	Event *OwlGameRoleGranted // Event containing the contract specifics and raw log

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
func (it *OwlGameRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameRoleGranted)
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
		it.Event = new(OwlGameRoleGranted)
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
func (it *OwlGameRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameRoleGranted represents a RoleGranted event raised by the OwlGame contract.
type OwlGameRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OwlGame *OwlGameFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OwlGameRoleGrantedIterator, error) {

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

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameRoleGrantedIterator{contract: _OwlGame.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OwlGame *OwlGameFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OwlGameRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameRoleGranted)
				if err := _OwlGame.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_OwlGame *OwlGameFilterer) ParseRoleGranted(log types.Log) (*OwlGameRoleGranted, error) {
	event := new(OwlGameRoleGranted)
	if err := _OwlGame.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the OwlGame contract.
type OwlGameRoleRevokedIterator struct {
	Event *OwlGameRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OwlGameRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameRoleRevoked)
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
		it.Event = new(OwlGameRoleRevoked)
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
func (it *OwlGameRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameRoleRevoked represents a RoleRevoked event raised by the OwlGame contract.
type OwlGameRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OwlGame *OwlGameFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OwlGameRoleRevokedIterator, error) {

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

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameRoleRevokedIterator{contract: _OwlGame.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OwlGame *OwlGameFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OwlGameRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameRoleRevoked)
				if err := _OwlGame.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_OwlGame *OwlGameFilterer) ParseRoleRevoked(log types.Log) (*OwlGameRoleRevoked, error) {
	event := new(OwlGameRoleRevoked)
	if err := _OwlGame.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
