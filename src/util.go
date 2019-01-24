package main

import (
	"crypto/sha256"
	"bytes"
	"encoding/binary"
	"math/big"
)

const BITS uint64 = 12

func GetHash(bytes []byte) [32]byte {
	sum256 := sha256.Sum256(bytes)
	return sum256
}

func BlockToHash(block *Block) [32]byte{

	temp := [][]byte{
		uint2Bytes(block.Version),
		block.PrevBlockHash,
		block.NowHash,
		block.MerkelRoot,
		uint2Bytes(block.TimeStamp),
		uint2Bytes(block.Bits),
		uint2Bytes(block.Nonce),
		block.Data,
	}

	tempBytes := bytes.Join(temp, []byte{})
	hash := GetHash(tempBytes)

	return hash
}

func uint2Bytes(num uint64) []byte {

	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func HashToBigInt(hash [32]byte) big.Int {
	var bitIntTemp big.Int
	bitIntTemp.SetBytes(hash[:])
	return bitIntTemp
}


