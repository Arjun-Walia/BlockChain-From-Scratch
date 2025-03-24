package network

import (
	"testing"
)

func TestStartServer(t *testing.T) {
	go StartServer("3000")
	// Add more test logic here (e.g., connect to the server and verify behavior)
}

func TestConnectToPeer(t *testing.T) {
	node := Node{Address: "localhost:3000", Peers: []string{}}
	node.ConnectToPeer("localhost:3001")

	if len(node.Peers) != 1 {
		t.Error("Peer was not added to the node")
	}
}
