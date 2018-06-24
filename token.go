// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_oldContract\",\"type\":\"address\"},{\"name\":\"_closeOld\",\"type\":\"bool\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgrading\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"userGet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceSell\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_oldContract\",\"type\":\"address\"}],\"name\":\"pullBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currencyFrom\",\"type\":\"address\"},{\"name\":\"_currencyTo\",\"type\":\"address\"},{\"name\":\"_tokenCount\",\"type\":\"uint256\"}],\"name\":\"fiatToFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newContract\",\"type\":\"address\"}],\"name\":\"lockForUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"holdingFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tierCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currency\",\"type\":\"address\"},{\"name\":\"_tokenCount\",\"type\":\"uint256\"}],\"name\":\"antToFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oldContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currency\",\"type\":\"address\"},{\"name\":\"_tokenCount\",\"type\":\"uint256\"}],\"name\":\"fiatToEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"referralFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceBuy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currenciesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"drainFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"crowdfundingFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tierSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currency\",\"type\":\"address\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferCurrencyOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"switched\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"usersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"getCurrencyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceProfitOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"ethToFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalDeposits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"antBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenCount\",\"type\":\"uint256\"}],\"name\":\"antToEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceStep\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_position\",\"type\":\"uint256\"}],\"name\":\"getCurrencyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawProfit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingsNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currency\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"updateCurrencyPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"increaseHoldingFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"ethToAnt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tierSupplyHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferAnyERC20Token\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currencies\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currency\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"registerCurrency\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currency\",\"type\":\"address\"},{\"name\":\"_tokenCount\",\"type\":\"uint256\"}],\"name\":\"fiatToAnt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialBalance\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Destroy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "_totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
// func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
// 	return _Token.Contract.TotalSupply(&_Token.CallOpts)
// }

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
// func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
// 	return _Token.Contract.TotalSupply(&_Token.CallOpts)
// }

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", tokenOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, tokenOwner, spender)
}

// AntBalance is a free data retrieval call binding the contract method 0x8205ac8c.
//
// Solidity: function antBalance() constant returns(uint256)
func (_Token *TokenCaller) AntBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "antBalance")
	return *ret0, err
}

// AntBalance is a free data retrieval call binding the contract method 0x8205ac8c.
//
// Solidity: function antBalance() constant returns(uint256)
func (_Token *TokenSession) AntBalance() (*big.Int, error) {
	return _Token.Contract.AntBalance(&_Token.CallOpts)
}

// AntBalance is a free data retrieval call binding the contract method 0x8205ac8c.
//
// Solidity: function antBalance() constant returns(uint256)
func (_Token *TokenCallerSession) AntBalance() (*big.Int, error) {
	return _Token.Contract.AntBalance(&_Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, tokenOwner)
}

// BalanceProfitOf is a free data retrieval call binding the contract method 0x707fdd12.
//
// Solidity: function balanceProfitOf(_tokenOwner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceProfitOf(opts *bind.CallOpts, _tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceProfitOf", _tokenOwner)
	return *ret0, err
}

// BalanceProfitOf is a free data retrieval call binding the contract method 0x707fdd12.
//
// Solidity: function balanceProfitOf(_tokenOwner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceProfitOf(_tokenOwner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceProfitOf(&_Token.CallOpts, _tokenOwner)
}

// BalanceProfitOf is a free data retrieval call binding the contract method 0x707fdd12.
//
// Solidity: function balanceProfitOf(_tokenOwner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceProfitOf(_tokenOwner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceProfitOf(&_Token.CallOpts, _tokenOwner)
}

// CrowdfundingFactor is a free data retrieval call binding the contract method 0x50037150.
//
// Solidity: function crowdfundingFactor() constant returns(uint256)
func (_Token *TokenCaller) CrowdfundingFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "crowdfundingFactor")
	return *ret0, err
}

// CrowdfundingFactor is a free data retrieval call binding the contract method 0x50037150.
//
// Solidity: function crowdfundingFactor() constant returns(uint256)
func (_Token *TokenSession) CrowdfundingFactor() (*big.Int, error) {
	return _Token.Contract.CrowdfundingFactor(&_Token.CallOpts)
}

// CrowdfundingFactor is a free data retrieval call binding the contract method 0x50037150.
//
// Solidity: function crowdfundingFactor() constant returns(uint256)
func (_Token *TokenCallerSession) CrowdfundingFactor() (*big.Int, error) {
	return _Token.Contract.CrowdfundingFactor(&_Token.CallOpts)
}

// Currencies is a free data retrieval call binding the contract method 0xf6d1c271.
//
// Solidity: function currencies( uint256) constant returns(address)
func (_Token *TokenCaller) Currencies(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "currencies", arg0)
	return *ret0, err
}

// Currencies is a free data retrieval call binding the contract method 0xf6d1c271.
//
// Solidity: function currencies( uint256) constant returns(address)
func (_Token *TokenSession) Currencies(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.Currencies(&_Token.CallOpts, arg0)
}

// Currencies is a free data retrieval call binding the contract method 0xf6d1c271.
//
// Solidity: function currencies( uint256) constant returns(address)
func (_Token *TokenCallerSession) Currencies(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.Currencies(&_Token.CallOpts, arg0)
}

// CurrenciesCount is a free data retrieval call binding the contract method 0x4c00f570.
//
// Solidity: function currenciesCount() constant returns(uint256)
func (_Token *TokenCaller) CurrenciesCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "currenciesCount")
	return *ret0, err
}

