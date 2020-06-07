package main

import (
	"github.com/boltdb/bolt"
	"log"
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

//迭代器属于区块链，next方法属于迭代器
//1.返回当前区块、2.指针前移
func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("迭代器遍历时bucket不应该为空,请检查！")
		}
		blockTmp := bucket.Get(it.currentHashPointer)
		//解码动作
		block = Deserialize(blockTmp)
		//fmt.Printf("block : %v\n", block)
		//hash左移
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}
