package blockchain

import "net/url"

type transaction struct {
	sender    string
	recipient string
	amount    int
}

type block struct {
	index        int
	timestamp    float64
	transactions []transaction
	proof        int
	previousHash string
}

type blockchain struct {
	currentTransactions []transaction
	chain               []block
	nodes               map[string]bool
}

func makeBlockchain() *blockchain {
	newBlockchain := new(blockchain)
	newBlockchain.currentTransactions = []transaction{}
	newBlockchain.chain = []block{}
	newBlockchain.nodes = make(map[string]bool)
	return newBlockchain
}

func (blockchain *blockchain) getLastBlock() *block {
	return &blockchain.chain[len(blockchain.chain)-1]
}

func (blockchain *blockchain) registerNode(urlAddress string) {
	parsedURL, _ := url.Parse(urlAddress)
	blockchain.nodes[parsedURL.Host] = true
}

func (blockchain *blockchain) createNewTransaction(sender string, recipient string, amount int) int {
	transaction := transaction{sender, recipient, amount}
	blockchain.currentTransactions = append(blockchain.currentTransactions, transaction)

	return blockchain.getLastBlock().index + 1
}
