package main

//6.引入区块链,将所有区块链接，定义一个区块链结构
type BlockChain struct {
	//定义一个区块链数组
	blocks []*Block
}

//7.定义一个区块链
func NewBlockChian() *BlockChain {
	genesisblock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisblock},
	}
}

//8.定义创世块
func GenesisBlock() *Block {
	return NewBlock("创世块定义", []byte{})
}

//9.添加区块

func (bc *BlockChain) AddBlock(data string) {
	//获取前区块hash
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	//a.创建创世区块
	block := NewBlock(data, prevHash)
	//b.添加到区块链数组中
	bc.blocks = append(bc.blocks, block)

}
