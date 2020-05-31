package main

import (
	"crypto/sha256"
	"fmt"
)

/**
0.定义结构
1.前区块hash
2.当前区块hash
3.数据
*/

type Block struct {
	PrevHash []byte
	Hash     []byte
	Data     []byte
}

//4.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, //先填空，后面在计算
		Data:     []byte(data),
	}
	block.SetHash()
	return &block
}

//5.生成hash
func (block *Block) SetHash() {
	//拼装数据
	blockInfo := append(block.PrevHash, block.Data...) //将数组打散
	//sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//6.引入区块链,将所有区块链接，定义一个区块链结构
type BlockChain struct {
	//定义一个区块链数组
	blocks []*Block
}

//7.定义一个区块链
func NewBlockChian() *BlockChain {
	genesisblock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisblock},
	}
}

//8.定义创世块
func GenesisBlock() *Block {
	return NewBlock("创世块定义", []byte{})
}

func main() {

	fmt.Printf("hello")

	bc := NewBlockChian()
	for i, block := range bc.blocks {
		fmt.Printf("======当前区块高度： %d===========\n", i)
		fmt.Printf("前区块hash值： %x\n", block.PrevHash)
		fmt.Printf("当区块hash值： %x\n", block.Hash)
		fmt.Printf("当前区块数据： %s\n", block.Data)
	}

	//block := NewBlock("one perseon hava a btc", []byte{})
	//fmt.Printf("前区块hash值： %x\n", block.PrevHash)
	//fmt.Printf("当区块hash值： %x\n", block.Hash)
	//fmt.Printf("当前区块数据： %s\n", block.Data)

}
