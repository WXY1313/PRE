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
	ABI: "[{\"inputs\":[],\"name\":\"GetFieldModulus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetVKResult\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"c0\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"c1\",\"type\":\"string\"}],\"name\":\"UploadCiphertext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"pkArray\",\"type\":\"tuple[]\"}],\"name\":\"UploadDRsKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"pk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"vk\",\"type\":\"tuple\"}],\"name\":\"UploadOwnerKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"rks0\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"rks1\",\"type\":\"tuple[]\"}],\"name\":\"UploadReKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"witness\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"a1\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"a2\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"c\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"z\",\"type\":\"uint256[]\"}],\"name\":\"UploadReKeysProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point[]\",\"name\":\"tau1\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point[]\",\"name\":\"tau2\",\"type\":\"tuple[]\"}],\"name\":\"UploadSystemKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"pk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"vk\",\"type\":\"tuple\"}],\"name\":\"UploadUserKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"a1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"a2\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"UploadVKoDLEQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structVerification.G1Point\",\"name\":\"a1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerification.G2Point\",\"name\":\"a2\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"UploadVKuDLEQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VKoVerify\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VKuVerify\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061217c8061001f6000396000f3fe6080604052600436106100c25760003560e01c80638959c8e51161007f578063c4bd0aba11610059578063c4bd0aba146101c7578063e29fb4d4146101e7578063f029c25e14610207578063f54de4651461022757600080fd5b80638959c8e51461017f57806390caa16f146101875780639da5cd00146101a757600080fd5b80630359a5dd146100c757806307267e5e146100e957806308ac6601146101095780633ae154901461011157806355a3e90f1461013157806385105c311461015f575b600080fd5b3480156100d357600080fd5b506100e76100e2366004611aa0565b610249565b005b3480156100f557600080fd5b506100e7610104366004611ad5565b610288565b6100e76102c8565b34801561011d57600080fd5b506100e761012c366004611c19565b610620565b34801561013d57600080fd5b5060405160008051602061212783398151915281526020015b60405180910390f35b34801561016b57600080fd5b506100e761017a366004611cb3565b61067a565b6100e76108c1565b34801561019357600080fd5b506100e76101a2366004611daa565b610a54565b3480156101b357600080fd5b506100e76101c2366004611def565b610a9b565b3480156101d357600080fd5b506100e76101e2366004611ebb565b610bbb565b3480156101f357600080fd5b506100e7610202366004611daa565b610cd2565b34801561021357600080fd5b506100e7610222366004611aa0565b610d19565b34801561023357600080fd5b5061023c610d3a565b6040516101569190611f22565b8151600155602082015160029081558151829160039161026b9183919061179c565b506020820151610281906002808401919061179c565b5050505050565b8151829060209061029c908290600261179c565b5060208201516102b2906002808401919061179c565b50602491506102c390508282611fdd565b505050565b6000610312600080815481106102e0576102e061209b565b600091825260209182902060408051808201909152600690920201805482526001015491810191909152601554610db2565b604080518082019091526001548152600254602082015260145491925060009161033c9190610db2565b9050600061034a8383610e07565b90506103546117da565b61035c6117da565b6103646117da565b610415600e60070154600080815481106103805761038061209b565b6000918252602082206002600690920201010154600080815481106103a7576103a761209b565b6000918252602090912060026006909202010160010154600080815481106103d1576103d161209b565b6000918252602082206004600690920201010154600080815481106103f8576103f861209b565b6000918252602090912060046006909202010160015b0154610e5f565b8651602080890151808201939093529290915290810191909152526014546003546004546005805461044b94939291600161040e565b85516020808801518082019390935292909152908101919091525282516105049060005b602002015184516001602002015185602001516000600281106104945761049461209b565b602002015186602001516001600281106104b0576104b061209b565b602002015186516000602002015187516001602002015188602001516000600281106104de576104de61209b565b602002015189602001516001600281106104fa576104fa61209b565b6020020151610f09565b8451602080870151808201939093529290915290810191909152528351600e5414158061053757506020840151600f5414155b80610546575080515160105414155b80610558575080516020015160115414155b8061056a575060208101515160125414155b8061058257506020818101510151601260015b015414155b156105ce57601f805460018101825560008290527fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d80760208204018054919092166101000a60ff02191690555b5050601f805460018101825560008290527fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d80760208204018054919092166101000a60ff81021990911617905550505050565b60005b815181101561067657600d8282815181106106405761064061209b565b60209081029190910181015182546001818101855560009485529383902082516002909202019081559101519082015501610623565b5050565b8551602655602086015160275561068f6117ff565b610697611853565b60005b87518110156108b6578681815181106106b5576106b561209b565b602002602001015183600001819052508581815181106106d7576106d761209b565b602002602001015183602001819052508481815181106106f9576106f961209b565b602002602001015183604001818152505083818151811061071c5761071c61209b565b602002602001015183606001818152505087818151811061073f5761073f61209b565b60209081029190910181015183528281018481526028805460018082018355600092909252855180517fe16da923a2d88192e5070f37b4571d58682c0d66212ec634d495f33de3f77ab56008909302928301558401517fe16da923a2d88192e5070f37b4571d58682c0d66212ec634d495f33de3f77ab68201559151805180517fe16da923a2d88192e5070f37b4571d58682c0d66212ec634d495f33de3f77ab78501558401517fe16da923a2d88192e5070f37b4571d58682c0d66212ec634d495f33de3f77ab88401558084015180517fe16da923a2d88192e5070f37b4571d58682c0d66212ec634d495f33de3f77ab9850155909301517fe16da923a2d88192e5070f37b4571d58682c0d66212ec634d495f33de3f77aba83015560408301517fe16da923a2d88192e5070f37b4571d58682c0d66212ec634d495f33de3f77abb8301556060909201517fe16da923a2d88192e5070f37b4571d58682c0d66212ec634d495f33de3f77abc909101550161069a565b505050505050505050565b600061090b600080815481106108d9576108d961209b565b600091825260209182902060408051808201909152600690920201805482526001015491810191909152601d54610db2565b6040805180820190915260075481526008546020820152601c549192506000916109359190610db2565b905060006109438383610e07565b905061094d6117da565b6109556117da565b61095d6117da565b610979601660070154600080815481106103805761038061209b565b865160208089015180820193909352929091529081019190915252601c54600954600a54600b80546109af94939291600161040e565b85516020808801518082019390935292909152908101919091525282516109d790600061046f565b8451602080870151808201939093529290915290810191909152528351601654141580610a0a5750602084015160175414155b80610a19575080515160185414155b80610a2b575080516020015160195414155b80610a3d5750602081015151601a5414155b8061058257506020818101510151601a600161057d565b8351600e556020840151600f5582518390601090610a75908290600261179c565b506020820151610a8b906002808401919061179c565b5050506014919091556015555050565b610aa361187c565b60005b8351811015610bb557838181518110610ac157610ac161209b565b60200260200101518260000181905250828181518110610ae357610ae361209b565b602090810291909101810151838201908152600080546001810182559080528451805160069092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56381019283559301517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5648401559051805185937f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5650190610b8e908290600261179c565b506020820151610ba4906002808401919061179c565b505060019093019250610aa6915050565b50505050565b610bc36118a5565b60005b8351811015610bb557838181518110610be157610be161209b565b60200260200101518260000181905250828181518110610c0357610c0361209b565b6020908102919091018101518382019081526025805460018082018355600092909252855180517f401968ff42a154441da5f6c4c935ac46b8671f0e062baaa62a7545ba53bb6e4c6004909302928301558401517f401968ff42a154441da5f6c4c935ac46b8671f0e062baaa62a7545ba53bb6e4d820155915180517f401968ff42a154441da5f6c4c935ac46b8671f0e062baaa62a7545ba53bb6e4e840155909201517f401968ff42a154441da5f6c4c935ac46b8671f0e062baaa62a7545ba53bb6e4f9091015501610bc6565b8351601655602084015160175582518390601890610cf3908290600261179c565b506020820151610d09906002808401919061179c565b505050601c91909155601d555050565b815160075560208201516008558051819060099061026b908290600261179c565b6060601f805480602002602001604051908101604052809291908181526020018280548015610da857602002820191906000526020600020906000905b825461010083900a900460ff161515815260206001928301818104948501949093039092029101808411610d775790505b5050505050905090565b6040805180820190915260008082526020820152610dce6118e4565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa905080610dff57600080fd5b505092915050565b6040805180820190915260008082526020820152610e23611902565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa905080610dff57600080fd5b6000808080600188158015610e72575087155b8015610e7c575086155b8015610e86575085155b15610e9a5750600197508795506000610eb2565b610ea68989898961103c565b610eb257610eb26120b1565b6000610ec48b8b8b8b8b8760006110f1565b9050610ef48160005b602090810291909101519083015160408401516060850151608086015160a0870151611174565b929e919d509b50909950975050505050505050565b60008080808b158015610f1a57508a155b8015610f24575089155b8015610f2e575088155b15610f7f5787158015610f3f575086155b8015610f49575085155b8015610f53575084155b610f6f57610f638888888861103c565b610f6f57610f6f6120b1565b508692508591508490508361102d565b87158015610f8b575086155b8015610f95575085155b8015610f9f575084155b15610fcc57610fb08c8c8c8c61103c565b610fbc57610fbc6120b1565b508a92508991508890508761102d565b610fd88c8c8c8c61103c565b610fe457610fe46120b1565b610ff08888888861103c565b610ffc57610ffc6120b1565b60006110168d8d8d8d600160008f8f8f8f600160006111be565b9050611023816000610ecd565b9450945094509450505b98509850985098945050505050565b600080600080600061105087878989611447565b909450925061106189898181611447565b909250905061107282828b8b611447565b9092509050611083848484846114b8565b90945092506110d384847f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e57e9713b03af0fed4cd2cafadeed8fdf4a74fa084e52d1852e4a2bd0685c315d26114b8565b9094509250831580156110e4575082155b9998505050505050505050565b6110f9611920565b871561116957600188161561113a578051602082015160408301516060840151608085015160a08601516111379594939291908d8d8d8d8d8d6111be565b90505b6111488787878787876114fa565b949b509299509097509550935091506111626002896120dd565b97506110f9565b979650505050505050565b6000806000806000806111878888611669565b90925090506111988c8c8484611447565b90965094506111a98a8a8484611447565b969d959c509a50949850929650505050505050565b6111c6611920565b881580156111d2575087155b15611214578686868686868660005b60a08901929092526080880192909252606087019290925260408601929092526020858101939093529091020152611437565b82158015611220575081155b15611233578c8c8c8c8c8c8660006111e1565b61123f85858b8b611447565b90955093506112508b8b8585611447565b6060830152604082015261126687878b8b611447565b90975095506112778d8d8585611447565b60a08301526080820181905287148015611294575060a081015186145b156112d9576040810151851480156112af5750606081015184145b156112ca576112c28d8d8d8d8d8d6114fa565b8660006111e1565b600160008181808086816111e1565b6112e589898585611447565b9093509150611305858583600260200201518460035b60200201516114b8565b909d509b5061131f878783600460200201518460056112fb565b909b5099506113308b8b8181611447565b9099509750611350898983600460200201518460055b6020020151611447565b909550935061136189898d8d611447565b909950975061137289898585611447565b60a083015260808201526113888d8d8181611447565b909750955061139987878585611447565b90975095506113aa87878b8b6114b8565b90975095506113bb858560026116f4565b90935091506113cc878785856114b8565b90975095506113dd8b8b8989611447565b602083015281526113f0858589896114b8565b909b5099506114018d8d8d8d611447565b909b50995061141b89898360026020020151846003611346565b909d509b5061142c8b8b8f8f6114b8565b606083015260408201525b9c9b505050505050505050505050565b600080611485600080516020612127833981519152858809600080516020612127833981519152858809600080516020612127833981519152611727565b60008051602061212783398151915280868809600080516020612127833981519152868a09089150915094509492505050565b6000806114d48685600080516020612127833981519152611727565b6114ed8685600080516020612127833981519152611727565b9150915094509492505050565b60008060008060008061150f8c8c60036116f4565b909650945061152086868e8e611447565b90965094506115318a8a8a8a611447565b90985096506115428c8c8c8c611447565b909450925061155384848a8a611447565b909450925061156486868181611447565b909c509a50611575848460086116f4565b90925090506115868c8c84846114b8565b909c509a5061159788888181611447565b90925090506115a8848460046116f4565b90945092506115b984848e8e6114b8565b90945092506115ca84848888611447565b90945092506115db8a8a60086116f4565b90965094506115ec86868c8c611447565b90965094506115fd86868484611447565b909650945061160e848488886114b8565b909450925061161f8c8c60026116f4565b909650945061163086868a8a611447565b909650945061164188888484611447565b9092509050611652828260086116f4565b809250819350505096509650965096509650969050565b600080806116aa600080516020612127833981519152808788096000805160206121278339815191528788090860008051602061212783398151915261174b565b90506000805160206121278339815191528186096000805160206121278339815191528286096116e8906000805160206121278339815191526120ff565b92509250509250929050565b60008060008051602061212783398151915283860960008051602061212783398151915284860991509150935093915050565b60008180611737576117376120c7565b61174184846120ff565b8508949350505050565b60008060405160208152602080820152602060408201528460608201526002840360808201528360a082015260208160c08360056107d05a03fa9051925090508061179557600080fd5b5092915050565b82600281019282156117ca579160200282015b828111156117ca5782518255916020019190600101906117af565b506117d692915061193e565b5090565b60405180604001604052806117ed611953565b81526020016117fa611953565b905290565b6040805160c0810190915260006080820181815260a08301919091528190815260200161183f604051806040016040528060008152602001600081525090565b815260200160008152602001600081525090565b6040805160808101825260009181018281526060820192909252908152602081016117fa6117ff565b6040805160808101825260009181018281526060820192909252908152602081016117fa6117da565b604080516080810182526000918101828152606082019290925290819081526020016117fa604051806040016040528060008152602001600081525090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b808211156117d6576000815560010161193f565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156119a9576119a9611971565b60405290565b604051601f8201601f191681016001600160401b03811182821017156119d7576119d7611971565b604052919050565b6000604082840312156119f157600080fd5b6119f9611987565b823581526020928301359281019290925250919050565b600082601f830112611a2157600080fd5b611a29611987565b806040840185811115611a3b57600080fd5b845b81811015611a55578035845260209384019301611a3d565b509095945050505050565b600060808284031215611a7257600080fd5b611a7a611987565b9050611a868383611a10565b8152611a958360408401611a10565b602082015292915050565b60008060c08385031215611ab357600080fd5b611abd84846119df565b9150611acc8460408501611a60565b90509250929050565b60008060a08385031215611ae857600080fd5b611af28484611a60565b915060808301356001600160401b03811115611b0d57600080fd5b8301601f81018513611b1e57600080fd5b80356001600160401b03811115611b3757611b37611971565b611b4a601f8201601f19166020016119af565b818152866020838501011115611b5f57600080fd5b816020840160208301376000602083830101528093505050509250929050565b60006001600160401b03821115611b9857611b98611971565b5060051b60200190565b600082601f830112611bb357600080fd5b8135611bc6611bc182611b7f565b6119af565b8082825260208201915060208360061b860101925085831115611be857600080fd5b602085015b83811015611c0f57611bff87826119df565b8352602090920191604001611bed565b5095945050505050565b600060208284031215611c2b57600080fd5b81356001600160401b03811115611c4157600080fd5b611c4d84828501611ba2565b949350505050565b600082601f830112611c6657600080fd5b8135611c74611bc182611b7f565b8082825260208201915060208360051b860101925085831115611c9657600080fd5b602085015b83811015611c0f578035835260209283019201611c9b565b60008060008060008060e08789031215611ccc57600080fd5b611cd688886119df565b955060408701356001600160401b03811115611cf157600080fd5b611cfd89828a01611ba2565b95505060608701356001600160401b03811115611d1957600080fd5b611d2589828a01611ba2565b94505060808701356001600160401b03811115611d4157600080fd5b611d4d89828a01611ba2565b93505060a08701356001600160401b03811115611d6957600080fd5b611d7589828a01611c55565b92505060c08701356001600160401b03811115611d9157600080fd5b611d9d89828a01611c55565b9150509295509295509295565b6000806000806101008587031215611dc157600080fd5b611dcb86866119df565b9350611dda8660408701611a60565b939693955050505060c08201359160e0013590565b60008060408385031215611e0257600080fd5b82356001600160401b03811115611e1857600080fd5b611e2485828601611ba2565b92505060208301356001600160401b03811115611e4057600080fd5b8301601f81018513611e5157600080fd5b8035611e5f611bc182611b7f565b8082825260208201915060208360071b850101925087831115611e8157600080fd5b6020840193505b82841015611ead57611e9a8885611a60565b8252602082019150608084019350611e88565b809450505050509250929050565b60008060408385031215611ece57600080fd5b82356001600160401b03811115611ee457600080fd5b611ef085828601611ba2565b92505060208301356001600160401b03811115611f0c57600080fd5b611f1885828601611ba2565b9150509250929050565b602080825282518282018190526000918401906040840190835b81811015611a555783511515835260209384019390920191600101611f3c565b600181811c90821680611f7057607f821691505b602082108103611f9057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156102c357806000526020600020601f840160051c81016020851015611fbd5750805b601f840160051c820191505b818110156102815760008155600101611fc9565b81516001600160401b03811115611ff657611ff6611971565b61200a816120048454611f5c565b84611f96565b6020601f82116001811461203e57600083156120265750848201515b600019600385901b1c1916600184901b178455610281565b600084815260208120601f198516915b8281101561206e578785015182556020948501946001909201910161204e565b508482101561208c5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b6000826120fa57634e487b7160e01b600052601260045260246000fd5b500490565b8181038181111561212057634e487b7160e01b600052601160045260246000fd5b9291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220b66172add7a819fd8a13fbe9840e8b19cd7d46e23bc352cdcf625e968d4f2e4264736f6c634300081a0033",
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

