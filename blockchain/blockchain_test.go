package blockchain

import (
	"testing"
)

func TestBlockchain_RegisterNode(t *testing.T) {
	bc := MakeBlockchain()
	urlAddress := "http://localhost:5000"
	url := "localhost:5000"
	bc.RegisterNode(urlAddress)

	if !bc.nodes[url] {
		t.Error("Failed to register node.")
	}
}

func TestBlockchain_CreateNewTransaction(t *testing.T) {
	bc := MakeBlockchain()
	bc.CreateNewTransaction("sender", "recipient", 30)

	if len(bc.currentTransactions) != 1 {
		t.Errorf("Current Transaction is not 1, but %d", len(bc.currentTransactions))
	}

	if bc.currentTransactions[0].sender != "sender" {
		t.Errorf("Sender is not \"sender\", but %s", bc.currentTransactions[0].sender)
	}
}

func TestBlockchain_CreateNewBlock(t *testing.T) {
	bc := MakeBlockchain()
	bc.CreateNewBlock("2", 200)

	if len(bc.chain) != 2 {
		t.Errorf("Chain length is not 2, but %d", len(bc.chain))
	}

	if &bc.chain[1].transactions == &bc.currentTransactions {
		t.Error("bc.chain[1].transactions == bc.currentTransactions.")
	}

	if len(bc.currentTransactions) != 0 {
		t.Errorf("Current transactions are not initialized.")
	}
}