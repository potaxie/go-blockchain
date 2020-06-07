package main

import (
	"fmt"
	"os"
)

//这是一个用来接受命令行参数，并且控制区块链操作的文件

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA "add data to blockchain"
	printChain			 "print all blockchain data"
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
			//cli.bc.AddBlock(data)
			cli.AddBlock(data)
		} else {
			fmt.Printf("添加区块数据使用不当，请检查\n")
		}
	case "printChain":
		fmt.Printf("打印区块\n")
		cli.PrintfBlockChain()
	default:
		fmt.Printf("无效的命令.请检查\n")
		fmt.Printf(Usage)
	}

}