// GetVKResult is a free data retrieval call binding the contract method 0xf54de465.
//
// Solidity: function GetVKResult() view returns(bool[])
func (_Contract *ContractCaller) GetVKResult(opts *bind.CallOpts) ([]bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "GetVKResult")

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// GetVKResult is a free data retrieval call binding the contract method 0xf54de465.
//
// Solidity: function GetVKResult() view returns(bool[])
func (_Contract *ContractSession) GetVKResult() ([]bool, error) {
	return _Contract.Contract.GetVKResult(&_Contract.CallOpts)
}

// GetVKResult is a free data retrieval call binding the contract method 0xf54de465.
//
// Solidity: function GetVKResult() view returns(bool[])
func (_Contract *ContractCallerSession) GetVKResult() ([]bool, error) {
	return _Contract.Contract.GetVKResult(&_Contract.CallOpts)
}

// UploadCiphertext is a paid mutator transaction binding the contract method 0x07267e5e.
//
// Solidity: function UploadCiphertext((uint256[2],uint256[2]) c0, string c1) returns()
func (_Contract *ContractTransactor) UploadCiphertext(opts *bind.TransactOpts, c0 VerificationG2Point, c1 string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "UploadCiphertext", c0, c1)
}

// UploadCiphertext is a paid mutator transaction binding the contract method 0x07267e5e.
//
// Solidity: function UploadCiphertext((uint256[2],uint256[2]) c0, string c1) returns()
func (_Contract *ContractSession) UploadCiphertext(c0 VerificationG2Point, c1 string) (*types.Transaction, error) {
	return _Contract.Contract.UploadCiphertext(&_Contract.TransactOpts, c0, c1)
}

