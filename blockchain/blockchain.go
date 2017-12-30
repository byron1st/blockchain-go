package blockchain

import (
	"net/url"
	"time"
)

type transaction struct {
	sender    string
	recipient string
	amount    int
}

type block struct {
	index        int
	timestamp    int64
	transactions []transaction
	proof        int
	previousHash string
}

type blockchain struct {
	currentTransactions []transaction
	chain               []block
	nodes               map[string]bool
}

func MakeBlockchain() *blockchain {
	blockchain := blockchain{
		[]transaction{},
		[]block{},
		make(map[string]bool),
	}

	blockchain.CreateNewBlock("1", 100)
	return &blockchain
}

func (blockchain *blockchain) getLastBlock() *block {
	return &blockchain.chain[len(blockchain.chain)-1]
}

func (blockchain *blockchain) RegisterNode(urlAddress string) {
	parsedURL, _ := url.Parse(urlAddress)
	blockchain.nodes[parsedURL.Host] = true
}

func (blockchain *blockchain) CreateNewTransaction(sender string, recipient string, amount int) int {
	transaction := transaction{sender, recipient, amount}
	blockchain.currentTransactions = append(blockchain.currentTransactions, transaction)

	return blockchain.getLastBlock().index + 1
}

func (blockchain *blockchain) CreateNewBlock(previousHash string, proof int) *block {
	block := block{
		len(blockchain.chain) + 1,
		time.Now().Unix(),
		blockchain.currentTransactions,
		proof,
		previousHash, // TODO: previousHash 값이 없을 경우, 앞 블록을 해싱.
	}

	blockchain.currentTransactions = []transaction{}
	blockchain.chain = append(blockchain.chain, block)
	return &block
}