// CurrenciesCount is a free data retrieval call binding the contract method 0x4c00f570.
//
// Solidity: function currenciesCount() constant returns(uint256)
func (_Token *TokenSession) CurrenciesCount() (*big.Int, error) {
	return _Token.Contract.CurrenciesCount(&_Token.CallOpts)
}

// CurrenciesCount is a free data retrieval call binding the contract method 0x4c00f570.
//
// Solidity: function currenciesCount() constant returns(uint256)
func (_Token *TokenCallerSession) CurrenciesCount() (*big.Int, error) {
	return _Token.Contract.CurrenciesCount(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// DrainFactor is a free data retrieval call binding the contract method 0x4d1b92cd.
//
// Solidity: function drainFactor() constant returns(uint256)
func (_Token *TokenCaller) DrainFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "drainFactor")
	return *ret0, err
}

// DrainFactor is a free data retrieval call binding the contract method 0x4d1b92cd.
//
// Solidity: function drainFactor() constant returns(uint256)
func (_Token *TokenSession) DrainFactor() (*big.Int, error) {
	return _Token.Contract.DrainFactor(&_Token.CallOpts)
}

// DrainFactor is a free data retrieval call binding the contract method 0x4d1b92cd.
//
// Solidity: function drainFactor() constant returns(uint256)
func (_Token *TokenCallerSession) DrainFactor() (*big.Int, error) {
	return _Token.Contract.DrainFactor(&_Token.CallOpts)
}

// FiatBalance is a free data retrieval call binding the contract method 0x649f0153.
//
// Solidity: function fiatBalance() constant returns(uint256)
func (_Token *TokenCaller) FiatBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "fiatBalance")
	return *ret0, err
}

// FiatBalance is a free data retrieval call binding the contract method 0x649f0153.
//
// Solidity: function fiatBalance() constant returns(uint256)
func (_Token *TokenSession) FiatBalance() (*big.Int, error) {
	return _Token.Contract.FiatBalance(&_Token.CallOpts)
}

// FiatBalance is a free data retrieval call binding the contract method 0x649f0153.
//
// Solidity: function fiatBalance() constant returns(uint256)
func (_Token *TokenCallerSession) FiatBalance() (*big.Int, error) {
	return _Token.Contract.FiatBalance(&_Token.CallOpts)
}

// FundingsNumber is a free data retrieval call binding the contract method 0xa6714020.
//
// Solidity: function fundingsNumber() constant returns(uint256)
func (_Token *TokenCaller) FundingsNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "fundingsNumber")
	return *ret0, err
}

// FundingsNumber is a free data retrieval call binding the contract method 0xa6714020.
//
// Solidity: function fundingsNumber() constant returns(uint256)
func (_Token *TokenSession) FundingsNumber() (*big.Int, error) {
	return _Token.Contract.FundingsNumber(&_Token.CallOpts)
}

// FundingsNumber is a free data retrieval call binding the contract method 0xa6714020.
//
// Solidity: function fundingsNumber() constant returns(uint256)
func (_Token *TokenCallerSession) FundingsNumber() (*big.Int, error) {
	return _Token.Contract.FundingsNumber(&_Token.CallOpts)
}

// GetCurrencyAddress is a free data retrieval call binding the contract method 0x8db0599d.
//
// Solidity: function getCurrencyAddress(_position uint256) constant returns(address)
func (_Token *TokenCaller) GetCurrencyAddress(opts *bind.CallOpts, _position *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getCurrencyAddress", _position)
	return *ret0, err
}

// GetCurrencyAddress is a free data retrieval call binding the contract method 0x8db0599d.
//
// Solidity: function getCurrencyAddress(_position uint256) constant returns(address)
func (_Token *TokenSession) GetCurrencyAddress(_position *big.Int) (common.Address, error) {
	return _Token.Contract.GetCurrencyAddress(&_Token.CallOpts, _position)
}

// GetCurrencyAddress is a free data retrieval call binding the contract method 0x8db0599d.
//
// Solidity: function getCurrencyAddress(_position uint256) constant returns(address)
func (_Token *TokenCallerSession) GetCurrencyAddress(_position *big.Int) (common.Address, error) {
	return _Token.Contract.GetCurrencyAddress(&_Token.CallOpts, _position)
}

// GetCurrencyPrice is a free data retrieval call binding the contract method 0x7058346f.
//
// Solidity: function getCurrencyPrice(_currency address) constant returns(uint256)
func (_Token *TokenCaller) GetCurrencyPrice(opts *bind.CallOpts, _currency common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getCurrencyPrice", _currency)
	return *ret0, err
}

// GetCurrencyPrice is a free data retrieval call binding the contract method 0x7058346f.
//
// Solidity: function getCurrencyPrice(_currency address) constant returns(uint256)
func (_Token *TokenSession) GetCurrencyPrice(_currency common.Address) (*big.Int, error) {
	return _Token.Contract.GetCurrencyPrice(&_Token.CallOpts, _currency)
}

// GetCurrencyPrice is a free data retrieval call binding the contract method 0x7058346f.
//
// Solidity: function getCurrencyPrice(_currency address) constant returns(uint256)
func (_Token *TokenCallerSession) GetCurrencyPrice(_currency common.Address) (*big.Int, error) {
	return _Token.Contract.GetCurrencyPrice(&_Token.CallOpts, _currency)
}

// HoldingFactor is a free data retrieval call binding the contract method 0x0d8880cd.
//
// Solidity: function holdingFactor() constant returns(uint256)
func (_Token *TokenCaller) HoldingFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "holdingFactor")
	return *ret0, err
}

