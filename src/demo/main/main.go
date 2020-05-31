package main

import (
	"fmt"
)

/**
0.定义结构
1.前区块hash
2.当前区块hash
3.数据
4.创建区块
5.生成hash
6.引入区块链
*/

type Block struct {
	PrevHash []byte
	Hash     []byte
	Data     []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, //先填空，后面在计算
		Data:     []byte(data),
	}

	return &block
}

func main() {

	fmt.Printf("hello")

	block := NewBlock("one perseon hava a btc", []byte{})

	fmt.Printf("前区块hash值： %x\n", block.PrevHash)
	fmt.Printf("当区块hash值： %x\n", block.Hash)
	fmt.Printf("当前区块数据： %s\n", block.Data)

}
