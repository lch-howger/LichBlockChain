package main

import "time"

type Block struct {
	//1. 版本号
	Version uint64

	//2. 前区块哈希
	PrevBlockHash []byte

	//3. 当前区块哈希, 这是为了方便加入的字段，正常区块中没有这个字段
	NowHash []byte

	//4. 梅克尔根
	MerkelRoot []byte //先忽略不管

	//5. 时间戳, 从1970.1.1至今描述，一个数字
	TimeStamp uint64

	//6. 难度值，一个数字，可以推导出难度哈希值
	Bits uint64

	//7. 随机数Nonce，挖矿要求得值
	Nonce uint64

	//8. 数据
	Data []byte
}

func NewBlock(data string, preBlockHash []byte) *Block {
	var block Block
	block.Version = 1
	block.PrevBlockHash = preBlockHash
	block.NowHash = []byte{}
	block.MerkelRoot = nil
	block.TimeStamp = uint64(time.Now().Unix())
	block.Bits = BITS
	block.Nonce = 0
	block.Data = []byte(data)

	hash, nonce := ProofOfWork(block)

	block.NowHash = hash
	block.Nonce = nonce

	return &block
}
