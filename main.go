package main

import (
	"log"
	"net/http"

	"github.com/bwehrle/indexer/internal/indexing"
	"github.com/bwehrle/indexer/internal/service"
	"github.com/bwehrle/indexer/internal/tokens"
)

func main() {
	indexer := indexing.NewMemIndexer()
	tokenizer := tokens.NewTextTokenizer()
	s := service.NewServiceProcessors(indexer, tokenizer)
	http.HandleFunc("POST /index", func (w http.ResponseWriter, r *http.Request) {
		s.ProcessStream(w, r)
	})
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}