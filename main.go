package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Hash      string
	Transactions []Transaction
	PrevHash  string
}

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

type Blockchain struct {
	Blocks []*Block
	mux    sync.Mutex
}

func CalculateHash(block Block) string {
	hashInput := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.PrevHash, block.Transactions)
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}

func CreateGenesisBlock() *Block {
	genesisBlock := &Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Transactions: []Transaction{},
		PrevHash:  "0",
	}
	genesisBlock.Hash = CalculateHash(*genesisBlock)
	return genesisBlock
}

func (blockchain *Blockchain) AddBlock(transactions []Transaction) {
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := &Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Transactions: transactions,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = CalculateHash(*newBlock)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func (blockchain *Blockchain) GetLastBlock() *Block {
	return blockchain.Blocks[len(blockchain.Blocks)-1]
}

func (blockchain *Blockchain) GetBlockchain(w http.ResponseWriter, r *http.Request) {
	blockchain.mux.Lock()
	defer blockchain.mux.Unlock()
	json.NewEncoder(w).Encode(blockchain.Blocks)
}

func (blockchain *Blockchain) AddTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions []Transaction
	if err := json.NewDecoder(r.Body).Decode(&transactions); err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}
	blockchain.AddBlock(transactions)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Block added successfully")
}

func main() {
	genesisBlock := CreateGenesisBlock()
	blockchain := &Blockchain{Blocks: []*Block{genesisBlock}}

	http.HandleFunc("/blockchain", blockchain.GetBlockchain)
	http.HandleFunc("/add", blockchain.AddTransaction)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}