// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.ConvertType
)

// ComplexTypePeople is an auto generated low-level Go binding around an user-defined struct.
type ComplexTypePeople struct {
	Id            *big.Int
	Name          []byte
	Ages          *big.Int
	Sex           bool
	BirthTime     *big.Int
	PlayMonths    []uint16
	PlayMonthDays [][]*big.Int
}

// ComplexTypePeopleBase is an auto generated low-level Go binding around an user-defined struct.
type ComplexTypePeopleBase struct {
	Ages *big.Int
	Sex  bool
}

// ComplexTypePeopleBaseConfig is an auto generated low-level Go binding around an user-defined struct.
type ComplexTypePeopleBaseConfig struct {
	Name []byte
	Ages *big.Int
	Sex  bool
}

// ComplexTypePlayDayConfig is an auto generated low-level Go binding around an user-defined struct.
type ComplexTypePlayDayConfig struct {
	BirthTime     *big.Int
	PlayMonths    []uint16
	PlayMonthDays [][]*big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ages\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sex\",\"type\":\"bool\"}],\"name\":\"AddUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"birthTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"playMonths\",\"type\":\"uint16[]\"},{\"indexed\":false,\"internalType\":\"uint256[][]\",\"name\":\"playMonthDays\",\"type\":\"uint256[][]\"}],\"name\":\"AddUserPlayDays\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"playMonths\",\"type\":\"uint16[]\"},{\"indexed\":false,\"internalType\":\"uint256[][]\",\"name\":\"playMonthDays\",\"type\":\"uint256[][]\"}],\"name\":\"UpdateUserPlayMonth\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"idList\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"ages\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sex\",\"type\":\"bool\"}],\"internalType\":\"structComplexType.PeopleBaseConfig[]\",\"name\":\"baseConfigList\",\"type\":\"tuple[]\"}],\"name\":\"addBatchUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"idList\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"birthTime\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"playMonths\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"playMonthDays\",\"type\":\"uint256[][]\"}],\"internalType\":\"structComplexType.PlayDayConfig[]\",\"name\":\"dayConfigList\",\"type\":\"tuple[]\"}],\"name\":\"addBatchUserPlayDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"ages\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sex\",\"type\":\"bool\"}],\"internalType\":\"structComplexType.PeopleBaseConfig\",\"name\":\"baseConfig\",\"type\":\"tuple\"}],\"name\":\"addUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ages\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sex\",\"type\":\"bool\"}],\"internalType\":\"structComplexType.PeopleBase\",\"name\":\"base\",\"type\":\"tuple\"}],\"name\":\"addUserBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"birthTime\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"playMonths\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"playMonthDays\",\"type\":\"uint256[][]\"}],\"internalType\":\"structComplexType.PlayDayConfig\",\"name\":\"dayConfig\",\"type\":\"tuple\"}],\"name\":\"addUserPlayDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"idList\",\"type\":\"uint256[]\"}],\"name\":\"getBatchUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"ages\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sex\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"birthTime\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"playMonths\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"playMonthDays\",\"type\":\"uint256[][]\"}],\"internalType\":\"structComplexType.People[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"idList\",\"type\":\"uint256[]\"}],\"name\":\"getBatchUserBase\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ages\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sex\",\"type\":\"bool\"}],\"internalType\":\"structComplexType.PeopleBase[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"idList\",\"type\":\"uint256[]\"}],\"name\":\"getBatchUserPlayDays\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"birthTime\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"playMonths\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"playMonthDays\",\"type\":\"uint256[][]\"}],\"internalType\":\"structComplexType.PlayDayConfig[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"ages\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sex\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"birthTime\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"playMonths\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"playMonthDays\",\"type\":\"uint256[][]\"}],\"internalType\":\"structComplexType.People\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getUserPlayDays\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"birthTime\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"playMonths\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"playMonthDays\",\"type\":\"uint256[][]\"}],\"internalType\":\"structComplexType.PlayDayConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getUserPlayMonthSingle\",\"outputs\":[{\"internalType\":\"uint16[]\",\"name\":\"\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"playMonths\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"playMonthDays\",\"type\":\"uint256[][]\"}],\"name\":\"updateUserPlayMonth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetBatchUser is a free data retrieval call binding the contract method 0x8d103977.
//
// Solidity: function getBatchUser(uint256[] idList) view returns((uint256,bytes,uint256,bool,uint256,uint16[],uint256[][])[])
func (_Contract *ContractCaller) GetBatchUser(opts *bind.CallOpts, idList []*big.Int) ([]ComplexTypePeople, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBatchUser", idList)

	if err != nil {
		return *new([]ComplexTypePeople), err
	}

	out0 := *abi.ConvertType(out[0], new([]ComplexTypePeople)).(*[]ComplexTypePeople)

	return out0, err

}

// GetBatchUser is a free data retrieval call binding the contract method 0x8d103977.
//
// Solidity: function getBatchUser(uint256[] idList) view returns((uint256,bytes,uint256,bool,uint256,uint16[],uint256[][])[])
func (_Contract *ContractSession) GetBatchUser(idList []*big.Int) ([]ComplexTypePeople, error) {
	return _Contract.Contract.GetBatchUser(&_Contract.CallOpts, idList)
}

// GetBatchUser is a free data retrieval call binding the contract method 0x8d103977.
//
// Solidity: function getBatchUser(uint256[] idList) view returns((uint256,bytes,uint256,bool,uint256,uint16[],uint256[][])[])
func (_Contract *ContractCallerSession) GetBatchUser(idList []*big.Int) ([]ComplexTypePeople, error) {
	return _Contract.Contract.GetBatchUser(&_Contract.CallOpts, idList)
}

// GetBatchUserBase is a free data retrieval call binding the contract method 0x6c713d96.
//
// Solidity: function getBatchUserBase(uint256[] idList) view returns((uint256,bool)[])
func (_Contract *ContractCaller) GetBatchUserBase(opts *bind.CallOpts, idList []*big.Int) ([]ComplexTypePeopleBase, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBatchUserBase", idList)

	if err != nil {
		return *new([]ComplexTypePeopleBase), err
	}

	out0 := *abi.ConvertType(out[0], new([]ComplexTypePeopleBase)).(*[]ComplexTypePeopleBase)

	return out0, err

}

// GetBatchUserBase is a free data retrieval call binding the contract method 0x6c713d96.
//
// Solidity: function getBatchUserBase(uint256[] idList) view returns((uint256,bool)[])
func (_Contract *ContractSession) GetBatchUserBase(idList []*big.Int) ([]ComplexTypePeopleBase, error) {
	return _Contract.Contract.GetBatchUserBase(&_Contract.CallOpts, idList)
}

// GetBatchUserBase is a free data retrieval call binding the contract method 0x6c713d96.
//
// Solidity: function getBatchUserBase(uint256[] idList) view returns((uint256,bool)[])
func (_Contract *ContractCallerSession) GetBatchUserBase(idList []*big.Int) ([]ComplexTypePeopleBase, error) {
	return _Contract.Contract.GetBatchUserBase(&_Contract.CallOpts, idList)
}

// GetBatchUserPlayDays is a free data retrieval call binding the contract method 0xa6b0cbc4.
//
// Solidity: function getBatchUserPlayDays(uint256[] idList) view returns((uint256,uint16[],uint256[][])[])
func (_Contract *ContractCaller) GetBatchUserPlayDays(opts *bind.CallOpts, idList []*big.Int) ([]ComplexTypePlayDayConfig, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBatchUserPlayDays", idList)

	if err != nil {
		return *new([]ComplexTypePlayDayConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]ComplexTypePlayDayConfig)).(*[]ComplexTypePlayDayConfig)

	return out0, err

}

