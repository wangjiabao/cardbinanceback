// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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

// BuySomethingMetaData contains all meta data concerning the BuySomething contract.
var BuySomethingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"setAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"account\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersAmountByIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usersAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BuySomethingABI is the input ABI used to generate the binding from.
// Deprecated: Use BuySomethingMetaData.ABI instead.
var BuySomethingABI = BuySomethingMetaData.ABI

// BuySomething is an auto generated Go binding around an Ethereum contract.
type BuySomething struct {
	BuySomethingCaller     // Read-only binding to the contract
	BuySomethingTransactor // Write-only binding to the contract
	BuySomethingFilterer   // Log filterer for contract events
}

// BuySomethingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuySomethingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuySomethingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuySomethingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuySomethingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuySomethingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuySomethingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuySomethingSession struct {
	Contract     *BuySomething     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuySomethingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuySomethingCallerSession struct {
	Contract *BuySomethingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BuySomethingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuySomethingTransactorSession struct {
	Contract     *BuySomethingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BuySomethingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuySomethingRaw struct {
	Contract *BuySomething // Generic contract binding to access the raw methods on
}

// BuySomethingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuySomethingCallerRaw struct {
	Contract *BuySomethingCaller // Generic read-only contract binding to access the raw methods on
}

// BuySomethingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuySomethingTransactorRaw struct {
	Contract *BuySomethingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuySomething creates a new instance of BuySomething, bound to a specific deployed contract.
func NewBuySomething(address common.Address, backend bind.ContractBackend) (*BuySomething, error) {
	contract, err := bindBuySomething(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BuySomething{BuySomethingCaller: BuySomethingCaller{contract: contract}, BuySomethingTransactor: BuySomethingTransactor{contract: contract}, BuySomethingFilterer: BuySomethingFilterer{contract: contract}}, nil
}

// NewBuySomethingCaller creates a new read-only instance of BuySomething, bound to a specific deployed contract.
func NewBuySomethingCaller(address common.Address, caller bind.ContractCaller) (*BuySomethingCaller, error) {
	contract, err := bindBuySomething(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuySomethingCaller{contract: contract}, nil
}

// NewBuySomethingTransactor creates a new write-only instance of BuySomething, bound to a specific deployed contract.
func NewBuySomethingTransactor(address common.Address, transactor bind.ContractTransactor) (*BuySomethingTransactor, error) {
	contract, err := bindBuySomething(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuySomethingTransactor{contract: contract}, nil
}

// NewBuySomethingFilterer creates a new log filterer instance of BuySomething, bound to a specific deployed contract.
func NewBuySomethingFilterer(address common.Address, filterer bind.ContractFilterer) (*BuySomethingFilterer, error) {
	contract, err := bindBuySomething(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuySomethingFilterer{contract: contract}, nil
}

// bindBuySomething binds a generic wrapper to an already deployed contract.
func bindBuySomething(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuySomethingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuySomething *BuySomethingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuySomething.Contract.BuySomethingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuySomething *BuySomethingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuySomething.Contract.BuySomethingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuySomething *BuySomethingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuySomething.Contract.BuySomethingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuySomething *BuySomethingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuySomething.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuySomething *BuySomethingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuySomething.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuySomething *BuySomethingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuySomething.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BuySomething *BuySomethingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BuySomething *BuySomethingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BuySomething.Contract.DEFAULTADMINROLE(&_BuySomething.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BuySomething *BuySomethingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BuySomething.Contract.DEFAULTADMINROLE(&_BuySomething.CallOpts)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() view returns(address)
func (_BuySomething *BuySomethingCaller) Account(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "account")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() view returns(address)
func (_BuySomething *BuySomethingSession) Account() (common.Address, error) {
	return _BuySomething.Contract.Account(&_BuySomething.CallOpts)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() view returns(address)
func (_BuySomething *BuySomethingCallerSession) Account() (common.Address, error) {
	return _BuySomething.Contract.Account(&_BuySomething.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BuySomething *BuySomethingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BuySomething *BuySomethingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BuySomething.Contract.GetRoleAdmin(&_BuySomething.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BuySomething *BuySomethingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BuySomething.Contract.GetRoleAdmin(&_BuySomething.CallOpts, role)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_BuySomething *BuySomethingCaller) GetUserLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "getUserLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_BuySomething *BuySomethingSession) GetUserLength() (*big.Int, error) {
	return _BuySomething.Contract.GetUserLength(&_BuySomething.CallOpts)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_BuySomething *BuySomethingCallerSession) GetUserLength() (*big.Int, error) {
	return _BuySomething.Contract.GetUserLength(&_BuySomething.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_BuySomething *BuySomethingCaller) GetUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "getUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_BuySomething *BuySomethingSession) GetUsers() ([]common.Address, error) {
	return _BuySomething.Contract.GetUsers(&_BuySomething.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_BuySomething *BuySomethingCallerSession) GetUsers() ([]common.Address, error) {
	return _BuySomething.Contract.GetUsers(&_BuySomething.CallOpts)
}

// GetUsersAmountByIndex is a free data retrieval call binding the contract method 0xadaf9e71.
//
// Solidity: function getUsersAmountByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_BuySomething *BuySomethingCaller) GetUsersAmountByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "getUsersAmountByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUsersAmountByIndex is a free data retrieval call binding the contract method 0xadaf9e71.
//
// Solidity: function getUsersAmountByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_BuySomething *BuySomethingSession) GetUsersAmountByIndex(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _BuySomething.Contract.GetUsersAmountByIndex(&_BuySomething.CallOpts, startIndex, endIndex)
}

// GetUsersAmountByIndex is a free data retrieval call binding the contract method 0xadaf9e71.
//
// Solidity: function getUsersAmountByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_BuySomething *BuySomethingCallerSession) GetUsersAmountByIndex(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _BuySomething.Contract.GetUsersAmountByIndex(&_BuySomething.CallOpts, startIndex, endIndex)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_BuySomething *BuySomethingCaller) GetUsersByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "getUsersByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_BuySomething *BuySomethingSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _BuySomething.Contract.GetUsersByIndex(&_BuySomething.CallOpts, startIndex, endIndex)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_BuySomething *BuySomethingCallerSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _BuySomething.Contract.GetUsersByIndex(&_BuySomething.CallOpts, startIndex, endIndex)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BuySomething *BuySomethingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BuySomething *BuySomethingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BuySomething.Contract.HasRole(&_BuySomething.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BuySomething *BuySomethingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BuySomething.Contract.HasRole(&_BuySomething.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BuySomething *BuySomethingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BuySomething *BuySomethingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BuySomething.Contract.SupportsInterface(&_BuySomething.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BuySomething *BuySomethingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BuySomething.Contract.SupportsInterface(&_BuySomething.CallOpts, interfaceId)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_BuySomething *BuySomethingCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_BuySomething *BuySomethingSession) Usdt() (common.Address, error) {
	return _BuySomething.Contract.Usdt(&_BuySomething.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_BuySomething *BuySomethingCallerSession) Usdt() (common.Address, error) {
	return _BuySomething.Contract.Usdt(&_BuySomething.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_BuySomething *BuySomethingCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "users", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_BuySomething *BuySomethingSession) Users(arg0 *big.Int) (common.Address, error) {
	return _BuySomething.Contract.Users(&_BuySomething.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_BuySomething *BuySomethingCallerSession) Users(arg0 *big.Int) (common.Address, error) {
	return _BuySomething.Contract.Users(&_BuySomething.CallOpts, arg0)
}

// UsersAmount is a free data retrieval call binding the contract method 0x0963b51e.
//
// Solidity: function usersAmount(uint256 ) view returns(uint256)
func (_BuySomething *BuySomethingCaller) UsersAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BuySomething.contract.Call(opts, &out, "usersAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsersAmount is a free data retrieval call binding the contract method 0x0963b51e.
//
// Solidity: function usersAmount(uint256 ) view returns(uint256)
func (_BuySomething *BuySomethingSession) UsersAmount(arg0 *big.Int) (*big.Int, error) {
	return _BuySomething.Contract.UsersAmount(&_BuySomething.CallOpts, arg0)
}

// UsersAmount is a free data retrieval call binding the contract method 0x0963b51e.
//
// Solidity: function usersAmount(uint256 ) view returns(uint256)
func (_BuySomething *BuySomethingCallerSession) UsersAmount(arg0 *big.Int) (*big.Int, error) {
	return _BuySomething.Contract.UsersAmount(&_BuySomething.CallOpts, arg0)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 num) returns()
func (_BuySomething *BuySomethingTransactor) Buy(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _BuySomething.contract.Transact(opts, "buy", num)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 num) returns()
func (_BuySomething *BuySomethingSession) Buy(num *big.Int) (*types.Transaction, error) {
	return _BuySomething.Contract.Buy(&_BuySomething.TransactOpts, num)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 num) returns()
func (_BuySomething *BuySomethingTransactorSession) Buy(num *big.Int) (*types.Transaction, error) {
	return _BuySomething.Contract.Buy(&_BuySomething.TransactOpts, num)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BuySomething *BuySomethingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuySomething.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BuySomething *BuySomethingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuySomething.Contract.GrantRole(&_BuySomething.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BuySomething *BuySomethingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuySomething.Contract.GrantRole(&_BuySomething.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BuySomething *BuySomethingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BuySomething.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BuySomething *BuySomethingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BuySomething.Contract.RenounceRole(&_BuySomething.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BuySomething *BuySomethingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BuySomething.Contract.RenounceRole(&_BuySomething.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BuySomething *BuySomethingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuySomething.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BuySomething *BuySomethingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuySomething.Contract.RevokeRole(&_BuySomething.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BuySomething *BuySomethingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuySomething.Contract.RevokeRole(&_BuySomething.TransactOpts, role, account)
}

// SetAccount is a paid mutator transaction binding the contract method 0x29f6d57c.
//
// Solidity: function setAccount(address account_) returns()
func (_BuySomething *BuySomethingTransactor) SetAccount(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _BuySomething.contract.Transact(opts, "setAccount", account_)
}

// SetAccount is a paid mutator transaction binding the contract method 0x29f6d57c.
//
// Solidity: function setAccount(address account_) returns()
func (_BuySomething *BuySomethingSession) SetAccount(account_ common.Address) (*types.Transaction, error) {
	return _BuySomething.Contract.SetAccount(&_BuySomething.TransactOpts, account_)
}

// SetAccount is a paid mutator transaction binding the contract method 0x29f6d57c.
//
// Solidity: function setAccount(address account_) returns()
func (_BuySomething *BuySomethingTransactorSession) SetAccount(account_ common.Address) (*types.Transaction, error) {
	return _BuySomething.Contract.SetAccount(&_BuySomething.TransactOpts, account_)
}

// BuySomethingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BuySomething contract.
type BuySomethingRoleAdminChangedIterator struct {
	Event *BuySomethingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BuySomethingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuySomethingRoleAdminChanged)
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
		it.Event = new(BuySomethingRoleAdminChanged)
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
func (it *BuySomethingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuySomethingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuySomethingRoleAdminChanged represents a RoleAdminChanged event raised by the BuySomething contract.
type BuySomethingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BuySomething *BuySomethingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BuySomethingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BuySomething.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BuySomethingRoleAdminChangedIterator{contract: _BuySomething.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BuySomething *BuySomethingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BuySomethingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BuySomething.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuySomethingRoleAdminChanged)
				if err := _BuySomething.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BuySomething *BuySomethingFilterer) ParseRoleAdminChanged(log types.Log) (*BuySomethingRoleAdminChanged, error) {
	event := new(BuySomethingRoleAdminChanged)
	if err := _BuySomething.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuySomethingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BuySomething contract.
type BuySomethingRoleGrantedIterator struct {
	Event *BuySomethingRoleGranted // Event containing the contract specifics and raw log

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
func (it *BuySomethingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuySomethingRoleGranted)
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
		it.Event = new(BuySomethingRoleGranted)
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
func (it *BuySomethingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuySomethingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuySomethingRoleGranted represents a RoleGranted event raised by the BuySomething contract.
type BuySomethingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuySomething *BuySomethingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BuySomethingRoleGrantedIterator, error) {

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

	logs, sub, err := _BuySomething.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BuySomethingRoleGrantedIterator{contract: _BuySomething.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuySomething *BuySomethingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BuySomethingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BuySomething.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuySomethingRoleGranted)
				if err := _BuySomething.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BuySomething *BuySomethingFilterer) ParseRoleGranted(log types.Log) (*BuySomethingRoleGranted, error) {
	event := new(BuySomethingRoleGranted)
	if err := _BuySomething.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuySomethingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BuySomething contract.
type BuySomethingRoleRevokedIterator struct {
	Event *BuySomethingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BuySomethingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuySomethingRoleRevoked)
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
		it.Event = new(BuySomethingRoleRevoked)
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
func (it *BuySomethingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuySomethingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuySomethingRoleRevoked represents a RoleRevoked event raised by the BuySomething contract.
type BuySomethingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuySomething *BuySomethingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BuySomethingRoleRevokedIterator, error) {

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

	logs, sub, err := _BuySomething.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BuySomethingRoleRevokedIterator{contract: _BuySomething.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuySomething *BuySomethingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BuySomethingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BuySomething.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuySomethingRoleRevoked)
				if err := _BuySomething.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BuySomething *BuySomethingFilterer) ParseRoleRevoked(log types.Log) (*BuySomethingRoleRevoked, error) {
	event := new(BuySomethingRoleRevoked)
	if err := _BuySomething.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
