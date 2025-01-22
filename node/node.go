package node

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/trotelalexandre/proto/blockchain"
)

type Node struct {
    Address string
    Peers   []string
    Blockchain *blockchain.Blockchain
}

func (n *Node) StartNode(node Node, blockchain *blockchain.Blockchain) {
    http.HandleFunc("/blockchain", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(blockchain)
    })

    log.Println("Node started at", node.Address)
    log.Fatal(http.ListenAndServe(node.Address, nil))
}

func (n *Node) ConnectToPeers(node Node) {
    for _, peer := range node.Peers {
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