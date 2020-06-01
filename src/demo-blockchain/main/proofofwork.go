package main

import "math/big"

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
	targetStr := "00001000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	//将难度值赋给big.int，指定16进制格式
	pow.target = &tmpInt
	return &pow
}

//提供不断计算hash的函数
//run()

//提供校验函数
//isVaild()