// UploadCiphertext is a paid mutator transaction binding the contract method 0x07267e5e.
//
// Solidity: function UploadCiphertext((uint256[2],uint256[2]) c0, string c1) returns()
func (_Contract *ContractTransactorSession) UploadCiphertext(c0 VerificationG2Point, c1 string) (*types.Transaction, error) {
	return _Contract.Contract.UploadCiphertext(&_Contract.TransactOpts, c0, c1)
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

// UploadReKeys is a paid mutator transaction binding the contract method 0xc4bd0aba.
//
// Solidity: function UploadReKeys((uint256,uint256)[] rks0, (uint256,uint256)[] rks1) returns()
func (_Contract *ContractTransactor) UploadReKeys(opts *bind.TransactOpts, rks0 []VerificationG1Point, rks1 []VerificationG1Point) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "UploadReKeys", rks0, rks1)
}

// UploadReKeys is a paid mutator transaction binding the contract method 0xc4bd0aba.
//
// Solidity: function UploadReKeys((uint256,uint256)[] rks0, (uint256,uint256)[] rks1) returns()
func (_Contract *ContractSession) UploadReKeys(rks0 []VerificationG1Point, rks1 []VerificationG1Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadReKeys(&_Contract.TransactOpts, rks0, rks1)
}

// UploadReKeys is a paid mutator transaction binding the contract method 0xc4bd0aba.
//
// Solidity: function UploadReKeys((uint256,uint256)[] rks0, (uint256,uint256)[] rks1) returns()
func (_Contract *ContractTransactorSession) UploadReKeys(rks0 []VerificationG1Point, rks1 []VerificationG1Point) (*types.Transaction, error) {
	return _Contract.Contract.UploadReKeys(&_Contract.TransactOpts, rks0, rks1)
}

