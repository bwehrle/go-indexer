package indexing

import (
	"testing"
)

func Test_import(t *testing.T) {
	i := NewMemIndexer()
	i.process("Hello, World!", 0, "test")
	i.process("Hello, World!", 1, "test")

	fileIndex := i.find("world")
	
	if _, ok := fileIndex["test"]; !ok {
		t.Errorf("Expected fileIndex contains 'test' source for 'world'")
	}
	if len(fileIndex["test"]) != 2 {
		t.Errorf("Expected fileIndex contains 2 pos")
	}
	for c, pos := range fileIndex["test"] {
		if pos.line != c || pos.col != 7 {
			t.Errorf("Got pos (%d, %d) wanted (%d, 7) for line %d", pos.line, pos.col, c, c)
		}
	}
}