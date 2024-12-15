package Convert

import (
	"PRE/compile/contract"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

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

func G2ToG2Point(point *bn256.G2) contract.VerificationG2Point {
	// Marshal the G1 point to get the X and Y coordinates as bytes
	pointBytes := point.Marshal()
	//fmt.Println(point.Marshal())

	// Create big.Int for X and Y coordinates
	a1 := new(big.Int).SetBytes(pointBytes[:32])
	a2 := new(big.Int).SetBytes(pointBytes[32:64])
	b1 := new(big.Int).SetBytes(pointBytes[64:96])
	b2 := new(big.Int).SetBytes(pointBytes[96:128])

	g2Point := contract.VerificationG2Point{
		X: [2]*big.Int{a2, a1},
		Y: [2]*big.Int{b2, b1},
	}
	return g2Point
}

// GTToString 将 bn256.GT 元素编码为 Base64 字符串
func GTToString(gt *bn256.GT) (string, error) {
	if gt == nil {
		return "", errors.New("invalid GT element: nil")
	}

	// 使用 Marshal 将 GT 序列化为字节切片
	gtBytes := gt.Marshal()

	// 使用 Base64 编码为字符串
	encoded := base64.StdEncoding.EncodeToString(gtBytes)
	return encoded, nil
}

// StringToGT 从 Base64 字符串解码还原 bn256.GT 元素
func StringToGT(encoded string) (*bn256.GT, error) {
	if encoded == "" {
		return nil, errors.New("input string is empty")
	}

	// 使用 Base64 解码为字节切片
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("base64 decoding failed: %v", err)
	}

	// 使用 Unmarshal 还原为 GT 元素
	gt := new(bn256.GT)
	if _, err := gt.Unmarshal(decodedBytes); err != nil {
		return nil, fmt.Errorf("unmarshalling to GT failed: %v", err)
	}
	return gt, nil
}

func G1ToBigIntArray(point *bn256.G1) [2]*big.Int {
	// Marshal the G1 point to get the X and Y coordinates as bytes
	pointBytes := point.Marshal()

	// Create big.Int for X and Y coordinates
	x := new(big.Int).SetBytes(pointBytes[:32])
	y := new(big.Int).SetBytes(pointBytes[32:64])

	return [2]*big.Int{x, y}
}
