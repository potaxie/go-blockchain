package main

import (
	"fmt"
)

func (cli *CLI) AddBlock(data string) {
	//cli.bc.AddBlock(data)
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
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("当区块hash值： %x\n", block.Hash)
		fmt.Printf("merkle根： %x\n", block.MerkleRoot)
		fmt.Printf("时间戳： %d\n", block.TimeStamp)
		fmt.Printf("难度值(随便写的)： %d\n", block.Difficulty)
		fmt.Printf("随机数： %d\n", block.Nonce)
		fmt.Printf("前区块hash值： %x\n", block.PrevHash)
		fmt.Printf("当前区块数据： %s\n", block.Transaction[0].TXInputs[0].Sig)

		if len(block.PrevHash) == 0 {
			fmt.Printf("区块链遍历结束")
			break
		}
	}
}

func (cli *CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)

	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("%s的余额为：%f\n", address, total)

}
func (cli *CLI) Send(from, to string, amount float64, miner, data string) {

	fmt.Printf("from :%s\n", from)
	fmt.Printf("to :%s\n", to)
	fmt.Printf("amount :%f\n", amount)
	fmt.Printf("miner :%s\n", miner)
	fmt.Printf("data :%s\n", data)

	//1.创建挖矿交易
	coinbase := NewCoinbaseTX(miner, data)
	//2.创建一个普通交易
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx == nil {
		fmt.Printf("无效的交易")
		return
	}
	//3.添加到区块
	cli.bc.AddBlock([]*Transaction{coinbase, tx})
	fmt.Printf("转账成功！")

}
