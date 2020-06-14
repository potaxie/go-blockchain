package main

import (
	"bytes"
	"fmt"
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
func NewBlockChian(address string) *BlockChain {

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
			genesisblock := GenesisBlock(address)
			fmt.Printf("genesisblock: %s\n", genesisblock)

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
func GenesisBlock(address string) *Block {
	coinbase := NewCoinbaseTX(address, "genesis")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

//5.添加区块
func (bc *BlockChain) AddBlock(txs []*Transaction) {
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
		block := NewBlock(txs, lastHash)
		//b.添加到区块db中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("lastHashKey"), block.Hash)
		//lastHash = block.Hash //更新内存
		//c.更新下内存中的区块链，指的是把最后的小尾巴更新一下
		bc.tail = block.Hash
		return nil
	})

}

func (bc *BlockChain) PrintChain() {
	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		//Assume bucket exists and has keys
		b := tx.Bucket([]byte("blockBucket"))

		//从第一个key->value进行遍历，到最后一个固定的key时直接返回
		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}
			block := Deserialize(v)
			fmt.Printf("============区块高度：%d======================\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号: %d\n", block.Version)
			fmt.Printf("当区块hash值： %x\n", block.Hash)
			fmt.Printf("merkle根： %x\n", block.MerkleRoot)
			fmt.Printf("时间戳： %d\n", block.TimeStamp)
			fmt.Printf("难度值(随便写的)： %d\n", block.Difficulty)
			fmt.Printf("随机数： %d\n", block.Nonce)
			fmt.Printf("前区块hash值： %x\n", block.PrevHash)
			fmt.Printf("当前区块数据： %s\n", block.Transaction[0].TXInputs[0].Sig)
			return nil
		})
		return nil
	})
}

//gebalance 找到指定地址的所有的utxo
func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	//map[交易id][]int64
	spentOutputs := make(map[string][]int64)

	//创建迭代器
	it := bc.NewIterator()

	for {
		//1.遍历区块
		block := it.Next()
		//2.遍历交易
		for _, tx := range block.Transaction {
			fmt.Printf("current txid : %x\n", tx.TXID)

		OUTPUT:

			//3.遍历output，找到和自己相关的utxo（在添加output之前检查一下是否意见消耗过）
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index :%d\n", i)
				//在这里做一个过滤，将所有消耗过的outputs和当前的所即添加的output对比一下
				//如果相同，则跳过，否则添加
				//如果当前的交易id存在于我们意见表示的map，那么说明这个交易里面有消耗过的output

				//map[2222]=[]int64{0}
				//map[3333]=[]int64{0,1}
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						//[]int64{0,1},j:0,1
						if int64(i) == j {
							//当前准备添加output已经消耗过了，不要在加了
							continue OUTPUT
						}
					}
				}

				//这个output和我们的目标的地址相同，满足条件，加到返回UTXO数组中
				if output.PubKeyHash == address {
					UTXO = append(UTXO, output)
				}
			}

			//如果当前交易是挖矿交易的话，那么不做遍历

			//4.遍历input，找到自己花费过的utxo的集合(把自己消耗过的标记出来)
			//我们定义一个map来保存消费过的output，key是这个output的交易id，value是这个交易中索引的数组

			//map[交易id][]int64
			for _, input := range tx.TXInputs {
				//判断一下当前这个input和目标（李四）是否一只，如果相同，说明这个是李四消耗过的output，就加进来
				if input.Sig == address {
					//spentOutputs :=make(map[string][]int64)
					indexArray := spentOutputs[string(input.TXid)]
					indexArray = append(indexArray, input.Index)
					//map[2222]=[]int64{0}
					//map[3333]=[]int64{0,1}
				}
			}
		}
		if len(block.PrevHash) == 0 {
			break
			fmt.Printf("区块遍历退出")
		}

	}

	return UTXO
}