// HoldingFactor is a free data retrieval call binding the contract method 0x0d8880cd.
//
// Solidity: function holdingFactor() constant returns(uint256)
func (_Token *TokenSession) HoldingFactor() (*big.Int, error) {
	return _Token.Contract.HoldingFactor(&_Token.CallOpts)
}

// HoldingFactor is a free data retrieval call binding the contract method 0x0d8880cd.
//
// Solidity: function holdingFactor() constant returns(uint256)
func (_Token *TokenCallerSession) HoldingFactor() (*big.Int, error) {
	return _Token.Contract.HoldingFactor(&_Token.CallOpts)
}

// IncreaseHoldingFactor is a free data retrieval call binding the contract method 0xbdf95907.
//
// Solidity: function increaseHoldingFactor() constant returns(bool)
func (_Token *TokenCaller) IncreaseHoldingFactor(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "increaseHoldingFactor")
	return *ret0, err
}

// IncreaseHoldingFactor is a free data retrieval call binding the contract method 0xbdf95907.
//
// Solidity: function increaseHoldingFactor() constant returns(bool)
func (_Token *TokenSession) IncreaseHoldingFactor() (bool, error) {
	return _Token.Contract.IncreaseHoldingFactor(&_Token.CallOpts)
}

// IncreaseHoldingFactor is a free data retrieval call binding the contract method 0xbdf95907.
//
// Solidity: function increaseHoldingFactor() constant returns(bool)
func (_Token *TokenCallerSession) IncreaseHoldingFactor() (bool, error) {
	return _Token.Contract.IncreaseHoldingFactor(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// NewContract is a free data retrieval call binding the contract method 0x4313b531.
//
// Solidity: function newContract() constant returns(address)
func (_Token *TokenCaller) NewContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "newContract")
	return *ret0, err
}

// NewContract is a free data retrieval call binding the contract method 0x4313b531.
//
// Solidity: function newContract() constant returns(address)
func (_Token *TokenSession) NewContract() (common.Address, error) {
	return _Token.Contract.NewContract(&_Token.CallOpts)
}

// NewContract is a free data retrieval call binding the contract method 0x4313b531.
//
// Solidity: function newContract() constant returns(address)
func (_Token *TokenCallerSession) NewContract() (common.Address, error) {
	return _Token.Contract.NewContract(&_Token.CallOpts)
}

// OldContract is a free data retrieval call binding the contract method 0x30503c4e.
//
// Solidity: function oldContract() constant returns(address)
func (_Token *TokenCaller) OldContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "oldContract")
	return *ret0, err
}

// OldContract is a free data retrieval call binding the contract method 0x30503c4e.
//
// Solidity: function oldContract() constant returns(address)
func (_Token *TokenSession) OldContract() (common.Address, error) {
	return _Token.Contract.OldContract(&_Token.CallOpts)
}

// OldContract is a free data retrieval call binding the contract method 0x30503c4e.
//
// Solidity: function oldContract() constant returns(address)
func (_Token *TokenCallerSession) OldContract() (common.Address, error) {
	return _Token.Contract.OldContract(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// PriceBuy is a free data retrieval call binding the contract method 0x488376f2.
//
// Solidity: function priceBuy() constant returns(uint256)
func (_Token *TokenCaller) PriceBuy(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "priceBuy")
	return *ret0, err
}

// PriceBuy is a free data retrieval call binding the contract method 0x488376f2.
//
// Solidity: function priceBuy() constant returns(uint256)
func (_Token *TokenSession) PriceBuy() (*big.Int, error) {
	return _Token.Contract.PriceBuy(&_Token.CallOpts)
}

// PriceBuy is a free data retrieval call binding the contract method 0x488376f2.
//
// Solidity: function priceBuy() constant returns(uint256)
func (_Token *TokenCallerSession) PriceBuy() (*big.Int, error) {
	return _Token.Contract.PriceBuy(&_Token.CallOpts)
}

// PriceSell is a free data retrieval call binding the contract method 0x04aa8692.
//
// Solidity: function priceSell() constant returns(uint256)
func (_Token *TokenCaller) PriceSell(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "priceSell")
	return *ret0, err
}

// PriceSell is a free data retrieval call binding the contract method 0x04aa8692.
//
// Solidity: function priceSell() constant returns(uint256)
func (_Token *TokenSession) PriceSell() (*big.Int, error) {
	return _Token.Contract.PriceSell(&_Token.CallOpts)
}

// PriceSell is a free data retrieval call binding the contract method 0x04aa8692.
//
// Solidity: function priceSell() constant returns(uint256)
func (_Token *TokenCallerSession) PriceSell() (*big.Int, error) {
	return _Token.Contract.PriceSell(&_Token.CallOpts)
}

// PriceStep is a free data retrieval call binding the contract method 0x88d519c4.
//
// Solidity: function priceStep() constant returns(uint256)
func (_Token *TokenCaller) PriceStep(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "priceStep")
	return *ret0, err
}

// PriceStep is a free data retrieval call binding the contract method 0x88d519c4.
//
// Solidity: function priceStep() constant returns(uint256)
func (_Token *TokenSession) PriceStep() (*big.Int, error) {
	return _Token.Contract.PriceStep(&_Token.CallOpts)
}

// PriceStep is a free data retrieval call binding the contract method 0x88d519c4.
//
// Solidity: function priceStep() constant returns(uint256)
func (_Token *TokenCallerSession) PriceStep() (*big.Int, error) {
	return _Token.Contract.PriceStep(&_Token.CallOpts)
}

