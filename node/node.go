package node

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/trotelalexandre/proto/blockchain"
)

type NodeConfig struct {
	Address string
	Port    int
	Peers   []string
}

type Node struct {
    Config    NodeConfig
    Blockchain *blockchain.Blockchain
}

func (n *Node) StartNode(node Node, blockchain *blockchain.Blockchain) {
    http.HandleFunc("/blockchain", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(blockchain)
    })

    log.Println("Node started at", node.Config.Address)
    log.Fatal(http.ListenAndServe(node.Config.Address, nil))
}

func (n *Node) ConnectToPeers(node Node) {
    for _, peer := range node.Config.Peers {
        go n.SyncWithPeer(peer)
    }
}

func (n *Node) SyncWithPeer(peer string) {
    resp, err := http.Get(peer + "/blockchain")
    if err != nil {
        log.Printf("Error connecting to peer %s: %v", peer, err)
        return
    }
    defer resp.Body.Close()

    var peerBlockchain blockchain.Blockchain
    err = json.NewDecoder(resp.Body).Decode(&peerBlockchain)
    if err != nil {
        log.Printf("Error decoding blockchain from peer %s: %v", peer, err)
        return
    }

    n.SyncBlockchainIfLonger(peerBlockchain)
}

func (n *Node) SyncBlockchainIfLonger(peerBlockchain blockchain.Blockchain) {
    if len(peerBlockchain.Blocks) > len(n.Blockchain.Blocks) {
        log.Println("Found a longer blockchain, syncing...")
        *n.Blockchain = peerBlockchain
    }
}

func (n *Node) SyncBlockchain(node Node) {
    ticker := time.NewTicker(30 * time.Second)
    for range ticker.C {
        n.ConnectToPeers(node)
    }
}

func (n *Node) BroadcastTransaction(transaction blockchain.Transaction) {
    for _, peer := range n.Config.Peers {
        go func(peer string) {
            data, err := json.Marshal(transaction)
            if err != nil {
                log.Printf("Error marshalling transaction: %v", err)
                return
            }
            _, err = http.Post(peer+"/transaction", "application/json", bytes.NewBuffer(data))
            if err != nil {
                log.Printf("Error broadcasting transaction to peer %s: %v", peer, err)
            }
        }(peer)
    }
}