package main

import (
	"blockchain/api"
	"blockchain/block"
	"blockchain/blockchain"
	"blockchain/network"
	"flag"
	"fmt"
)

func main() {
	// Parse command-line arguments
	p2pPort := flag.String("p2p-port", "3000", "Port for P2P communication")
	apiPort := flag.String("api-port", "8080", "Port for REST API")
	flag.Parse()

	// Initialize the blockchain
	bc := blockchain.Blockchain{
		Blocks:     []*block.Block{blockchain.CreateGenesisBlock()},
		Difficulty: 4,
	}

	// Initialize the P2P node
	node := network.Node{
		Address: fmt.Sprintf("localhost:%s", *p2pPort),
		Peers:   []string{},
	}

	// Start the P2P server
	go network.StartServer(*p2pPort)

	// Connect to a peer (example)
	if *p2pPort == "3000" {
		node.ConnectToPeer("localhost:3001")
	} else if *p2pPort == "3001" {
		node.ConnectToPeer("localhost:3000")
	}

	// Start the REST API
	api.RunAPI(&bc, *apiPort)
}
