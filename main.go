package main

import (
	"fmt"
	"PRE/crypto/KZG"
	"PRE/crypto/ElGamal"
	"math/big"
	"crypto/rand"
	"PRE/crypto/DLEQ"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	/*
	"context"
	"PRE/compile/contract"
	"PRE/utils"
	"log"
	
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/core/types"
	*/
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
	C0  *bn256.G2
	C1  *bn256.GT
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



/*
func G1ToG1Point(bn256Point *bn256.G1) contract.VerificationG1Point {
	// Marshal the G1 point to get the X and Y coordinates as bytes
	point := bn256Point.Marshal()

	// Create big.Int for X and Y coordinates
	x := new(big.Int).SetBytes(point[:32])
	y := new(big.Int).SetBytes(point[32:64])

	g1Point := contract.VerificationG1Point{
		X: x,
		Y: y,
	}
	return g1Point
}

func G1ToBigIntArray(point *bn256.G1) [2]*big.Int {
	// Marshal the G1 point to get the X and Y coordinates as bytes
	pointBytes := point.Marshal()

	// Create big.Int for X and Y coordinates
	x := new(big.Int).SetBytes(pointBytes[:32])
	y := new(big.Int).SetBytes(pointBytes[32:64])

	return [2]*big.Int{x, y}
}
*/

func main() {
//==========================================System Initialization===========================================//
	// the number of key shares
	numShares := 10
	// threshold value
	threshold := numShares/2 + 1
	//threshold := 2*numShares/3 + 1

	fmt.Printf("The number of shares is %v\n", numShares)
	fmt.Printf("The threshold value is %v\n", threshold)

	//Setup Algorithm
	PK,_:= KZG.NewTrustedSetup(numShares)
	//fmt.Printf("The system public key is %v\n",PK)

//===========================================User Registration===========================================//
	//KeyGen_u
	//The key pair of data owner
	sko,pko:=ElGamal.EGKeyGen()
	vko:=new(bn256.G2).ScalarBaseMult(sko)	
	fmt.Printf("The key pair of data owner is %v || %v || %v\n", sko,pko,vko)
	//DLEQProof(g1,pko,g2,vko,sko)
	var prf_sko DLEQProofG1_G2
	c,z,rg,rh,_:= DLEQ.DLEQProofG1_G2(PK.Tau1[0], PK.Tau2[0], pko, vko, sko)
	prf_sko.C=c
	prf_sko.Z=z
	prf_sko.RG=rg
	prf_sko.RH=rh
	fmt.Printf("The DLEQ proof of vko is %v\n",prf_sko)
	//KeyGen_u
	//The key pair of data user
	sku,pku:=ElGamal.EGKeyGen()
	vku:=new(bn256.G2).ScalarBaseMult(sku)
	fmt.Printf("The key pair of data user is %v || %v || %v\n", sku,pku,vku)
	var prf_sku DLEQProofG1_G2
	c,z,rg,rh,_= DLEQ.DLEQProofG1_G2(PK.Tau1[0], PK.Tau2[0], pku, vku, sku)
	prf_sku.C=c
	prf_sku.Z=z
	prf_sku.RG=rg
	prf_sku.RH=rh
	fmt.Printf("The DLEQ proof of vku is %v\n",prf_sku)
	//KeyGen_DR	
	//The key pair of data regulators
	sk:=make([]*big.Int,numShares)
	pk:=make([]*bn256.G1,numShares)
	for i:=0;i<numShares;i++{
		sk[i],pk[i]=ElGamal.EGKeyGen()
		fmt.Printf("The key pair of data regulator%v is %v || %v\n", i,sk[i],pk[i])
	}
	
	//VKVerify
	//vko verify
	DLEQResult := DLEQ.VerifyG1_G2(prf_sko.C, prf_sko.Z, PK.Tau1[0], PK.Tau2[0], pko, vko, prf_sko.RG, prf_sko.RH)

	if DLEQResult == nil {
		fmt.Printf("\n\nThe verification of vko passes!\n\n")
	}else{
		fmt.Printf("\n\nThe verification of vko fails to pass!\n\n")
	}
	//vku verify
	DLEQResult = DLEQ.VerifyG1_G2(prf_sku.C, prf_sku.Z, PK.Tau1[0], PK.Tau2[0], pku, vku, prf_sku.RG, prf_sku.RH)

	if DLEQResult == nil {
		fmt.Printf("\n\nThe verification of vku passes!\n\n")
	}else{
		fmt.Printf("\n\nThe verification of vku fails to pass!\n\n")
	}
//======================================Sensitive Message Encryption========================================//
	//Data owner encrypts the sensitive message M which is the AES key.
	//secret=H(sk||Nonce)
	secret:=big.NewInt(int64(2))
	m,_ := rand.Int(rand.Reader, Q)
	//gT=e(g1,g2)
	M:=bn256.Pair(new(bn256.G1).ScalarBaseMult(m),new(bn256.G2).ScalarBaseMult(big.NewInt(int64(1))))
	fmt.Printf("The plaintext is %v\n",M)
	Cipher:= ElGamal.EGEncrypt(M,pko,secret)
	fmt.Printf("The data ciphertext is %v\n",Cipher)
	/*
	//Pair Check
	K:=new(bn256.G1).ScalarMult(pko,s)
	result:=ElGamal.EGDecrypt(C,K)
	fmt.Printf("The decryption result is %v\n",result)
	*/
//======================================Re-encrypted Key Generation========================================//

	//Generate the polynomial coefficients
	coefficients := make([]*big.Int, threshold)
	coefficients[0] = secret
	for i := 1; i < threshold; i++ {
		coefficients[i], _ = rand.Int(rand.Reader, Q)
	}

	//Compute the shares of secret
	shares := make([]*big.Int, numShares)
	for i := 0; i < numShares; i++ {
		x := big.NewInt(int64(i + 1))
		shares[i] = KZG.EvaluatePolynomial(coefficients, x)
	}
	fmt.Printf("The shares is %v\n",shares)	
	
	//Generate RK
	var ReKey RK
	ReKey0:=make([]*bn256.G1,numShares)
	ReKey1:=make([]*bn256.G1,numShares)
	Base:=make([]*bn256.G1,numShares)
	for i:=0;i<numShares;i++{
		ReKey0[i]=new(bn256.G1).ScalarMult(PK.Tau1[0],shares[i])
		Base[i]=new(bn256.G1).Add(pko,new(bn256.G1).Add(pk[i],pku))
		ReKey1[i]=new(bn256.G1).ScalarMult(Base[i],shares[i])
	}
	ReKey.RK0=ReKey0
	ReKey.RK1=ReKey1
	fmt.Printf("The re-encrypted key shares are %v\n",ReKey)
	
	
	
	//Generate Proof including a KZG commitment, n part KZG witnesses, and n part DLEQ Proofs
	//KZG commitment
	Commit:=KZG.Commit(PK,coefficients)
	fmt.Printf("The commitment is %v\n",Commit)
	//KZG witness
	witness:=make([]*bn256.G1,numShares)
	for i:=0;i<numShares;i++{
		x := big.NewInt(int64(i + 1))
		witness[i],_=KZG.EvaluationProof(PK, coefficients, x, shares[i])
	}
	fmt.Printf("The witness is %v\n",witness)
	//DLEQ Proof
	prf_si := make([] DLEQProofG1,numShares)
	for i:=0;i<numShares;i++{
		_c,_z,_rg,_rh,_ := DLEQ.DLEQProofG1(PK.Tau1[0], Base[i], ReKey.RK0[i], ReKey.RK1[i], shares[i])
		prf_si[i].C=_c
		prf_si[i].Z=_z
		prf_si[i].RG=_rg
		prf_si[i].RH=_rh
	}
	fmt.Printf("The DLEQ proof of re-encrypted key shares are %v\n",prf_si)
	
	//ReKeyVerify: Verify the corresponding proof of ReKey
	for i:=0;i<numShares;i++{
		x := big.NewInt(int64(i + 1))
		if KZG.Verify(PK, Commit, witness[i], x, ReKey.RK0[i]) == true && DLEQ.VerifyG1(prf_si[i].C, prf_si[i].Z, PK.Tau1[0], new(bn256.G1).Add(pko,new(bn256.G1).Add(pk[i],pku)), ReKey.RK0[i], ReKey.RK1[i], prf_si[i].RG, prf_si[i].RH) == nil{
			fmt.Printf("The index %v of re-encrypted key shares passes the check!\n",i)
		}else{
			fmt.Printf("The index %v of re-encrypted key shares fails to pass the check!\n",i)
		}
		
	}
	
//======================================Data Ciphertext Re-encryption========================================//	
//Compute re-encrypted ciphertext
	var ReCipher RC
	c2:=make([]*bn256.G1,numShares)
	c3:=make([]*bn256.G1,numShares)
	for i:=0;i<numShares;i++{
		c2[i]=ReKey.RK0[i]
		c3[i]=new(bn256.G1).Add(ReKey.RK1[i], new(bn256.G1).Neg(new(bn256.G1).ScalarMult(ReKey.RK0[i],sk[i])))
	}
	ReCipher.C0=Cipher.C0
	ReCipher.C1=Cipher.C1
	ReCipher.C2=c2
	ReCipher.C3=c3
	fmt.Printf("The re-encypted ciphertext is %v\n",ReCipher)

//Verify re-encrypted ciphertext
	var I []*big.Int
	for i:=0;i<numShares;i++{
		e1:=bn256.Pair(ReCipher.C3[i],PK.Tau2[0])
		e2:=bn256.Pair(ReCipher.C2[i],new(bn256.G2).Add(vko,vku))
		if  e1.String() == e2.String() {
			fmt.Printf("The index %v of ReCipher passes the check!\n",i)
			x := big.NewInt(int64(i + 1))
			I=append(I,x)
		}else{
		fmt.Printf("The index %v of ReCipher fails to pass the check!\n",i)
		}
		if len(I)==threshold+1{
			break
		}
		
	}
	fmt.Printf("The index of correct re-encrypted ciphertext shares is %v\n",I)
//===================================Re-encrypted Ciphertext Decryption======================================//
	//Data onwer decrypts the data ciphertext
	fmt.Printf("The plaintext is %v\n",M)
	DO_M:=new(bn256.GT).Add(Cipher.C1,new(bn256.GT).Neg(bn256.Pair(new(bn256.G1).ScalarMult(pko,secret),Cipher.C0)))
	fmt.Printf("The decryption result of data onwer is %v\n",DO_M)
	//Data user decrypts the re-encrypted ciphertext	
	KeyShares:=make([]*bn256.G1,len(I))
	for i:=0;i<len(I);i++{
		KeyShares[i]=new(bn256.G1).Add(ReCipher.C3[i],new(bn256.G1).Neg(new(bn256.G1).ScalarMult(ReCipher.C2[i],sku)))
	}
	Key:=recoverKey(KeyShares, I, threshold)
	DU_M:=ElGamal.EGDecrypt(ReCipher.C0, ReCipher.C1, Key)
	fmt.Printf("The decryption result of data user is %v\n",DU_M)
}

