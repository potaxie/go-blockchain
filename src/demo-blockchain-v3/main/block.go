package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

//1.定义结构
type Block struct {
	//1.版本号
	Version uint64
	//2.前区块hash
	PrevHash []byte
	//3.Merkle根
	MerkleRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机数，也就是挖矿要找的数据
	Nonce uint64
	//a.当前区块hash,正常比特币区块中没有当前的hash
	Hash []byte
	//b.数据
	Data []byte
}

//实现一个辅助函数,将uint64转成[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{}, //先填空，后面在计算
		Data:       []byte(data),
	}
	//block.SetHash()

	//创建一个pow对象
	pow := NewProofOfWork(&block)

	//查找随机数，不停的进行hash计算
	hash, nonce := pow.Run()

	//根据挖矿结果对区块数据进行更新
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

//序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer

	//使用gob进行序列化，得到字节流
	//定义一个编码器
	//使用编码器进行编码
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic("编码出错，小明不知去向")
	}
	//fmt.Printf("编码%v\n", buffer.Bytes())
	return buffer.Bytes()
}

//反序列化
func Deserialize(data []byte) Block {

	decoder := gob.NewDecoder(bytes.NewReader(data))
	//fmt.Printf("解码 \n")

	var block Block
	//使用解码器进行解码
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错")
	}
	return block
}

//3.生成hash
func (block *Block) SetHash() {
	//var blockInfo []byte
	//拼装数据
	//blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	//blockInfo = append(blockInfo, block.PrevHash...)
	//blockInfo = append(blockInfo, block.MerkleRoot...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	//blockInfo = append(blockInfo, block.Data...)

	//优化，将二维的切片数组连接起来，返回一个一唯的切片
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkleRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}

	blockInfo := bytes.Join(tmp, []byte{})
	//sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
