package blockchain

import (
	"net/url"
	"time"
	"fmt"
	"encoding/json"
	"github.com/byron1st/blockchain-go/util"
)

type Transaction struct {
	sender    string
	recipient string
	amount    int
}

type Block struct {
	index        int
	timestamp    int64
	transactions []*Transaction
	proof        int
	previousHash string
}

type Blockchain struct {
	currentTransactions []*Transaction
	chain               []*Block
	nodes               map[string]bool
}

func MakeBlockchain() *Blockchain {
	blockchain := Blockchain{
		[]*Transaction{},
		[]*Block{},
		make(map[string]bool),
	}

	blockchain.CreateNewBlock("1", 100)
	return &blockchain
}

func (blockchain *Blockchain) GetLastBlock() *Block {
	return blockchain.chain[len(blockchain.chain)-1]
}

func (blockchain *Blockchain) GetLastBlockProof() int {
	return blockchain.GetLastBlock().proof
}

func (blockchain *Blockchain) RegisterNode(urlAddress string) {
	parsedURL, _ := url.Parse(urlAddress)
	blockchain.nodes[parsedURL.Host] = true
}

func (blockchain *Blockchain) CreateNewTransaction(sender string, recipient string, amount int) int {
	transaction := &Transaction{sender, recipient, amount}
	blockchain.currentTransactions = append(blockchain.currentTransactions, transaction)

	return blockchain.GetLastBlock().index + 1
}

func (blockchain *Blockchain) CreateNewBlock(previousHash string, proof int) *Block {
	block := &Block{
		len(blockchain.chain) + 1,
		time.Now().Unix(),
		blockchain.currentTransactions,
		proof,
		previousHash, // TODO: previousHash 값이 없을 경우, 앞 블록을 해싱.
	}

	blockchain.currentTransactions = []*Transaction{}
	blockchain.chain = append(blockchain.chain, block)
	return block
}

func (blockchain *Blockchain) ResolveConflicts() (bool, error) {
	var newChain []*Block
	maxLength := len(blockchain.chain)
	var isReplaced bool

	for nodeUrl, _ := range blockchain.nodes {
		responseObj := &util.FullChainResponse{}
		error := util.GetChainFromRemote(nodeUrl, responseObj)
		if error != nil {
			return false, error
		}

		if responseObj.Length > maxLength && ValidChain(responseObj.Chain) {
			maxLength = responseObj.Length
			newChain = responseObj.Chain
		}
	}

	if newChain != nil {
		blockchain.chain = newChain
		isReplaced = true
	} else {
		isReplaced = false
	}

	return isReplaced, nil
}

func ProofOfWork(lastProof int) int {
	proof := 0
	for !ValidProof(lastProof, proof) {
		proof += 1
	}
	return proof
}

func ValidProof(lastBlockProof int, currentBlockProof int) bool {
	guess := fmt.Sprintf("%d%d", lastBlockProof, currentBlockProof)
	return util.Hash(guess) == "0000"
}

func ValidChain(chain []*Block) bool {
	for index, lastBlock := range chain {
		if index + 1 == len(chain) {
			break
		}

		currentBlock := chain[index + 1]
		resultThisTime := currentBlock.previousHash == util.Hash(BlockStringify(lastBlock)) && ValidProof(lastBlock.proof, currentBlock.proof)
		if resultThisTime == false {
			return false
		}
	}

	return true
}

func BlockStringify(block *Block) string {
	out, _ := json.Marshal(block)

	return string(out)
}
