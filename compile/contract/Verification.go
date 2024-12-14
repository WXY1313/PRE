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
	ABI: "[{\"inputs\":[],\"name\":\"GetFieldModulus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetTestMul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetVKoResult\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TestECTwistMul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"pkArray\",\"type\":\"tuple[]\"}],\"name\":\"UploadDRsKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"pk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"vk\",\"type\":\"tuple\"}],\"name\":\"UploadOwnerKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"tau1\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point[]\",\"name\":\"tau2\",\"type\":\"tuple[]\"}],\"name\":\"UploadSystemKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"pk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"vk\",\"type\":\"tuple\"}],\"name\":\"UploadUserKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"a1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"a2\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"UploadVKoDLEQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"a1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"a2\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"UploadVKuDLEQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VKoVerify\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061173e8061001f6000396000f3fe60806040526004361061009c5760003560e01c806355a3e90f1161006457806355a3e90f1461015057806390caa16f146101795780639da5cd0014610199578063a695a6b4146101b9578063e29fb4d4146101ce578063f029c25e146101ee57600080fd5b80630359a5dd146100a157806308ac6601146100c35780631357d59b146100cb578063184eecf0146100f65780633ae1549014610130575b600080fd5b3480156100ad57600080fd5b506100c16100bc366004611403565b61020e565b005b6100c161024d565b3480156100d757600080fd5b506100e06104f4565b6040516100ed9190611438565b60405180910390f35b34801561010257600080fd5b506020546021546022546023545b6040805194855260208501939093529183015260608201526080016100ed565b34801561013c57600080fd5b506100c161014b36600461150d565b61056c565b34801561015c57600080fd5b506040516000805160206116e983398151915281526020016100ed565b34801561018557600080fd5b506100c161019436600461154a565b6105c6565b3480156101a557600080fd5b506100c16101b436600461158f565b61060d565b3480156101c557600080fd5b5061011061072d565b3480156101da57600080fd5b506100c16101e936600461154a565b610767565b3480156101fa57600080fd5b506100c1610209366004611403565b6107ae565b81516001556020820151600290815581518291600391610230918391906111b9565b50602082015161024690600280840191906111b9565b5050505050565b6000610297600080815481106102655761026561165d565b6000918252602091829020604080518082019091526006909202018054825260010154918101919091526015546107cf565b60408051808201909152600154815260025460208201526014549192506000916102c191906107cf565b905060006102cf8383610824565b8051600e549192501415806102ea57506020810151600f5414155b1561033657601f805460018101825560008290527fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d80760208204018054919092166101000a60ff02191690555b601f805460018101825560008290527fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d80760208204018054919092166101000a60ff8102199091161790556103886111f7565b6103906111f7565b6103986111f7565b610449600e60070154600080815481106103b4576103b461165d565b6000918252602082206002600690920201010154600080815481106103db576103db61165d565b6000918252602090912060026006909202010160010154600080815481106104055761040561165d565b60009182526020822060046006909202010101546000808154811061042c5761042c61165d565b6000918252602090912060046006909202010160015b015461087c565b8651602080890151808201939093529290915290810191909152526014546003546004546005805461047f949392916001610442565b8551602080880180518083019490945293909252818101939093529290915284518051908301518684015180519085015187518051960151945180516104d5979596949593949293929160016020020151610926565b8451602095860151808701929092529190529283015290525050505050565b6060601f80548060200260200160405190810160405280929190818152602001828054801561056257602002820191906000526020600020906000905b825461010083900a900460ff1615158152602060019283018181049485019490930390920291018084116105315790505b5050505050905090565b60005b81518110156105c257600d82828151811061058c5761058c61165d565b6020908102919091018101518254600181810185556000948552938390208251600290920201908155910151908201550161056f565b5050565b8351600e556020840151600f55825183906010906105e790829060026111b9565b5060208201516105fd90600280840191906111b9565b5050506014919091556015555050565b61061561121c565b60005b8351811015610727578381815181106106335761063361165d565b602002602001015182600001819052508281815181106106555761065561165d565b602090810291909101810151838201908152600080546001810182559080528451805160069092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56381019283559301517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5648401559051805185937f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e565019061070090829060026111b9565b50602082015161071690600280840191906111b9565b505060019093019250610618915050565b50505050565b6000806000806107456005600160026003600461087c565b6023819055602282905560218390556020849055929791965094509092509050565b835160165560208401516017558251839060189061078890829060026111b9565b50602082015161079e90600280840191906111b9565b505050601c91909155601d555050565b815160075560208201516008558051819060099061023090829060026111b9565b60408051808201909152600080825260208201526107eb611245565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa90508061081c57600080fd5b505092915050565b6040805180820190915260008082526020820152610840611263565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa90508061081c57600080fd5b600080808060018815801561088f575087155b8015610899575086155b80156108a3575085155b156108b757506001975087955060006108cf565b6108c389898989610a59565b6108cf576108cf611673565b60006108e18b8b8b8b8b876000610b0e565b90506109118160005b602090810291909101519083015160408401516060850151608086015160a0870151610b91565b929e919d509b50909950975050505050505050565b60008080808b15801561093757508a155b8015610941575089155b801561094b575088155b1561099c578715801561095c575086155b8015610966575085155b8015610970575084155b61098c5761098088888888610a59565b61098c5761098c611673565b5086925085915084905083610a4a565b871580156109a8575086155b80156109b2575085155b80156109bc575084155b156109e9576109cd8c8c8c8c610a59565b6109d9576109d9611673565b508a925089915088905087610a4a565b6109f58c8c8c8c610a59565b610a0157610a01611673565b610a0d88888888610a59565b610a1957610a19611673565b6000610a338d8d8d8d600160008f8f8f8f60016000610bdb565b9050610a408160006108ea565b9450945094509450505b98509850985098945050505050565b6000806000806000610a6d87878989610e64565b9094509250610a7e89898181610e64565b9092509050610a8f82828b8b610e64565b9092509050610aa084848484610ed5565b9094509250610af084847f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e57e9713b03af0fed4cd2cafadeed8fdf4a74fa084e52d1852e4a2bd0685c315d2610ed5565b909450925083158015610b01575082155b9998505050505050505050565b610b16611281565b8715610b86576001881615610b57578051602082015160408301516060840151608085015160a0860151610b549594939291908d8d8d8d8d8d610bdb565b90505b610b65878787878787610f17565b949b50929950909750955093509150610b7f60028961169f565b9750610b16565b979650505050505050565b600080600080600080610ba48888611086565b9092509050610bb58c8c8484610e64565b9096509450610bc68a8a8484610e64565b969d959c509a50949850929650505050505050565b610be3611281565b88158015610bef575087155b15610c31578686868686868660005b60a08901929092526080880192909252606087019290925260408601929092526020858101939093529091020152610e54565b82158015610c3d575081155b15610c50578c8c8c8c8c8c866000610bfe565b610c5c85858b8b610e64565b9095509350610c6d8b8b8585610e64565b60608301526040820152610c8387878b8b610e64565b9097509550610c948d8d8585610e64565b60a08301526080820181905287148015610cb1575060a081015186145b15610cf657604081015185148015610ccc5750606081015184145b15610ce757610cdf8d8d8d8d8d8d610f17565b866000610bfe565b60016000818180808681610bfe565b610d0289898585610e64565b9093509150610d22858583600260200201518460035b6020020151610ed5565b909d509b50610d3c87878360046020020151846005610d18565b909b509950610d4d8b8b8181610e64565b9099509750610d6d898983600460200201518460055b6020020151610e64565b9095509350610d7e89898d8d610e64565b9099509750610d8f89898585610e64565b60a08301526080820152610da58d8d8181610e64565b9097509550610db687878585610e64565b9097509550610dc787878b8b610ed5565b9097509550610dd885856002611111565b9093509150610de987878585610ed5565b9097509550610dfa8b8b8989610e64565b60208301528152610e0d85858989610ed5565b909b509950610e1e8d8d8d8d610e64565b909b509950610e3889898360026020020151846003610d63565b909d509b50610e498b8b8f8f610ed5565b606083015260408201525b9c9b505050505050505050505050565b600080610ea26000805160206116e98339815191528588096000805160206116e98339815191528588096000805160206116e9833981519152611144565b6000805160206116e9833981519152808688096000805160206116e9833981519152868a09089150915094509492505050565b600080610ef186856000805160206116e9833981519152611144565b610f0a86856000805160206116e9833981519152611144565b9150915094509492505050565b600080600080600080610f2c8c8c6003611111565b9096509450610f3d86868e8e610e64565b9096509450610f4e8a8a8a8a610e64565b9098509650610f5f8c8c8c8c610e64565b9094509250610f7084848a8a610e64565b9094509250610f8186868181610e64565b909c509a50610f9284846008611111565b9092509050610fa38c8c8484610ed5565b909c509a50610fb488888181610e64565b9092509050610fc584846004611111565b9094509250610fd684848e8e610ed5565b9094509250610fe784848888610e64565b9094509250610ff88a8a6008611111565b909650945061100986868c8c610e64565b909650945061101a86868484610e64565b909650945061102b84848888610ed5565b909450925061103c8c8c6002611111565b909650945061104d86868a8a610e64565b909650945061105e88888484610e64565b909250905061106f82826008611111565b809250819350505096509650965096509650969050565b600080806110c76000805160206116e9833981519152808788096000805160206116e9833981519152878809086000805160206116e9833981519152611168565b90506000805160206116e98339815191528186096000805160206116e9833981519152828609611105906000805160206116e98339815191526116c1565b92509250509250929050565b6000806000805160206116e98339815191528386096000805160206116e983398151915284860991509150935093915050565b6000818061115457611154611689565b61115e84846116c1565b8508949350505050565b60008060405160208152602080820152602060408201528460608201526002840360808201528360a082015260208160c08360056107d05a03fa905192509050806111b257600080fd5b5092915050565b82600281019282156111e7579160200282015b828111156111e75782518255916020019190600101906111cc565b506111f392915061129f565b5090565b604051806040016040528061120a6112b4565b81526020016112176112b4565b905290565b6040805160808101825260009181018281526060820192909252908152602081016112176111f7565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b808211156111f357600081556001016112a0565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561130b5761130b6112d2565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561133a5761133a6112d2565b604052919050565b60006040828403121561135457600080fd5b61135c6112e8565b823581526020928301359281019290925250919050565b600082601f83011261138457600080fd5b61138c6112e8565b80604084018581111561139e57600080fd5b845b818110156113b85780358452602093840193016113a0565b509095945050505050565b6000608082840312156113d557600080fd5b6113dd6112e8565b90506113e98383611373565b81526113f88360408401611373565b602082015292915050565b60008060c0838503121561141657600080fd5b6114208484611342565b915061142f84604085016113c3565b90509250929050565b602080825282518282018190526000918401906040840190835b818110156113b85783511515835260209384019390920191600101611452565b600067ffffffffffffffff82111561148c5761148c6112d2565b5060051b60200190565b600082601f8301126114a757600080fd5b81356114ba6114b582611472565b611311565b8082825260208201915060208360061b8601019250858311156114dc57600080fd5b602085015b83811015611503576114f38782611342565b83526020909201916040016114e1565b5095945050505050565b60006020828403121561151f57600080fd5b813567ffffffffffffffff81111561153657600080fd5b61154284828501611496565b949350505050565b600080600080610100858703121561156157600080fd5b61156b8686611342565b935061157a86604087016113c3565b939693955050505060c08201359160e0013590565b600080604083850312156115a257600080fd5b823567ffffffffffffffff8111156115b957600080fd5b6115c585828601611496565b925050602083013567ffffffffffffffff8111156115e257600080fd5b8301601f810185136115f357600080fd5b80356116016114b582611472565b8082825260208201915060208360071b85010192508783111561162357600080fd5b6020840193505b8284101561164f5761163c88856113c3565b825260208201915060808401935061162a565b809450505050509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b6000826116bc57634e487b7160e01b600052601260045260246000fd5b500490565b818103818111156116e257634e487b7160e01b600052601160045260246000fd5b9291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220fe2e70568e30b7c9a5bb2cc032ce457b371d64e371285221afff73805d21c01364736f6c634300081a0033",
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

