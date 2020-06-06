package main

//10.重构代码 go run *.go
func main() {
	bc := NewBlockChian()

	bc.AddBlock("first tx 10 btc")
	bc.AddBlock("second tx 10 btc")

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
