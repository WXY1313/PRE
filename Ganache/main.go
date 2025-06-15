package main

import (
	"PRE/crypto/Convert"
	"PRE/crypto/DLEQ"
	"PRE/crypto/ElGamal"
	"PRE/crypto/KZG"
	"crypto/rand"

	"fmt"
	"math/big"
	"time"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"

	"PRE/compile/contract"
	"PRE/utils"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/core/types"
)

// Q is the order of the integer field (Zq) that fits inside the snark
var Q, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)

// R is the mod of the finite field
var R, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)

type RK struct {
	RK0 []*bn256.G1
	RK1 []*bn256.G1
}

type DLEQProofG1 struct {
	C  *big.Int
	Z  *big.Int
	RG *bn256.G1
	RH *bn256.G1
}

type DLEQProofG1_G2 struct {
	C  *big.Int
	Z  *big.Int
	RG *bn256.G1
	RH *bn256.G2
}

type RC struct {
	C0 *bn256.G2
	C1 *bn256.GT
	C2 []*bn256.G1
	C3 []*bn256.G1
}

// Compute Lagrangian interpolation on exponential
func recoverKey(Key []*bn256.G1, indices []*big.Int, threshold int) *bn256.G1 {

	k := threshold

	Recover_Key := new(bn256.G1).ScalarBaseMult(big.NewInt(0))

	for i := 0; i < k; i++ {

		num := big.NewInt(1)
		den := big.NewInt(1)

		for j := 0; j < k; j++ {
			if i != j {

				num.Mul(num, new(big.Int).Neg(indices[j]))
				num.Mod(num, R)

				den.Mul(den, new(big.Int).Sub(indices[i], indices[j]))
				den.Mod(den, R)
			}
		}

		den.ModInverse(den, R)

		term := new(big.Int).Mul(big.NewInt(1), num)
		term.Mul(term, den)
		term.Mod(term, R)
		Recover_Key = new(bn256.G1).Add(Recover_Key, new(bn256.G1).ScalarMult(Key[i], term))
	}
	return Recover_Key
}

