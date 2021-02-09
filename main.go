package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain representation as slice of blocks
type BlockChain struct {
	blocks []*Block
}

// Block represents a block in blockchain
type Block struct {
	data     []byte
	hash     []byte
	prevHash []byte
}

func (b *Block) generateHash() {
	info := bytes.Join([][]byte{b.data, b.prevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.hash = hash[:]
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte(data), []byte{}, prevHash}
	block.generateHash()
	return block
}

func (bc *BlockChain) addBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := createBlock(data, prevBlock.hash)
	bc.blocks = append(bc.blocks, block)
}

func genesis() *Block {
	return createBlock("Genesis Block", []byte{})
}

func initBlockChain() *BlockChain {
	return &BlockChain{[]*Block{genesis()}}
}

func main() {
	chain := initBlockChain()

	chain.addBlock("First Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.addBlock("Fourth Block")
	chain.addBlock("Fifth Block")

	for _, v := range chain.blocks {
		fmt.Printf("Data : %s\n", v.data)
		fmt.Printf("Hash : %x\n", v.hash)
		fmt.Printf("Prev Hash : %x\n", v.prevHash)
	}
}
