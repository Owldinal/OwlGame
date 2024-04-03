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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"server\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"invitee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"name\":\"BindInvitation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"ClaimInviterReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"JoinGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizePoolDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizePoolIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"}],\"name\":\"StakeMysteryBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"}],\"name\":\"StakeOwldinalNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumBoxType\",\"name\":\"boxType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"UnstakeMysteryBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"}],\"name\":\"UnstakeOwldinalNft\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRUIT_REWARD_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_REBATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"}],\"name\":\"addPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"name\":\"claimAndUnstakeMysteryBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimInviterReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"inviteCode\",\"type\":\"uint32\"}],\"name\":\"handleInviteCode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owlTokenAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owldinalNftAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mysteryBoxAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMoonBoostEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"mintMysteryBox\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mysteryBoxContract\",\"outputs\":[{\"internalType\":\"contractOwldinalGenOneBox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owlToken\",\"outputs\":[{\"internalType\":\"contractERC20Burnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owldinalNftContract\",\"outputs\":[{\"internalType\":\"contractOwldinal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proportion\",\"type\":\"uint256\"}],\"name\":\"setGlobalFruitRewardsProportion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isEnable\",\"type\":\"bool\"}],\"name\":\"setMoonBoost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"name\":\"stakeMysteryBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"name\":\"stakeOwldinalNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"name\":\"unstakeOwldinalNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAllFruitRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
// Solidity: function MINT_REBATE() view returns(uint256)
func (_OwlGame *OwlGameCaller) MINTREBATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "MINT_REBATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTREBATE is a free data retrieval call binding the contract method 0xd84b9ac2.
//
// Solidity: function MINT_REBATE() view returns(uint256)
func (_OwlGame *OwlGameSession) MINTREBATE() (*big.Int, error) {
	return _OwlGame.Contract.MINTREBATE(&_OwlGame.CallOpts)
}

// MINTREBATE is a free data retrieval call binding the contract method 0xd84b9ac2.
//
// Solidity: function MINT_REBATE() view returns(uint256)
func (_OwlGame *OwlGameCallerSession) MINTREBATE() (*big.Int, error) {
	return _OwlGame.Contract.MINTREBATE(&_OwlGame.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_OwlGame *OwlGameCaller) SERVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "SERVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_OwlGame *OwlGameSession) SERVERROLE() ([32]byte, error) {
	return _OwlGame.Contract.SERVERROLE(&_OwlGame.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_OwlGame *OwlGameCallerSession) SERVERROLE() ([32]byte, error) {
	return _OwlGame.Contract.SERVERROLE(&_OwlGame.CallOpts)
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

// IsMoonBoostEnable is a free data retrieval call binding the contract method 0x1a8d3ad1.
//
// Solidity: function isMoonBoostEnable() view returns(bool)
func (_OwlGame *OwlGameCaller) IsMoonBoostEnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "isMoonBoostEnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMoonBoostEnable is a free data retrieval call binding the contract method 0x1a8d3ad1.
//
// Solidity: function isMoonBoostEnable() view returns(bool)
func (_OwlGame *OwlGameSession) IsMoonBoostEnable() (bool, error) {
	return _OwlGame.Contract.IsMoonBoostEnable(&_OwlGame.CallOpts)
}

// IsMoonBoostEnable is a free data retrieval call binding the contract method 0x1a8d3ad1.
//
// Solidity: function isMoonBoostEnable() view returns(bool)
func (_OwlGame *OwlGameCallerSession) IsMoonBoostEnable() (bool, error) {
	return _OwlGame.Contract.IsMoonBoostEnable(&_OwlGame.CallOpts)
}

// MysteryBoxContract is a free data retrieval call binding the contract method 0x95e84de4.
//
// Solidity: function mysteryBoxContract() view returns(address)
func (_OwlGame *OwlGameCaller) MysteryBoxContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "mysteryBoxContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MysteryBoxContract is a free data retrieval call binding the contract method 0x95e84de4.
//
// Solidity: function mysteryBoxContract() view returns(address)
func (_OwlGame *OwlGameSession) MysteryBoxContract() (common.Address, error) {
	return _OwlGame.Contract.MysteryBoxContract(&_OwlGame.CallOpts)
}

// MysteryBoxContract is a free data retrieval call binding the contract method 0x95e84de4.
//
// Solidity: function mysteryBoxContract() view returns(address)
func (_OwlGame *OwlGameCallerSession) MysteryBoxContract() (common.Address, error) {
	return _OwlGame.Contract.MysteryBoxContract(&_OwlGame.CallOpts)
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

// OwldinalNftContract is a free data retrieval call binding the contract method 0x0325d3ed.
//
// Solidity: function owldinalNftContract() view returns(address)
func (_OwlGame *OwlGameCaller) OwldinalNftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "owldinalNftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwldinalNftContract is a free data retrieval call binding the contract method 0x0325d3ed.
//
// Solidity: function owldinalNftContract() view returns(address)
func (_OwlGame *OwlGameSession) OwldinalNftContract() (common.Address, error) {
	return _OwlGame.Contract.OwldinalNftContract(&_OwlGame.CallOpts)
}

// OwldinalNftContract is a free data retrieval call binding the contract method 0x0325d3ed.
//
// Solidity: function owldinalNftContract() view returns(address)
func (_OwlGame *OwlGameCallerSession) OwldinalNftContract() (common.Address, error) {
	return _OwlGame.Contract.OwldinalNftContract(&_OwlGame.CallOpts)
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

// ClaimAndUnstakeMysteryBox is a paid mutator transaction binding the contract method 0xf280cd26.
//
// Solidity: function claimAndUnstakeMysteryBox(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameTransactor) ClaimAndUnstakeMysteryBox(opts *bind.TransactOpts, tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "claimAndUnstakeMysteryBox", tokenIdList)
}

// ClaimAndUnstakeMysteryBox is a paid mutator transaction binding the contract method 0xf280cd26.
//
// Solidity: function claimAndUnstakeMysteryBox(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameSession) ClaimAndUnstakeMysteryBox(tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.ClaimAndUnstakeMysteryBox(&_OwlGame.TransactOpts, tokenIdList)
}

// ClaimAndUnstakeMysteryBox is a paid mutator transaction binding the contract method 0xf280cd26.
//
// Solidity: function claimAndUnstakeMysteryBox(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameTransactorSession) ClaimAndUnstakeMysteryBox(tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.ClaimAndUnstakeMysteryBox(&_OwlGame.TransactOpts, tokenIdList)
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

// HandleInviteCode is a paid mutator transaction binding the contract method 0x857ac2fe.
//
// Solidity: function handleInviteCode(uint32 inviteCode) returns(address inviter)
func (_OwlGame *OwlGameTransactor) HandleInviteCode(opts *bind.TransactOpts, inviteCode uint32) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "handleInviteCode", inviteCode)
}

