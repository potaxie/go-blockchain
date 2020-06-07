package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Printf("添加区块成功!\n")
}

func (cli *CLI) PrintfBlockChain() {
	bc := cli.bc
	//创建迭代器
	it := bc.NewIterator()

	////调用迭代器，返回我们的每一个区块数据
	for {
		//返回区块，左移
		block := it.Next()
		fmt.Printf("==================================\n")
		fmt.Printf("前区块hash值： %x\n", block.PrevHash)
		fmt.Printf("当区块hash值： %x\n", block.Hash)
		fmt.Printf("当前区块数据： %s\n", block.Data)

		if len(block.PrevHash) == 0 {
			fmt.Printf("区块链遍历结束")
			break
		}
	}
}
