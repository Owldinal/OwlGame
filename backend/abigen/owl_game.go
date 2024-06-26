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

// OwlGameTokenStakingInfo is an auto generated low-level Go binding around an user-defined struct.
type OwlGameTokenStakingInfo struct {
	TokenId     *big.Int
	Reward      *big.Int
	Owner       common.Address
	StakingTime uint64
	BoxType     uint8
	BuffLevel   uint16
}

// OwlGameMetaData contains all meta data concerning the OwlGame contract.
var OwlGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"server\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"invitee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"name\":\"BindInvitation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"ElfRewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFruitCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalElfCount\",\"type\":\"uint256\"}],\"name\":\"FruitRewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"inviteCode\",\"type\":\"uint32\"}],\"name\":\"JoinGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"}],\"name\":\"MintMysteryBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OwlTokenBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizePoolDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizePoolIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RebateClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"invitee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RebateRewardsIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"RequestMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"buffLevel\",\"type\":\"uint16\"}],\"name\":\"StakeMysteryBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"}],\"name\":\"StakeOwldinalNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockableRebateIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumBoxType\",\"name\":\"boxType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"UnstakeMysteryBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"}],\"name\":\"UnstakeOwldinalNft\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRUIT_REWARD_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_REBATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"}],\"name\":\"addPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"name\":\"claimAndUnstakeMysteryBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimInviterReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"name\":\"getInviteCode\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"inviteCode\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"stakingTime\",\"type\":\"uint64\"},{\"internalType\":\"enumBoxType\",\"name\":\"boxType\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"buffLevel\",\"type\":\"uint16\"}],\"internalType\":\"structOwlGame.TokenStakingInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"inviteCode\",\"type\":\"uint32\"}],\"name\":\"handleInviteCode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owlTokenAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owldinalNftAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mysteryBoxAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"inviteCodeToInviterMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviteeToInviterMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviterRebateMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRebateEarned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebatePendingWithdrawal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedRebateToClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintedBoxCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviterToInviteCodeMap\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inviterToInviteesMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMoonBoostEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintRequestId\",\"type\":\"uint256\"}],\"name\":\"mintMysteryBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mysteryBoxContract\",\"outputs\":[{\"internalType\":\"contractOwldinalGenOneBox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owlToken\",\"outputs\":[{\"internalType\":\"contractERC20Burnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owldinalNftContract\",\"outputs\":[{\"internalType\":\"contractOwldinal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"requestMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestMintMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proportion\",\"type\":\"uint256\"}],\"name\":\"setGlobalFruitRewardsProportion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isEnable\",\"type\":\"bool\"}],\"name\":\"setMoonBoost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"server\",\"type\":\"address\"}],\"name\":\"setServer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"name\":\"stakeMysteryBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"name\":\"stakeOwldinalNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenInfoMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"stakingTime\",\"type\":\"uint64\"},{\"internalType\":\"enumBoxType\",\"name\":\"boxType\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"buffLevel\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdList\",\"type\":\"uint256[]\"}],\"name\":\"unstakeOwldinalNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAllFruitRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetInviteCode is a free data retrieval call binding the contract method 0x9bd86c34.
//
// Solidity: function getInviteCode(address inviter) view returns(uint32 inviteCode)
func (_OwlGame *OwlGameCaller) GetInviteCode(opts *bind.CallOpts, inviter common.Address) (uint32, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "getInviteCode", inviter)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetInviteCode is a free data retrieval call binding the contract method 0x9bd86c34.
//
// Solidity: function getInviteCode(address inviter) view returns(uint32 inviteCode)
func (_OwlGame *OwlGameSession) GetInviteCode(inviter common.Address) (uint32, error) {
	return _OwlGame.Contract.GetInviteCode(&_OwlGame.CallOpts, inviter)
}

// GetInviteCode is a free data retrieval call binding the contract method 0x9bd86c34.
//
// Solidity: function getInviteCode(address inviter) view returns(uint32 inviteCode)
func (_OwlGame *OwlGameCallerSession) GetInviteCode(inviter common.Address) (uint32, error) {
	return _OwlGame.Contract.GetInviteCode(&_OwlGame.CallOpts, inviter)
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

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 tokenId) view returns((uint256,uint256,address,uint64,uint8,uint16))
func (_OwlGame *OwlGameCaller) GetTokenInfo(opts *bind.CallOpts, tokenId *big.Int) (OwlGameTokenStakingInfo, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "getTokenInfo", tokenId)

	if err != nil {
		return *new(OwlGameTokenStakingInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(OwlGameTokenStakingInfo)).(*OwlGameTokenStakingInfo)

	return out0, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 tokenId) view returns((uint256,uint256,address,uint64,uint8,uint16))
func (_OwlGame *OwlGameSession) GetTokenInfo(tokenId *big.Int) (OwlGameTokenStakingInfo, error) {
	return _OwlGame.Contract.GetTokenInfo(&_OwlGame.CallOpts, tokenId)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 tokenId) view returns((uint256,uint256,address,uint64,uint8,uint16))