// HandleInviteCode is a paid mutator transaction binding the contract method 0x857ac2fe.
//
// Solidity: function handleInviteCode(uint32 inviteCode) returns(address inviter)
func (_OwlGame *OwlGameSession) HandleInviteCode(inviteCode uint32) (*types.Transaction, error) {
	return _OwlGame.Contract.HandleInviteCode(&_OwlGame.TransactOpts, inviteCode)
}

// HandleInviteCode is a paid mutator transaction binding the contract method 0x857ac2fe.
//
// Solidity: function handleInviteCode(uint32 inviteCode) returns(address inviter)
func (_OwlGame *OwlGameTransactorSession) HandleInviteCode(inviteCode uint32) (*types.Transaction, error) {
	return _OwlGame.Contract.HandleInviteCode(&_OwlGame.TransactOpts, inviteCode)
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

// MintMysteryBox is a paid mutator transaction binding the contract method 0x896b4d40.
//
// Solidity: function mintMysteryBox(uint256 count) returns(uint256[] tokenIdList)
func (_OwlGame *OwlGameTransactor) MintMysteryBox(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "mintMysteryBox", count)
}

// MintMysteryBox is a paid mutator transaction binding the contract method 0x896b4d40.
//
// Solidity: function mintMysteryBox(uint256 count) returns(uint256[] tokenIdList)
func (_OwlGame *OwlGameSession) MintMysteryBox(count *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.MintMysteryBox(&_OwlGame.TransactOpts, count)
}

// MintMysteryBox is a paid mutator transaction binding the contract method 0x896b4d40.
//
// Solidity: function mintMysteryBox(uint256 count) returns(uint256[] tokenIdList)
func (_OwlGame *OwlGameTransactorSession) MintMysteryBox(count *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.MintMysteryBox(&_OwlGame.TransactOpts, count)
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

// SetGlobalFruitRewardsProportion is a paid mutator transaction binding the contract method 0x68a83a87.
//
// Solidity: function setGlobalFruitRewardsProportion(uint256 proportion) returns()
func (_OwlGame *OwlGameTransactor) SetGlobalFruitRewardsProportion(opts *bind.TransactOpts, proportion *big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "setGlobalFruitRewardsProportion", proportion)
}