func main() {
	contract_name := "Verification"
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privatekey := utils.GetENV("PRIVATE_KEY_1")

	auth := utils.Transact(client, privatekey, big.NewInt(0))

	address, tx := utils.Deploy(client, contract_name, auth)

	receipt, _ := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("Deploy Gas used: %d\n", receipt.GasUsed)

	Contract, err := contract.NewContract(common.HexToAddress(address.Hex()), client)
	if err != nil {
		fmt.Println(err)
	}

	//==========================================System Initialization===========================================//
	// the number of key shares
	numShares := 5
	// threshold value
	threshold := 2
	//threshold := numShares/2 + 1
	//threshold := 2*numShares/3 + 1
	n := 1

	fmt.Printf("The number of shares is %v\n", numShares)
	fmt.Printf("The threshold value is %v\n", threshold)

	//Setup Algorithm(off-chain)
	PK, _ := KZG.NewTrustedSetup(numShares)
	//fmt.Printf("The system public key is %v\n",PK)
	PKTau1 := make([]contract.VerificationG1Point, numShares)
	PKTau2 := make([]contract.VerificationG2Point, numShares)
	//PKG2i := make([]contract.VerificationG2Point, numShares)
	for i := 0; i < numShares; i++ {
		PKTau1[i] = Convert.G1ToG1Point(PK.Tau1[i])
		PKTau2[i] = Convert.G2ToG2Point(PK.Tau2[i])
		//PKG2i[i] = Convert.G2ToG2Point(new(bn256.G2).Add(PK.Tau2[1], new(bn256.G2).Neg(new(bn256.G2).ScalarBaseMult(big.NewInt(int64(i+1))))))
		//PKG2i[i] = Convert.G2ToG2Point(new(bn256.G2).Neg(new(bn256.G2).ScalarBaseMult(big.NewInt(int64(i + 1)))))
	}

	// PKTau2[0] = Convert.G2ToG2Point(PK.Tau2[0])
	// PKTau2[1] = Convert.G2ToG2Point(PK.Tau2[1])

	//fmt.Printf("\n\nThe result of G2i on the off-chain is %v\n", PKG2i)

	auth0 := utils.Transact(client, privatekey, big.NewInt(0))
	tx0, _ := Contract.UploadSystemKey(auth0, PKTau1, PKTau2)

	receipt0, err := bind.WaitMined(context.Background(), client, tx0)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("UploadSystemKey Gas used: %d\n", receipt0.GasUsed)

	//TestResult, _ := Contract.TestReturn(&bind.CallOpts{})
	//fmt.Printf("Test result is: %v\n", TestResult)

	//=====================================================Test=====================================================//
	mm, _ := rand.Int(rand.Reader, Q)
	auth555 := utils.Transact(client, privatekey, big.NewInt(0))
	tx555, _ := Contract.Empty(auth555, PKTau2[0], PKTau2[1])
	receipt555, err := bind.WaitMined(context.Background(), client, tx555)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}

	auth556 := utils.Transact(client, privatekey, big.NewInt(0))
	tx556, _ := Contract.Mul(auth556, PKTau2[0], mm)
	receipt556, err := bind.WaitMined(context.Background(), client, tx556)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("Mul Gas used: %d\n", receipt556.GasUsed-receipt555.GasUsed)
	//fmt.Printf("Empty Gas used: %d\n", receipt555.GasUsed)

	//==============================================================================================================//

	//===========================================User Registration===========================================//
	//KeyGen_u(off-chain)
	//The key pair of data owner
	sko, pko := ElGamal.EGKeyGen()
	vko := new(bn256.G2).ScalarBaseMult(sko)

	var _sko *big.Int
	var _pko *bn256.G1
	var _vko *bn256.G2
	starttime1 := time.Now().UnixMicro()
	for i := 0; i < int(n); i++ {
		_sko, _pko = ElGamal.EGKeyGen()
		_vko = new(bn256.G2).ScalarBaseMult(_sko)
	}
	endtime1 := time.Now().UnixMicro()

	fmt.Printf("%v\n%v\n", _pko, _vko)
	fmt.Printf("The KeyGen_u algorithm Time Used is %v us\n", (endtime1-starttime1)/int64(n))

	var _sk *big.Int
	var _pk *bn256.G1
	starttime1 = time.Now().UnixMicro()
	for i := 0; i < int(n); i++ {
		_sk, _pk = ElGamal.EGKeyGen()
	}
	endtime1 = time.Now().UnixMicro()
	fmt.Printf("%v\n%v\n", _sk, _pk)
	fmt.Printf("The KeyGen_DR algorithm Time Used is %v us\n", (endtime1-starttime1)/int64(n))

	//fmt.Printf("The key pair of data owner is %v || %v || %v\n", sko, pko, vko)
	//Upload pko and vko on the blockchain
	auth1 := utils.Transact(client, privatekey, big.NewInt(0))
	tx1, _ := Contract.UploadOwnerKey(auth1, Convert.G1ToG1Point(pko), Convert.G2ToG2Point(vko))
	receipt1, err := bind.WaitMined(context.Background(), client, tx1)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("UploadOwnerKey Gas used: %d\n", receipt1.GasUsed)
	//vko verify(on-chain)
	auth6 := utils.Transact(client, privatekey, big.NewInt(0))
	tx6, _ := Contract.VKoVerify(auth6)
	receipt6, err := bind.WaitMined(context.Background(), client, tx6)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("VKoVerify Gas used: %d\n", receipt6.GasUsed)

	VKoResult, _ := Contract.GetVKoResult(&bind.CallOpts{})
	fmt.Printf("The Verification results of vko is %v\n", VKoResult)

	//fmt.Printf("The DLEQ proof of vko is %v\n", prf_sko)
	//KeyGen_u(off-chain)
	//The key pair of data user
	sku, pku := ElGamal.EGKeyGen()
	vku := new(bn256.G2).ScalarBaseMult(sku)
	//fmt.Printf("The key pair of data user is %v || %v || %v\n", sku, pku, vku)
	//Upload pku and vku on the blockchain
	auth3 := utils.Transact(client, privatekey, big.NewInt(0))
	tx3, _ := Contract.UploadUserKey(auth3, Convert.G1ToG1Point(pku), Convert.G2ToG2Point(vku))
	receipt3, err := bind.WaitMined(context.Background(), client, tx3)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("UploadUserKey Gas used: %d\n", receipt3.GasUsed)

	//vku verify(on-chain)
	auth7 := utils.Transact(client, privatekey, big.NewInt(0))
	tx7, _ := Contract.VKuVerify(auth7)
	receipt7, err := bind.WaitMined(context.Background(), client, tx7)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("VKuVerify Gas used: %d\n", receipt7.GasUsed)
	VKuResult, _ := Contract.GetVKuResult(&bind.CallOpts{})
	fmt.Printf("The Verification results of vku is %v\n", VKuResult)
	//KeyGen_DR(off-chain)
	//The key pair of data regulators
	sk := make([]*big.Int, numShares)
	pk := make([]*bn256.G1, numShares)
	DRsPK := make([]contract.VerificationG1Point, numShares)
	for i := 0; i < numShares; i++ {
		sk[i], pk[i] = ElGamal.EGKeyGen()
		DRsPK[i] = Convert.G1ToG1Point(pk[i])
		//fmt.Printf("The key pair of data regulator%v is %v || %v\n", i, sk[i], pk[i])
	}
	//Upload all public keys of data regulators on the blockchain
	auth5 := utils.Transact(client, privatekey, big.NewInt(0))
	tx5, _ := Contract.UploadDRsKey(auth5, DRsPK)
	receipt5, err := bind.WaitMined(context.Background(), client, tx5)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("UploadDRsKey Gas used: %d\n", receipt5.GasUsed)

	fmt.Printf("The total gas used in user registration phase is %v\n", receipt1.GasUsed+receipt3.GasUsed+receipt5.GasUsed+receipt6.GasUsed+receipt7.GasUsed)

	//======================================Sensitive Message Encryption========================================//
	//Data owner encrypts the sensitive message M which is the AES key.(off-chain)
	//secret=H(sk||Nonce)
	Nonce := "QmQRmBnYFqkmun15nusG3Pj5LKaxmc2jZ4HMqGNsqqRKMY"
	secretString := sko.String() + Nonce
	secret := Convert.StringToBigInt(secretString)
	fmt.Printf("The secret value is %v\n", secret)
	//secret := big.NewInt(2)
	m, _ := rand.Int(rand.Reader, Q)
	//gT=e(g1,g2)
	M := bn256.Pair(new(bn256.G1).ScalarBaseMult(m), new(bn256.G2).ScalarBaseMult(big.NewInt(int64(1))))
	var Cipher *ElGamal.C
	//fmt.Printf("The plaintext is %v\n", M)
	starttime := time.Now().UnixMicro()
	for i := 0; i < int(n); i++ {
		Cipher = ElGamal.EGEncrypt(M, pko, secret)
	}
	endtime := time.Now().UnixMicro()
	fmt.Printf("Enc algorithm Time Used is %v us\n", (endtime-starttime)/int64(n))

	//fmt.Printf("The data ciphertext is %v\n", Cipher)

	//Upload the data ciphertext on the blockchain
	CipherC0 := Convert.G2ToG2Point(Cipher.C0)
	CipherC1 := Convert.GTToString(Cipher.C1)

	auth8 := utils.Transact(client, privatekey, big.NewInt(0))
	tx8, _ := Contract.UploadCiphertext(auth8, CipherC0, CipherC1)
	receipt8, err := bind.WaitMined(context.Background(), client, tx8)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("UploadCiphertext Gas used: %d\n", receipt8.GasUsed)
	fmt.Printf("The total gas used in sensitive message encryption phase is %v\n", receipt8.GasUsed)

	//======================================Re-encrypted Key Generation========================================//
	//Generate the polynomial coefficients(off-chain)
	coefficients := make([]*big.Int, threshold)
	coefficients[0] = secret
	for i := 1; i < threshold; i++ {
		coefficients[i], _ = rand.Int(rand.Reader, Q)
	}
	//Compute the shares of secret(off-chain)
	shares := make([]*big.Int, numShares)
	for i := 0; i < numShares; i++ {
		x := big.NewInt(int64(i + 1))
		shares[i] = KZG.EvaluatePolynomial(coefficients, x)
	}
	//fmt.Printf("The shares is %v\n", shares)

	//Generate RK(off-chain)
	var ReKey RK

	ReKey0 := make([]*bn256.G1, numShares)
	ReKey1 := make([]*bn256.G1, numShares)
	Base := make([]*bn256.G1, numShares)

	var Commit *bn256.G1
	witness := make([]*bn256.G1, numShares)
	prf_si := make([]DLEQProofG1, numShares)

	starttime = time.Now().UnixMicro()
	for j := 0; j < int(n); j++ {
		for i := 0; i < numShares; i++ {
			ReKey0[i] = new(bn256.G1).ScalarMult(PK.Tau1[0], shares[i])
			Base[i] = new(bn256.G1).Add(pko, new(bn256.G1).Add(pk[i], pku))
			ReKey1[i] = new(bn256.G1).ScalarMult(Base[i], shares[i])
		}
		ReKey.RK0 = ReKey0
		ReKey.RK1 = ReKey1
		//fmt.Printf("The re-encrypted key shares are %v\n", ReKey)

		//Generate Proof including a KZG commitment, n part KZG witnesses, and n part DLEQ Proofs
		//KZG commitment(off-chain)
		Commit = KZG.Commit(PK, coefficients)
		//fmt.Printf("The commitment is %v\n", Commit)
		//KZG witness(off-chain)
		for i := 0; i < numShares; i++ {
			x := big.NewInt(int64(i + 1))
			witness[i], _ = KZG.EvaluationProof(PK, coefficients, x, shares[i])
		}
		//fmt.Printf("The witness is %v\n", witness)
		//DLEQ Proof(on-chain)
		for i := 0; i < numShares; i++ {
			_c, _z, _rg, _rh, _ := DLEQ.DLEQProofG1(PK.Tau1[0], Base[i], ReKey.RK0[i], ReKey.RK1[i], shares[i])
			prf_si[i].C = _c
			prf_si[i].Z = _z
			prf_si[i].RG = _rg
			prf_si[i].RH = _rh
		}
		//fmt.Printf("The DLEQ proof of re-encrypted key shares are %v\n", prf_si)
	}
	endtime = time.Now().UnixMicro()
	fmt.Printf("Re-encrypted Key Generation Time Used is %v us\n", (endtime-starttime)/int64(n))

	//Upload re-encrypted key shares
	RK0s := make([]contract.VerificationG1Point, numShares)
	RK1s := make([]contract.VerificationG1Point, numShares)
	for i := 0; i < numShares; i++ {
		RK0s[i] = Convert.G1ToG1Point(ReKey.RK0[i])
		RK1s[i] = Convert.G1ToG1Point(ReKey.RK1[i])
		//fmt.Printf("The key pair of data regulator%v is %v || %v\n", i, sk[i], pk[i])
	}
	//Upload all public keys of data regulators on the blockchain
	auth9 := utils.Transact(client, privatekey, big.NewInt(0))
	tx9, _ := Contract.UploadReKeys(auth9, RK0s, RK1s)
	receipt9, err := bind.WaitMined(context.Background(), client, tx9)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("UploadReKeys Gas used: %d\n", receipt9.GasUsed)

	//Upload corresponding proofs of re-encrypted key shares.
	I := make([]*big.Int, numShares)
	w := make([]contract.VerificationG1Point, numShares)
	a1 := make([]contract.VerificationG1Point, numShares)
	a2 := make([]contract.VerificationG1Point, numShares)
	cc := make([]*big.Int, numShares)
	zz := make([]*big.Int, numShares)
	for i := 0; i < numShares; i++ {
		x := big.NewInt(int64(i + 1))
		I[i] = x
		a1[i] = Convert.G1ToG1Point(prf_si[i].RG)
		a2[i] = Convert.G1ToG1Point(prf_si[i].RH)
		cc[i] = prf_si[i].C
		zz[i] = prf_si[i].Z
		w[i] = Convert.G1ToG1Point(witness[i])
	}

	auth10 := utils.Transact(client, privatekey, big.NewInt(0))
	tx10, _ := Contract.UploadReKeysProof(auth10, I, Convert.G1ToG1Point(Commit), w, a1, a2, cc, zz)
	receipt10, err := bind.WaitMined(context.Background(), client, tx10)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("UploadReKeysProof Gas used: %d\n", receipt10.GasUsed)

	auth11 := utils.Transact(client, privatekey, big.NewInt(0))
	tx11, _ := Contract.ReKeysVerify(auth11)
	receipt11, err := bind.WaitMined(context.Background(), client, tx11)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("ReKeysVerify Gas used: %d\n", receipt11.GasUsed)

	ReKeyResult, _ := Contract.GetReKeysResult(&bind.CallOpts{})
	ReKeyResultIndex, _ := Contract.GetIndex(&bind.CallOpts{})
	fmt.Printf("The Verification result of re-encrytped key shares is %v\n", ReKeyResult)
	fmt.Printf("The index of correct re-encrytped key shares  is %v\n", ReKeyResultIndex)
	fmt.Printf("The total gas used in re-encrypted key generation phase is %v\n", receipt9.GasUsed+receipt10.GasUsed+receipt11.GasUsed)

	//======================================Data Ciphertext Re-encryption========================================//
	//Compute re-encrypted ciphertext(off-chain)
	//var ReCipher RC
	c3 := make([]*bn256.G1, numShares)

	starttime = time.Now().UnixMicro()
	for j := 0; j < int(n); j++ {
		for i := 0; i < numShares; i++ {
			c3[i] = new(bn256.G1).Add(ReKey.RK1[i], new(bn256.G1).Neg(new(bn256.G1).ScalarMult(ReKey.RK0[i], sk[i])))
		}
	}
	endtime = time.Now().UnixMicro()
	fmt.Printf("Data Ciphertext Re-encryption Time Used is %v us\n", (endtime-starttime)/int64(n))

	// ReCipher.C0 = Cipher.C0
	// ReCipher.C1 = Cipher.C1
	// ReCipher.C2 = c2
	// ReCipher.C3 = c3
	//fmt.Printf("The re-encypted ciphertext is %v\n", ReCipher)
	Convertc3 := make([]contract.VerificationG1Point, numShares)
	for i := 0; i < numShares; i++ {
		Convertc3[i] = Convert.G1ToG1Point(c3[i])
	}
	auth12 := utils.Transact(client, privatekey, big.NewInt(0))
	tx12, _ := Contract.UploadReCipher(auth12, Convertc3)
	receipt12, err := bind.WaitMined(context.Background(), client, tx12)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("UploadReCiphertext Gas used: %d\n", receipt12.GasUsed)

	auth13 := utils.Transact(client, privatekey, big.NewInt(0))
	tx13, _ := Contract.ReCipherVerify(auth13)
	receipt13, err := bind.WaitMined(context.Background(), client, tx13)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}
	fmt.Printf("ReCipherVerify Gas used: %d\n", receipt13.GasUsed)

	index, _ := Contract.GetIndex(&bind.CallOpts{})
	fmt.Printf("The correct re-ciphertext index is %v\n", index)

	C0, C1, reciphertext, _ := Contract.GetReCipher(&bind.CallOpts{})
	fmt.Printf("The total gas used in data ciphertext re-encryption phase is %v\n", receipt12.GasUsed+receipt13.GasUsed)

	//===================================Re-encrypted Ciphertext Decryption======================================//
	//Data onwer decrypts the data ciphertext(off-chain)
	fmt.Printf("\n\nThe plaintext is %v\n", M)
	var DO_M *bn256.GT
	starttime1 = time.Now().UnixMicro()
	for i := 0; i < int(n); i++ {
		secretString := sko.String() + Nonce
		secret := Convert.StringToBigInt(secretString)
		DO_M = new(bn256.GT).Add(Cipher.C1, new(bn256.GT).Neg(bn256.Pair(new(bn256.G1).ScalarMult(pko, secret), Cipher.C0)))
	}
	endtime1 = time.Now().UnixMicro()
	fmt.Printf("The Dec_2 algorithm Time Used is %v us\n", (endtime1-starttime1)/int64(n))
	fmt.Printf("The decryption result of data onwer is %v\n", DO_M)

	// //Data user decrypts the re-encrypted ciphertext
	var DU_M *bn256.GT
	for i := 0; ; i++ {
		if len(index) == threshold {
			break
		}
		index = index[:len(index)-1]
	}
	fmt.Printf("The index set is %v\n", index)
	KeyShares := make([]*bn256.G1, threshold)

	starttime = time.Now().UnixMicro()
	for j := 0; j < int(n); j++ {
		for i := 0; i < len(index); i++ {
			RecipherC3, errC3 := Convert.G1PointToG1(reciphertext[i].C3)
			if errC3 != nil {
				log.Fatalf("RecipherC3 convert failed: %v", errC3)
			}
			RecipherC2, errC2 := Convert.G1PointToG1(reciphertext[i].C2)
			if errC2 != nil {
				log.Fatalf("RecipherC2 convert failed: %v", errC2)
			}
			KeyShares[i] = new(bn256.G1).Add(RecipherC3, new(bn256.G1).Neg(new(bn256.G1).ScalarMult(RecipherC2, sku)))
		}
		Key := recoverKey(KeyShares, index, threshold)
		RecipherC0, _ := Convert.G2PointToG2(C0)
		RecipherC1, _ := Convert.StringToGT(C1)

		DU_M = ElGamal.EGDecrypt(RecipherC0, RecipherC1, Key)
		//fmt.Printf("The Dec_1 algorithm Time Used is %v us\n", (endtime1-starttime1)/int64(n))
	}
	endtime = time.Now().UnixMicro()
	fmt.Printf("Re-encrypted Ciphertext Decryption Time Used is %v us\n", (endtime-starttime)/int64(n))
	fmt.Printf("The decryption result of data user is %v\n", DU_M)
}
