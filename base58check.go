package base58check

import (
	"bytes"
	"math/big"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func b58encode(data string) string {
	var encoded string
	decimalData := new(big.Int)
	decimalData.SetBytes([]byte(data))
	divisor, zero := big.NewInt(58), big.NewInt(0)
	for decimalData.Cmp(zero) > 0 {
		mod := new(big.Int)
		decimalData.DivMod(decimalData, divisor, mod)
		encoded = string(alphabet[mod.Int64()]) + encoded
	}
	return encoded
}

func b58decode(data string) string {
	decimalData := new(big.Int)
	alphabetBytes := []byte(alphabet)
	multiplier := big.NewInt(58)
	for _, value := range data {
		pos := bytes.IndexByte(alphabetBytes, byte(value))
		if pos == -1 {
			panic("Fuck off")
		}
		decimalData.Mul(decimalData, multiplier)
		decimalData.Add(decimalData, big.NewInt(int64(pos)))
	}
	return string(decimalData.Bytes())
}
