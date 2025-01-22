package node

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (n *Node) StartNode() {
    http.HandleFunc("/protochain", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(n.Blockchain)
    })

    log.Println("Node started at", n.Config.Address, "on port", n.Config.Port)
    err := http.ListenAndServe(fmt.Sprintf("%s:%d", n.Config.Address, n.Config.Port), nil)
	if err != nil {
		log.Printf("Error starting server: %v", err)
	}
}

func (n *Node) ConnectToPeers() {
    for _, peer := range n.Config.Peers {
        go n.SyncWithPeer(peer)
    }
}

func (n *Node) SyncWithPeer(peer string) {
    resp, err := http.Get(peer + "/protochain")
    if err != nil {
        log.Printf("Error connecting to peer %s: %v", peer, err)
        return
    }
    defer resp.Body.Close()

    var peerBlockchain blockchain.Blockchain
    err = json.NewDecoder(resp.Body).Decode(&peerBlockchain)
    if err != nil {
        log.Printf("Error decoding Protochain from peer %s: %v", peer, err)
        return
    }

    n.SyncBlockchainIfLonger(peerBlockchain)
}

func (n *Node) SyncBlockchainIfLonger(peerBlockchain blockchain.Blockchain) {
    if len(peerBlockchain.Blocks) > len(n.Blockchain.Blocks) {
        log.Println("Found a longer Protochain, syncing...")
        *n.Blockchain = peerBlockchain
    } else {
        log.Println("Protochain is up to date")
    }
}

func (n *Node) SyncBlockchain() {
    n.ConnectToPeers()
    ticker := time.NewTicker(30 * time.Second)
    for range ticker.C {
        n.ConnectToPeers()
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