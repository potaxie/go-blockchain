package main

//10.重构代码 go run *.go
func main() {

	bc := NewBlockChian("banzhang")
	cli := CLI{bc}
	cli.Run()

	//bc := NewBlockChian()
	//bc.AddBlock("1111111111111111")
	//bc.AddBlock("2222222222222222")
	//
	////创建迭代器
	//it := bc.NewIterator()
	//
	//////调用迭代器，返回我们的每一个区块数据
	//for {
	//	//返回区块，左移
	//	block := it.Next()
	//	fmt.Printf("==================================\n")
	//	fmt.Printf("前区块hash值： %x\n", block.PrevHash)
	//	fmt.Printf("当区块hash值： %x\n", block.Hash)
	//	fmt.Printf("当前区块数据： %s\n", block.Data)
	//
	//	if len(block.PrevHash) == 0 {
	//		fmt.Printf("区块链遍历结束")
	//		break
	//	}
	//}

	//for i, block := range bc.blocks {
	//	fmt.Printf("======当前区块高度： %d===========\n", i)
	//	fmt.Printf("前区块hash值： %x\n", block.PrevHash)
	//	fmt.Printf("当区块hash值： %x\n", block.Hash)
	//	fmt.Printf("当前区块数据： %s\n", block.Data)
	//}

	//block := NewBlock("one perseon hava a btc", []byte{})
	//fmt.Printf("前区块hash值： %x\n", block.PrevHash)
	//fmt.Printf("当区块hash值： %x\n", block.Hash)
	//fmt.Printf("当前区块数据： %s\n", block.Data)

}
