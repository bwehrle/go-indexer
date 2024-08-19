package indexing

import (
	"github.com/bwehrle/indexer/internal/tokens"
)

type memIndexer struct {
	index TokenIndex[string]
}

func NewMemIndexer() *memIndexer {
	return &memIndexer{index: make(TokenIndex[string])}
}

func (m *memIndexer) process(line string, lineNo int, source string, tokenizer tokens.Tokenizer) {
	col := 0
	tokens, _ := tokenizer.Tokenize(line)
	for _, token := range tokens {
		if srcMap, ok := m.index[token]; ok {
			if _, ok := srcMap[source]; ok {
				srcMap[source] = append(m.index[token][source], Location{line: lineNo, col: col})
			} else {
				m.index[token][source] = []Location{
					{line: lineNo, col: col},
				}
			}
		} else {
			m.index[token] = map[string][]Location{
				source: {
					{line: lineNo, col: col},
				},
			}
		}
		col += len(token) + 1
	}
}

func (m *memIndexer) find(word string) TokenLocationIndex[string] {
	return m.index[word]
}