// GetBatchUserPlayDays is a free data retrieval call binding the contract method 0xa6b0cbc4.
//
// Solidity: function getBatchUserPlayDays(uint256[] idList) view returns((uint256,uint16[],uint256[][])[])
func (_Contract *ContractSession) GetBatchUserPlayDays(idList []*big.Int) ([]ComplexTypePlayDayConfig, error) {
	return _Contract.Contract.GetBatchUserPlayDays(&_Contract.CallOpts, idList)
}

// GetBatchUserPlayDays is a free data retrieval call binding the contract method 0xa6b0cbc4.
//
// Solidity: function getBatchUserPlayDays(uint256[] idList) view returns((uint256,uint16[],uint256[][])[])
func (_Contract *ContractCallerSession) GetBatchUserPlayDays(idList []*big.Int) ([]ComplexTypePlayDayConfig, error) {
	return _Contract.Contract.GetBatchUserPlayDays(&_Contract.CallOpts, idList)
}

// GetUser is a free data retrieval call binding the contract method 0xb0467deb.
//
// Solidity: function getUser(uint256 id) view returns((uint256,bytes,uint256,bool,uint256,uint16[],uint256[][]))
func (_Contract *ContractCaller) GetUser(opts *bind.CallOpts, id *big.Int) (ComplexTypePeople, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUser", id)

	if err != nil {
		return *new(ComplexTypePeople), err
	}

	out0 := *abi.ConvertType(out[0], new(ComplexTypePeople)).(*ComplexTypePeople)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0xb0467deb.
//
// Solidity: function getUser(uint256 id) view returns((uint256,bytes,uint256,bool,uint256,uint16[],uint256[][]))
func (_Contract *ContractSession) GetUser(id *big.Int) (ComplexTypePeople, error) {
	return _Contract.Contract.GetUser(&_Contract.CallOpts, id)
}

// GetUser is a free data retrieval call binding the contract method 0xb0467deb.
//
// Solidity: function getUser(uint256 id) view returns((uint256,bytes,uint256,bool,uint256,uint16[],uint256[][]))
func (_Contract *ContractCallerSession) GetUser(id *big.Int) (ComplexTypePeople, error) {
	return _Contract.Contract.GetUser(&_Contract.CallOpts, id)
}

// GetUserPlayDays is a free data retrieval call binding the contract method 0xef032201.
//
// Solidity: function getUserPlayDays(uint256 id) view returns((uint256,uint16[],uint256[][]))
func (_Contract *ContractCaller) GetUserPlayDays(opts *bind.CallOpts, id *big.Int) (ComplexTypePlayDayConfig, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUserPlayDays", id)

	if err != nil {
		return *new(ComplexTypePlayDayConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ComplexTypePlayDayConfig)).(*ComplexTypePlayDayConfig)

	return out0, err

}

// GetUserPlayDays is a free data retrieval call binding the contract method 0xef032201.
//
// Solidity: function getUserPlayDays(uint256 id) view returns((uint256,uint16[],uint256[][]))
func (_Contract *ContractSession) GetUserPlayDays(id *big.Int) (ComplexTypePlayDayConfig, error) {
	return _Contract.Contract.GetUserPlayDays(&_Contract.CallOpts, id)
}

// GetUserPlayDays is a free data retrieval call binding the contract method 0xef032201.
//
// Solidity: function getUserPlayDays(uint256 id) view returns((uint256,uint16[],uint256[][]))
func (_Contract *ContractCallerSession) GetUserPlayDays(id *big.Int) (ComplexTypePlayDayConfig, error) {
	return _Contract.Contract.GetUserPlayDays(&_Contract.CallOpts, id)
}

// GetUserPlayMonthSingle is a free data retrieval call binding the contract method 0xa5920d5c.
//
// Solidity: function getUserPlayMonthSingle(uint256 id) view returns(uint16[], uint256[][])
func (_Contract *ContractCaller) GetUserPlayMonthSingle(opts *bind.CallOpts, id *big.Int) ([]uint16, [][]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUserPlayMonthSingle", id)

	if err != nil {
		return *new([]uint16), *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)
	out1 := *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return out0, out1, err

}

// GetUserPlayMonthSingle is a free data retrieval call binding the contract method 0xa5920d5c.
//
// Solidity: function getUserPlayMonthSingle(uint256 id) view returns(uint16[], uint256[][])
func (_Contract *ContractSession) GetUserPlayMonthSingle(id *big.Int) ([]uint16, [][]*big.Int, error) {
	return _Contract.Contract.GetUserPlayMonthSingle(&_Contract.CallOpts, id)
}

// GetUserPlayMonthSingle is a free data retrieval call binding the contract method 0xa5920d5c.
//
// Solidity: function getUserPlayMonthSingle(uint256 id) view returns(uint16[], uint256[][])
func (_Contract *ContractCallerSession) GetUserPlayMonthSingle(id *big.Int) ([]uint16, [][]*big.Int, error) {
	return _Contract.Contract.GetUserPlayMonthSingle(&_Contract.CallOpts, id)
}

// AddBatchUser is a paid mutator transaction binding the contract method 0xa3d366cc.
//
// Solidity: function addBatchUser(uint256[] idList, (bytes,uint256,bool)[] baseConfigList) returns()
func (_Contract *ContractTransactor) AddBatchUser(opts *bind.TransactOpts, idList []*big.Int, baseConfigList []ComplexTypePeopleBaseConfig) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addBatchUser", idList, baseConfigList)
}