// GetFieldModulus is a free data retrieval call binding the contract method 0x55a3e90f.
//
// Solidity: function GetFieldModulus() pure returns(uint256)
func (_Contract *ContractCaller) GetFieldModulus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "GetFieldModulus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldModulus is a free data retrieval call binding the contract method 0x55a3e90f.
//
// Solidity: function GetFieldModulus() pure returns(uint256)
func (_Contract *ContractSession) GetFieldModulus() (*big.Int, error) {
	return _Contract.Contract.GetFieldModulus(&_Contract.CallOpts)
}

// GetFieldModulus is a free data retrieval call binding the contract method 0x55a3e90f.
//
// Solidity: function GetFieldModulus() pure returns(uint256)
func (_Contract *ContractCallerSession) GetFieldModulus() (*big.Int, error) {
	return _Contract.Contract.GetFieldModulus(&_Contract.CallOpts)
}

// GetTestMul is a free data retrieval call binding the contract method 0x184eecf0.
//
// Solidity: function GetTestMul() view returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractCaller) GetTestMul(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "GetTestMul")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetTestMul is a free data retrieval call binding the contract method 0x184eecf0.
//
// Solidity: function GetTestMul() view returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractSession) GetTestMul() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetTestMul(&_Contract.CallOpts)
}

