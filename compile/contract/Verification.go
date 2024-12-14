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
	ABI: "[{\"inputs\":[],\"name\":\"GetFieldModulus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPKTau2\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetTestAdd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetVKoResult\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TestECTwistAdd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"pkArray\",\"type\":\"tuple[]\"}],\"name\":\"UploadDRsKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"pk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"vk\",\"type\":\"tuple\"}],\"name\":\"UploadOwnerKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"tau1\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point[]\",\"name\":\"tau2\",\"type\":\"tuple[]\"}],\"name\":\"UploadSystemKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"pk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"vk\",\"type\":\"tuple\"}],\"name\":\"UploadUserKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"a1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"a2\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"UploadVKoDLEQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"a1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"a2\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"UploadVKuDLEQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VKoVerify\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001602455600160255560016026556001602755348015602357600080fd5b50611991806100336000396000f3fe6080604052600436106100a75760003560e01c80639da5cd00116100645780639da5cd001461016a578063ab095d5b1461018a578063bae0794c146101bf578063e29fb4d4146101e1578063eb3891b414610201578063f029c25e1461021f57600080fd5b80630359a5dd146100ac57806308ac6601146100ce5780631357d59b146100d65780633ae154901461010157806355a3e90f1461012157806390caa16f1461014a575b600080fd5b3480156100b857600080fd5b506100cc6100c736600461160e565b61023f565b005b6100cc61027e565b3480156100e257600080fd5b506100eb610525565b6040516100f89190611643565b60405180910390f35b34801561010d57600080fd5b506100cc61011c366004611718565b61059d565b34801561012d57600080fd5b5060405160008051602061193c83398151915281526020016100f8565b34801561015657600080fd5b506100cc610165366004611755565b6105f7565b34801561017657600080fd5b506100cc61018536600461179a565b61063e565b34801561019657600080fd5b5061019f61075e565b6040805194855260208501939093529183015260608201526080016100f8565b3480156101cb57600080fd5b506101d46108d3565b6040516100f8919061188b565b3480156101ed57600080fd5b506100cc6101fc366004611755565b610972565b34801561020d57600080fd5b5060245460255460265460275461019f565b34801561022b57600080fd5b506100cc61023a36600461160e565b6109b9565b81516001556020820151600290815581518291600391610261918391906113c4565b50602082015161027790600280840191906113c4565b5050505050565b60006102c860008081548110610296576102966118b0565b6000918252602091829020604080518082019091526006909202018054825260010154918101919091526015546109da565b60408051808201909152600154815260025460208201526014549192506000916102f291906109da565b905060006103008383610a2f565b8051600e5491925014158061031b57506020810151600f5414155b1561036757601f805460018101825560008290527fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d80760208204018054919092166101000a60ff02191690555b601f805460018101825560008290527fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d80760208204018054919092166101000a60ff8102199091161790556103b9611402565b6103c1611402565b6103c9611402565b61047a600e60070154600080815481106103e5576103e56118b0565b60009182526020822060026006909202010101546000808154811061040c5761040c6118b0565b600091825260209091206002600690920201016001015460008081548110610436576104366118b0565b60009182526020822060046006909202010101546000808154811061045d5761045d6118b0565b6000918252602090912060046006909202010160015b0154610a87565b865160208089015180820193909352929091529081019190915252601454600354600454600580546104b0949392916001610473565b855160208088018051808301949094529390925281810193909352929091528451805190830151868401518051908501518751805196015194518051610506979596949593949293929160016020020151610b31565b8451602095860151808701929092529190529283015290525050505050565b6060601f80548060200260200160405190810160405280929190818152602001828054801561059357602002820191906000526020600020906000905b825461010083900a900460ff1615158152602060019283018181049485019490930390920291018084116105625790505b5050505050905090565b60005b81518110156105f357600d8282815181106105bd576105bd6118b0565b602090810291909101810151825460018181018555600094855293839020825160029092020190815591015190820155016105a0565b5050565b8351600e556020840151600f558251839060109061061890829060026113c4565b50602082015161062e90600280840191906113c4565b5050506014919091556015555050565b610646611427565b60005b835181101561075857838181518110610664576106646118b0565b60200260200101518260000181905250828181518110610686576106866118b0565b602090810291909101810151838201908152600080546001810182559080528451805160069092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56381019283559301517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5648401559051805185937f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e565019061073190829060026113c4565b50602082015161074790600280840191906113c4565b505060019093019250610649915050565b50505050565b6000806000806108b16000808154811061077a5761077a6118b0565b6000918252602082206002600690920201010154600080815481106107a1576107a16118b0565b6000918252602090912060026006909202010160010154600080815481106107cb576107cb6118b0565b6000918252602082206004600690920201010154600080815481106107f2576107f26118b0565b6000918252602090912060046006909202010160010154600060018154811061081d5761081d6118b0565b60009182526020822060026006909202010101546000600181548110610845576108456118b0565b60009182526020909120600260069092020101600101546000600181548110610870576108706118b0565b60009182526020822060046006909202010101546000600181548110610898576108986118b0565b6000918252602090912060056006909202010154610b31565b6027819055602682905560258390556024849055929791965094509092509050565b6108db611402565b600080815481106108ee576108ee6118b0565b60009182526020909120604080516080810180835290936006029092016002908101928492830191849182845b81548152602001906001019080831161091b57505050918352505060408051808201918290526020909201919060028481019182845b81548152602001906001019080831161095157505050505081525050905090565b835160165560208401516017558251839060189061099390829060026113c4565b5060208201516109a990600280840191906113c4565b505050601c91909155601d555050565b815160075560208201516008558051819060099061026190829060026113c4565b60408051808201909152600080825260208201526109f6611450565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa905080610a2757600080fd5b505092915050565b6040805180820190915260008082526020820152610a4b61146e565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa905080610a2757600080fd5b6000808080600188158015610a9a575087155b8015610aa4575086155b8015610aae575085155b15610ac25750600197508795506000610ada565b610ace89898989610c64565b610ada57610ada6118c6565b6000610aec8b8b8b8b8b876000610d19565b9050610b1c8160005b602090810291909101519083015160408401516060850151608086015160a0870151610d9c565b929e919d509b50909950975050505050505050565b60008080808b158015610b4257508a155b8015610b4c575089155b8015610b56575088155b15610ba75787158015610b67575086155b8015610b71575085155b8015610b7b575084155b610b9757610b8b88888888610c64565b610b9757610b976118c6565b5086925085915084905083610c55565b87158015610bb3575086155b8015610bbd575085155b8015610bc7575084155b15610bf457610bd88c8c8c8c610c64565b610be457610be46118c6565b508a925089915088905087610c55565b610c008c8c8c8c610c64565b610c0c57610c0c6118c6565b610c1888888888610c64565b610c2457610c246118c6565b6000610c3e8d8d8d8d600160008f8f8f8f60016000610de6565b9050610c4b816000610af5565b9450945094509450505b98509850985098945050505050565b6000806000806000610c788787898961106f565b9094509250610c898989818161106f565b9092509050610c9a82828b8b61106f565b9092509050610cab848484846110e0565b9094509250610cfb84847f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e57e9713b03af0fed4cd2cafadeed8fdf4a74fa084e52d1852e4a2bd0685c315d26110e0565b909450925083158015610d0c575082155b9998505050505050505050565b610d2161148c565b8715610d91576001881615610d62578051602082015160408301516060840151608085015160a0860151610d5f9594939291908d8d8d8d8d8d610de6565b90505b610d70878787878787611122565b949b50929950909750955093509150610d8a6002896118f2565b9750610d21565b979650505050505050565b600080600080600080610daf8888611291565b9092509050610dc08c8c848461106f565b9096509450610dd18a8a848461106f565b969d959c509a50949850929650505050505050565b610dee61148c565b88158015610dfa575087155b15610e3c578686868686868660005b60a0890192909252608088019290925260608701929092526040860192909252602085810193909352909102015261105f565b82158015610e48575081155b15610e5b578c8c8c8c8c8c866000610e09565b610e6785858b8b61106f565b9095509350610e788b8b858561106f565b60608301526040820152610e8e87878b8b61106f565b9097509550610e9f8d8d858561106f565b60a08301526080820181905287148015610ebc575060a081015186145b15610f0157604081015185148015610ed75750606081015184145b15610ef257610eea8d8d8d8d8d8d611122565b866000610e09565b60016000818180808681610e09565b610f0d8989858561106f565b9093509150610f2d858583600260200201518460035b60200201516110e0565b909d509b50610f4787878360046020020151846005610f23565b909b509950610f588b8b818161106f565b9099509750610f78898983600460200201518460055b602002015161106f565b9095509350610f8989898d8d61106f565b9099509750610f9a8989858561106f565b60a08301526080820152610fb08d8d818161106f565b9097509550610fc18787858561106f565b9097509550610fd287878b8b6110e0565b9097509550610fe38585600261131c565b9093509150610ff4878785856110e0565b90975095506110058b8b898961106f565b60208301528152611018858589896110e0565b909b5099506110298d8d8d8d61106f565b909b50995061104389898360026020020151846003610f6e565b909d509b506110548b8b8f8f6110e0565b606083015260408201525b9c9b505050505050505050505050565b6000806110ad60008051602061193c83398151915285880960008051602061193c83398151915285880960008051602061193c83398151915261134f565b60008051602061193c8339815191528086880960008051602061193c833981519152868a09089150915094509492505050565b6000806110fc868560008051602061193c83398151915261134f565b611115868560008051602061193c83398151915261134f565b9150915094509492505050565b6000806000806000806111378c8c600361131c565b909650945061114886868e8e61106f565b90965094506111598a8a8a8a61106f565b909850965061116a8c8c8c8c61106f565b909450925061117b84848a8a61106f565b909450925061118c8686818161106f565b909c509a5061119d8484600861131c565b90925090506111ae8c8c84846110e0565b909c509a506111bf8888818161106f565b90925090506111d08484600461131c565b90945092506111e184848e8e6110e0565b90945092506111f28484888861106f565b90945092506112038a8a600861131c565b909650945061121486868c8c61106f565b90965094506112258686848461106f565b9096509450611236848488886110e0565b90945092506112478c8c600261131c565b909650945061125886868a8a61106f565b90965094506112698888848461106f565b909250905061127a8282600861131c565b809250819350505096509650965096509650969050565b600080806112d260008051602061193c8339815191528087880960008051602061193c8339815191528788090860008051602061193c833981519152611373565b905060008051602061193c83398151915281860960008051602061193c8339815191528286096113109060008051602061193c833981519152611914565b92509250509250929050565b60008060008051602061193c83398151915283860960008051602061193c83398151915284860991509150935093915050565b6000818061135f5761135f6118dc565b6113698484611914565b8508949350505050565b60008060405160208152602080820152602060408201528460608201526002840360808201528360a082015260208160c08360056107d05a03fa905192509050806113bd57600080fd5b5092915050565b82600281019282156113f2579160200282015b828111156113f25782518255916020019190600101906113d7565b506113fe9291506114aa565b5090565b60405180604001604052806114156114bf565b81526020016114226114bf565b905290565b604080516080810182526000918101828152606082019290925290815260208101611422611402565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b808211156113fe57600081556001016114ab565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611516576115166114dd565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611545576115456114dd565b604052919050565b60006040828403121561155f57600080fd5b6115676114f3565b823581526020928301359281019290925250919050565b600082601f83011261158f57600080fd5b6115976114f3565b8060408401858111156115a957600080fd5b845b818110156115c35780358452602093840193016115ab565b509095945050505050565b6000608082840312156115e057600080fd5b6115e86114f3565b90506115f4838361157e565b8152611603836040840161157e565b602082015292915050565b60008060c0838503121561162157600080fd5b61162b848461154d565b915061163a84604085016115ce565b90509250929050565b602080825282518282018190526000918401906040840190835b818110156115c3578351151583526020938401939092019160010161165d565b600067ffffffffffffffff821115611697576116976114dd565b5060051b60200190565b600082601f8301126116b257600080fd5b81356116c56116c08261167d565b61151c565b8082825260208201915060208360061b8601019250858311156116e757600080fd5b602085015b8381101561170e576116fe878261154d565b83526020909201916040016116ec565b5095945050505050565b60006020828403121561172a57600080fd5b813567ffffffffffffffff81111561174157600080fd5b61174d848285016116a1565b949350505050565b600080600080610100858703121561176c57600080fd5b611776868661154d565b935061178586604087016115ce565b939693955050505060c08201359160e0013590565b600080604083850312156117ad57600080fd5b823567ffffffffffffffff8111156117c457600080fd5b6117d0858286016116a1565b925050602083013567ffffffffffffffff8111156117ed57600080fd5b8301601f810185136117fe57600080fd5b803561180c6116c08261167d565b8082825260208201915060208360071b85010192508783111561182e57600080fd5b6020840193505b8284101561185a5761184788856115ce565b8252602082019150608084019350611835565b809450505050509250929050565b8060005b600281101561075857815184526020938401939091019060010161186c565b600060808201905061189e828451611868565b60208301516113bd6040840182611868565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b60008261190f57634e487b7160e01b600052601260045260246000fd5b500490565b8181038181111561193557634e487b7160e01b600052601160045260246000fd5b9291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122018e6f77720acc31512dd90e4dceb2390f45e884ef1dbe8bddcb88ef2e6f64ebb64736f6c634300081a0033",
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