// AddBatchUser is a paid mutator transaction binding the contract method 0xa3d366cc.
//
// Solidity: function addBatchUser(uint256[] idList, (bytes,uint256,bool)[] baseConfigList) returns()
func (_Contract *ContractSession) AddBatchUser(idList []*big.Int, baseConfigList []ComplexTypePeopleBaseConfig) (*types.Transaction, error) {
	return _Contract.Contract.AddBatchUser(&_Contract.TransactOpts, idList, baseConfigList)
}

// AddBatchUser is a paid mutator transaction binding the contract method 0xa3d366cc.
//
// Solidity: function addBatchUser(uint256[] idList, (bytes,uint256,bool)[] baseConfigList) returns()
func (_Contract *ContractTransactorSession) AddBatchUser(idList []*big.Int, baseConfigList []ComplexTypePeopleBaseConfig) (*types.Transaction, error) {
	return _Contract.Contract.AddBatchUser(&_Contract.TransactOpts, idList, baseConfigList)
}

// AddBatchUserPlayDays is a paid mutator transaction binding the contract method 0x66ea21f5.
//
// Solidity: function addBatchUserPlayDays(uint256[] idList, (uint256,uint16[],uint256[][])[] dayConfigList) returns()
func (_Contract *ContractTransactor) AddBatchUserPlayDays(opts *bind.TransactOpts, idList []*big.Int, dayConfigList []ComplexTypePlayDayConfig) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addBatchUserPlayDays", idList, dayConfigList)
}

