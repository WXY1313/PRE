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

// VerificationG1Point is an auto generated low-level Go binding around an user-defined struct.
type VerificationG1Point struct {
	X *big.Int
	Y *big.Int
}

// VerificationG2Point is an auto generated low-level Go binding around an user-defined struct.
type VerificationG2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"pkArray\",\"type\":\"tuple[]\"}],\"name\":\"UploadDRsKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"pk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"vk\",\"type\":\"tuple\"}],\"name\":\"UploadOwnerKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"tau1\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point[]\",\"name\":\"tau2\",\"type\":\"tuple[]\"}],\"name\":\"UploadSystemKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"pk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"vk\",\"type\":\"tuple\"}],\"name\":\"UploadUserKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506106888061001f6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630359a5dd146100515780633ae15490146100665780639da5cd0014610079578063f029c25e1461008c575b600080fd5b61006461005f366004610461565b61009f565b005b610064610074366004610531565b6100de565b61006461008736600461056e565b610138565b61006461009a366004610461565b610258565b815160015560208201516002908155815182916003916100c191839190610275565b5060208201516100d79060028084019190610275565b5050505050565b60005b815181101561013457600d8282815181106100fe576100fe61063c565b602090810291909101810151825460018181018555600094855293839020825160029092020190815591015190820155016100e1565b5050565b6101406102b3565b60005b83518110156102525783818151811061015e5761015e61063c565b602002602001015182600001819052508281815181106101805761018061063c565b602090810291909101810151838201908152600080546001810182559080528451805160069092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56381019283559301517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5648401559051805185937f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e565019061022b9082906002610275565b5060208201516102419060028084019190610275565b505060019093019250610143915050565b50505050565b81516007556020820151600855805181906009906100c190829060025b82600281019282156102a3579160200282015b828111156102a3578251825591602001919060010190610288565b506102af9291506102e1565b5090565b6040805160808101825260009181018281526060820192909252908152602081016102dc6102f6565b905290565b5b808211156102af57600081556001016102e2565b6040518060400160405280610309610312565b81526020016102dc5b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561036957610369610330565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561039857610398610330565b604052919050565b6000604082840312156103b257600080fd5b6103ba610346565b823581526020928301359281019290925250919050565b600082601f8301126103e257600080fd5b6103ea610346565b8060408401858111156103fc57600080fd5b845b818110156104165780358452602093840193016103fe565b509095945050505050565b60006080828403121561043357600080fd5b61043b610346565b905061044783836103d1565b815261045683604084016103d1565b602082015292915050565b60008060c0838503121561047457600080fd5b61047e84846103a0565b915061048d8460408501610421565b90509250929050565b600067ffffffffffffffff8211156104b0576104b0610330565b5060051b60200190565b600082601f8301126104cb57600080fd5b81356104de6104d982610496565b61036f565b8082825260208201915060208360061b86010192508583111561050057600080fd5b602085015b838110156105275761051787826103a0565b8352602090920191604001610505565b5095945050505050565b60006020828403121561054357600080fd5b813567ffffffffffffffff81111561055a57600080fd5b610566848285016104ba565b949350505050565b6000806040838503121561058157600080fd5b823567ffffffffffffffff81111561059857600080fd5b6105a4858286016104ba565b925050602083013567ffffffffffffffff8111156105c157600080fd5b8301601f810185136105d257600080fd5b80356105e06104d982610496565b8082825260208201915060208360071b85010192508783111561060257600080fd5b6020840193505b8284101561062e5761061b8885610421565b8252602082019150608084019350610609565b809450505050509250929050565b634e487b7160e01b600052603260045260246000fdfea2646970667358221220503fd1c0ab3053e7cdfa19535a871b07035b94f14e1dec0d8b161021b326e2b264736f6c634300081a0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

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

// UploadDRsKey is a paid mutator transaction binding the contract method 0x3ae15490.
//
// Solidity: function UploadDRsKey((uint256,uint256)[] pkArray) returns()
func (_Contract *ContractTransactor) UploadDRsKey(opts *bind.TransactOpts, pkArray []VerificationG1Point) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "UploadDRsKey", pkArray)
}

