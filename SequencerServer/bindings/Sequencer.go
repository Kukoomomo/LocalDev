// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// SequencerSequencerInfo is an auto generated low-level Go binding around an user-defined struct.
type SequencerSequencerInfo struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}

// SequencerMetaData contains all meta data concerning the Sequencer contract.
var SequencerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"SequencerCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"SequencerDelete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SequencerUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bitToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_nodeID\",\"type\":\"bytes\"}],\"name\":\"createSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getSequencer\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"internalType\":\"structSequencer.SequencerInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencers\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"internalType\":\"structSequencer.SequencerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bitToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerLimit\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sequencers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bitToken\",\"type\":\"address\"}],\"name\":\"updateBitAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"updateEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_scheduler\",\"type\":\"address\"}],\"name\":\"updateScheduler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_limit\",\"type\":\"uint8\"}],\"name\":\"updateSequencerLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SequencerABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerMetaData.ABI instead.
var SequencerABI = SequencerMetaData.ABI

// Sequencer is an auto generated Go binding around an Ethereum contract.
type Sequencer struct {
	SequencerCaller     // Read-only binding to the contract
	SequencerTransactor // Write-only binding to the contract
	SequencerFilterer   // Log filterer for contract events
}

// SequencerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerSession struct {
	Contract     *Sequencer        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SequencerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerCallerSession struct {
	Contract *SequencerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SequencerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerTransactorSession struct {
	Contract     *SequencerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SequencerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerRaw struct {
	Contract *Sequencer // Generic contract binding to access the raw methods on
}

// SequencerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerCallerRaw struct {
	Contract *SequencerCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerTransactorRaw struct {
	Contract *SequencerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencer creates a new instance of Sequencer, bound to a specific deployed contract.
func NewSequencer(address common.Address, backend bind.ContractBackend) (*Sequencer, error) {
	contract, err := bindSequencer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sequencer{SequencerCaller: SequencerCaller{contract: contract}, SequencerTransactor: SequencerTransactor{contract: contract}, SequencerFilterer: SequencerFilterer{contract: contract}}, nil
}

// NewSequencerCaller creates a new read-only instance of Sequencer, bound to a specific deployed contract.
func NewSequencerCaller(address common.Address, caller bind.ContractCaller) (*SequencerCaller, error) {
	contract, err := bindSequencer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerCaller{contract: contract}, nil
}

// NewSequencerTransactor creates a new write-only instance of Sequencer, bound to a specific deployed contract.
func NewSequencerTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerTransactor, error) {
	contract, err := bindSequencer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerTransactor{contract: contract}, nil
}

// NewSequencerFilterer creates a new log filterer instance of Sequencer, bound to a specific deployed contract.
func NewSequencerFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerFilterer, error) {
	contract, err := bindSequencer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerFilterer{contract: contract}, nil
}

// bindSequencer binds a generic wrapper to an already deployed contract.
func bindSequencer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sequencer *SequencerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sequencer.Contract.SequencerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sequencer *SequencerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.Contract.SequencerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sequencer *SequencerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sequencer.Contract.SequencerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sequencer *SequencerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sequencer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sequencer *SequencerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sequencer *SequencerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sequencer.Contract.contract.Transact(opts, method, params...)
}

// BitToken is a free data retrieval call binding the contract method 0xd84e9f92.
//
// Solidity: function bitToken() view returns(address)
func (_Sequencer *SequencerCaller) BitToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "bitToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BitToken is a free data retrieval call binding the contract method 0xd84e9f92.
//
// Solidity: function bitToken() view returns(address)
func (_Sequencer *SequencerSession) BitToken() (common.Address, error) {
	return _Sequencer.Contract.BitToken(&_Sequencer.CallOpts)
}