// AddBatchUserPlayDays is a paid mutator transaction binding the contract method 0x66ea21f5.
//
// Solidity: function addBatchUserPlayDays(uint256[] idList, (uint256,uint16[],uint256[][])[] dayConfigList) returns()
func (_Contract *ContractSession) AddBatchUserPlayDays(idList []*big.Int, dayConfigList []ComplexTypePlayDayConfig) (*types.Transaction, error) {
	return _Contract.Contract.AddBatchUserPlayDays(&_Contract.TransactOpts, idList, dayConfigList)
}

// AddBatchUserPlayDays is a paid mutator transaction binding the contract method 0x66ea21f5.
//
// Solidity: function addBatchUserPlayDays(uint256[] idList, (uint256,uint16[],uint256[][])[] dayConfigList) returns()
func (_Contract *ContractTransactorSession) AddBatchUserPlayDays(idList []*big.Int, dayConfigList []ComplexTypePlayDayConfig) (*types.Transaction, error) {
	return _Contract.Contract.AddBatchUserPlayDays(&_Contract.TransactOpts, idList, dayConfigList)
}

// AddUser is a paid mutator transaction binding the contract method 0x579ec2fe.
//
// Solidity: function addUser(uint256 id, (bytes,uint256,bool) baseConfig) returns()
func (_Contract *ContractTransactor) AddUser(opts *bind.TransactOpts, id *big.Int, baseConfig ComplexTypePeopleBaseConfig) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addUser", id, baseConfig)
}

