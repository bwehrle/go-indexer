package indexing

import (
	"reflect"
	"testing"

	"github.com/bwehrle/indexer/internal/tokens"
)

func Test_import(t *testing.T) {

	tests := []struct {
		name     string
		text     []string
		word     string
		expected TokenLocationIndex[string]
	}{
		{
			name: "Repeated words",
			text: []string{"Hello, World!", "Hello, World!"},
			word: "world",
			expected: TokenLocationIndex[string]{
				"test": []Location{{0, 6}, {1, 6}},
			},
		},
	}

	for _, tt := range tests {
		i := NewMemIndexer()
		for c, line := range tt.text {
			i.process(line, c, "test", tokens.NewTextTokenizer())
		}
		fileIndex := i.find("world")
		if eq := reflect.DeepEqual(fileIndex, tt.expected); !eq {
			t.Errorf("Got %+v, wanted %+v", fileIndex, tt.expected)
		}
	}
}
