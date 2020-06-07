package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	//定义一个区块链数组
	db   *bolt.DB
	tail []byte //存储最后一个区块的hash
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

//5.定义一个区块链
func NewBlockChian() *BlockChain {

	//最后一个区块的hash，从数据库里读出来
	var lastHash []byte
	//1.打开数据库, 创建bolt数据库本地文件,strings test.db
	db, err := bolt.Open(blockChainDb, 0600, nil)

	//defer db.Close() //go特性

	if err != nil {
		log.Panic("打开数据库失败！")
	}

	//将要改写数据库
	db.Update(func(tx *bolt.Tx) error {
		//2.找到抽屉bucket，如果没有就创建
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//没有抽屉，需要创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket(b1)失败！")
			}
			genesisblock := GenesisBlock()

			//3.写数据
			//hash 作为key，block的字节流作为value
			bucket.Put(genesisblock.Hash, genesisblock.Serialize())
			//存hash
			bucket.Put([]byte("lastHashKey"), genesisblock.Hash)
			lastHash = genesisblock.Hash

			//这是为了读数据测试，马上删掉
			//blockBytes := bucket.Get(genesisblock.Hash)
			//block := Deserialize(blockBytes)
			//fmt.Printf("block info : %v\n", block)

		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

//定义创世块
func GenesisBlock() *Block {
	return NewBlock("创世块定义", []byte{})
}

//5.添加区块
func (bc *BlockChain) AddBlock(data string) {
	//如何获取前区块hash

	db := bc.db         //区块数据库
	lastHash := bc.tail //最后一个区块的hash

	db.Update(func(tx *bolt.Tx) error {
		//完成数据添加
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("不应该为空，请检查！")
		}
		//a.创建新的区块
		block := NewBlock(data, lastHash)
		//b.添加到区块db中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("lastHashKey"), block.Hash)
		//lastHash = block.Hash //更新内存
		//c.更新下内存中的区块链，指的是把最后的小尾巴更新一下
		bc.tail = block.Hash
		return nil
	})

}
