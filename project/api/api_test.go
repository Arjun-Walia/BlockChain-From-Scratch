package api

import (
	"blockchain/block"
	"blockchain/blockchain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBlocks(t *testing.T) {
	bc := blockchain.Blockchain{
		Blocks:     []*block.Block{blockchain.CreateGenesisBlock()},
		Difficulty: 4,
	}

	req, _ := http.NewRequest("GET", "/blocks", nil)
	rr := httptest.NewRecorder()

	RunAPI(&bc, "8080")
	http.DefaultServeMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}
}