func (_OwlGame *OwlGameCallerSession) GetTokenInfo(tokenId *big.Int) (OwlGameTokenStakingInfo, error) {
	return _OwlGame.Contract.GetTokenInfo(&_OwlGame.CallOpts, tokenId)
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

// InviteCodeToInviterMap is a free data retrieval call binding the contract method 0x45d42adb.
//
// Solidity: function inviteCodeToInviterMap(uint32 ) view returns(address)
func (_OwlGame *OwlGameCaller) InviteCodeToInviterMap(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "inviteCodeToInviterMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InviteCodeToInviterMap is a free data retrieval call binding the contract method 0x45d42adb.
//
// Solidity: function inviteCodeToInviterMap(uint32 ) view returns(address)
func (_OwlGame *OwlGameSession) InviteCodeToInviterMap(arg0 uint32) (common.Address, error) {
	return _OwlGame.Contract.InviteCodeToInviterMap(&_OwlGame.CallOpts, arg0)
}

// InviteCodeToInviterMap is a free data retrieval call binding the contract method 0x45d42adb.
//
// Solidity: function inviteCodeToInviterMap(uint32 ) view returns(address)
func (_OwlGame *OwlGameCallerSession) InviteCodeToInviterMap(arg0 uint32) (common.Address, error) {
	return _OwlGame.Contract.InviteCodeToInviterMap(&_OwlGame.CallOpts, arg0)
}

// InviteeToInviterMap is a free data retrieval call binding the contract method 0xbfb6b3ac.
//
// Solidity: function inviteeToInviterMap(address ) view returns(address)
func (_OwlGame *OwlGameCaller) InviteeToInviterMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "inviteeToInviterMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InviteeToInviterMap is a free data retrieval call binding the contract method 0xbfb6b3ac.
//
// Solidity: function inviteeToInviterMap(address ) view returns(address)
func (_OwlGame *OwlGameSession) InviteeToInviterMap(arg0 common.Address) (common.Address, error) {
	return _OwlGame.Contract.InviteeToInviterMap(&_OwlGame.CallOpts, arg0)
}

// InviteeToInviterMap is a free data retrieval call binding the contract method 0xbfb6b3ac.
//
// Solidity: function inviteeToInviterMap(address ) view returns(address)
func (_OwlGame *OwlGameCallerSession) InviteeToInviterMap(arg0 common.Address) (common.Address, error) {
	return _OwlGame.Contract.InviteeToInviterMap(&_OwlGame.CallOpts, arg0)
}

// InviterRebateMap is a free data retrieval call binding the contract method 0xf47893cf.
//
// Solidity: function inviterRebateMap(address ) view returns(uint256 totalRebateEarned, uint256 rebatePendingWithdrawal, uint256 unlockedRebateToClaim, uint256 mintedBoxCount)
func (_OwlGame *OwlGameCaller) InviterRebateMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalRebateEarned       *big.Int
	RebatePendingWithdrawal *big.Int
	UnlockedRebateToClaim   *big.Int
	MintedBoxCount          *big.Int
}, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "inviterRebateMap", arg0)

	outstruct := new(struct {
		TotalRebateEarned       *big.Int
		RebatePendingWithdrawal *big.Int
		UnlockedRebateToClaim   *big.Int
		MintedBoxCount          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalRebateEarned = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RebatePendingWithdrawal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnlockedRebateToClaim = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MintedBoxCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// InviterRebateMap is a free data retrieval call binding the contract method 0xf47893cf.
//
// Solidity: function inviterRebateMap(address ) view returns(uint256 totalRebateEarned, uint256 rebatePendingWithdrawal, uint256 unlockedRebateToClaim, uint256 mintedBoxCount)
func (_OwlGame *OwlGameSession) InviterRebateMap(arg0 common.Address) (struct {
	TotalRebateEarned       *big.Int
	RebatePendingWithdrawal *big.Int
	UnlockedRebateToClaim   *big.Int
	MintedBoxCount          *big.Int
}, error) {
	return _OwlGame.Contract.InviterRebateMap(&_OwlGame.CallOpts, arg0)
}

// InviterRebateMap is a free data retrieval call binding the contract method 0xf47893cf.
//
// Solidity: function inviterRebateMap(address ) view returns(uint256 totalRebateEarned, uint256 rebatePendingWithdrawal, uint256 unlockedRebateToClaim, uint256 mintedBoxCount)
func (_OwlGame *OwlGameCallerSession) InviterRebateMap(arg0 common.Address) (struct {
	TotalRebateEarned       *big.Int
	RebatePendingWithdrawal *big.Int
	UnlockedRebateToClaim   *big.Int
	MintedBoxCount          *big.Int
}, error) {
	return _OwlGame.Contract.InviterRebateMap(&_OwlGame.CallOpts, arg0)
}

// InviterToInviteCodeMap is a free data retrieval call binding the contract method 0x041586b6.
//
// Solidity: function inviterToInviteCodeMap(address ) view returns(uint32)
func (_OwlGame *OwlGameCaller) InviterToInviteCodeMap(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "inviterToInviteCodeMap", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// InviterToInviteCodeMap is a free data retrieval call binding the contract method 0x041586b6.
//
// Solidity: function inviterToInviteCodeMap(address ) view returns(uint32)
func (_OwlGame *OwlGameSession) InviterToInviteCodeMap(arg0 common.Address) (uint32, error) {
	return _OwlGame.Contract.InviterToInviteCodeMap(&_OwlGame.CallOpts, arg0)
}

// InviterToInviteCodeMap is a free data retrieval call binding the contract method 0x041586b6.
//
// Solidity: function inviterToInviteCodeMap(address ) view returns(uint32)
func (_OwlGame *OwlGameCallerSession) InviterToInviteCodeMap(arg0 common.Address) (uint32, error) {
	return _OwlGame.Contract.InviterToInviteCodeMap(&_OwlGame.CallOpts, arg0)
}

// InviterToInviteesMap is a free data retrieval call binding the contract method 0x8b659d9c.
//
// Solidity: function inviterToInviteesMap(address , uint256 ) view returns(address)
func (_OwlGame *OwlGameCaller) InviterToInviteesMap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "inviterToInviteesMap", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InviterToInviteesMap is a free data retrieval call binding the contract method 0x8b659d9c.
//
// Solidity: function inviterToInviteesMap(address , uint256 ) view returns(address)
func (_OwlGame *OwlGameSession) InviterToInviteesMap(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _OwlGame.Contract.InviterToInviteesMap(&_OwlGame.CallOpts, arg0, arg1)
}

// InviterToInviteesMap is a free data retrieval call binding the contract method 0x8b659d9c.
//
// Solidity: function inviterToInviteesMap(address , uint256 ) view returns(address)
func (_OwlGame *OwlGameCallerSession) InviterToInviteesMap(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _OwlGame.Contract.InviterToInviteesMap(&_OwlGame.CallOpts, arg0, arg1)
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

// RequestMintMap is a free data retrieval call binding the contract method 0x84f49897.
//
// Solidity: function requestMintMap(uint256 ) view returns(address user, uint256 count)
func (_OwlGame *OwlGameCaller) RequestMintMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User  common.Address
	Count *big.Int
}, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "requestMintMap", arg0)

	outstruct := new(struct {
		User  common.Address
		Count *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Count = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RequestMintMap is a free data retrieval call binding the contract method 0x84f49897.
//
// Solidity: function requestMintMap(uint256 ) view returns(address user, uint256 count)
func (_OwlGame *OwlGameSession) RequestMintMap(arg0 *big.Int) (struct {
	User  common.Address
	Count *big.Int
}, error) {
	return _OwlGame.Contract.RequestMintMap(&_OwlGame.CallOpts, arg0)
}

// RequestMintMap is a free data retrieval call binding the contract method 0x84f49897.
//
// Solidity: function requestMintMap(uint256 ) view returns(address user, uint256 count)
func (_OwlGame *OwlGameCallerSession) RequestMintMap(arg0 *big.Int) (struct {
	User  common.Address
	Count *big.Int
}, error) {
	return _OwlGame.Contract.RequestMintMap(&_OwlGame.CallOpts, arg0)
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

// TokenInfoMap is a free data retrieval call binding the contract method 0x6fa700bf.
//
// Solidity: function tokenInfoMap(uint256 ) view returns(uint256 tokenId, uint256 reward, address owner, uint64 stakingTime, uint8 boxType, uint16 buffLevel)
func (_OwlGame *OwlGameCaller) TokenInfoMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId     *big.Int
	Reward      *big.Int
	Owner       common.Address
	StakingTime uint64
	BoxType     uint8
	BuffLevel   uint16
}, error) {
	var out []interface{}
	err := _OwlGame.contract.Call(opts, &out, "tokenInfoMap", arg0)

	outstruct := new(struct {
		TokenId     *big.Int
		Reward      *big.Int
		Owner       common.Address
		StakingTime uint64
		BoxType     uint8
		BuffLevel   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.StakingTime = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.BoxType = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.BuffLevel = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// TokenInfoMap is a free data retrieval call binding the contract method 0x6fa700bf.
//
// Solidity: function tokenInfoMap(uint256 ) view returns(uint256 tokenId, uint256 reward, address owner, uint64 stakingTime, uint8 boxType, uint16 buffLevel)
func (_OwlGame *OwlGameSession) TokenInfoMap(arg0 *big.Int) (struct {
	TokenId     *big.Int
	Reward      *big.Int
	Owner       common.Address
	StakingTime uint64
	BoxType     uint8
	BuffLevel   uint16
}, error) {
	return _OwlGame.Contract.TokenInfoMap(&_OwlGame.CallOpts, arg0)
}

// TokenInfoMap is a free data retrieval call binding the contract method 0x6fa700bf.
//
// Solidity: function tokenInfoMap(uint256 ) view returns(uint256 tokenId, uint256 reward, address owner, uint64 stakingTime, uint8 boxType, uint16 buffLevel)
func (_OwlGame *OwlGameCallerSession) TokenInfoMap(arg0 *big.Int) (struct {
	TokenId     *big.Int
	Reward      *big.Int
	Owner       common.Address
	StakingTime uint64
	BoxType     uint8
	BuffLevel   uint16
}, error) {
	return _OwlGame.Contract.TokenInfoMap(&_OwlGame.CallOpts, arg0)
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
// Solidity: function mintMysteryBox(uint256 mintRequestId) returns()
func (_OwlGame *OwlGameTransactor) MintMysteryBox(opts *bind.TransactOpts, mintRequestId *big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "mintMysteryBox", mintRequestId)
}

// MintMysteryBox is a paid mutator transaction binding the contract method 0x896b4d40.
//
// Solidity: function mintMysteryBox(uint256 mintRequestId) returns()
func (_OwlGame *OwlGameSession) MintMysteryBox(mintRequestId *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.MintMysteryBox(&_OwlGame.TransactOpts, mintRequestId)
}

// MintMysteryBox is a paid mutator transaction binding the contract method 0x896b4d40.
//
// Solidity: function mintMysteryBox(uint256 mintRequestId) returns()
func (_OwlGame *OwlGameTransactorSession) MintMysteryBox(mintRequestId *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.MintMysteryBox(&_OwlGame.TransactOpts, mintRequestId)
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

// RequestMint is a paid mutator transaction binding the contract method 0x49733d04.
//
// Solidity: function requestMint(uint256 count) returns()
func (_OwlGame *OwlGameTransactor) RequestMint(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "requestMint", count)
}

// RequestMint is a paid mutator transaction binding the contract method 0x49733d04.
//
// Solidity: function requestMint(uint256 count) returns()
func (_OwlGame *OwlGameSession) RequestMint(count *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.RequestMint(&_OwlGame.TransactOpts, count)
}

// RequestMint is a paid mutator transaction binding the contract method 0x49733d04.
//
// Solidity: function requestMint(uint256 count) returns()
func (_OwlGame *OwlGameTransactorSession) RequestMint(count *big.Int) (*types.Transaction, error) {
	return _OwlGame.Contract.RequestMint(&_OwlGame.TransactOpts, count)
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

// SetServer is a paid mutator transaction binding the contract method 0xf19a0b54.
//
// Solidity: function setServer(address server) returns()
func (_OwlGame *OwlGameTransactor) SetServer(opts *bind.TransactOpts, server common.Address) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "setServer", server)
}

// SetServer is a paid mutator transaction binding the contract method 0xf19a0b54.
//
// Solidity: function setServer(address server) returns()
func (_OwlGame *OwlGameSession) SetServer(server common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.SetServer(&_OwlGame.TransactOpts, server)
}

// SetServer is a paid mutator transaction binding the contract method 0xf19a0b54.
//
// Solidity: function setServer(address server) returns()
func (_OwlGame *OwlGameTransactorSession) SetServer(server common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.SetServer(&_OwlGame.TransactOpts, server)
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

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns(uint256)
func (_OwlGame *OwlGameTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _OwlGame.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns(uint256)
func (_OwlGame *OwlGameSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.Withdraw(&_OwlGame.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns(uint256)
func (_OwlGame *OwlGameTransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _OwlGame.Contract.Withdraw(&_OwlGame.TransactOpts, recipient)
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

// OwlGameElfRewardUpdatedIterator is returned from FilterElfRewardUpdated and is used to iterate over the raw logs and unpacked data for ElfRewardUpdated events raised by the OwlGame contract.
type OwlGameElfRewardUpdatedIterator struct {
	Event *OwlGameElfRewardUpdated // Event containing the contract specifics and raw log

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
func (it *OwlGameElfRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameElfRewardUpdated)
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
		it.Event = new(OwlGameElfRewardUpdated)
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
func (it *OwlGameElfRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameElfRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameElfRewardUpdated represents a ElfRewardUpdated event raised by the OwlGame contract.
type OwlGameElfRewardUpdated struct {
	Amount *big.Int
	Count  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterElfRewardUpdated is a free log retrieval operation binding the contract event 0x9b964cbd458fdb8541cd6349fd809bc91ecc1a53602693a74fd37f5d3916363c.
//
// Solidity: event ElfRewardUpdated(uint256 amount, uint256 count)
func (_OwlGame *OwlGameFilterer) FilterElfRewardUpdated(opts *bind.FilterOpts) (*OwlGameElfRewardUpdatedIterator, error) {

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "ElfRewardUpdated")
	if err != nil {
		return nil, err
	}
	return &OwlGameElfRewardUpdatedIterator{contract: _OwlGame.contract, event: "ElfRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchElfRewardUpdated is a free log subscription operation binding the contract event 0x9b964cbd458fdb8541cd6349fd809bc91ecc1a53602693a74fd37f5d3916363c.
//
// Solidity: event ElfRewardUpdated(uint256 amount, uint256 count)
func (_OwlGame *OwlGameFilterer) WatchElfRewardUpdated(opts *bind.WatchOpts, sink chan<- *OwlGameElfRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "ElfRewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameElfRewardUpdated)
				if err := _OwlGame.contract.UnpackLog(event, "ElfRewardUpdated", log); err != nil {
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

// ParseElfRewardUpdated is a log parse operation binding the contract event 0x9b964cbd458fdb8541cd6349fd809bc91ecc1a53602693a74fd37f5d3916363c.
//
// Solidity: event ElfRewardUpdated(uint256 amount, uint256 count)
func (_OwlGame *OwlGameFilterer) ParseElfRewardUpdated(log types.Log) (*OwlGameElfRewardUpdated, error) {
	event := new(OwlGameElfRewardUpdated)
	if err := _OwlGame.contract.UnpackLog(event, "ElfRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameFruitRewardUpdatedIterator is returned from FilterFruitRewardUpdated and is used to iterate over the raw logs and unpacked data for FruitRewardUpdated events raised by the OwlGame contract.
type OwlGameFruitRewardUpdatedIterator struct {
	Event *OwlGameFruitRewardUpdated // Event containing the contract specifics and raw log

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
func (it *OwlGameFruitRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameFruitRewardUpdated)
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
		it.Event = new(OwlGameFruitRewardUpdated)
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
func (it *OwlGameFruitRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameFruitRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameFruitRewardUpdated represents a FruitRewardUpdated event raised by the OwlGame contract.
type OwlGameFruitRewardUpdated struct {
	Amount          *big.Int
	Count           *big.Int
	TotalFruitCount *big.Int
	TotalElfCount   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFruitRewardUpdated is a free log retrieval operation binding the contract event 0x77cbc09843659e18ed105aec15e30ade634e6ca7b6d13e5ac92bb82cea91f8ee.
//
// Solidity: event FruitRewardUpdated(uint256 amount, uint256 count, uint256 totalFruitCount, uint256 totalElfCount)
func (_OwlGame *OwlGameFilterer) FilterFruitRewardUpdated(opts *bind.FilterOpts) (*OwlGameFruitRewardUpdatedIterator, error) {

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "FruitRewardUpdated")
	if err != nil {
		return nil, err
	}
	return &OwlGameFruitRewardUpdatedIterator{contract: _OwlGame.contract, event: "FruitRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchFruitRewardUpdated is a free log subscription operation binding the contract event 0x77cbc09843659e18ed105aec15e30ade634e6ca7b6d13e5ac92bb82cea91f8ee.
//
// Solidity: event FruitRewardUpdated(uint256 amount, uint256 count, uint256 totalFruitCount, uint256 totalElfCount)
func (_OwlGame *OwlGameFilterer) WatchFruitRewardUpdated(opts *bind.WatchOpts, sink chan<- *OwlGameFruitRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "FruitRewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameFruitRewardUpdated)
				if err := _OwlGame.contract.UnpackLog(event, "FruitRewardUpdated", log); err != nil {
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

// ParseFruitRewardUpdated is a log parse operation binding the contract event 0x77cbc09843659e18ed105aec15e30ade634e6ca7b6d13e5ac92bb82cea91f8ee.
//
// Solidity: event FruitRewardUpdated(uint256 amount, uint256 count, uint256 totalFruitCount, uint256 totalElfCount)
func (_OwlGame *OwlGameFilterer) ParseFruitRewardUpdated(log types.Log) (*OwlGameFruitRewardUpdated, error) {
	event := new(OwlGameFruitRewardUpdated)
	if err := _OwlGame.contract.UnpackLog(event, "FruitRewardUpdated", log); err != nil {
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
	User       common.Address
	InviteCode uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterJoinGame is a free log retrieval operation binding the contract event 0x3518b1830e6ec9f510e24a95f032e27013c745334b813b19f15405c90923f4bc.
//
// Solidity: event JoinGame(address indexed user, uint32 inviteCode)
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

// WatchJoinGame is a free log subscription operation binding the contract event 0x3518b1830e6ec9f510e24a95f032e27013c745334b813b19f15405c90923f4bc.
//
// Solidity: event JoinGame(address indexed user, uint32 inviteCode)
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

// ParseJoinGame is a log parse operation binding the contract event 0x3518b1830e6ec9f510e24a95f032e27013c745334b813b19f15405c90923f4bc.
//
// Solidity: event JoinGame(address indexed user, uint32 inviteCode)
func (_OwlGame *OwlGameFilterer) ParseJoinGame(log types.Log) (*OwlGameJoinGame, error) {
	event := new(OwlGameJoinGame)
	if err := _OwlGame.contract.UnpackLog(event, "JoinGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameMintMysteryBoxIterator is returned from FilterMintMysteryBox and is used to iterate over the raw logs and unpacked data for MintMysteryBox events raised by the OwlGame contract.
type OwlGameMintMysteryBoxIterator struct {
	Event *OwlGameMintMysteryBox // Event containing the contract specifics and raw log

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
func (it *OwlGameMintMysteryBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameMintMysteryBox)
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
		it.Event = new(OwlGameMintMysteryBox)
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
func (it *OwlGameMintMysteryBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameMintMysteryBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameMintMysteryBox represents a MintMysteryBox event raised by the OwlGame contract.
type OwlGameMintMysteryBox struct {
	User      common.Address
	RequestId *big.Int
	Count     *big.Int
	TokenId   []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintMysteryBox is a free log retrieval operation binding the contract event 0xe8d0b37cb79abcf29fd9f6b0cb6005971fd5480295fc5a1f3b25f90c5faf727e.
//
// Solidity: event MintMysteryBox(address indexed user, uint256 requestId, uint256 count, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) FilterMintMysteryBox(opts *bind.FilterOpts, user []common.Address) (*OwlGameMintMysteryBoxIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "MintMysteryBox", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameMintMysteryBoxIterator{contract: _OwlGame.contract, event: "MintMysteryBox", logs: logs, sub: sub}, nil
}

// WatchMintMysteryBox is a free log subscription operation binding the contract event 0xe8d0b37cb79abcf29fd9f6b0cb6005971fd5480295fc5a1f3b25f90c5faf727e.
//
// Solidity: event MintMysteryBox(address indexed user, uint256 requestId, uint256 count, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) WatchMintMysteryBox(opts *bind.WatchOpts, sink chan<- *OwlGameMintMysteryBox, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "MintMysteryBox", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameMintMysteryBox)
				if err := _OwlGame.contract.UnpackLog(event, "MintMysteryBox", log); err != nil {
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

// ParseMintMysteryBox is a log parse operation binding the contract event 0xe8d0b37cb79abcf29fd9f6b0cb6005971fd5480295fc5a1f3b25f90c5faf727e.
//
// Solidity: event MintMysteryBox(address indexed user, uint256 requestId, uint256 count, uint256[] tokenId)
func (_OwlGame *OwlGameFilterer) ParseMintMysteryBox(log types.Log) (*OwlGameMintMysteryBox, error) {
	event := new(OwlGameMintMysteryBox)
	if err := _OwlGame.contract.UnpackLog(event, "MintMysteryBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameOwlTokenBurnedIterator is returned from FilterOwlTokenBurned and is used to iterate over the raw logs and unpacked data for OwlTokenBurned events raised by the OwlGame contract.
type OwlGameOwlTokenBurnedIterator struct {
	Event *OwlGameOwlTokenBurned // Event containing the contract specifics and raw log

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
func (it *OwlGameOwlTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameOwlTokenBurned)
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
		it.Event = new(OwlGameOwlTokenBurned)
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
func (it *OwlGameOwlTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameOwlTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameOwlTokenBurned represents a OwlTokenBurned event raised by the OwlGame contract.
type OwlGameOwlTokenBurned struct {
	User      common.Address
	MintCount *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwlTokenBurned is a free log retrieval operation binding the contract event 0x818c9d2e3949d2ab3a0c76d31b137496e1b14d58e2b2d746301c2926a31c435d.
//
// Solidity: event OwlTokenBurned(address user, uint256 mintCount, uint256 amount)
func (_OwlGame *OwlGameFilterer) FilterOwlTokenBurned(opts *bind.FilterOpts) (*OwlGameOwlTokenBurnedIterator, error) {

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "OwlTokenBurned")
	if err != nil {
		return nil, err
	}
	return &OwlGameOwlTokenBurnedIterator{contract: _OwlGame.contract, event: "OwlTokenBurned", logs: logs, sub: sub}, nil
}

// WatchOwlTokenBurned is a free log subscription operation binding the contract event 0x818c9d2e3949d2ab3a0c76d31b137496e1b14d58e2b2d746301c2926a31c435d.
//
// Solidity: event OwlTokenBurned(address user, uint256 mintCount, uint256 amount)
func (_OwlGame *OwlGameFilterer) WatchOwlTokenBurned(opts *bind.WatchOpts, sink chan<- *OwlGameOwlTokenBurned) (event.Subscription, error) {

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "OwlTokenBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameOwlTokenBurned)
				if err := _OwlGame.contract.UnpackLog(event, "OwlTokenBurned", log); err != nil {
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

// ParseOwlTokenBurned is a log parse operation binding the contract event 0x818c9d2e3949d2ab3a0c76d31b137496e1b14d58e2b2d746301c2926a31c435d.
//
// Solidity: event OwlTokenBurned(address user, uint256 mintCount, uint256 amount)
func (_OwlGame *OwlGameFilterer) ParseOwlTokenBurned(log types.Log) (*OwlGameOwlTokenBurned, error) {
	event := new(OwlGameOwlTokenBurned)
	if err := _OwlGame.contract.UnpackLog(event, "OwlTokenBurned", log); err != nil {
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

// OwlGameRebateClaimedIterator is returned from FilterRebateClaimed and is used to iterate over the raw logs and unpacked data for RebateClaimed events raised by the OwlGame contract.
type OwlGameRebateClaimedIterator struct {
	Event *OwlGameRebateClaimed // Event containing the contract specifics and raw log

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
func (it *OwlGameRebateClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameRebateClaimed)
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
		it.Event = new(OwlGameRebateClaimed)
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
func (it *OwlGameRebateClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameRebateClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameRebateClaimed represents a RebateClaimed event raised by the OwlGame contract.
type OwlGameRebateClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRebateClaimed is a free log retrieval operation binding the contract event 0x8ec2a9ed236322c625516e43eab59fcb8145e38ee8d489a28d0aacaeecc298d6.
//
// Solidity: event RebateClaimed(address indexed user, uint256 amount)
func (_OwlGame *OwlGameFilterer) FilterRebateClaimed(opts *bind.FilterOpts, user []common.Address) (*OwlGameRebateClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "RebateClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameRebateClaimedIterator{contract: _OwlGame.contract, event: "RebateClaimed", logs: logs, sub: sub}, nil
}

// WatchRebateClaimed is a free log subscription operation binding the contract event 0x8ec2a9ed236322c625516e43eab59fcb8145e38ee8d489a28d0aacaeecc298d6.
//
// Solidity: event RebateClaimed(address indexed user, uint256 amount)
func (_OwlGame *OwlGameFilterer) WatchRebateClaimed(opts *bind.WatchOpts, sink chan<- *OwlGameRebateClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "RebateClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameRebateClaimed)
				if err := _OwlGame.contract.UnpackLog(event, "RebateClaimed", log); err != nil {
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

// ParseRebateClaimed is a log parse operation binding the contract event 0x8ec2a9ed236322c625516e43eab59fcb8145e38ee8d489a28d0aacaeecc298d6.
//
// Solidity: event RebateClaimed(address indexed user, uint256 amount)
func (_OwlGame *OwlGameFilterer) ParseRebateClaimed(log types.Log) (*OwlGameRebateClaimed, error) {
	event := new(OwlGameRebateClaimed)
	if err := _OwlGame.contract.UnpackLog(event, "RebateClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameRebateRewardsIncreasedIterator is returned from FilterRebateRewardsIncreased and is used to iterate over the raw logs and unpacked data for RebateRewardsIncreased events raised by the OwlGame contract.
type OwlGameRebateRewardsIncreasedIterator struct {
	Event *OwlGameRebateRewardsIncreased // Event containing the contract specifics and raw log

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
func (it *OwlGameRebateRewardsIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameRebateRewardsIncreased)
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
		it.Event = new(OwlGameRebateRewardsIncreased)
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
func (it *OwlGameRebateRewardsIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameRebateRewardsIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameRebateRewardsIncreased represents a RebateRewardsIncreased event raised by the OwlGame contract.
type OwlGameRebateRewardsIncreased struct {
	User      common.Address
	Invitee   common.Address
	MintCount *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRebateRewardsIncreased is a free log retrieval operation binding the contract event 0x07e9fcef60e55f4c1c03cfeb57d85db49ecae00bc0367f586d1d67750dcebbe6.
//
// Solidity: event RebateRewardsIncreased(address indexed user, address invitee, uint256 mintCount, uint256 amount)
func (_OwlGame *OwlGameFilterer) FilterRebateRewardsIncreased(opts *bind.FilterOpts, user []common.Address) (*OwlGameRebateRewardsIncreasedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "RebateRewardsIncreased", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameRebateRewardsIncreasedIterator{contract: _OwlGame.contract, event: "RebateRewardsIncreased", logs: logs, sub: sub}, nil
}

// WatchRebateRewardsIncreased is a free log subscription operation binding the contract event 0x07e9fcef60e55f4c1c03cfeb57d85db49ecae00bc0367f586d1d67750dcebbe6.
//
// Solidity: event RebateRewardsIncreased(address indexed user, address invitee, uint256 mintCount, uint256 amount)
func (_OwlGame *OwlGameFilterer) WatchRebateRewardsIncreased(opts *bind.WatchOpts, sink chan<- *OwlGameRebateRewardsIncreased, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "RebateRewardsIncreased", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameRebateRewardsIncreased)
				if err := _OwlGame.contract.UnpackLog(event, "RebateRewardsIncreased", log); err != nil {
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

// ParseRebateRewardsIncreased is a log parse operation binding the contract event 0x07e9fcef60e55f4c1c03cfeb57d85db49ecae00bc0367f586d1d67750dcebbe6.
//
// Solidity: event RebateRewardsIncreased(address indexed user, address invitee, uint256 mintCount, uint256 amount)
func (_OwlGame *OwlGameFilterer) ParseRebateRewardsIncreased(log types.Log) (*OwlGameRebateRewardsIncreased, error) {
	event := new(OwlGameRebateRewardsIncreased)
	if err := _OwlGame.contract.UnpackLog(event, "RebateRewardsIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwlGameRequestMintIterator is returned from FilterRequestMint and is used to iterate over the raw logs and unpacked data for RequestMint events raised by the OwlGame contract.
type OwlGameRequestMintIterator struct {
	Event *OwlGameRequestMint // Event containing the contract specifics and raw log

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
func (it *OwlGameRequestMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameRequestMint)
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
		it.Event = new(OwlGameRequestMint)
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
func (it *OwlGameRequestMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameRequestMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameRequestMint represents a RequestMint event raised by the OwlGame contract.
type OwlGameRequestMint struct {
	User      common.Address
	RequestId *big.Int
	Count     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestMint is a free log retrieval operation binding the contract event 0x79ab85c2314dba73749962c93d6920c946a5f6f0532f5e675fd7b1ff530ee058.
//
// Solidity: event RequestMint(address indexed user, uint256 requestId, uint256 count)
func (_OwlGame *OwlGameFilterer) FilterRequestMint(opts *bind.FilterOpts, user []common.Address) (*OwlGameRequestMintIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "RequestMint", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameRequestMintIterator{contract: _OwlGame.contract, event: "RequestMint", logs: logs, sub: sub}, nil
}

// WatchRequestMint is a free log subscription operation binding the contract event 0x79ab85c2314dba73749962c93d6920c946a5f6f0532f5e675fd7b1ff530ee058.
//
// Solidity: event RequestMint(address indexed user, uint256 requestId, uint256 count)
func (_OwlGame *OwlGameFilterer) WatchRequestMint(opts *bind.WatchOpts, sink chan<- *OwlGameRequestMint, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "RequestMint", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameRequestMint)
				if err := _OwlGame.contract.UnpackLog(event, "RequestMint", log); err != nil {
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

// ParseRequestMint is a log parse operation binding the contract event 0x79ab85c2314dba73749962c93d6920c946a5f6f0532f5e675fd7b1ff530ee058.
//
// Solidity: event RequestMint(address indexed user, uint256 requestId, uint256 count)
func (_OwlGame *OwlGameFilterer) ParseRequestMint(log types.Log) (*OwlGameRequestMint, error) {
	event := new(OwlGameRequestMint)
	if err := _OwlGame.contract.UnpackLog(event, "RequestMint", log); err != nil {
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
	User      common.Address
	TokenId   []*big.Int
	BuffLevel uint16
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeMysteryBox is a free log retrieval operation binding the contract event 0xca6d3ba2c5c2c5da541c637a6da9bccafef99671140cd5fd444322bdaca03336.
//
// Solidity: event StakeMysteryBox(address indexed user, uint256[] tokenId, uint16 buffLevel)
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

// WatchStakeMysteryBox is a free log subscription operation binding the contract event 0xca6d3ba2c5c2c5da541c637a6da9bccafef99671140cd5fd444322bdaca03336.
//
// Solidity: event StakeMysteryBox(address indexed user, uint256[] tokenId, uint16 buffLevel)
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

// ParseStakeMysteryBox is a log parse operation binding the contract event 0xca6d3ba2c5c2c5da541c637a6da9bccafef99671140cd5fd444322bdaca03336.
//
// Solidity: event StakeMysteryBox(address indexed user, uint256[] tokenId, uint16 buffLevel)
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

// OwlGameUnlockableRebateIncreasedIterator is returned from FilterUnlockableRebateIncreased and is used to iterate over the raw logs and unpacked data for UnlockableRebateIncreased events raised by the OwlGame contract.
type OwlGameUnlockableRebateIncreasedIterator struct {
	Event *OwlGameUnlockableRebateIncreased // Event containing the contract specifics and raw log

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
func (it *OwlGameUnlockableRebateIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwlGameUnlockableRebateIncreased)
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
		it.Event = new(OwlGameUnlockableRebateIncreased)
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
func (it *OwlGameUnlockableRebateIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwlGameUnlockableRebateIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwlGameUnlockableRebateIncreased represents a UnlockableRebateIncreased event raised by the OwlGame contract.
type OwlGameUnlockableRebateIncreased struct {
	User      common.Address
	MintCount *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockableRebateIncreased is a free log retrieval operation binding the contract event 0x1f464096fc7638b46885d25d4cfc646c289ce897eeb853b8c4c4ace63285b77f.
//
// Solidity: event UnlockableRebateIncreased(address indexed user, uint256 mintCount, uint256 amount)
func (_OwlGame *OwlGameFilterer) FilterUnlockableRebateIncreased(opts *bind.FilterOpts, user []common.Address) (*OwlGameUnlockableRebateIncreasedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.FilterLogs(opts, "UnlockableRebateIncreased", userRule)
	if err != nil {
		return nil, err
	}
	return &OwlGameUnlockableRebateIncreasedIterator{contract: _OwlGame.contract, event: "UnlockableRebateIncreased", logs: logs, sub: sub}, nil
}

// WatchUnlockableRebateIncreased is a free log subscription operation binding the contract event 0x1f464096fc7638b46885d25d4cfc646c289ce897eeb853b8c4c4ace63285b77f.
//
// Solidity: event UnlockableRebateIncreased(address indexed user, uint256 mintCount, uint256 amount)
func (_OwlGame *OwlGameFilterer) WatchUnlockableRebateIncreased(opts *bind.WatchOpts, sink chan<- *OwlGameUnlockableRebateIncreased, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OwlGame.contract.WatchLogs(opts, "UnlockableRebateIncreased", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwlGameUnlockableRebateIncreased)
				if err := _OwlGame.contract.UnpackLog(event, "UnlockableRebateIncreased", log); err != nil {
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

// ParseUnlockableRebateIncreased is a log parse operation binding the contract event 0x1f464096fc7638b46885d25d4cfc646c289ce897eeb853b8c4c4ace63285b77f.
//
// Solidity: event UnlockableRebateIncreased(address indexed user, uint256 mintCount, uint256 amount)
func (_OwlGame *OwlGameFilterer) ParseUnlockableRebateIncreased(log types.Log) (*OwlGameUnlockableRebateIncreased, error) {
	event := new(OwlGameUnlockableRebateIncreased)
	if err := _OwlGame.contract.UnpackLog(event, "UnlockableRebateIncreased", log); err != nil {
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