// AddUser is a paid mutator transaction binding the contract method 0x579ec2fe.
//
// Solidity: function addUser(uint256 id, (bytes,uint256,bool) baseConfig) returns()
func (_Contract *ContractSession) AddUser(id *big.Int, baseConfig ComplexTypePeopleBaseConfig) (*types.Transaction, error) {
	return _Contract.Contract.AddUser(&_Contract.TransactOpts, id, baseConfig)
}

// AddUser is a paid mutator transaction binding the contract method 0x579ec2fe.
//
// Solidity: function addUser(uint256 id, (bytes,uint256,bool) baseConfig) returns()
func (_Contract *ContractTransactorSession) AddUser(id *big.Int, baseConfig ComplexTypePeopleBaseConfig) (*types.Transaction, error) {
	return _Contract.Contract.AddUser(&_Contract.TransactOpts, id, baseConfig)
}

// AddUserBase is a paid mutator transaction binding the contract method 0x3f570223.
//
// Solidity: function addUserBase(uint256 id, (uint256,bool) base) returns()
func (_Contract *ContractTransactor) AddUserBase(opts *bind.TransactOpts, id *big.Int, base ComplexTypePeopleBase) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addUserBase", id, base)
}

// AddUserBase is a paid mutator transaction binding the contract method 0x3f570223.
//
// Solidity: function addUserBase(uint256 id, (uint256,bool) base) returns()
func (_Contract *ContractSession) AddUserBase(id *big.Int, base ComplexTypePeopleBase) (*types.Transaction, error) {
	return _Contract.Contract.AddUserBase(&_Contract.TransactOpts, id, base)
}

// AddUserBase is a paid mutator transaction binding the contract method 0x3f570223.
//
// Solidity: function addUserBase(uint256 id, (uint256,bool) base) returns()
func (_Contract *ContractTransactorSession) AddUserBase(id *big.Int, base ComplexTypePeopleBase) (*types.Transaction, error) {
	return _Contract.Contract.AddUserBase(&_Contract.TransactOpts, id, base)
}

// AddUserPlayDays is a paid mutator transaction binding the contract method 0xbdf57ae5.
//
// Solidity: function addUserPlayDays(uint256 id, (uint256,uint16[],uint256[][]) dayConfig) returns()
func (_Contract *ContractTransactor) AddUserPlayDays(opts *bind.TransactOpts, id *big.Int, dayConfig ComplexTypePlayDayConfig) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addUserPlayDays", id, dayConfig)
}

// AddUserPlayDays is a paid mutator transaction binding the contract method 0xbdf57ae5.
//
// Solidity: function addUserPlayDays(uint256 id, (uint256,uint16[],uint256[][]) dayConfig) returns()
func (_Contract *ContractSession) AddUserPlayDays(id *big.Int, dayConfig ComplexTypePlayDayConfig) (*types.Transaction, error) {
	return _Contract.Contract.AddUserPlayDays(&_Contract.TransactOpts, id, dayConfig)
}

// AddUserPlayDays is a paid mutator transaction binding the contract method 0xbdf57ae5.
//
// Solidity: function addUserPlayDays(uint256 id, (uint256,uint16[],uint256[][]) dayConfig) returns()
func (_Contract *ContractTransactorSession) AddUserPlayDays(id *big.Int, dayConfig ComplexTypePlayDayConfig) (*types.Transaction, error) {
	return _Contract.Contract.AddUserPlayDays(&_Contract.TransactOpts, id, dayConfig)
}

// UpdateUserPlayMonth is a paid mutator transaction binding the contract method 0xa9ee5e88.
//
// Solidity: function updateUserPlayMonth(uint256 id, uint16[] playMonths, uint256[][] playMonthDays) returns()
func (_Contract *ContractTransactor) UpdateUserPlayMonth(opts *bind.TransactOpts, id *big.Int, playMonths []uint16, playMonthDays [][]*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateUserPlayMonth", id, playMonths, playMonthDays)
}

// UpdateUserPlayMonth is a paid mutator transaction binding the contract method 0xa9ee5e88.
//
// Solidity: function updateUserPlayMonth(uint256 id, uint16[] playMonths, uint256[][] playMonthDays) returns()
func (_Contract *ContractSession) UpdateUserPlayMonth(id *big.Int, playMonths []uint16, playMonthDays [][]*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateUserPlayMonth(&_Contract.TransactOpts, id, playMonths, playMonthDays)
}

