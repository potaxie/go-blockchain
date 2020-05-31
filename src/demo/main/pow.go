package main

import (
	"crypto/sha256"
	"fmt"
)

/*
暴力计算hash
*/

func main() {

	//交易数据
	data := "helloworld"

	for i := 0; i < 10000; i++ {
		hash := sha256.Sum256([]byte(data + string(i)))
		fmt.Printf("hash:%x\n", hash[:])
	}

}
