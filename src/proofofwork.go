package main

import (
	"math/big"
	"fmt"
)

func ProofOfWork(block Block) ([]byte, uint64) {

	target := GetTarget(block)

	hash, nonce := GetHashAndNonce(&block, target)

	return hash[:], nonce
}

func GetHashAndNonce(block *Block, target big.Int) ([32]byte, uint64) {

	var result [32]byte

	for {
		hash := BlockToHash(block)

		fmt.Printf("挖矿中...%x\n", hash)

		bigInt := HashToBigInt(hash)

		if bigInt.Cmp(&target) < 0 {
			result = hash
			break
		} else {
			block.Nonce++
		}

	}

	return result, block.Nonce

}

func GetTarget(block Block) big.Int {
	bits := block.Bits
	bigIntTemp := big.NewInt(1)
	bigIntTemp.Lsh(bigIntTemp, uint(256-bits))

	return *bigIntTemp
}
