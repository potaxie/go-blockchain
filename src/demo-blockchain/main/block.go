package main

import (
	"crypto/sha256"
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
