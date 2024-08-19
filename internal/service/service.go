package service

import (
	"net/http"
	"io/ioutil"
	
	"github.com/bwehrle/indexer/internal/indexing"
	"github.com/bwehrle/indexer/internal/service"
	"github.com/bwehrle/indexer/internal/tokens"
)

type Text struct {
	Source string `json:"source"`
	Text   string `json:"text"`
}

type ServiceProcessor[T indexing.Token] struct {
	indexer   indexing.Index[T]
	tokenizer tokens.Tokenizer
}

func NewServiceProcessors[T indexing.Token](
	indexer indexing.Index[T],
	tokenizer tokens.Tokenizer) *ServiceProcessor[T] {

	return &service.ServiceProcessor[T]{indexer, tokenizer}
}

func (s *ServiceProcessor[T]) ProcessStream(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	t := &Text{}
	err := json.Unmarshal(b, &t)
	s.indexer.Process(t.Text, 0, t.Source, s.tokenizer)
}
