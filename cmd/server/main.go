package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wijayasf/docnote/notary"
)

func main() {
	fmt.Println("docnote server starting on http://localhost:8080")

	http.HandleFunc("/notarize", func(w http.ResponseWriter, r *http.Request) {
		var req struct{ Data string }
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		block := notary.Notarize(req.Data)

		log.Printf("New block: Index=%d, Hash=%s\n", block.Index, block.Hash)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(block)
	})

	http.HandleFunc("/chain", func(w http.ResponseWriter, r *http.Request) {
		chain := notary.GetChain()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chain)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