// GetPKTau2 is a free data retrieval call binding the contract method 0xbae0794c.
//
// Solidity: function GetPKTau2() view returns((uint256[2],uint256[2]))
func (_Contract *ContractCaller) GetPKTau2(opts *bind.CallOpts) (VerificationG2Point, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "GetPKTau2")

	if err != nil {
		return *new(VerificationG2Point), err
	}

	out0 := *abi.ConvertType(out[0], new(VerificationG2Point)).(*VerificationG2Point)

	return out0, err

}

// GetPKTau2 is a free data retrieval call binding the contract method 0xbae0794c.
//
// Solidity: function GetPKTau2() view returns((uint256[2],uint256[2]))
func (_Contract *ContractSession) GetPKTau2() (VerificationG2Point, error) {
	return _Contract.Contract.GetPKTau2(&_Contract.CallOpts)
}

// GetPKTau2 is a free data retrieval call binding the contract method 0xbae0794c.
//
// Solidity: function GetPKTau2() view returns((uint256[2],uint256[2]))
func (_Contract *ContractCallerSession) GetPKTau2() (VerificationG2Point, error) {
	return _Contract.Contract.GetPKTau2(&_Contract.CallOpts)
}

// GetTestAdd is a free data retrieval call binding the contract method 0xeb3891b4.
//
// Solidity: function GetTestAdd() view returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractCaller) GetTestAdd(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "GetTestAdd")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetTestAdd is a free data retrieval call binding the contract method 0xeb3891b4.
//
// Solidity: function GetTestAdd() view returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractSession) GetTestAdd() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetTestAdd(&_Contract.CallOpts)
}

// GetTestAdd is a free data retrieval call binding the contract method 0xeb3891b4.
//
// Solidity: function GetTestAdd() view returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractCallerSession) GetTestAdd() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetTestAdd(&_Contract.CallOpts)
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

// TestECTwistAdd is a paid mutator transaction binding the contract method 0xab095d5b.
//
// Solidity: function TestECTwistAdd() returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractTransactor) TestECTwistAdd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "TestECTwistAdd")
}

// TestECTwistAdd is a paid mutator transaction binding the contract method 0xab095d5b.
//
// Solidity: function TestECTwistAdd() returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractSession) TestECTwistAdd() (*types.Transaction, error) {
	return _Contract.Contract.TestECTwistAdd(&_Contract.TransactOpts)
}

// TestECTwistAdd is a paid mutator transaction binding the contract method 0xab095d5b.
//
// Solidity: function TestECTwistAdd() returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractTransactorSession) TestECTwistAdd() (*types.Transaction, error) {
	return _Contract.Contract.TestECTwistAdd(&_Contract.TransactOpts)
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
