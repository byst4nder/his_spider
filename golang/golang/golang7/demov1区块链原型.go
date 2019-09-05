package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)
/*
区块结构
*/
type Block struct {
	Timestamp     int64  //时间戳
	Date          []byte // 当前区块 存放信息
	PrevBlockHash []byte //上一个区块的加密Hash
	Hash          []byte // 当前的区块的Hash
}

/*t
给Block 绑定一个生成区块的方法
*/

func (this *Block) Sethash() {
	//将本区块的时间戳和上一个区块进行拼接
	timestamp := []byte(strconv.FormatInt(this.Timestamp, 10)) // 时间转化成二进制
	// 将三个二进制进行拼接
	hesders := bytes.Join([][]byte{this.PrevBlockHash, this.Date, timestamp}, []byte{})
	// 拼接之后的hesders 进行SHA256加密
	hash := sha256.Sum256(hesders)
	// 依照前区块和新区块  添加到区块链blocks中
	this.Hash = hash[:]
}

/*
新建一个区块链的 API
data 当前区块所保存的数据
prevBlockHash 表示当前区块的前区
*/
func NewBlock(data string, prevBlockHash []byte) *Block {
	// 生成一个区块
	block := Block{}
	// 给当前的区块创建时间
	block.Timestamp = time.Now().Unix()
	//给当前的区块赋值
	block.Date = []byte(data)
	// 绑定前驱hash
	block.PrevBlockHash = prevBlockHash
	// 给当前区块进行hash加密
	block.Sethash()
	// 将已经赋值好的区块  返回
	return &block
}

/*
定义一个区块链的结构
*/
type BlockChain struct {
	Blocks []*Block
}

// 将区块添加到一个区块链中
func (this *BlockChain) AddBlock(data string) {
	//得到 新添加的区块的 前驱快
	prevBlock := this.Blocks[len(this.Blocks)-1]
	//根据data 创建一个新的区块
	newBlock := NewBlock(data, prevBlock.Hash)
	//依照前区块和新区块 添加到区块链blocks中
	this.Blocks = append(this.Blocks, newBlock)
}

// 区块链 = 创世块> 区块>区块
// 新建一个创世块
func NewGenesisBlock() *Block {
	genesisBlock := Block{}
	genesisBlock.Date =[]byte("Genesis block")
	genesisBlock.PrevBlockHash = []byte{}
	return &genesisBlock
}

// 新建一个区块链11
func NewBlockChain() *BlockChain {
	//blockChain :=BlockChain{}
	//blockChain.Blocks = append(blockChain.Blocks)
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func main() {
	//fmt.Println("测试这个程序1")
	//block := NewBlock("哈哈", []byte{})
	//fmt.Printf("block.hash = %x", block.Hash)
	// 创建一个区块链bc
	bc := NewBlockChain()
	// 输入指令
	for {
		fmt.Println("按1添加一条信息数据到区块链中")
		fmt.Println("按2遍历当前的区块的所有信息")
		fmt.Println("按3表示退出")
		var cmd string
		fmt.Scanf("%s", &cmd)
		cmd_str :=string(cmd)
		switch cmd_str {
		case "1":
			input := make([]byte, 1024)
			fmt.Println("请输入区块链的行为数据：")
			fmt.Scanf("%s", &input)
			//fmt.Println("已经输入了区块链行为数据")
			fmt.Println(input)
			bc.AddBlock(string(input))
		case "2":
			// 遍历整个区块链，输出整个区块链的值
			for i, block := range bc.Blocks {
				fmt.Println("===========")
				fmt.Println("第", i, "个区块的信息：")
				fmt.Printf("PrevHash:%x\n", block.PrevBlockHash)
				fmt.Printf("Data:%s\n", block.Date)
				fmt.Printf("Hash:%x\n", block.Hash)
				fmt.Println("=============")
			}
		default:
			fmt.Println("已经退出了")
			return
		}
	}
}
