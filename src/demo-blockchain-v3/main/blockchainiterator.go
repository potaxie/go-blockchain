package main

import (
	"github.com/boltdb/bolt"
)

type BlockChainIterator struct {
	//定义一个区块链数组
	db *bolt.DB
	//游标，用于不断索引
	currentHashPointer []byte
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {

	return &BlockChainIterator{
		bc.db,
		//最初指向区块链的最后一个区块，随着next的调用，不断变化
		bc.tail,
	}

}