// ReferralFactor is a free data retrieval call binding the contract method 0x421dce17.
//
// Solidity: function referralFactor() constant returns(uint256)
func (_Token *TokenCaller) ReferralFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "referralFactor")
	return *ret0, err
}

// ReferralFactor is a free data retrieval call binding the contract method 0x421dce17.
//
// Solidity: function referralFactor() constant returns(uint256)
func (_Token *TokenSession) ReferralFactor() (*big.Int, error) {
	return _Token.Contract.ReferralFactor(&_Token.CallOpts)
}

// ReferralFactor is a free data retrieval call binding the contract method 0x421dce17.
//
// Solidity: function referralFactor() constant returns(uint256)
func (_Token *TokenCallerSession) ReferralFactor() (*big.Int, error) {
	return _Token.Contract.ReferralFactor(&_Token.CallOpts)
}

// Switched is a free data retrieval call binding the contract method 0x612be8b1.
//
// Solidity: function switched() constant returns(bool)
func (_Token *TokenCaller) Switched(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "switched")
	return *ret0, err
}

// Switched is a free data retrieval call binding the contract method 0x612be8b1.
//
// Solidity: function switched() constant returns(bool)
func (_Token *TokenSession) Switched() (bool, error) {
	return _Token.Contract.Switched(&_Token.CallOpts)
}

// Switched is a free data retrieval call binding the contract method 0x612be8b1.
//
// Solidity: function switched() constant returns(bool)
func (_Token *TokenCallerSession) Switched() (bool, error) {
	return _Token.Contract.Switched(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TierCounter is a free data retrieval call binding the contract method 0x11d6a107.
//
// Solidity: function tierCounter() constant returns(uint256)
func (_Token *TokenCaller) TierCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "tierCounter")
	return *ret0, err
}

// TierCounter is a free data retrieval call binding the contract method 0x11d6a107.
//
// Solidity: function tierCounter() constant returns(uint256)
func (_Token *TokenSession) TierCounter() (*big.Int, error) {
	return _Token.Contract.TierCounter(&_Token.CallOpts)
}

// TierCounter is a free data retrieval call binding the contract method 0x11d6a107.
//
// Solidity: function tierCounter() constant returns(uint256)
func (_Token *TokenCallerSession) TierCounter() (*big.Int, error) {
	return _Token.Contract.TierCounter(&_Token.CallOpts)
}

// TierSupply is a free data retrieval call binding the contract method 0x549b9da5.
//
// Solidity: function tierSupply() constant returns(uint256)
func (_Token *TokenCaller) TierSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "tierSupply")
	return *ret0, err
}

// TierSupply is a free data retrieval call binding the contract method 0x549b9da5.
//
// Solidity: function tierSupply() constant returns(uint256)
func (_Token *TokenSession) TierSupply() (*big.Int, error) {
	return _Token.Contract.TierSupply(&_Token.CallOpts)
}

// TierSupply is a free data retrieval call binding the contract method 0x549b9da5.
//
// Solidity: function tierSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TierSupply() (*big.Int, error) {
	return _Token.Contract.TierSupply(&_Token.CallOpts)
}

// TierSupplyHolder is a free data retrieval call binding the contract method 0xdacc4782.
//
// Solidity: function tierSupplyHolder() constant returns(uint256)
func (_Token *TokenCaller) TierSupplyHolder(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "tierSupplyHolder")
	return *ret0, err
}

// TierSupplyHolder is a free data retrieval call binding the contract method 0xdacc4782.
//
// Solidity: function tierSupplyHolder() constant returns(uint256)
func (_Token *TokenSession) TierSupplyHolder() (*big.Int, error) {
	return _Token.Contract.TierSupplyHolder(&_Token.CallOpts)
}

// TierSupplyHolder is a free data retrieval call binding the contract method 0xdacc4782.
//
// Solidity: function tierSupplyHolder() constant returns(uint256)
func (_Token *TokenCallerSession) TierSupplyHolder() (*big.Int, error) {
	return _Token.Contract.TierSupplyHolder(&_Token.CallOpts)
}

// TotalDeposits is a free data retrieval call binding the contract method 0x7d882097.
//
// Solidity: function totalDeposits() constant returns(uint256)
func (_Token *TokenCaller) TotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalDeposits")
	return *ret0, err
}

// TotalDeposits is a free data retrieval call binding the contract method 0x7d882097.
//
// Solidity: function totalDeposits() constant returns(uint256)
func (_Token *TokenSession) TotalDeposits() (*big.Int, error) {
	return _Token.Contract.TotalDeposits(&_Token.CallOpts)
}

