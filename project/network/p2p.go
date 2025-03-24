package network

import (
	"fmt"
	"net"
)

type Node struct {
	Address string
	Peers   []string
}

func StartServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Node listening on port", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Handle peer messages here
	fmt.Println("New connection from:", conn.RemoteAddr())
}

func (n *Node) ConnectToPeer(peerAddress string) {
	conn, err := net.Dial("tcp", peerAddress)
	if err != nil {
		fmt.Println("Error connecting to peer:", err)
		return
	}
	defer conn.Close()

	// Add the peer to the list of peers
	n.Peers = append(n.Peers, peerAddress)
	fmt.Println("Connected to peer:", peerAddress)
}
