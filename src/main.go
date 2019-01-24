package main

import "fmt"

func main() {

	blockChain := NewBlockChain("aaa")
	blockChain.AddBlock("bbb")
	blockChain.AddBlock("ccc")
	blockChain.AddBlock("ddd")
	blockChain.AddBlock("eee")

	for key, block := range blockChain.Blocks {
		fmt.Printf("挖到了第%d个,hash为:%x\n", key, block.NowHash)
	}

}