// TotalDeposits is a free data retrieval call binding the contract method 0x7d882097.
//
// Solidity: function totalDeposits() constant returns(uint256)
func (_Token *TokenCallerSession) TotalDeposits() (*big.Int, error) {
	return _Token.Contract.TotalDeposits(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
// func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
// 	var (
// 		ret0 = new(*big.Int)
// 	)
// 	out := ret0
// 	err := _Token.contract.Call(opts, out, "totalSupply")
// 	return *ret0, err
// }

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
// func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
// 	return _Token.Contract.TotalSupply(&_Token.CallOpts)
// }

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
// func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
// 	return _Token.Contract.TotalSupply(&_Token.CallOpts)
// }

// Upgrading is a free data retrieval call binding the contract method 0x02836f24.
//
// Solidity: function upgrading() constant returns(bool)
func (_Token *TokenCaller) Upgrading(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "upgrading")
	return *ret0, err
}

// Upgrading is a free data retrieval call binding the contract method 0x02836f24.
//
// Solidity: function upgrading() constant returns(bool)
func (_Token *TokenSession) Upgrading() (bool, error) {
	return _Token.Contract.Upgrading(&_Token.CallOpts)
}

// Upgrading is a free data retrieval call binding the contract method 0x02836f24.
//
// Solidity: function upgrading() constant returns(bool)
func (_Token *TokenCallerSession) Upgrading() (bool, error) {
	return _Token.Contract.Upgrading(&_Token.CallOpts)
}

// UserExists is a free data retrieval call binding the contract method 0x0e666e49.
//
// Solidity: function userExists(_user address) constant returns(bool)
func (_Token *TokenCaller) UserExists(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "userExists", _user)
	return *ret0, err
}

// UserExists is a free data retrieval call binding the contract method 0x0e666e49.
//
// Solidity: function userExists(_user address) constant returns(bool)
func (_Token *TokenSession) UserExists(_user common.Address) (bool, error) {
	return _Token.Contract.UserExists(&_Token.CallOpts, _user)
}

// UserExists is a free data retrieval call binding the contract method 0x0e666e49.
//
// Solidity: function userExists(_user address) constant returns(bool)
func (_Token *TokenCallerSession) UserExists(_user common.Address) (bool, error) {
	return _Token.Contract.UserExists(&_Token.CallOpts, _user)
}

// UserGet is a free data retrieval call binding the contract method 0x0381eb1e.
//
// Solidity: function userGet(_index uint256) constant returns(address)
func (_Token *TokenCaller) UserGet(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "userGet", _index)
	return *ret0, err
}

// UserGet is a free data retrieval call binding the contract method 0x0381eb1e.
//
// Solidity: function userGet(_index uint256) constant returns(address)
func (_Token *TokenSession) UserGet(_index *big.Int) (common.Address, error) {
	return _Token.Contract.UserGet(&_Token.CallOpts, _index)
}

// UserGet is a free data retrieval call binding the contract method 0x0381eb1e.
//
// Solidity: function userGet(_index uint256) constant returns(address)
func (_Token *TokenCallerSession) UserGet(_index *big.Int) (common.Address, error) {
	return _Token.Contract.UserGet(&_Token.CallOpts, _index)
}

// UsersCount is a free data retrieval call binding the contract method 0x6ba13a82.
//
// Solidity: function usersCount() constant returns(uint256)
func (_Token *TokenCaller) UsersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "usersCount")
	return *ret0, err
}

// UsersCount is a free data retrieval call binding the contract method 0x6ba13a82.
//
// Solidity: function usersCount() constant returns(uint256)
func (_Token *TokenSession) UsersCount() (*big.Int, error) {
	return _Token.Contract.UsersCount(&_Token.CallOpts)
}

// UsersCount is a free data retrieval call binding the contract method 0x6ba13a82.
//
// Solidity: function usersCount() constant returns(uint256)
func (_Token *TokenCallerSession) UsersCount() (*big.Int, error) {
	return _Token.Contract.UsersCount(&_Token.CallOpts)
}

// AntToEth is a paid mutator transaction binding the contract method 0x872de5bc.
//
// Solidity: function antToEth(_tokenCount uint256) returns(bool)
func (_Token *TokenTransactor) AntToEth(opts *bind.TransactOpts, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "antToEth", _tokenCount)
}

// AntToEth is a paid mutator transaction binding the contract method 0x872de5bc.
//
// Solidity: function antToEth(_tokenCount uint256) returns(bool)
func (_Token *TokenSession) AntToEth(_tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.AntToEth(&_Token.TransactOpts, _tokenCount)
}

// AntToEth is a paid mutator transaction binding the contract method 0x872de5bc.
//
// Solidity: function antToEth(_tokenCount uint256) returns(bool)
func (_Token *TokenTransactorSession) AntToEth(_tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.AntToEth(&_Token.TransactOpts, _tokenCount)
}

// AntToFiat is a paid mutator transaction binding the contract method 0x2115d50b.
//
// Solidity: function antToFiat(_currency address, _tokenCount uint256) returns(bool)
func (_Token *TokenTransactor) AntToFiat(opts *bind.TransactOpts, _currency common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "antToFiat", _currency, _tokenCount)
}

// AntToFiat is a paid mutator transaction binding the contract method 0x2115d50b.
//
// Solidity: function antToFiat(_currency address, _tokenCount uint256) returns(bool)
func (_Token *TokenSession) AntToFiat(_currency common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.AntToFiat(&_Token.TransactOpts, _currency, _tokenCount)
}

// AntToFiat is a paid mutator transaction binding the contract method 0x2115d50b.
//
// Solidity: function antToFiat(_currency address, _tokenCount uint256) returns(bool)
func (_Token *TokenTransactorSession) AntToFiat(_currency common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.AntToFiat(&_Token.TransactOpts, _currency, _tokenCount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Token *TokenSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, tokens)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_Token *TokenTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approveAndCall", spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_Token *TokenSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_Token *TokenTransactorSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, spender, tokens, data)
}

// Close is a paid mutator transaction binding the contract method 0xc74073a1.
//
// Solidity: function close(_destination address) returns()
func (_Token *TokenTransactor) Close(opts *bind.TransactOpts, _destination common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "close", _destination)
}

