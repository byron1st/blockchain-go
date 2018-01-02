package server

import (
	"net/http"
	"github.com/byron1st/blockchain-go/blockchain"
	"github.com/byron1st/blockchain-go/util"
	"fmt"
	"encoding/json"
)

type App struct {
	blockchain *blockchain.Blockchain
	nodeId string
	Server *http.Server
}

func CreateServer(nodeId string, port int) *App {
	return &App{
		&blockchain.Blockchain{},
		nodeId,
		&http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
	}
}

func (app *App) mine(writer http.ResponseWriter, request *http.Request) {
	proof := blockchain.ProofOfWork(app.blockchain.GetLastBlockProof())
	app.blockchain.CreateNewTransaction("0", app.nodeId, 1)

	previousHash := util.Hash(blockchain.BlockStringify(app.blockchain.GetLastBlock()))
	block := app.blockchain.CreateNewBlock(previousHash, proof)

	responseObj := &util.MineResponse{
		"New Block Forged",
		block,
	}

	jsonData, error := json.Marshal(responseObj)
	if error != nil {
		panic(error)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonData)
}