// UploadReKeysProof is a paid mutator transaction binding the contract method 0x85105c31.
//
// Solidity: function UploadReKeysProof((uint256,uint256) commitment, (uint256,uint256)[] witness, (uint256,uint256)[] a1, (uint256,uint256)[] a2, uint256[] c, uint256[] z) returns()
func (_Contract *ContractTransactor) UploadReKeysProof(opts *bind.TransactOpts, commitment VerificationG1Point, witness []VerificationG1Point, a1 []VerificationG1Point, a2 []VerificationG1Point, c []*big.Int, z []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "UploadReKeysProof", commitment, witness, a1, a2, c, z)
}

// UploadReKeysProof is a paid mutator transaction binding the contract method 0x85105c31.
//
// Solidity: function UploadReKeysProof((uint256,uint256) commitment, (uint256,uint256)[] witness, (uint256,uint256)[] a1, (uint256,uint256)[] a2, uint256[] c, uint256[] z) returns()
func (_Contract *ContractSession) UploadReKeysProof(commitment VerificationG1Point, witness []VerificationG1Point, a1 []VerificationG1Point, a2 []VerificationG1Point, c []*big.Int, z []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UploadReKeysProof(&_Contract.TransactOpts, commitment, witness, a1, a2, c, z)
}

// UploadReKeysProof is a paid mutator transaction binding the contract method 0x85105c31.
//
// Solidity: function UploadReKeysProof((uint256,uint256) commitment, (uint256,uint256)[] witness, (uint256,uint256)[] a1, (uint256,uint256)[] a2, uint256[] c, uint256[] z) returns()
func (_Contract *ContractTransactorSession) UploadReKeysProof(commitment VerificationG1Point, witness []VerificationG1Point, a1 []VerificationG1Point, a2 []VerificationG1Point, c []*big.Int, z []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UploadReKeysProof(&_Contract.TransactOpts, commitment, witness, a1, a2, c, z)
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

// VKuVerify is a paid mutator transaction binding the contract method 0x8959c8e5.
//
// Solidity: function VKuVerify() payable returns()
func (_Contract *ContractTransactor) VKuVerify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "VKuVerify")
}

// VKuVerify is a paid mutator transaction binding the contract method 0x8959c8e5.
//
// Solidity: function VKuVerify() payable returns()
func (_Contract *ContractSession) VKuVerify() (*types.Transaction, error) {
	return _Contract.Contract.VKuVerify(&_Contract.TransactOpts)
}

// VKuVerify is a paid mutator transaction binding the contract method 0x8959c8e5.
//
// Solidity: function VKuVerify() payable returns()
func (_Contract *ContractTransactorSession) VKuVerify() (*types.Transaction, error) {
	return _Contract.Contract.VKuVerify(&_Contract.TransactOpts)
}
