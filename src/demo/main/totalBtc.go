package main

import "fmt"

/**
20200531
比特币总量
*/

func main() {

	fmt.Println("hello: ")

	//1.首先21万个块减半
	//2.最初奖励50btc
	//3.用一个循环来判断，累加

	total := 0.0
	blockInterval := 21.0 //单位万
	currentReward := 50.0

	for currentReward > 0 {

		//每一个区间的总量
		amount1 := blockInterval * currentReward

		currentReward *= 0.5 //除的效率低，用乘法替换

		total += amount1
	}

	fmt.Println("btc summary amout", total, "w")

}
