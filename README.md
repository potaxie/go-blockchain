# go-blockchain
use golang to learn blockchain 



## demo

    简易版区块链实现
    
* 随机数和难度值随便填写
* 区块的哈希值是无规则的

### totalBtc

    实现计算btc总量
    
 
## demo-blockchain-v2

* Pow介绍
    * 定义一个工作量证明的结构proofofwork（block、目标值）
* 提供创建Pow的函数
* 提供不断计算hash的函数
    * run()
* 提供校验函数
    * isVaild()
    
### block

* 定义区块结构


### blockchain

* 定义区块链结构
* 添加区块方法

### ProofOfWork

* 定义一个工作量证明的结构proofofwork


## demo-blockchain-v3

* blot数据库
    * []byte -> []byte
* 使用数据库改写区块链结构
* 调整代码
* 添加命令
    * 闯进区块
    * 添加区块
    * 打印区块

### block

* 实现序列化与反序列化
* 实现hash

### blockchain

* 利用blot定义区块链