// Close is a paid mutator transaction binding the contract method 0xc74073a1.
//
// Solidity: function close(_destination address) returns()
func (_Token *TokenSession) Close(_destination common.Address) (*types.Transaction, error) {
	return _Token.Contract.Close(&_Token.TransactOpts, _destination)
}

// Close is a paid mutator transaction binding the contract method 0xc74073a1.
//
// Solidity: function close(_destination address) returns()
func (_Token *TokenTransactorSession) Close(_destination common.Address) (*types.Transaction, error) {
	return _Token.Contract.Close(&_Token.TransactOpts, _destination)
}

// EthToAnt is a paid mutator transaction binding the contract method 0xd749f4ec.
//
// Solidity: function ethToAnt(_referral address) returns(bool)
func (_Token *TokenTransactor) EthToAnt(opts *bind.TransactOpts, _referral common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "ethToAnt", _referral)
}

// EthToAnt is a paid mutator transaction binding the contract method 0xd749f4ec.
//
// Solidity: function ethToAnt(_referral address) returns(bool)
func (_Token *TokenSession) EthToAnt(_referral common.Address) (*types.Transaction, error) {
	return _Token.Contract.EthToAnt(&_Token.TransactOpts, _referral)
}

// EthToAnt is a paid mutator transaction binding the contract method 0xd749f4ec.
//
// Solidity: function ethToAnt(_referral address) returns(bool)
func (_Token *TokenTransactorSession) EthToAnt(_referral common.Address) (*types.Transaction, error) {
	return _Token.Contract.EthToAnt(&_Token.TransactOpts, _referral)
}

// EthToFiat is a paid mutator transaction binding the contract method 0x7c66c396.
//
// Solidity: function ethToFiat(_currency address) returns(bool)
func (_Token *TokenTransactor) EthToFiat(opts *bind.TransactOpts, _currency common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "ethToFiat", _currency)
}

// EthToFiat is a paid mutator transaction binding the contract method 0x7c66c396.
//
// Solidity: function ethToFiat(_currency address) returns(bool)
func (_Token *TokenSession) EthToFiat(_currency common.Address) (*types.Transaction, error) {
	return _Token.Contract.EthToFiat(&_Token.TransactOpts, _currency)
}

// EthToFiat is a paid mutator transaction binding the contract method 0x7c66c396.
//
// Solidity: function ethToFiat(_currency address) returns(bool)
func (_Token *TokenTransactorSession) EthToFiat(_currency common.Address) (*types.Transaction, error) {
	return _Token.Contract.EthToFiat(&_Token.TransactOpts, _currency)
}

// FiatToAnt is a paid mutator transaction binding the contract method 0xfef76e47.
//
// Solidity: function fiatToAnt(_currency address, _tokenCount uint256) returns(bool)
func (_Token *TokenTransactor) FiatToAnt(opts *bind.TransactOpts, _currency common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "fiatToAnt", _currency, _tokenCount)
}

// FiatToAnt is a paid mutator transaction binding the contract method 0xfef76e47.
//
// Solidity: function fiatToAnt(_currency address, _tokenCount uint256) returns(bool)
func (_Token *TokenSession) FiatToAnt(_currency common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FiatToAnt(&_Token.TransactOpts, _currency, _tokenCount)
}

// FiatToAnt is a paid mutator transaction binding the contract method 0xfef76e47.
//
// Solidity: function fiatToAnt(_currency address, _tokenCount uint256) returns(bool)
func (_Token *TokenTransactorSession) FiatToAnt(_currency common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FiatToAnt(&_Token.TransactOpts, _currency, _tokenCount)
}

// FiatToEth is a paid mutator transaction binding the contract method 0x3ae2e35b.
//
// Solidity: function fiatToEth(_currency address, _tokenCount uint256) returns(bool)
func (_Token *TokenTransactor) FiatToEth(opts *bind.TransactOpts, _currency common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "fiatToEth", _currency, _tokenCount)
}

// FiatToEth is a paid mutator transaction binding the contract method 0x3ae2e35b.
//
// Solidity: function fiatToEth(_currency address, _tokenCount uint256) returns(bool)
func (_Token *TokenSession) FiatToEth(_currency common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FiatToEth(&_Token.TransactOpts, _currency, _tokenCount)
}

// FiatToEth is a paid mutator transaction binding the contract method 0x3ae2e35b.
//
// Solidity: function fiatToEth(_currency address, _tokenCount uint256) returns(bool)
func (_Token *TokenTransactorSession) FiatToEth(_currency common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FiatToEth(&_Token.TransactOpts, _currency, _tokenCount)
}

// FiatToFiat is a paid mutator transaction binding the contract method 0x05fb2d9c.
//
// Solidity: function fiatToFiat(_currencyFrom address, _currencyTo address, _tokenCount uint256) returns(bool)
func (_Token *TokenTransactor) FiatToFiat(opts *bind.TransactOpts, _currencyFrom common.Address, _currencyTo common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "fiatToFiat", _currencyFrom, _currencyTo, _tokenCount)
}

// FiatToFiat is a paid mutator transaction binding the contract method 0x05fb2d9c.
//
// Solidity: function fiatToFiat(_currencyFrom address, _currencyTo address, _tokenCount uint256) returns(bool)
func (_Token *TokenSession) FiatToFiat(_currencyFrom common.Address, _currencyTo common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FiatToFiat(&_Token.TransactOpts, _currencyFrom, _currencyTo, _tokenCount)
}