// GetTestMul is a free data retrieval call binding the contract method 0x184eecf0.
//
// Solidity: function GetTestMul() view returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractCallerSession) GetTestMul() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetTestMul(&_Contract.CallOpts)
}

// GetVKoResult is a free data retrieval call binding the contract method 0x1357d59b.
//
// Solidity: function GetVKoResult() view returns(bool[])
func (_Contract *ContractCaller) GetVKoResult(opts *bind.CallOpts) ([]bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "GetVKoResult")

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// GetVKoResult is a free data retrieval call binding the contract method 0x1357d59b.
//
// Solidity: function GetVKoResult() view returns(bool[])
func (_Contract *ContractSession) GetVKoResult() ([]bool, error) {
	return _Contract.Contract.GetVKoResult(&_Contract.CallOpts)
}

// GetVKoResult is a free data retrieval call binding the contract method 0x1357d59b.
//
// Solidity: function GetVKoResult() view returns(bool[])
func (_Contract *ContractCallerSession) GetVKoResult() ([]bool, error) {
	return _Contract.Contract.GetVKoResult(&_Contract.CallOpts)
}

// TestECTwistMul is a paid mutator transaction binding the contract method 0xa695a6b4.
//
// Solidity: function TestECTwistMul() returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractTransactor) TestECTwistMul(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "TestECTwistMul")
}

