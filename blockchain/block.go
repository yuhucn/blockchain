package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"
)

//0. 定义结构
type Block struct {
	//1. 版本号
	Version uint64
	//2. 前区块哈希
	PrevHash []byte
	//3. Merkle根
	MerkleRoot []byte
	//4. 时间戳
	TimeStamp uint64
	//5. 难度值
	Difficulty uint64
	//6. 随机数
	Nonce uint64

	//a. 当前区块哈希
	Hash []byte
	//b. 数据
	Data []byte
}

//1. 补充区块字段
//2. 更新计算哈希函数
//3. 优化代码

//实现一个辅助函数，功能是将unit64转成[]byte

func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}
	//block.SetHash()
	//创建一个pow对象
	pow := NewProofWork(&block)
	//查找随机数，不停地进行哈希运算
	hash, nonce := pow.Run()
	//根据挖矿结果对区块数据进行更新（补充）
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

//3. 生成哈希
/*
func (block *Block) SetHash() {
	//var blockInfo []byte
	//1. 拼装数据
	/*
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkleRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)
*/
/*
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkleRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}

	//将二维的切片数组链接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte{})

	//2. sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
*/
