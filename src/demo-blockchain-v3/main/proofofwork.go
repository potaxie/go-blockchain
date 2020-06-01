package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//定义一个工作量证明的结构proofofwork

type ProofOfWork struct {
	//a.block
	block *Block
	//b.目标值 ,一个非常大的数
	target *big.Int
}

//提供创建Pow的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	//我们指定的难度值，现在是String，需要转换
	targetStr := "0000300000000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	//将难度值赋给big.int，指定16进制格式
	pow.target = &tmpInt
	return &pow
}

//提供不断计算hash的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {

	var nonce uint64
	block := pow.block

	for {
		//1.拼装数据（区块的数据，还有不断变化的随机数）
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkleRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		//将二维切片数组拼接起来，返回一个一维切片
		blockInfo := bytes.Join(tmp, []byte{})

		//2.做hash计算
		hash := sha256.Sum256(blockInfo)

		//3.与pow中target进行比较（找到退出、没找到继续，随机数+1）
		tmpInt := big.Int{}

		//将得到的hash数组转换成big.int
		tmpInt.SetBytes(hash[:])

		if tmpInt.Cmp(pow.target) == -1 {
			//找到了，退出
			fmt.Printf("挖矿成功 hash：%x,nonce:%d \n", hash, nonce)
			//break
			return hash[:], nonce
		} else {
			//没找到，随机数+1
			fmt.Printf("继续寻找 hash：%x,nonce:%d \n", hash, nonce)
			nonce++
		}
	}
	//return []byte("hello world"), 10
}
