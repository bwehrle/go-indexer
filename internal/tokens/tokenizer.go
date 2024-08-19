package tokens

import (
	"regexp"
	"strings"
)

type Tokenizer interface {
	Tokenize(text string) ([]string, error)
}


type TextTokenizer struct{
	wordRegex *regexp.Regexp
}

func (t*TextTokenizer) Tokenize(test string) ([]string, error) {
	words := t.wordRegex.FindAllString(test, -1)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words, nil
}

func NewTextTokenizer() Tokenizer{
	return &TextTokenizer{
		wordRegex: regexp.MustCompile(`\w+`),
	}
}

