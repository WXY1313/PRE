// used to interact with the smart contract
package main

import (
	Contract "PRE/gen"
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"

	"PRE/crypto/Convert"
	"PRE/crypto/DLEQ"
	"PRE/crypto/ElGamal"
	"PRE/crypto/KZG"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"

	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
	// 1. Load encrypted wallet-file
	walletPath := "./wallet/UTC--2025-05-29T13-51-04.984946200Z--9dda53414c9a26b1054427718cd991ec14bd5fd4"
	walletData, err := os.ReadFile(walletPath)
	if err != nil {
		log.Fatalf("Read wallet-file fail: %v", err)
	}

	// 2. Decrypt wallet
	key, err := keystore.DecryptKey(walletData, "password")
	if err != nil {
		log.Fatalf("Decrypt wallet fail: %v", err)
	}

	// 3. Link Sepolia network
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/de28dcce3b8f4d319a904bfab58d1e1a")
	if err != nil {
		log.Fatalf("Fail to link Sepolia nodes: %v", err)
	}
	defer client.Close()

	// 4. Obtain chain-ID and gas price
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Fail to obtain network-ID: %v", err)
	}

	// 5. Build contract instance
	contractAddress := common.HexToAddress("0x198c4e01792E3975F6d332381AA03f7CC97740AB")
	contract, err := Contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Fail to build contract instance: %v", err)
	}

	// 6. prepare transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to build transaction signature: %v", err)
	}

	// set gas parameters
	auth.GasLimit = 300000000
	auth.Context = context.Background()

	//==========================================System Initialization===========================================//
	// the number of key shares
	numShares := 50
	// threshold value
	//threshold := 2
	threshold := numShares/2 + 1
	n := 100

	fmt.Printf("The number of shares is %v\n", numShares)
	fmt.Printf("The threshold value is %v\n", threshold)

	//Setup Algorithm(off-chain)
	PK, _ := KZG.NewTrustedSetup(numShares)
	//fmt.Printf("The system public key is %v\n",PK)
	PKTau1 := make([]Contract.VerificationG1Point, numShares)
	PKTau2 := make([]Contract.VerificationG2Point, numShares)
	//PKG2i := make([]contract.VerificationG2Point, numShares)
	for i := 0; i < numShares; i++ {
		PKTau1[i] = Convert.G1ToG1Point(PK.Tau1[i])
		PKTau2[i] = Convert.G2ToG2Point(PK.Tau2[i])
	}

	code, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal("Fail to obtain contract code：", err)
	}
	if len(code) == 0 {
		log.Fatalf("Fail to deploy contract code", contractAddress.Hex())
	}

	//Send UploadSystemKey transaction
	auth1, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to build transaction signature: %v", err)
	}
	//invoke function
	tx1, err := contract.UploadSystemKey(auth1, PKTau1, PKTau2)
	if err != nil {
		log.Fatalf("Fail to invoke UploadSystemKey: %v", err)
	}
	//wait for the transaction to succeed
	receipt1, err := bind.WaitMined(context.Background(), client, tx1)
	if err != nil {
		log.Fatalf("Fail to ensure transaction: %v", err)
	}
	if receipt1.Status == 0 {
		log.Fatal("Fial to excute transaction")
	}
	fmt.Printf("UploadSystemKey Gas used: %d\n", receipt1.GasUsed)

	//===========================================User Registration===========================================//
	//KeyGen_u(off-chain)
	//The key pair of data owner
	sko, pko := ElGamal.EGKeyGen()
	vko := new(bn256.G2).ScalarBaseMult(sko)
	//fmt.Printf("%v\n%v\n", pko, vko)

	//Upload pko and vko on the blockchain
	//Send UploadOwnerKey transaction
	auth2, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to build transaction signature: %v", err)
	}
	tx2, err := contract.UploadOwnerKey(auth2, Convert.G1ToG1Point(pko), Convert.G2ToG2Point(vko))
	if err != nil {
		log.Fatalf("Fail to invoke UploadOwnerKey: %v", err)
	}
	receipt2, err := bind.WaitMined(context.Background(), client, tx2)
	if err != nil {
		log.Fatalf("Fail to ensure transaction: %v", err)
	}
	if receipt2.Status == 0 {
		log.Fatal("Fail to excute transaction")
	}
	fmt.Printf("UploadOwnerKey Gas used: %d\n", receipt2.GasUsed)

	//Verify verification key VKo
	//Send VKoVerify transaction
	auth3, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to build transaction signature: %v", err)
	}
	tx3, err := contract.VKoVerify(auth3)
	if err != nil {
		log.Fatalf("Fail to invoke VKoVerify: %v", err)
	}
	receipt3, err := bind.WaitMined(context.Background(), client, tx3)
	if err != nil {
		log.Fatalf("Fail to ensure transaction: %v", err)
	}
	if receipt3.Status == 0 {
		log.Fatal("Fail to excute transaction")
	}
	fmt.Printf("VKoVerify used: %d\n", receipt3.GasUsed)
	VKoResult, _ := contract.GetVKoResult(&bind.CallOpts{})
	fmt.Printf("The Verification results of vko is %v\n", VKoResult)

	//The key pair of data user
	sku, pku := ElGamal.EGKeyGen()
	vku := new(bn256.G2).ScalarBaseMult(sku)
	//fmt.Printf("%v\n%v\n", pku, vku)
	//Upload pku and vku on the blockchain
	//Send UploadUserKey transaction
	auth4, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to build transaction signature: %v", err)
	}
	tx4, err := contract.UploadUserKey(auth4, Convert.G1ToG1Point(pku), Convert.G2ToG2Point(vku))
	if err != nil {
		log.Fatalf("Fail to invoke UploadUserK: %v", err)
	}
	receipt4, err := bind.WaitMined(context.Background(), client, tx4)
	if err != nil {
		log.Fatalf("Fail to ensure transaction: %v", err)
	}
	if receipt4.Status == 0 {
		log.Fatal("Fail to excute transaction")
	}
	fmt.Printf("UploadUserKey Gas used: %d\n", receipt4.GasUsed)
	//Verify verification key VKu
	//Send VKuVerify transaction
	auth5, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to transaction signature: %v", err)
	}
	tx5, err := contract.VKuVerify(auth5)
	if err != nil {
		log.Fatalf("Faic to invoke VKuVerify: %v", err)
	}
	receipt5, err := bind.WaitMined(context.Background(), client, tx5)
	if err != nil {
		log.Fatalf("Fail to ensure transaction: %v", err)
	}
	if receipt5.Status == 0 {
		log.Fatal("Fail to excute transaction")
	}
	fmt.Printf("VKuVerify used: %d\n", receipt5.GasUsed)
	VKuResult, _ := contract.GetVKuResult(&bind.CallOpts{})
	fmt.Printf("The Verification results of vku is %v\n", VKuResult)

	//KeyGen_DR(off-chain)
	//The key pair of data regulators
	sk := make([]*big.Int, numShares)
	pk := make([]*bn256.G1, numShares)
	DRsPK := make([]Contract.VerificationG1Point, numShares)
	for i := 0; i < numShares; i++ {
		sk[i], pk[i] = ElGamal.EGKeyGen()
		DRsPK[i] = Convert.G1ToG1Point(pk[i])
	}
	//Upload all public keys of data regulators on the blockchain
	auth6, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to build transaction signature: %v", err)
	}
	tx6, err := contract.UploadDRsKey(auth6, DRsPK)
	if err != nil {
		log.Fatalf("Fail to invoke UploadDRsKey: %v", err)
	}
	//8. 等待交易被挖出
	receipt6, err := bind.WaitMined(context.Background(), client, tx6)
	if err != nil {
		log.Fatalf("Fail to ensure transaction: %v", err)
	}
	if receipt4.Status == 0 {
		log.Fatal("Fail to excute transaction")
	}
	fmt.Printf("UploadDRsKey Gas used: %d\n", receipt6.GasUsed)
	fmt.Printf("The total gas used in user registration phase is %v\n", receipt2.GasUsed+receipt3.GasUsed+receipt4.GasUsed+receipt5.GasUsed+receipt6.GasUsed)

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
	Cipher = ElGamal.EGEncrypt(M, pko, secret)
	//Upload the data ciphertext on the blockchain
	CipherC0 := Convert.G2ToG2Point(Cipher.C0)
	CipherC1 := Convert.GTToString(Cipher.C1)

	//Upload all public keys of data regulators on the blockchain
	auth7, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to build transaction signature: %v", err)
	}
	tx7, err := contract.UploadCiphertext(auth7, CipherC0, CipherC1)
	if err != nil {
		log.Fatalf("Fail to invoke UploadCiphertext: %v", err)
	}
	//8. 等待交易被挖出
	receipt7, err := bind.WaitMined(context.Background(), client, tx7)
	if err != nil {
		log.Fatalf("Fail to ensure transaction: %v", err)
	}
	if receipt4.Status == 0 {
		log.Fatal("Fail to exccute transaction")
	}
	fmt.Printf("UploadCiphertext Gas used: %d\n", receipt7.GasUsed)
	fmt.Printf("The total gas used in sensitive message encryption phase is %v\n", receipt7.GasUsed)

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
	for i := 0; i < numShares; i++ {
		ReKey0[i] = new(bn256.G1).ScalarMult(PK.Tau1[0], shares[i])
		Base[i] = new(bn256.G1).Add(pko, new(bn256.G1).Add(pk[i], pku))
		ReKey1[i] = new(bn256.G1).ScalarMult(Base[i], shares[i])
	}
	ReKey.RK0 = ReKey0
	ReKey.RK1 = ReKey1
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

	//Upload re-encrypted key shares
	RK0s := make([]Contract.VerificationG1Point, numShares)
	RK1s := make([]Contract.VerificationG1Point, numShares)
	for i := 0; i < numShares; i++ {
		RK0s[i] = Convert.G1ToG1Point(ReKey.RK0[i])
		RK1s[i] = Convert.G1ToG1Point(ReKey.RK1[i])
		//fmt.Printf("The key pair of data regulator%v is %v || %v\n", i, sk[i], pk[i])
	}

	//Upload corresponding proofs of re-encrypted key shares.
	I := make([]*big.Int, numShares)
	w := make([]Contract.VerificationG1Point, numShares)
	a1 := make([]Contract.VerificationG1Point, numShares)
	a2 := make([]Contract.VerificationG1Point, numShares)
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

	var ReKeyVerifyResult bool = true
	for j := 0; j < n; j++ {
		for i := 0; i < numShares; i++ {
			if !KZG.Verify(PK, Commit, witness[i], big.NewInt(int64(i+1)), ReKey0[i]) {
				ReKeyVerifyResult = false
				fmt.Printf("%vth share is fault\n", i)
				break
			} else {
				if !DLEQ.VerifyG1(prf_si[i].C, prf_si[i].Z, PK.Tau1[0], Base[i], ReKey.RK0[i], ReKey.RK1[i], prf_si[i].RG, prf_si[i].RH) {
					fmt.Printf("%vth share is fault\n", i)
					ReKeyVerifyResult = false
					break
				}
			}
		}
	}
	fmt.Printf("ReKeyVerifyResult Result:%v\n", ReKeyVerifyResult)
	//Upload all public keys of data regulators on the blockchain
	auth8, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to build transaction signature: %v", err)
	}
	tx8, err := contract.UploadReKeys(auth8, RK0s, RK1s, I, Convert.G1ToG1Point(Commit), w, a1, a2, cc, zz)
	if err != nil {
		log.Fatalf("Fail to invoke UploadReKeys: %v", err)
	}
	//8. 等待交易被挖出
	receipt8, err := bind.WaitMined(context.Background(), client, tx8)
	if err != nil {
		log.Fatalf("Fail to ensure transaction: %v", err)
	}
	if receipt8.Status == 0 {
		log.Fatal("Fail to excute transaction")
	}
	fmt.Printf("UploadReKeys Gas used: %d\n", receipt8.GasUsed)

	ReKeyResult, _ := contract.GetReKeysResult(&bind.CallOpts{})
	//ReKeyResultIndex, _ := contract.GetIndex(&bind.CallOpts{})
	fmt.Printf("The Verification result of re-encrytped key shares is %v\n", ReKeyResult)
	//fmt.Printf("The index of correct re-encrytped key shares  is %v\n", ReKeyResultIndex)
	fmt.Printf("The total gas used in re-encrypted key generation phase is %v\n", receipt8.GasUsed)

	//ReKeyResultIndex, _ := contract.GetIndex(&bind.CallOpts{})
	//fmt.Printf("The only Verification result of re-encrytped key shares is %v\n", OnlyReKeyResult)
	//======================================Data Ciphertext Re-encryption========================================//
	//Compute re-encrypted ciphertext(off-chain)
	//var ReCipher RC
	c3 := make([]*bn256.G1, numShares)

	for i := 0; i < numShares; i++ {
		c3[i] = new(bn256.G1).Add(ReKey.RK1[i], new(bn256.G1).Neg(new(bn256.G1).ScalarMult(ReKey.RK0[i], sk[i])))
	}

	var ReEncVerifyResult bool = true
	for j := 0; j < n; j++ {
		for i := 0; i < numShares; i++ {
			left := bn256.Pair(c3[i], PK.Tau2[0])
			right := new(bn256.GT).Add(bn256.Pair(ReKey0[i], vko), bn256.Pair(ReKey0[i], vku))
			if left.String() != right.String() {
				ReEncVerifyResult = false
				break
			}
		}
	}
	fmt.Printf("ReEncVerify Result:%v\n", ReEncVerifyResult)

	Convertc3 := make([]Contract.VerificationG1Point, numShares)
	for i := 0; i < numShares; i++ {
		Convertc3[i] = Convert.G1ToG1Point(c3[i])
	}

	auth11, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Fail to ensure transaction signature: %v", err)
	}
	tx11, err := contract.UploadReCipher(auth11, Convertc3)
	if err != nil {
		log.Fatalf("Fail to invoke UploadReCipher: %v", err)
	}
	//8. 等待交易被挖出
	receipt11, err := bind.WaitMined(context.Background(), client, tx11)
	if err != nil {
		log.Fatalf("Fail to ensure transaction: %v", err)
	}
	if receipt11.Status == 0 {
		log.Fatal("Fail to excute transaction")
	}
	fmt.Printf("UploadReCipher Gas used: %d\n", receipt11.GasUsed)

	index, _ := contract.GetIndex(&bind.CallOpts{})
	fmt.Printf("The correct re-ciphertext index is %v\n", index)

	C0, C1, reciphertext, _ := contract.GetReCipher(&bind.CallOpts{})
	fmt.Printf("The total gas used in data ciphertext re-encryption phase is %v\n", receipt11.GasUsed)

	//===================================Re-encrypted Ciphertext Decryption======================================//
	//Data onwer decrypts the data ciphertext(off-chain)
	//fmt.Printf("\n\nThe plaintext is %v\n", M)
	var DO_M *bn256.GT
	secretString = sko.String() + Nonce
	secret = Convert.StringToBigInt(secretString)
	DO_M = new(bn256.GT).Add(Cipher.C1, new(bn256.GT).Neg(bn256.Pair(new(bn256.G1).ScalarMult(pko, secret), Cipher.C0)))
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
	fmt.Printf("The decryption result of data user is %v\n", DU_M)
}
