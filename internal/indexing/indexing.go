package indexing

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type pos struct {
	line, col int
}

type FileIndex map[string][]pos

type indexing interface {
	process(line string, lineNo int, source string)
	find(string) FileIndex
}

var indexer indexing

func ProcessFile(path string) {
	processFile(path, indexer)
}

func processFile(path string, indexer indexing) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		var lineNo = 0
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		indexer.process(line, lineNo, path)
		lineNo++
	}
}

type filePosIndex map[string]FileIndex

type memIndexer struct {
	index filePosIndex
}

func NewMemIndexer() *memIndexer {
	return &memIndexer{index: make(filePosIndex)}
}

func (m *memIndexer) process(line string, lineNo int, source string) {
	col := 0
	for _, token := range strings.Split(line, " ") {
		if strings.ContainsAny(token, "&-+=%*^~|<>") {
			continue
		}

		entry := strings.ToLower(strings.Trim(token, "\"':;.,!?-()\n"))

		if srcMap, ok := m.index[entry]; ok {
			if _, ok := srcMap[source]; ok {
				srcMap[source] = append(m.index[entry][source], pos{line: lineNo, col: col})
			} else {
				m.index[entry][source] = []pos{
					{line: lineNo, col: col},
				}
			}
		} else {
				m.index[entry] = map[string][]pos {
					source: {
						{line: lineNo, col: col},
					},
				}
		}
		col += len(token) + 1
	}
}

func (m *memIndexer) find(word string) FileIndex {
	return m.index[word]
}