// TestECTwistMul is a paid mutator transaction binding the contract method 0xa695a6b4.
//
// Solidity: function TestECTwistMul() returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractSession) TestECTwistMul() (*types.Transaction, error) {
	return _Contract.Contract.TestECTwistMul(&_Contract.TransactOpts)
}

// TestECTwistMul is a paid mutator transaction binding the contract method 0xa695a6b4.
//
// Solidity: function TestECTwistMul() returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractTransactorSession) TestECTwistMul() (*types.Transaction, error) {
	return _Contract.Contract.TestECTwistMul(&_Contract.TransactOpts)
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

// UploadVKoDLEQ is a paid mutator transaction binding the contract method 0x90caa16f.
//
// Solidity: function UploadVKoDLEQ((uint256,uint256) a1, (uint256[2],uint256[2]) a2, uint256 c, uint256 z) returns()
func (_Contract *ContractTransactor) UploadVKoDLEQ(opts *bind.TransactOpts, a1 VerificationG1Point, a2 VerificationG2Point, c *big.Int, z *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "UploadVKoDLEQ", a1, a2, c, z)
}

// UploadVKoDLEQ is a paid mutator transaction binding the contract method 0x90caa16f.
//
// Solidity: function UploadVKoDLEQ((uint256,uint256) a1, (uint256[2],uint256[2]) a2, uint256 c, uint256 z) returns()
func (_Contract *ContractSession) UploadVKoDLEQ(a1 VerificationG1Point, a2 VerificationG2Point, c *big.Int, z *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UploadVKoDLEQ(&_Contract.TransactOpts, a1, a2, c, z)
}