// SetGlobalFruitRewardsProportion is a paid mutator transaction binding the contract method 0x68a83a87.
//
// Solidity: function setGlobalFruitRewardsProportion(uint256 proportion) returns()
func (_OwlGame *OwlGameSession) SetGlobalFruitRewardsProportion(proportion *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.SetGlobalFruitRewardsProportion(&_OwlGame.TransactOpts, proportion)
}

// SetGlobalFruitRewardsProportion is a paid mutator transaction binding the contract method 0x68a83a87.
//
// Solidity: function setGlobalFruitRewardsProportion(uint256 proportion) returns()
func (_OwlGame *OwlGameTransactorSession) SetGlobalFruitRewardsProportion(proportion *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.SetGlobalFruitRewardsProportion(&_OwlGame.TransactOpts, proportion)
}

// SetMoonBoost is a paid mutator transaction binding the contract method 0xbb1cf85a.
//
// Solidity: function setMoonBoost(bool isEnable) returns()
func (_OwlGame *OwlGameTransactor) SetMoonBoost(opts *bind.TransactOpts, isEnable bool) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "setMoonBoost", isEnable)
}

// SetMoonBoost is a paid mutator transaction binding the contract method 0xbb1cf85a.
//
// Solidity: function setMoonBoost(bool isEnable) returns()
func (_OwlGame *OwlGameSession) SetMoonBoost(isEnable bool) (*types.Transaction, error) {
	return _OwlGame.Contract.SetMoonBoost(&_OwlGame.TransactOpts, isEnable)
}

// SetMoonBoost is a paid mutator transaction binding the contract method 0xbb1cf85a.
//
// Solidity: function setMoonBoost(bool isEnable) returns()
func (_OwlGame *OwlGameTransactorSession) SetMoonBoost(isEnable bool) (*types.Transaction, error) {
	return _OwlGame.Contract.SetMoonBoost(&_OwlGame.TransactOpts, isEnable)
}

// StakeMysteryBox is a paid mutator transaction binding the contract method 0x95898b00.
//
// Solidity: function stakeMysteryBox(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameTransactor) StakeMysteryBox(opts *bind.TransactOpts, tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "stakeMysteryBox", tokenIdList)
}

// StakeMysteryBox is a paid mutator transaction binding the contract method 0x95898b00.
//
// Solidity: function stakeMysteryBox(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameSession) StakeMysteryBox(tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.StakeMysteryBox(&_OwlGame.TransactOpts, tokenIdList)
}

// StakeMysteryBox is a paid mutator transaction binding the contract method 0x95898b00.
//
// Solidity: function stakeMysteryBox(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameTransactorSession) StakeMysteryBox(tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.StakeMysteryBox(&_OwlGame.TransactOpts, tokenIdList)
}

// StakeOwldinalNft is a paid mutator transaction binding the contract method 0x655a4c62.
//
// Solidity: function stakeOwldinalNft(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameTransactor) StakeOwldinalNft(opts *bind.TransactOpts, tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "stakeOwldinalNft", tokenIdList)
}

// StakeOwldinalNft is a paid mutator transaction binding the contract method 0x655a4c62.
//
// Solidity: function stakeOwldinalNft(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameSession) StakeOwldinalNft(tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.StakeOwldinalNft(&_OwlGame.TransactOpts, tokenIdList)
}

// StakeOwldinalNft is a paid mutator transaction binding the contract method 0x655a4c62.
//
// Solidity: function stakeOwldinalNft(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameTransactorSession) StakeOwldinalNft(tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.StakeOwldinalNft(&_OwlGame.TransactOpts, tokenIdList)
}

// UnstakeOwldinalNft is a paid mutator transaction binding the contract method 0xc2886a2f.
//
// Solidity: function unstakeOwldinalNft(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameTransactor) UnstakeOwldinalNft(opts *bind.TransactOpts, tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "unstakeOwldinalNft", tokenIdList)
}

// UnstakeOwldinalNft is a paid mutator transaction binding the contract method 0xc2886a2f.
//
// Solidity: function unstakeOwldinalNft(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameSession) UnstakeOwldinalNft(tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.UnstakeOwldinalNft(&_OwlGame.TransactOpts, tokenIdList)
}

