package indexing

import (
	"bufio"
	"log"
	"os"

	"github.com/bwehrle/indexer/internal/tokens"
)

type Token interface {
	~string
}

type Location struct {
	line, col int
}

type TokenLocationIndex[T Token] map[T][]Location

type TokenIndex[T Token] map[T]TokenLocationIndex[T]

type Index[T Token] interface {
	Process(line string, lineNo int, source string, tokenizer tokens.Tokenizer)
	Find(T) TokenLocationIndex[T]
}

type FileIndexProcessor[T Token] struct {
	indexer   Index[T]
	tokenizer tokens.Tokenizer
}

func NewFileProcessor[T Token](indexer Index[T], tokenizer tokens.Tokenizer) *FileIndexProcessor[T] {
	return &FileIndexProcessor[T]{indexer, tokenizer}
}

func (f *FileIndexProcessor[T]) ProcessFile(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		var lineNo = 0
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		f.indexer.Process(line, lineNo, path, f.tokenizer)
		lineNo++
	}
}
