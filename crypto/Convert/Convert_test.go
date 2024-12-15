package Convert_test

import(
      "PRE/crypto/Convert"
      "fmt"
      "testing"
      "crypto/rand"
      bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

var Q, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)

func TestConvertFunction(t *testing.T){
      //Convert between bn256.G1 and G1Point
      k,_:=rand.Int(rand.Reader, Q)
      g1Point= new(bn256.G1).ScalarBaseMult(k)
      
      //Encrypt the plaintext to ciphertext
      EK:=ElGamal.EGEncrypt(K,pk,len(K))
      //Decrypt the ciphertext
	  _K:=ElGamal.EGDecrypt(EK,sk,len(K))
	  fmt.Printf("解密后的明文信息：%s\n",_K)
}
