package api

import (
	"blockchain/block"
	"blockchain/blockchain"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func RunAPI(bc *blockchain.Blockchain, port string) {
	r := mux.NewRouter()

	r.HandleFunc("/mine", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request to mine a new block")

		newBlock := &block.Block{
			Index:        len(bc.Blocks),
			Timestamp:    time.Now(),
			Transactions: bc.PendingTrans,
			PrevHash:     bc.Blocks[len(bc.Blocks)-1].Hash,
		}
		bc.AddBlock(newBlock)
		bc.PendingTrans = nil

		fmt.Println("New block mined:", newBlock)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "New block mined",
			"block":   newBlock,
		})
	}).Methods("POST")

	http.ListenAndServe(":"+port, r)
}