// FiatToFiat is a paid mutator transaction binding the contract method 0x05fb2d9c.
//
// Solidity: function fiatToFiat(_currencyFrom address, _currencyTo address, _tokenCount uint256) returns(bool)
func (_Token *TokenTransactorSession) FiatToFiat(_currencyFrom common.Address, _currencyTo common.Address, _tokenCount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FiatToFiat(&_Token.TransactOpts, _currencyFrom, _currencyTo, _tokenCount)
}

// FinishUpgrade is a paid mutator transaction binding the contract method 0x340df28f.
//
// Solidity: function finishUpgrade() returns()
func (_Token *TokenTransactor) FinishUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "finishUpgrade")
}

// FinishUpgrade is a paid mutator transaction binding the contract method 0x340df28f.
//
// Solidity: function finishUpgrade() returns()
func (_Token *TokenSession) FinishUpgrade() (*types.Transaction, error) {
	return _Token.Contract.FinishUpgrade(&_Token.TransactOpts)
}

// FinishUpgrade is a paid mutator transaction binding the contract method 0x340df28f.
//
// Solidity: function finishUpgrade() returns()
func (_Token *TokenTransactorSession) FinishUpgrade() (*types.Transaction, error) {
	return _Token.Contract.FinishUpgrade(&_Token.TransactOpts)
}

// LockForUpgrade is a paid mutator transaction binding the contract method 0x0be1e877.
//
// Solidity: function lockForUpgrade(_newContract address) returns()
func (_Token *TokenTransactor) LockForUpgrade(opts *bind.TransactOpts, _newContract common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "lockForUpgrade", _newContract)
}

// LockForUpgrade is a paid mutator transaction binding the contract method 0x0be1e877.
//
// Solidity: function lockForUpgrade(_newContract address) returns()
func (_Token *TokenSession) LockForUpgrade(_newContract common.Address) (*types.Transaction, error) {
	return _Token.Contract.LockForUpgrade(&_Token.TransactOpts, _newContract)
}

// LockForUpgrade is a paid mutator transaction binding the contract method 0x0be1e877.
//
// Solidity: function lockForUpgrade(_newContract address) returns()
func (_Token *TokenTransactorSession) LockForUpgrade(_newContract common.Address) (*types.Transaction, error) {
	return _Token.Contract.LockForUpgrade(&_Token.TransactOpts, _newContract)
}

// PullBalance is a paid mutator transaction binding the contract method 0x058f3ed2.
//
// Solidity: function pullBalance(_owner address, _oldContract address) returns()
func (_Token *TokenTransactor) PullBalance(opts *bind.TransactOpts, _owner common.Address, _oldContract common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pullBalance", _owner, _oldContract)
}

// PullBalance is a paid mutator transaction binding the contract method 0x058f3ed2.
//
// Solidity: function pullBalance(_owner address, _oldContract address) returns()
func (_Token *TokenSession) PullBalance(_owner common.Address, _oldContract common.Address) (*types.Transaction, error) {
	return _Token.Contract.PullBalance(&_Token.TransactOpts, _owner, _oldContract)
}

// PullBalance is a paid mutator transaction binding the contract method 0x058f3ed2.
//
// Solidity: function pullBalance(_owner address, _oldContract address) returns()
func (_Token *TokenTransactorSession) PullBalance(_owner common.Address, _oldContract common.Address) (*types.Transaction, error) {
	return _Token.Contract.PullBalance(&_Token.TransactOpts, _owner, _oldContract)
}

// RegisterCurrency is a paid mutator transaction binding the contract method 0xf7d44a29.
//
// Solidity: function registerCurrency(_currency address, _price uint256) returns(bool)
func (_Token *TokenTransactor) RegisterCurrency(opts *bind.TransactOpts, _currency common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "registerCurrency", _currency, _price)
}

// RegisterCurrency is a paid mutator transaction binding the contract method 0xf7d44a29.
//
// Solidity: function registerCurrency(_currency address, _price uint256) returns(bool)
func (_Token *TokenSession) RegisterCurrency(_currency common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Token.Contract.RegisterCurrency(&_Token.TransactOpts, _currency, _price)
}