// BitToken is a free data retrieval call binding the contract method 0xd84e9f92.
//
// Solidity: function bitToken() view returns(address)
func (_Sequencer *SequencerCallerSession) BitToken() (common.Address, error) {
	return _Sequencer.Contract.BitToken(&_Sequencer.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Sequencer *SequencerCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Sequencer *SequencerSession) Epoch() (*big.Int, error) {
	return _Sequencer.Contract.Epoch(&_Sequencer.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Sequencer *SequencerCallerSession) Epoch() (*big.Int, error) {
	return _Sequencer.Contract.Epoch(&_Sequencer.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Sequencer *SequencerCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Sequencer *SequencerSession) GetOwners() ([]common.Address, error) {
	return _Sequencer.Contract.GetOwners(&_Sequencer.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Sequencer *SequencerCallerSession) GetOwners() ([]common.Address, error) {
	return _Sequencer.Contract.GetOwners(&_Sequencer.CallOpts)
}

// GetSequencer is a free data retrieval call binding the contract method 0xe90f218f.
//
// Solidity: function getSequencer(address signer) view returns((address,address,bytes,uint256,uint256))
func (_Sequencer *SequencerCaller) GetSequencer(opts *bind.CallOpts, signer common.Address) (SequencerSequencerInfo, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencer", signer)

	if err != nil {
		return *new(SequencerSequencerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SequencerSequencerInfo)).(*SequencerSequencerInfo)

	return out0, err

}

// GetSequencer is a free data retrieval call binding the contract method 0xe90f218f.
//
// Solidity: function getSequencer(address signer) view returns((address,address,bytes,uint256,uint256))
func (_Sequencer *SequencerSession) GetSequencer(signer common.Address) (SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencer(&_Sequencer.CallOpts, signer)
}

// GetSequencer is a free data retrieval call binding the contract method 0xe90f218f.
//
// Solidity: function getSequencer(address signer) view returns((address,address,bytes,uint256,uint256))
func (_Sequencer *SequencerCallerSession) GetSequencer(signer common.Address) (SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencer(&_Sequencer.CallOpts, signer)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns((address,address,bytes,uint256,uint256)[])
func (_Sequencer *SequencerCaller) GetSequencers(opts *bind.CallOpts) ([]SequencerSequencerInfo, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencers")

	if err != nil {
		return *new([]SequencerSequencerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]SequencerSequencerInfo)).(*[]SequencerSequencerInfo)

	return out0, err

}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns((address,address,bytes,uint256,uint256)[])
func (_Sequencer *SequencerSession) GetSequencers() ([]SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencers(&_Sequencer.CallOpts)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns((address,address,bytes,uint256,uint256)[])
func (_Sequencer *SequencerCallerSession) GetSequencers() ([]SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencers(&_Sequencer.CallOpts)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address signer) view returns(bool)
func (_Sequencer *SequencerCaller) IsSequencer(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "isSequencer", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address signer) view returns(bool)
func (_Sequencer *SequencerSession) IsSequencer(signer common.Address) (bool, error) {
	return _Sequencer.Contract.IsSequencer(&_Sequencer.CallOpts, signer)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address signer) view returns(bool)
func (_Sequencer *SequencerCallerSession) IsSequencer(signer common.Address) (bool, error) {
	return _Sequencer.Contract.IsSequencer(&_Sequencer.CallOpts, signer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sequencer *SequencerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sequencer *SequencerSession) Owner() (common.Address, error) {
	return _Sequencer.Contract.Owner(&_Sequencer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sequencer *SequencerCallerSession) Owner() (common.Address, error) {
	return _Sequencer.Contract.Owner(&_Sequencer.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Sequencer *SequencerCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Sequencer *SequencerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.Owners(&_Sequencer.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Sequencer *SequencerCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.Owners(&_Sequencer.CallOpts, arg0)
}

// Rel is a free data retrieval call binding the contract method 0xcab2ea2a.
//
// Solidity: function rel(address ) view returns(address)
func (_Sequencer *SequencerCaller) Rel(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "rel", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xcab2ea2a.
//
// Solidity: function rel(address ) view returns(address)
func (_Sequencer *SequencerSession) Rel(arg0 common.Address) (common.Address, error) {
	return _Sequencer.Contract.Rel(&_Sequencer.CallOpts, arg0)
}

// Rel is a free data retrieval call binding the contract method 0xcab2ea2a.
//
// Solidity: function rel(address ) view returns(address)
func (_Sequencer *SequencerCallerSession) Rel(arg0 common.Address) (common.Address, error) {
	return _Sequencer.Contract.Rel(&_Sequencer.CallOpts, arg0)
}

// Scheduler is a free data retrieval call binding the contract method 0xd1ad17bf.
//
// Solidity: function scheduler() view returns(address)
func (_Sequencer *SequencerCaller) Scheduler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "scheduler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Scheduler is a free data retrieval call binding the contract method 0xd1ad17bf.
//
// Solidity: function scheduler() view returns(address)
func (_Sequencer *SequencerSession) Scheduler() (common.Address, error) {
	return _Sequencer.Contract.Scheduler(&_Sequencer.CallOpts)
}

// Scheduler is a free data retrieval call binding the contract method 0xd1ad17bf.
//
// Solidity: function scheduler() view returns(address)
func (_Sequencer *SequencerCallerSession) Scheduler() (common.Address, error) {
	return _Sequencer.Contract.Scheduler(&_Sequencer.CallOpts)
}

// SequencerLimit is a free data retrieval call binding the contract method 0x9c13b6b5.
//
// Solidity: function sequencerLimit() view returns(uint8)
func (_Sequencer *SequencerCaller) SequencerLimit(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "sequencerLimit")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SequencerLimit is a free data retrieval call binding the contract method 0x9c13b6b5.
//
// Solidity: function sequencerLimit() view returns(uint8)
func (_Sequencer *SequencerSession) SequencerLimit() (uint8, error) {
	return _Sequencer.Contract.SequencerLimit(&_Sequencer.CallOpts)
}

// SequencerLimit is a free data retrieval call binding the contract method 0x9c13b6b5.
//
// Solidity: function sequencerLimit() view returns(uint8)
func (_Sequencer *SequencerCallerSession) SequencerLimit() (uint8, error) {
	return _Sequencer.Contract.SequencerLimit(&_Sequencer.CallOpts)
}

// Sequencers is a free data retrieval call binding the contract method 0x1c7a07ee.
//
// Solidity: function sequencers(address ) view returns(address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
func (_Sequencer *SequencerCaller) Sequencers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "sequencers", arg0)

	outstruct := new(struct {
		Owner       common.Address
		MintAddress common.Address
		NodeID      []byte
		Amount      *big.Int
		KeyIndex    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MintAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NodeID = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.KeyIndex = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Sequencers is a free data retrieval call binding the contract method 0x1c7a07ee.
//
// Solidity: function sequencers(address ) view returns(address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
func (_Sequencer *SequencerSession) Sequencers(arg0 common.Address) (struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}, error) {
	return _Sequencer.Contract.Sequencers(&_Sequencer.CallOpts, arg0)
}

// Sequencers is a free data retrieval call binding the contract method 0x1c7a07ee.
//
// Solidity: function sequencers(address ) view returns(address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
func (_Sequencer *SequencerCallerSession) Sequencers(arg0 common.Address) (struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}, error) {
	return _Sequencer.Contract.Sequencers(&_Sequencer.CallOpts, arg0)
}

// CreateSequencer is a paid mutator transaction binding the contract method 0x43dfc471.
//
// Solidity: function createSequencer(uint256 _amount, address _mintAddress, bytes _nodeID) returns()
func (_Sequencer *SequencerTransactor) CreateSequencer(opts *bind.TransactOpts, _amount *big.Int, _mintAddress common.Address, _nodeID []byte) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "createSequencer", _amount, _mintAddress, _nodeID)
}

// CreateSequencer is a paid mutator transaction binding the contract method 0x43dfc471.
//
// Solidity: function createSequencer(uint256 _amount, address _mintAddress, bytes _nodeID) returns()
func (_Sequencer *SequencerSession) CreateSequencer(_amount *big.Int, _mintAddress common.Address, _nodeID []byte) (*types.Transaction, error) {
	return _Sequencer.Contract.CreateSequencer(&_Sequencer.TransactOpts, _amount, _mintAddress, _nodeID)
}

// CreateSequencer is a paid mutator transaction binding the contract method 0x43dfc471.
//
// Solidity: function createSequencer(uint256 _amount, address _mintAddress, bytes _nodeID) returns()
func (_Sequencer *SequencerTransactorSession) CreateSequencer(_amount *big.Int, _mintAddress common.Address, _nodeID []byte) (*types.Transaction, error) {
	return _Sequencer.Contract.CreateSequencer(&_Sequencer.TransactOpts, _amount, _mintAddress, _nodeID)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Sequencer *SequencerTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Sequencer *SequencerSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.Deposit(&_Sequencer.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Sequencer *SequencerTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.Deposit(&_Sequencer.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bitToken) returns()
func (_Sequencer *SequencerTransactor) Initialize(opts *bind.TransactOpts, _bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "initialize", _bitToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bitToken) returns()
func (_Sequencer *SequencerSession) Initialize(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, _bitToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bitToken) returns()
func (_Sequencer *SequencerTransactorSession) Initialize(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, _bitToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sequencer *SequencerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sequencer *SequencerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sequencer.Contract.RenounceOwnership(&_Sequencer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sequencer *SequencerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sequencer.Contract.RenounceOwnership(&_Sequencer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sequencer *SequencerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sequencer *SequencerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.TransferOwnership(&_Sequencer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sequencer *SequencerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.TransferOwnership(&_Sequencer.TransactOpts, newOwner)
}

// UpdateBitAddress is a paid mutator transaction binding the contract method 0xee43b5d9.
//
// Solidity: function updateBitAddress(address _bitToken) returns()
func (_Sequencer *SequencerTransactor) UpdateBitAddress(opts *bind.TransactOpts, _bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "updateBitAddress", _bitToken)
}

// UpdateBitAddress is a paid mutator transaction binding the contract method 0xee43b5d9.
//
// Solidity: function updateBitAddress(address _bitToken) returns()
func (_Sequencer *SequencerSession) UpdateBitAddress(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateBitAddress(&_Sequencer.TransactOpts, _bitToken)
}

// UpdateBitAddress is a paid mutator transaction binding the contract method 0xee43b5d9.
//
// Solidity: function updateBitAddress(address _bitToken) returns()
func (_Sequencer *SequencerTransactorSession) UpdateBitAddress(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateBitAddress(&_Sequencer.TransactOpts, _bitToken)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x15ca0dc0.
//
// Solidity: function updateEpoch(uint256 _epoch) returns()
func (_Sequencer *SequencerTransactor) UpdateEpoch(opts *bind.TransactOpts, _epoch *big.Int) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "updateEpoch", _epoch)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x15ca0dc0.
//
// Solidity: function updateEpoch(uint256 _epoch) returns()
func (_Sequencer *SequencerSession) UpdateEpoch(_epoch *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateEpoch(&_Sequencer.TransactOpts, _epoch)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x15ca0dc0.
//
// Solidity: function updateEpoch(uint256 _epoch) returns()
func (_Sequencer *SequencerTransactorSession) UpdateEpoch(_epoch *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateEpoch(&_Sequencer.TransactOpts, _epoch)
}

// UpdateScheduler is a paid mutator transaction binding the contract method 0xbea0051d.
//
// Solidity: function updateScheduler(address _scheduler) returns()
func (_Sequencer *SequencerTransactor) UpdateScheduler(opts *bind.TransactOpts, _scheduler common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "updateScheduler", _scheduler)
}

// UpdateScheduler is a paid mutator transaction binding the contract method 0xbea0051d.
//
// Solidity: function updateScheduler(address _scheduler) returns()
func (_Sequencer *SequencerSession) UpdateScheduler(_scheduler common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateScheduler(&_Sequencer.TransactOpts, _scheduler)
}

// UpdateScheduler is a paid mutator transaction binding the contract method 0xbea0051d.
//
// Solidity: function updateScheduler(address _scheduler) returns()
func (_Sequencer *SequencerTransactorSession) UpdateScheduler(_scheduler common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateScheduler(&_Sequencer.TransactOpts, _scheduler)
}

// UpdateSequencerLimit is a paid mutator transaction binding the contract method 0x26b0c114.
//
// Solidity: function updateSequencerLimit(uint8 _limit) returns()
func (_Sequencer *SequencerTransactor) UpdateSequencerLimit(opts *bind.TransactOpts, _limit uint8) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "updateSequencerLimit", _limit)
}

// UpdateSequencerLimit is a paid mutator transaction binding the contract method 0x26b0c114.
//
// Solidity: function updateSequencerLimit(uint8 _limit) returns()
func (_Sequencer *SequencerSession) UpdateSequencerLimit(_limit uint8) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateSequencerLimit(&_Sequencer.TransactOpts, _limit)
}

// UpdateSequencerLimit is a paid mutator transaction binding the contract method 0x26b0c114.
//
// Solidity: function updateSequencerLimit(uint8 _limit) returns()
func (_Sequencer *SequencerTransactorSession) UpdateSequencerLimit(_limit uint8) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateSequencerLimit(&_Sequencer.TransactOpts, _limit)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Sequencer *SequencerTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Sequencer *SequencerSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.Withdraw(&_Sequencer.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Sequencer *SequencerTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.Withdraw(&_Sequencer.TransactOpts, _amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Sequencer *SequencerTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Sequencer *SequencerSession) WithdrawAll() (*types.Transaction, error) {
	return _Sequencer.Contract.WithdrawAll(&_Sequencer.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Sequencer *SequencerTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Sequencer.Contract.WithdrawAll(&_Sequencer.TransactOpts)
}

// SequencerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Sequencer contract.
type SequencerInitializedIterator struct {
	Event *SequencerInitialized // Event containing the contract specifics and raw log

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
func (it *SequencerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInitialized)
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
		it.Event = new(SequencerInitialized)
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
func (it *SequencerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInitialized represents a Initialized event raised by the Sequencer contract.
type SequencerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Sequencer *SequencerFilterer) FilterInitialized(opts *bind.FilterOpts) (*SequencerInitializedIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SequencerInitializedIterator{contract: _Sequencer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Sequencer *SequencerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SequencerInitialized) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInitialized)
				if err := _Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Sequencer *SequencerFilterer) ParseInitialized(log types.Log) (*SequencerInitialized, error) {
	event := new(SequencerInitialized)
	if err := _Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Sequencer contract.
type SequencerOwnershipTransferredIterator struct {
	Event *SequencerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SequencerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerOwnershipTransferred)
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
		it.Event = new(SequencerOwnershipTransferred)
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
func (it *SequencerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerOwnershipTransferred represents a OwnershipTransferred event raised by the Sequencer contract.
type SequencerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sequencer *SequencerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SequencerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerOwnershipTransferredIterator{contract: _Sequencer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sequencer *SequencerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SequencerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerOwnershipTransferred)
				if err := _Sequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sequencer *SequencerFilterer) ParseOwnershipTransferred(log types.Log) (*SequencerOwnershipTransferred, error) {
	event := new(SequencerOwnershipTransferred)
	if err := _Sequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerSequencerCreateIterator is returned from FilterSequencerCreate and is used to iterate over the raw logs and unpacked data for SequencerCreate events raised by the Sequencer contract.
type SequencerSequencerCreateIterator struct {
	Event *SequencerSequencerCreate // Event containing the contract specifics and raw log

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
func (it *SequencerSequencerCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerSequencerCreate)
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
		it.Event = new(SequencerSequencerCreate)
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
func (it *SequencerSequencerCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerSequencerCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerSequencerCreate represents a SequencerCreate event raised by the Sequencer contract.
type SequencerSequencerCreate struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSequencerCreate is a free log retrieval operation binding the contract event 0x4bc782d7487c41941075eb09650b2eb45a57e23e2241db6a958b8af8485324f6.
//
// Solidity: event SequencerCreate(address arg0, address arg1, bytes arg2)
func (_Sequencer *SequencerFilterer) FilterSequencerCreate(opts *bind.FilterOpts) (*SequencerSequencerCreateIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerCreate")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerCreateIterator{contract: _Sequencer.contract, event: "SequencerCreate", logs: logs, sub: sub}, nil
}

// WatchSequencerCreate is a free log subscription operation binding the contract event 0x4bc782d7487c41941075eb09650b2eb45a57e23e2241db6a958b8af8485324f6.
//
// Solidity: event SequencerCreate(address arg0, address arg1, bytes arg2)
func (_Sequencer *SequencerFilterer) WatchSequencerCreate(opts *bind.WatchOpts, sink chan<- *SequencerSequencerCreate) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "SequencerCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerSequencerCreate)
				if err := _Sequencer.contract.UnpackLog(event, "SequencerCreate", log); err != nil {
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

// ParseSequencerCreate is a log parse operation binding the contract event 0x4bc782d7487c41941075eb09650b2eb45a57e23e2241db6a958b8af8485324f6.
//
// Solidity: event SequencerCreate(address arg0, address arg1, bytes arg2)
func (_Sequencer *SequencerFilterer) ParseSequencerCreate(log types.Log) (*SequencerSequencerCreate, error) {
	event := new(SequencerSequencerCreate)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerCreate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerSequencerDeleteIterator is returned from FilterSequencerDelete and is used to iterate over the raw logs and unpacked data for SequencerDelete events raised by the Sequencer contract.
type SequencerSequencerDeleteIterator struct {
	Event *SequencerSequencerDelete // Event containing the contract specifics and raw log

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
func (it *SequencerSequencerDeleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerSequencerDelete)
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
		it.Event = new(SequencerSequencerDelete)
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
func (it *SequencerSequencerDeleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerSequencerDeleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerSequencerDelete represents a SequencerDelete event raised by the Sequencer contract.
type SequencerSequencerDelete struct {
	Arg0 common.Address
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSequencerDelete is a free log retrieval operation binding the contract event 0x37c49a6eae060065fbdeff05623918ea0969b4b1146b4c2fd33658caa72ed0db.
//
// Solidity: event SequencerDelete(address arg0, bytes arg1)
func (_Sequencer *SequencerFilterer) FilterSequencerDelete(opts *bind.FilterOpts) (*SequencerSequencerDeleteIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerDelete")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerDeleteIterator{contract: _Sequencer.contract, event: "SequencerDelete", logs: logs, sub: sub}, nil
}

// WatchSequencerDelete is a free log subscription operation binding the contract event 0x37c49a6eae060065fbdeff05623918ea0969b4b1146b4c2fd33658caa72ed0db.
//
// Solidity: event SequencerDelete(address arg0, bytes arg1)
func (_Sequencer *SequencerFilterer) WatchSequencerDelete(opts *bind.WatchOpts, sink chan<- *SequencerSequencerDelete) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "SequencerDelete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerSequencerDelete)
				if err := _Sequencer.contract.UnpackLog(event, "SequencerDelete", log); err != nil {
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

// ParseSequencerDelete is a log parse operation binding the contract event 0x37c49a6eae060065fbdeff05623918ea0969b4b1146b4c2fd33658caa72ed0db.
//
// Solidity: event SequencerDelete(address arg0, bytes arg1)
func (_Sequencer *SequencerFilterer) ParseSequencerDelete(log types.Log) (*SequencerSequencerDelete, error) {
	event := new(SequencerSequencerDelete)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerDelete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerSequencerUpdateIterator is returned from FilterSequencerUpdate and is used to iterate over the raw logs and unpacked data for SequencerUpdate events raised by the Sequencer contract.
type SequencerSequencerUpdateIterator struct {
	Event *SequencerSequencerUpdate // Event containing the contract specifics and raw log

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
func (it *SequencerSequencerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerSequencerUpdate)
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
		it.Event = new(SequencerSequencerUpdate)
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
func (it *SequencerSequencerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerSequencerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerSequencerUpdate represents a SequencerUpdate event raised by the Sequencer contract.
type SequencerSequencerUpdate struct {
	Arg0 common.Address
	Arg1 []byte
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSequencerUpdate is a free log retrieval operation binding the contract event 0xb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b16824.
//
// Solidity: event SequencerUpdate(address arg0, bytes arg1, uint256 arg2)
func (_Sequencer *SequencerFilterer) FilterSequencerUpdate(opts *bind.FilterOpts) (*SequencerSequencerUpdateIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerUpdate")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerUpdateIterator{contract: _Sequencer.contract, event: "SequencerUpdate", logs: logs, sub: sub}, nil
}

// WatchSequencerUpdate is a free log subscription operation binding the contract event 0xb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b16824.
//
// Solidity: event SequencerUpdate(address arg0, bytes arg1, uint256 arg2)
func (_Sequencer *SequencerFilterer) WatchSequencerUpdate(opts *bind.WatchOpts, sink chan<- *SequencerSequencerUpdate) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "SequencerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerSequencerUpdate)
				if err := _Sequencer.contract.UnpackLog(event, "SequencerUpdate", log); err != nil {
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

// ParseSequencerUpdate is a log parse operation binding the contract event 0xb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b16824.
//
// Solidity: event SequencerUpdate(address arg0, bytes arg1, uint256 arg2)
func (_Sequencer *SequencerFilterer) ParseSequencerUpdate(log types.Log) (*SequencerSequencerUpdate, error) {
	event := new(SequencerSequencerUpdate)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
