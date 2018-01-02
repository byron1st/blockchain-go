package util

import "github.com/byron1st/blockchain-go/blockchain"

type FullChainResponse struct {
	Chain []*blockchain.Block
	Length int
}

type MineResponse struct {
	message string
	block *blockchain.Block
}