// UpdateUserPlayMonth is a paid mutator transaction binding the contract method 0xa9ee5e88.
//
// Solidity: function updateUserPlayMonth(uint256 id, uint16[] playMonths, uint256[][] playMonthDays) returns()
func (_Contract *ContractTransactorSession) UpdateUserPlayMonth(id *big.Int, playMonths []uint16, playMonthDays [][]*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateUserPlayMonth(&_Contract.TransactOpts, id, playMonths, playMonthDays)
}

// ContractAddUserIterator is returned from FilterAddUser and is used to iterate over the raw logs and unpacked data for AddUser events raised by the Contract contract.
type ContractAddUserIterator struct {
	Event *ContractAddUser // Event containing the contract specifics and raw log

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
func (it *ContractAddUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddUser)
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
		it.Event = new(ContractAddUser)
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
func (it *ContractAddUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddUser represents a AddUser event raised by the Contract contract.
type ContractAddUser struct {
	Id   *big.Int
	Name []byte
	Ages *big.Int
	Sex  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddUser is a free log retrieval operation binding the contract event 0x30f474585b7e52f14d19a576697cf807f23026834a0be4a258153b935dd07795.
//
// Solidity: event AddUser(uint256 indexed id, bytes name, uint256 ages, bool sex)
func (_Contract *ContractFilterer) FilterAddUser(opts *bind.FilterOpts, id []*big.Int) (*ContractAddUserIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddUser", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractAddUserIterator{contract: _Contract.contract, event: "AddUser", logs: logs, sub: sub}, nil
}

// WatchAddUser is a free log subscription operation binding the contract event 0x30f474585b7e52f14d19a576697cf807f23026834a0be4a258153b935dd07795.
//
// Solidity: event AddUser(uint256 indexed id, bytes name, uint256 ages, bool sex)
func (_Contract *ContractFilterer) WatchAddUser(opts *bind.WatchOpts, sink chan<- *ContractAddUser, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddUser", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddUser)
				if err := _Contract.contract.UnpackLog(event, "AddUser", log); err != nil {
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

// ParseAddUser is a log parse operation binding the contract event 0x30f474585b7e52f14d19a576697cf807f23026834a0be4a258153b935dd07795.
//
// Solidity: event AddUser(uint256 indexed id, bytes name, uint256 ages, bool sex)
func (_Contract *ContractFilterer) ParseAddUser(log types.Log) (*ContractAddUser, error) {
	event := new(ContractAddUser)
	if err := _Contract.contract.UnpackLog(event, "AddUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAddUserPlayDaysIterator is returned from FilterAddUserPlayDays and is used to iterate over the raw logs and unpacked data for AddUserPlayDays events raised by the Contract contract.
type ContractAddUserPlayDaysIterator struct {
	Event *ContractAddUserPlayDays // Event containing the contract specifics and raw log

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
func (it *ContractAddUserPlayDaysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddUserPlayDays)
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
		it.Event = new(ContractAddUserPlayDays)
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
func (it *ContractAddUserPlayDaysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddUserPlayDaysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddUserPlayDays represents a AddUserPlayDays event raised by the Contract contract.
type ContractAddUserPlayDays struct {
	Id            *big.Int
	BirthTime     *big.Int
	PlayMonths    []uint16
	PlayMonthDays [][]*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddUserPlayDays is a free log retrieval operation binding the contract event 0x4a2287ccd406bfce29a4ae415024db55e84b5dd7459150d7b5c913f9d1a4e4f2.
//
// Solidity: event AddUserPlayDays(uint256 indexed id, uint256 birthTime, uint16[] playMonths, uint256[][] playMonthDays)
func (_Contract *ContractFilterer) FilterAddUserPlayDays(opts *bind.FilterOpts, id []*big.Int) (*ContractAddUserPlayDaysIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddUserPlayDays", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractAddUserPlayDaysIterator{contract: _Contract.contract, event: "AddUserPlayDays", logs: logs, sub: sub}, nil
}

// WatchAddUserPlayDays is a free log subscription operation binding the contract event 0x4a2287ccd406bfce29a4ae415024db55e84b5dd7459150d7b5c913f9d1a4e4f2.
//
// Solidity: event AddUserPlayDays(uint256 indexed id, uint256 birthTime, uint16[] playMonths, uint256[][] playMonthDays)
func (_Contract *ContractFilterer) WatchAddUserPlayDays(opts *bind.WatchOpts, sink chan<- *ContractAddUserPlayDays, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddUserPlayDays", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddUserPlayDays)
				if err := _Contract.contract.UnpackLog(event, "AddUserPlayDays", log); err != nil {
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

// ParseAddUserPlayDays is a log parse operation binding the contract event 0x4a2287ccd406bfce29a4ae415024db55e84b5dd7459150d7b5c913f9d1a4e4f2.
//
// Solidity: event AddUserPlayDays(uint256 indexed id, uint256 birthTime, uint16[] playMonths, uint256[][] playMonthDays)
func (_Contract *ContractFilterer) ParseAddUserPlayDays(log types.Log) (*ContractAddUserPlayDays, error) {
	event := new(ContractAddUserPlayDays)
	if err := _Contract.contract.UnpackLog(event, "AddUserPlayDays", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateUserPlayMonthIterator is returned from FilterUpdateUserPlayMonth and is used to iterate over the raw logs and unpacked data for UpdateUserPlayMonth events raised by the Contract contract.
type ContractUpdateUserPlayMonthIterator struct {
	Event *ContractUpdateUserPlayMonth // Event containing the contract specifics and raw log

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
func (it *ContractUpdateUserPlayMonthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateUserPlayMonth)
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
		it.Event = new(ContractUpdateUserPlayMonth)
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
func (it *ContractUpdateUserPlayMonthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateUserPlayMonthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateUserPlayMonth represents a UpdateUserPlayMonth event raised by the Contract contract.
type ContractUpdateUserPlayMonth struct {
	Id            *big.Int
	PlayMonths    []uint16
	PlayMonthDays [][]*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateUserPlayMonth is a free log retrieval operation binding the contract event 0x301d29edacc20c15bc4369418d5521c90dd124eb3a04828a045824ff02a1f7be.
//
// Solidity: event UpdateUserPlayMonth(uint256 indexed id, uint16[] playMonths, uint256[][] playMonthDays)
func (_Contract *ContractFilterer) FilterUpdateUserPlayMonth(opts *bind.FilterOpts, id []*big.Int) (*ContractUpdateUserPlayMonthIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateUserPlayMonth", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdateUserPlayMonthIterator{contract: _Contract.contract, event: "UpdateUserPlayMonth", logs: logs, sub: sub}, nil
}

// WatchUpdateUserPlayMonth is a free log subscription operation binding the contract event 0x301d29edacc20c15bc4369418d5521c90dd124eb3a04828a045824ff02a1f7be.
//
// Solidity: event UpdateUserPlayMonth(uint256 indexed id, uint16[] playMonths, uint256[][] playMonthDays)
func (_Contract *ContractFilterer) WatchUpdateUserPlayMonth(opts *bind.WatchOpts, sink chan<- *ContractUpdateUserPlayMonth, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateUserPlayMonth", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateUserPlayMonth)
				if err := _Contract.contract.UnpackLog(event, "UpdateUserPlayMonth", log); err != nil {
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

// ParseUpdateUserPlayMonth is a log parse operation binding the contract event 0x301d29edacc20c15bc4369418d5521c90dd124eb3a04828a045824ff02a1f7be.
//
// Solidity: event UpdateUserPlayMonth(uint256 indexed id, uint16[] playMonths, uint256[][] playMonthDays)
func (_Contract *ContractFilterer) ParseUpdateUserPlayMonth(log types.Log) (*ContractUpdateUserPlayMonth, error) {
	event := new(ContractUpdateUserPlayMonth)
	if err := _Contract.contract.UnpackLog(event, "UpdateUserPlayMonth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