// UploadDRsKey is a paid mutator transaction binding the contract method 0x3ae15490.
//
// Solidity: function UploadDRsKey((uint256,uint256)[] pkArray) returns()
func (_Contract *ContractSession) UploadDRsKey(pkArray []VerificationG1Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadDRsKey(&_Contract.TransactOpts, pkArray)
}

// UploadDRsKey is a paid mutator transaction binding the contract method 0x3ae15490.
//
// Solidity: function UploadDRsKey((uint256,uint256)[] pkArray) returns()
func (_Contract *ContractTransactorSession) UploadDRsKey(pkArray []VerificationG1Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadDRsKey(&_Contract.TransactOpts, pkArray)
}

// UploadOwnerKey is a paid mutator transaction binding the contract method 0x0359a5dd.
//
// Solidity: function UploadOwnerKey((uint256,uint256) pk, (uint256[2],uint256[2]) vk) returns()
func (_Contract *ContractTransactor) UploadOwnerKey(opts *bind.TransactOpts, pk VerificationG1Point, vk VerificationG2Point) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "UploadOwnerKey", pk, vk)
}

// UploadOwnerKey is a paid mutator transaction binding the contract method 0x0359a5dd.
//
// Solidity: function UploadOwnerKey((uint256,uint256) pk, (uint256[2],uint256[2]) vk) returns()
func (_Contract *ContractSession) UploadOwnerKey(pk VerificationG1Point, vk VerificationG2Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadOwnerKey(&_Contract.TransactOpts, pk, vk)
}

// UploadOwnerKey is a paid mutator transaction binding the contract method 0x0359a5dd.
//
// Solidity: function UploadOwnerKey((uint256,uint256) pk, (uint256[2],uint256[2]) vk) returns()
func (_Contract *ContractTransactorSession) UploadOwnerKey(pk VerificationG1Point, vk VerificationG2Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadOwnerKey(&_Contract.TransactOpts, pk, vk)
}

// UploadSystemKey is a paid mutator transaction binding the contract method 0x9da5cd00.
//
// Solidity: function UploadSystemKey((uint256,uint256)[] tau1, (uint256[2],uint256[2])[] tau2) returns()
func (_Contract *ContractTransactor) UploadSystemKey(opts *bind.TransactOpts, tau1 []VerificationG1Point, tau2 []VerificationG2Point) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "UploadSystemKey", tau1, tau2)
}

// UploadSystemKey is a paid mutator transaction binding the contract method 0x9da5cd00.
//
// Solidity: function UploadSystemKey((uint256,uint256)[] tau1, (uint256[2],uint256[2])[] tau2) returns()
func (_Contract *ContractSession) UploadSystemKey(tau1 []VerificationG1Point, tau2 []VerificationG2Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadSystemKey(&_Contract.TransactOpts, tau1, tau2)
}

// UploadSystemKey is a paid mutator transaction binding the contract method 0x9da5cd00.
//
// Solidity: function UploadSystemKey((uint256,uint256)[] tau1, (uint256[2],uint256[2])[] tau2) returns()
func (_Contract *ContractTransactorSession) UploadSystemKey(tau1 []VerificationG1Point, tau2 []VerificationG2Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadSystemKey(&_Contract.TransactOpts, tau1, tau2)
}

// UploadUserKey is a paid mutator transaction binding the contract method 0xf029c25e.
//
// Solidity: function UploadUserKey((uint256,uint256) pk, (uint256[2],uint256[2]) vk) returns()
func (_Contract *ContractTransactor) UploadUserKey(opts *bind.TransactOpts, pk VerificationG1Point, vk VerificationG2Point) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "UploadUserKey", pk, vk)
}

// UploadUserKey is a paid mutator transaction binding the contract method 0xf029c25e.
//
// Solidity: function UploadUserKey((uint256,uint256) pk, (uint256[2],uint256[2]) vk) returns()
func (_Contract *ContractSession) UploadUserKey(pk VerificationG1Point, vk VerificationG2Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadUserKey(&_Contract.TransactOpts, pk, vk)
}

// UploadUserKey is a paid mutator transaction binding the contract method 0xf029c25e.
//
// Solidity: function UploadUserKey((uint256,uint256) pk, (uint256[2],uint256[2]) vk) returns()
func (_Contract *ContractTransactorSession) UploadUserKey(pk VerificationG1Point, vk VerificationG2Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadUserKey(&_Contract.TransactOpts, pk, vk)
}
