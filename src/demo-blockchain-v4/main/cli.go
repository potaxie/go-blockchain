package main

import (
	"fmt"
	"os"
	"strconv"
)

//这是一个用来接受命令行参数，并且控制区块链操作的文件

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA "add data to blockchain"
	printChain			 "正向打印区块链"
	printChainR			 "反向打印区块链"
	getBalance --address ADDRESS "获取指定地址的余额"
	send FROM TO AMOUT MINER DATA "由FROM转AMOUT给TO，由MINER挖矿，写入DATA"
`

//接受参数的动作，我们放到一个函数中

func (cli *CLI) Run() {

	//./block printChain
	//./block addBlock --data "helloworld"

	//1.得到所有命令
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("参数过少:%s\n", Usage)
		return
	}

	//2.分析命令
	cmd := args[1]

	//3.执行相应的动作
	switch cmd {
	case "addBlock":
		fmt.Printf("添加区块\n")
		if len(args) == 4 && args[2] == "--data" {
			//a.获取数据
			data := args[3]
			//b.使用bc添加区块addBlock
			cli.AddBlock(data)
		} else {
			fmt.Printf("添加区块数据使用不当，请检查\n")
		}
	case "printChain":
		fmt.Printf("正向打印区块\n")
		cli.PrintfBlockChain()
	case "printChainR":
		fmt.Printf("反向打印区块\n")
		//cli.PrintfBlockChainReverse()
	case "getBalance":
		fmt.Printf("获取余额\n")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			cli.GetBalance(address)
		}
	case "send":
		if len(args) != 7 {
			fmt.Printf("参数个数错误")
			fmt.Printf(Usage)
			return
		}
		fmt.Printf("转账开始\n")
		//	send FROM TO AMOUT MINER DATA "由FROM转AMOUT给TO，由MINER挖矿，写入DATA"
		from := args[2]
		to := args[3]
		amount, _ := strconv.ParseFloat(args[4], 64)
		miner := args[5]
		data := args[6]
		cli.Send(from, to, amount, miner, data)

	default:
		fmt.Printf("无效的命令.请检查\n")
		fmt.Printf(Usage)
	}

}