// UploadVKoDLEQ is a paid mutator transaction binding the contract method 0x90caa16f.
//
// Solidity: function UploadVKoDLEQ((uint256,uint256) a1, (uint256[2],uint256[2]) a2, uint256 c, uint256 z) returns()
func (_Contract *ContractTransactorSession) UploadVKoDLEQ(a1 VerificationG1Point, a2 VerificationG2Point, c *big.Int, z *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UploadVKoDLEQ(&_Contract.TransactOpts, a1, a2, c, z)
}

// UploadVKuDLEQ is a paid mutator transaction binding the contract method 0xe29fb4d4.
//
// Solidity: function UploadVKuDLEQ((uint256,uint256) a1, (uint256[2],uint256[2]) a2, uint256 c, uint256 z) returns()
func (_Contract *ContractTransactor) UploadVKuDLEQ(opts *bind.TransactOpts, a1 VerificationG1Point, a2 VerificationG2Point, c *big.Int, z *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "UploadVKuDLEQ", a1, a2, c, z)
}

// UploadVKuDLEQ is a paid mutator transaction binding the contract method 0xe29fb4d4.
//
// Solidity: function UploadVKuDLEQ((uint256,uint256) a1, (uint256[2],uint256[2]) a2, uint256 c, uint256 z) returns()
func (_Contract *ContractSession) UploadVKuDLEQ(a1 VerificationG1Point, a2 VerificationG2Point, c *big.Int, z *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UploadVKuDLEQ(&_Contract.TransactOpts, a1, a2, c, z)
}

// UploadVKuDLEQ is a paid mutator transaction binding the contract method 0xe29fb4d4.
//
// Solidity: function UploadVKuDLEQ((uint256,uint256) a1, (uint256[2],uint256[2]) a2, uint256 c, uint256 z) returns()
func (_Contract *ContractTransactorSession) UploadVKuDLEQ(a1 VerificationG1Point, a2 VerificationG2Point, c *big.Int, z *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UploadVKuDLEQ(&_Contract.TransactOpts, a1, a2, c, z)
}

// VKoVerify is a paid mutator transaction binding the contract method 0x08ac6601.
//
// Solidity: function VKoVerify() payable returns()
func (_Contract *ContractTransactor) VKoVerify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "VKoVerify")
}

// VKoVerify is a paid mutator transaction binding the contract method 0x08ac6601.
//
// Solidity: function VKoVerify() payable returns()
func (_Contract *ContractSession) VKoVerify() (*types.Transaction, error) {
	return _Contract.Contract.VKoVerify(&_Contract.TransactOpts)
}

// VKoVerify is a paid mutator transaction binding the contract method 0x08ac6601.
//
// Solidity: function VKoVerify() payable returns()
func (_Contract *ContractTransactorSession) VKoVerify() (*types.Transaction, error) {
	return _Contract.Contract.VKoVerify(&_Contract.TransactOpts)
}