// UnstakeOwldinalNft is a paid mutator transaction binding the contract method 0xc2886a2f.
//
// Solidity: function unstakeOwldinalNft(uint256[] tokenIdList) returns()
func (_OwlGame *OwlGameTransactorSession) UnstakeOwldinalNft(tokenIdList []*big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.UnstakeOwldinalNft(&_OwlGame.TransactOpts, tokenIdList)
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

// OwlGameBindInvitationIterator is returned from FilterBindInvitation and is used to iterate over the raw logs and unpacked data for BindInvitation events raised by the OwlGame contract.
type OwlGameBindInvitationIterator struct {
	Event *OwlGameBindInvitation // Event containing the contract specifics and raw log

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
func (it *OwlGameBindInvitationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameBindInvitation)
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
		it.Event = new(OwlGameBindInvitation)
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
func (it *OwlGameBindInvitationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameBindInvitationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameBindInvitation represents a BindInvitation event raised by the OwlGame contract.
type OwlGameBindInvitation struct {
	Invitee common.Address
	Inviter common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBindInvitation is a free log retrieval operation binding the contract event 0x3144fc7320adf238514c761c91814e301f210c0e4e0bb5f9fd88bec051b4f100.
//
// Solidity: event BindInvitation(address indexed invitee, address inviter)
func (_OwlGame *OwlGameFilterer) FilterBindInvitation(opts *bind.FilterOpts, invitee []common.Address) (*OwlGameBindInvitationIterator, error) {

	var inviteeRule []interface{}
	for _, inviteeItem := range invitee {
		inviteeRule = append(inviteeRule, inviteeItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "BindInvitation", inviteeRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameBindInvitationIterator{contract: _OwlGame.contract, event: "BindInvitation", logs: logs, sub: sub}, nil
}

// WatchBindInvitation is a free log subscription operation binding the contract event 0x3144fc7320adf238514c761c91814e301f210c0e4e0bb5f9fd88bec051b4f100.
//
// Solidity: event BindInvitation(address indexed invitee, address inviter)
func (_OwlGame *OwlGameFilterer) WatchBindInvitation(opts *bind.WatchOpts, sink chan<- *OwlGameBindInvitation, invitee []common.Address) (event.Subscription, error) {

	var inviteeRule []interface{}
	for _, inviteeItem := range invitee {
		inviteeRule = append(inviteeRule, inviteeItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "BindInvitation", inviteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameBindInvitation)
				if err := _OwlGame.contract.UnpackLog(event, "BindInvitation", log); err != nil {
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

// ParseBindInvitation is a log parse operation binding the contract event 0x3144fc7320adf238514c761c91814e301f210c0e4e0bb5f9fd88bec051b4f100.
//
// Solidity: event BindInvitation(address indexed invitee, address inviter)
func (_OwlGame *OwlGameFilterer) ParseBindInvitation(log types.Log) (*OwlGameBindInvitation, error) {
	event := new(OwlGameBindInvitation)
	if err := _OwlGame.contract.UnpackLog(event, "BindInvitation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameClaimInviterRewardIterator is returned from FilterClaimInviterReward and is used to iterate over the raw logs and unpacked data for ClaimInviterReward events raised by the OwlGame contract.
type OwlGameClaimInviterRewardIterator struct {
	Event *OwlGameClaimInviterReward // Event containing the contract specifics and raw log

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
func (it *OwlGameClaimInviterRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameClaimInviterReward)
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
		it.Event = new(OwlGameClaimInviterReward)
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
func (it *OwlGameClaimInviterRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameClaimInviterRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameClaimInviterReward represents a ClaimInviterReward event raised by the OwlGame contract.
type OwlGameClaimInviterReward struct {
	User           common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterClaimInviterReward is a free log retrieval operation binding the contract event 0x1b066a4c1595fdf3bf1c300bca29f51f237bbc15ee85e815b39befc1bc630cfc.
//
// Solidity: event ClaimInviterReward(address indexed user, uint256 withdrawAmount)
func (_OwlGame *OwlGameFilterer) FilterClaimInviterReward(opts *bind.FilterOpts, user []common.Address) (*OwlGameClaimInviterRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "ClaimInviterReward", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameClaimInviterRewardIterator{contract: _OwlGame.contract, event: "ClaimInviterReward", logs: logs, sub: sub}, nil
}

// WatchClaimInviterReward is a free log subscription operation binding the contract event 0x1b066a4c1595fdf3bf1c300bca29f51f237bbc15ee85e815b39befc1bc630cfc.
//
// Solidity: event ClaimInviterReward(address indexed user, uint256 withdrawAmount)
func (_OwlGame *OwlGameFilterer) WatchClaimInviterReward(opts *bind.WatchOpts, sink chan<- *OwlGameClaimInviterReward, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "ClaimInviterReward", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameClaimInviterReward)
				if err := _OwlGame.contract.UnpackLog(event, "ClaimInviterReward", log); err != nil {
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

// ParseClaimInviterReward is a log parse operation binding the contract event 0x1b066a4c1595fdf3bf1c300bca29f51f237bbc15ee85e815b39befc1bc630cfc.
//
// Solidity: event ClaimInviterReward(address indexed user, uint256 withdrawAmount)
func (_OwlGame *OwlGameFilterer) ParseClaimInviterReward(log types.Log) (*OwlGameClaimInviterReward, error) {
	event := new(OwlGameClaimInviterReward)
	if err := _OwlGame.contract.UnpackLog(event, "ClaimInviterReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameJoinGameIterator is returned from FilterJoinGame and is used to iterate over the raw logs and unpacked data for JoinGame events raised by the OwlGame contract.
type OwlGameJoinGameIterator struct {
	Event *OwlGameJoinGame // Event containing the contract specifics and raw log

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
func (it *OwlGameJoinGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameJoinGame)
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
		it.Event = new(OwlGameJoinGame)
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
func (it *OwlGameJoinGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameJoinGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameJoinGame represents a JoinGame event raised by the OwlGame contract.
type OwlGameJoinGame struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterJoinGame is a free log retrieval operation binding the contract event 0x9d148569af2a4ae8c34122247102efb7bb91bf1b595c37c539b852954707d482.
//
// Solidity: event JoinGame(address indexed user)
func (_OwlGame *OwlGameFilterer) FilterJoinGame(opts *bind.FilterOpts, user []common.Address) (*OwlGameJoinGameIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "JoinGame", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameJoinGameIterator{contract: _OwlGame.contract, event: "JoinGame", logs: logs, sub: sub}, nil
}

// WatchJoinGame is a free log subscription operation binding the contract event 0x9d148569af2a4ae8c34122247102efb7bb91bf1b595c37c539b852954707d482.
//
// Solidity: event JoinGame(address indexed user)
func (_OwlGame *OwlGameFilterer) WatchJoinGame(opts *bind.WatchOpts, sink chan<- *OwlGameJoinGame, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "JoinGame", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameJoinGame)
				if err := _OwlGame.contract.UnpackLog(event, "JoinGame", log); err != nil {
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

// ParseJoinGame is a log parse operation binding the contract event 0x9d148569af2a4ae8c34122247102efb7bb91bf1b595c37c539b852954707d482.
//
// Solidity: event JoinGame(address indexed user)
func (_OwlGame *OwlGameFilterer) ParseJoinGame(log types.Log) (*OwlGameJoinGame, error) {
	event := new(OwlGameJoinGame)
	if err := _OwlGame.contract.UnpackLog(event, "JoinGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGamePrizePoolDecreasedIterator is returned from FilterPrizePoolDecreased and is used to iterate over the raw logs and unpacked data for PrizePoolDecreased events raised by the OwlGame contract.
type OwlGamePrizePoolDecreasedIterator struct {
	Event *OwlGamePrizePoolDecreased // Event containing the contract specifics and raw log

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
func (it *OwlGamePrizePoolDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGamePrizePoolDecreased)
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
		it.Event = new(OwlGamePrizePoolDecreased)
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
func (it *OwlGamePrizePoolDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGamePrizePoolDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGamePrizePoolDecreased represents a PrizePoolDecreased event raised by the OwlGame contract.
type OwlGamePrizePoolDecreased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPrizePoolDecreased is a free log retrieval operation binding the contract event 0xf0870791f79c9e860000c2e3005d55ab8944ec1a18a6b61d4ad60ce184654f42.
//
// Solidity: event PrizePoolDecreased(uint256 amount)
func (_OwlGame *OwlGameFilterer) FilterPrizePoolDecreased(opts *bind.FilterOpts) (*OwlGamePrizePoolDecreasedIterator, error) {

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "PrizePoolDecreased")
	if err != nil {
		return nil, err
	}
	return &OwlGamePrizePoolDecreasedIterator{contract: _OwlGame.contract, event: "PrizePoolDecreased", logs: logs, sub: sub}, nil
}

// WatchPrizePoolDecreased is a free log subscription operation binding the contract event 0xf0870791f79c9e860000c2e3005d55ab8944ec1a18a6b61d4ad60ce184654f42.
//
// Solidity: event PrizePoolDecreased(uint256 amount)
func (_OwlGame *OwlGameFilterer) WatchPrizePoolDecreased(opts *bind.WatchOpts, sink chan<- *OwlGamePrizePoolDecreased) (event.Subscription, error) {

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "PrizePoolDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGamePrizePoolDecreased)
				if err := _OwlGame.contract.UnpackLog(event, "PrizePoolDecreased", log); err != nil {
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

// ParsePrizePoolDecreased is a log parse operation binding the contract event 0xf0870791f79c9e860000c2e3005d55ab8944ec1a18a6b61d4ad60ce184654f42.
//
// Solidity: event PrizePoolDecreased(uint256 amount)
func (_OwlGame *OwlGameFilterer) ParsePrizePoolDecreased(log types.Log) (*OwlGamePrizePoolDecreased, error) {
	event := new(OwlGamePrizePoolDecreased)
	if err := _OwlGame.contract.UnpackLog(event, "PrizePoolDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGamePrizePoolIncreasedIterator is returned from FilterPrizePoolIncreased and is used to iterate over the raw logs and unpacked data for PrizePoolIncreased events raised by the OwlGame contract.
type OwlGamePrizePoolIncreasedIterator struct {
	Event *OwlGamePrizePoolIncreased // Event containing the contract specifics and raw log

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
func (it *OwlGamePrizePoolIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGamePrizePoolIncreased)
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
		it.Event = new(OwlGamePrizePoolIncreased)
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
func (it *OwlGamePrizePoolIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGamePrizePoolIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGamePrizePoolIncreased represents a PrizePoolIncreased event raised by the OwlGame contract.
type OwlGamePrizePoolIncreased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPrizePoolIncreased is a free log retrieval operation binding the contract event 0x3c231f64b16483d16ca517ccd881f34f86dd04ccc1305e65b854fe82189b7625.
//
// Solidity: event PrizePoolIncreased(uint256 amount)
func (_OwlGame *OwlGameFilterer) FilterPrizePoolIncreased(opts *bind.FilterOpts) (*OwlGamePrizePoolIncreasedIterator, error) {

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "PrizePoolIncreased")
	if err != nil {
		return nil, err
	}
	return &OwlGamePrizePoolIncreasedIterator{contract: _OwlGame.contract, event: "PrizePoolIncreased", logs: logs, sub: sub}, nil
}

// WatchPrizePoolIncreased is a free log subscription operation binding the contract event 0x3c231f64b16483d16ca517ccd881f34f86dd04ccc1305e65b854fe82189b7625.
//
// Solidity: event PrizePoolIncreased(uint256 amount)
func (_OwlGame *OwlGameFilterer) WatchPrizePoolIncreased(opts *bind.WatchOpts, sink chan<- *OwlGamePrizePoolIncreased) (event.Subscription, error) {

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "PrizePoolIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGamePrizePoolIncreased)
				if err := _OwlGame.contract.UnpackLog(event, "PrizePoolIncreased", log); err != nil {
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

// ParsePrizePoolIncreased is a log parse operation binding the contract event 0x3c231f64b16483d16ca517ccd881f34f86dd04ccc1305e65b854fe82189b7625.
//
// Solidity: event PrizePoolIncreased(uint256 amount)
func (_OwlGame *OwlGameFilterer) ParsePrizePoolIncreased(log types.Log) (*OwlGamePrizePoolIncreased, error) {
	event := new(OwlGamePrizePoolIncreased)
	if err := _OwlGame.contract.UnpackLog(event, "PrizePoolIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// OwlGameStakeMysteryBoxIterator is returned from FilterStakeMysteryBox and is used to iterate over the raw logs and unpacked data for StakeMysteryBox events raised by the OwlGame contract.
type OwlGameStakeMysteryBoxIterator struct {
	Event *OwlGameStakeMysteryBox // Event containing the contract specifics and raw log

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
func (it *OwlGameStakeMysteryBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameStakeMysteryBox)
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
		it.Event = new(OwlGameStakeMysteryBox)
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
func (it *OwlGameStakeMysteryBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameStakeMysteryBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameStakeMysteryBox represents a StakeMysteryBox event raised by the OwlGame contract.
type OwlGameStakeMysteryBox struct {
	User    common.Address
	TokenId []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeMysteryBox is a free log retrieval operation binding the contract event 0x85ee4e66817b80797ec9b3286b977e5b1e7c0890c0341916ff9acaa6f8210b0d.
//
// Solidity: event StakeMysteryBox(address indexed user, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) FilterStakeMysteryBox(opts *bind.FilterOpts, user []common.Address) (*OwlGameStakeMysteryBoxIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "StakeMysteryBox", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameStakeMysteryBoxIterator{contract: _OwlGame.contract, event: "StakeMysteryBox", logs: logs, sub: sub}, nil
}

// WatchStakeMysteryBox is a free log subscription operation binding the contract event 0x85ee4e66817b80797ec9b3286b977e5b1e7c0890c0341916ff9acaa6f8210b0d.
//
// Solidity: event StakeMysteryBox(address indexed user, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) WatchStakeMysteryBox(opts *bind.WatchOpts, sink chan<- *OwlGameStakeMysteryBox, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "StakeMysteryBox", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameStakeMysteryBox)
				if err := _OwlGame.contract.UnpackLog(event, "StakeMysteryBox", log); err != nil {
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

// ParseStakeMysteryBox is a log parse operation binding the contract event 0x85ee4e66817b80797ec9b3286b977e5b1e7c0890c0341916ff9acaa6f8210b0d.
//
// Solidity: event StakeMysteryBox(address indexed user, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) ParseStakeMysteryBox(log types.Log) (*OwlGameStakeMysteryBox, error) {
	event := new(OwlGameStakeMysteryBox)
	if err := _OwlGame.contract.UnpackLog(event, "StakeMysteryBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameStakeOwldinalNftIterator is returned from FilterStakeOwldinalNft and is used to iterate over the raw logs and unpacked data for StakeOwldinalNft events raised by the OwlGame contract.
type OwlGameStakeOwldinalNftIterator struct {
	Event *OwlGameStakeOwldinalNft // Event containing the contract specifics and raw log

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
func (it *OwlGameStakeOwldinalNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameStakeOwldinalNft)
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
		it.Event = new(OwlGameStakeOwldinalNft)
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
func (it *OwlGameStakeOwldinalNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameStakeOwldinalNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameStakeOwldinalNft represents a StakeOwldinalNft event raised by the OwlGame contract.
type OwlGameStakeOwldinalNft struct {
	User    common.Address
	TokenId []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeOwldinalNft is a free log retrieval operation binding the contract event 0x292b69a5590aefdf5de5c9da21ea45b29afd0635e4d0c7d149d1d84be9224106.
//
// Solidity: event StakeOwldinalNft(address indexed user, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) FilterStakeOwldinalNft(opts *bind.FilterOpts, user []common.Address) (*OwlGameStakeOwldinalNftIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "StakeOwldinalNft", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameStakeOwldinalNftIterator{contract: _OwlGame.contract, event: "StakeOwldinalNft", logs: logs, sub: sub}, nil
}

// WatchStakeOwldinalNft is a free log subscription operation binding the contract event 0x292b69a5590aefdf5de5c9da21ea45b29afd0635e4d0c7d149d1d84be9224106.
//
// Solidity: event StakeOwldinalNft(address indexed user, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) WatchStakeOwldinalNft(opts *bind.WatchOpts, sink chan<- *OwlGameStakeOwldinalNft, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "StakeOwldinalNft", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameStakeOwldinalNft)
				if err := _OwlGame.contract.UnpackLog(event, "StakeOwldinalNft", log); err != nil {
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

// ParseStakeOwldinalNft is a log parse operation binding the contract event 0x292b69a5590aefdf5de5c9da21ea45b29afd0635e4d0c7d149d1d84be9224106.
//
// Solidity: event StakeOwldinalNft(address indexed user, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) ParseStakeOwldinalNft(log types.Log) (*OwlGameStakeOwldinalNft, error) {
	event := new(OwlGameStakeOwldinalNft)
	if err := _OwlGame.contract.UnpackLog(event, "StakeOwldinalNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameUnstakeMysteryBoxIterator is returned from FilterUnstakeMysteryBox and is used to iterate over the raw logs and unpacked data for UnstakeMysteryBox events raised by the OwlGame contract.
type OwlGameUnstakeMysteryBoxIterator struct {
	Event *OwlGameUnstakeMysteryBox // Event containing the contract specifics and raw log

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
func (it *OwlGameUnstakeMysteryBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameUnstakeMysteryBox)
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
		it.Event = new(OwlGameUnstakeMysteryBox)
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
func (it *OwlGameUnstakeMysteryBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameUnstakeMysteryBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameUnstakeMysteryBox represents a UnstakeMysteryBox event raised by the OwlGame contract.
type OwlGameUnstakeMysteryBox struct {
	User    common.Address
	TokenId *big.Int
	BoxType uint8
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeMysteryBox is a free log retrieval operation binding the contract event 0x363717deff436618c48fd125aa63e987b64c96ed5976f67cce76223bc8ab2a29.
//
// Solidity: event UnstakeMysteryBox(address indexed user, uint256 tokenId, uint8 boxType, uint256 rewards)
func (_OwlGame *OwlGameFilterer) FilterUnstakeMysteryBox(opts *bind.FilterOpts, user []common.Address) (*OwlGameUnstakeMysteryBoxIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "UnstakeMysteryBox", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameUnstakeMysteryBoxIterator{contract: _OwlGame.contract, event: "UnstakeMysteryBox", logs: logs, sub: sub}, nil
}

// WatchUnstakeMysteryBox is a free log subscription operation binding the contract event 0x363717deff436618c48fd125aa63e987b64c96ed5976f67cce76223bc8ab2a29.
//
// Solidity: event UnstakeMysteryBox(address indexed user, uint256 tokenId, uint8 boxType, uint256 rewards)
func (_OwlGame *OwlGameFilterer) WatchUnstakeMysteryBox(opts *bind.WatchOpts, sink chan<- *OwlGameUnstakeMysteryBox, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "UnstakeMysteryBox", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameUnstakeMysteryBox)
				if err := _OwlGame.contract.UnpackLog(event, "UnstakeMysteryBox", log); err != nil {
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

// ParseUnstakeMysteryBox is a log parse operation binding the contract event 0x363717deff436618c48fd125aa63e987b64c96ed5976f67cce76223bc8ab2a29.
//
// Solidity: event UnstakeMysteryBox(address indexed user, uint256 tokenId, uint8 boxType, uint256 rewards)
func (_OwlGame *OwlGameFilterer) ParseUnstakeMysteryBox(log types.Log) (*OwlGameUnstakeMysteryBox, error) {
	event := new(OwlGameUnstakeMysteryBox)
	if err := _OwlGame.contract.UnpackLog(event, "UnstakeMysteryBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameUnstakeOwldinalNftIterator is returned from FilterUnstakeOwldinalNft and is used to iterate over the raw logs and unpacked data for UnstakeOwldinalNft events raised by the OwlGame contract.
type OwlGameUnstakeOwldinalNftIterator struct {
	Event *OwlGameUnstakeOwldinalNft // Event containing the contract specifics and raw log

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
func (it *OwlGameUnstakeOwldinalNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameUnstakeOwldinalNft)
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
		it.Event = new(OwlGameUnstakeOwldinalNft)
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
func (it *OwlGameUnstakeOwldinalNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameUnstakeOwldinalNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameUnstakeOwldinalNft represents a UnstakeOwldinalNft event raised by the OwlGame contract.
type OwlGameUnstakeOwldinalNft struct {
	User    common.Address
	TokenId []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeOwldinalNft is a free log retrieval operation binding the contract event 0x953947fb8c8d5e1bd1fb1561e4a4077c9f02d4952ad1185d28133c0d764d9a5e.
//
// Solidity: event UnstakeOwldinalNft(address indexed user, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) FilterUnstakeOwldinalNft(opts *bind.FilterOpts, user []common.Address) (*OwlGameUnstakeOwldinalNftIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "UnstakeOwldinalNft", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameUnstakeOwldinalNftIterator{contract: _OwlGame.contract, event: "UnstakeOwldinalNft", logs: logs, sub: sub}, nil
}

// WatchUnstakeOwldinalNft is a free log subscription operation binding the contract event 0x953947fb8c8d5e1bd1fb1561e4a4077c9f02d4952ad1185d28133c0d764d9a5e.
//
// Solidity: event UnstakeOwldinalNft(address indexed user, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) WatchUnstakeOwldinalNft(opts *bind.WatchOpts, sink chan<- *OwlGameUnstakeOwldinalNft, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "UnstakeOwldinalNft", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameUnstakeOwldinalNft)
				if err := _OwlGame.contract.UnpackLog(event, "UnstakeOwldinalNft", log); err != nil {
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

// ParseUnstakeOwldinalNft is a log parse operation binding the contract event 0x953947fb8c8d5e1bd1fb1561e4a4077c9f02d4952ad1185d28133c0d764d9a5e.
//
// Solidity: event UnstakeOwldinalNft(address indexed user, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) ParseUnstakeOwldinalNft(log types.Log) (*OwlGameUnstakeOwldinalNft, error) {
	event := new(OwlGameUnstakeOwldinalNft)
	if err := _OwlGame.contract.UnpackLog(event, "UnstakeOwldinalNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