// RegisterCurrency is a paid mutator transaction binding the contract method 0xf7d44a29.
//
// Solidity: function registerCurrency(_currency address, _price uint256) returns(bool)
func (_Token *TokenTransactorSession) RegisterCurrency(_currency common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Token.Contract.RegisterCurrency(&_Token.TransactOpts, _currency, _price)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Token *TokenSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(tokenAddress address, tokens uint256) returns(success bool)
func (_Token *TokenTransactor) TransferAnyERC20Token(opts *bind.TransactOpts, tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferAnyERC20Token", tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(tokenAddress address, tokens uint256) returns(success bool)
func (_Token *TokenSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferAnyERC20Token(&_Token.TransactOpts, tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(tokenAddress address, tokens uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferAnyERC20Token(&_Token.TransactOpts, tokenAddress, tokens)
}

// TransferCurrencyOwnership is a paid mutator transaction binding the contract method 0x580b2875.
//
// Solidity: function transferCurrencyOwnership(_currency address, _newOwner address) returns()
func (_Token *TokenTransactor) TransferCurrencyOwnership(opts *bind.TransactOpts, _currency common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferCurrencyOwnership", _currency, _newOwner)
}

// TransferCurrencyOwnership is a paid mutator transaction binding the contract method 0x580b2875.
//
// Solidity: function transferCurrencyOwnership(_currency address, _newOwner address) returns()
func (_Token *TokenSession) TransferCurrencyOwnership(_currency common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferCurrencyOwnership(&_Token.TransactOpts, _currency, _newOwner)
}

// TransferCurrencyOwnership is a paid mutator transaction binding the contract method 0x580b2875.
//
// Solidity: function transferCurrencyOwnership(_currency address, _newOwner address) returns()
func (_Token *TokenTransactorSession) TransferCurrencyOwnership(_currency common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferCurrencyOwnership(&_Token.TransactOpts, _currency, _newOwner)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Token *TokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Token *TokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, _newOwner)
}

// UpdateCurrencyPrice is a paid mutator transaction binding the contract method 0xb8846c1b.
//
// Solidity: function updateCurrencyPrice(_currency address, _price uint256) returns()
func (_Token *TokenTransactor) UpdateCurrencyPrice(opts *bind.TransactOpts, _currency common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateCurrencyPrice", _currency, _price)
}

// UpdateCurrencyPrice is a paid mutator transaction binding the contract method 0xb8846c1b.
//
// Solidity: function updateCurrencyPrice(_currency address, _price uint256) returns()
func (_Token *TokenSession) UpdateCurrencyPrice(_currency common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateCurrencyPrice(&_Token.TransactOpts, _currency, _price)
}

// UpdateCurrencyPrice is a paid mutator transaction binding the contract method 0xb8846c1b.
//
// Solidity: function updateCurrencyPrice(_currency address, _price uint256) returns()
func (_Token *TokenTransactorSession) UpdateCurrencyPrice(_currency common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateCurrencyPrice(&_Token.TransactOpts, _currency, _price)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0110c892.
//
// Solidity: function upgrade(_oldContract address, _closeOld bool) returns()
func (_Token *TokenTransactor) Upgrade(opts *bind.TransactOpts, _oldContract common.Address, _closeOld bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "upgrade", _oldContract, _closeOld)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0110c892.
//
// Solidity: function upgrade(_oldContract address, _closeOld bool) returns()
func (_Token *TokenSession) Upgrade(_oldContract common.Address, _closeOld bool) (*types.Transaction, error) {
	return _Token.Contract.Upgrade(&_Token.TransactOpts, _oldContract, _closeOld)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0110c892.
//
// Solidity: function upgrade(_oldContract address, _closeOld bool) returns()
func (_Token *TokenTransactorSession) Upgrade(_oldContract common.Address, _closeOld bool) (*types.Transaction, error) {
	return _Token.Contract.Upgrade(&_Token.TransactOpts, _oldContract, _closeOld)
}

// WithdrawProfit is a paid mutator transaction binding the contract method 0x959499b6.
//
// Solidity: function withdrawProfit() returns()
func (_Token *TokenTransactor) WithdrawProfit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawProfit")
}

// WithdrawProfit is a paid mutator transaction binding the contract method 0x959499b6.
//
// Solidity: function withdrawProfit() returns()
func (_Token *TokenSession) WithdrawProfit() (*types.Transaction, error) {
	return _Token.Contract.WithdrawProfit(&_Token.TransactOpts)
}

// WithdrawProfit is a paid mutator transaction binding the contract method 0x959499b6.
//
// Solidity: function withdrawProfit() returns()
func (_Token *TokenTransactorSession) WithdrawProfit() (*types.Transaction, error) {
	return _Token.Contract.WithdrawProfit(&_Token.TransactOpts)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TokenDestroyIterator is returned from FilterDestroy and is used to iterate over the raw logs and unpacked data for Destroy events raised by the Token contract.
type TokenDestroyIterator struct {
	Event *TokenDestroy // Event containing the contract specifics and raw log

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
func (it *TokenDestroyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDestroy)
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
		it.Event = new(TokenDestroy)
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
func (it *TokenDestroyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDestroyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDestroy represents a Destroy event raised by the Token contract.
type TokenDestroy struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDestroy is a free log retrieval operation binding the contract event 0x81325e2a6c442af9d36e4ee9697f38d5f4bf0837ade0f6c411c6a40af7c057ee.
//
// Solidity: e Destroy(from indexed address, amount uint256)
func (_Token *TokenFilterer) FilterDestroy(opts *bind.FilterOpts, from []common.Address) (*TokenDestroyIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Destroy", fromRule)
	if err != nil {
		return nil, err
	}
	return &TokenDestroyIterator{contract: _Token.contract, event: "Destroy", logs: logs, sub: sub}, nil
}

// WatchDestroy is a free log subscription operation binding the contract event 0x81325e2a6c442af9d36e4ee9697f38d5f4bf0837ade0f6c411c6a40af7c057ee.
//
// Solidity: e Destroy(from indexed address, amount uint256)
func (_Token *TokenFilterer) WatchDestroy(opts *bind.WatchOpts, sink chan<- *TokenDestroy, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Destroy", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDestroy)
				if err := _Token.contract.UnpackLog(event, "Destroy", log); err != nil {
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

// TokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Token contract.
type TokenMintIterator struct {
	Event *TokenMint // Event containing the contract specifics and raw log

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
func (it *TokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMint)
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
		it.Event = new(TokenMint)
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
func (it *TokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMint represents a Mint event raised by the Token contract.
type TokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_Token *TokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*TokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &TokenMintIterator{contract: _Token.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_Token *TokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *TokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMint)
				if err := _Token.contract.UnpackLog(event, "Mint", log); err != nil {
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

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_from indexed address, _to indexed address)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_from indexed address, _to indexed address)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokens uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokens uint256)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
