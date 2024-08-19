package tokens

import (
	"testing"

	"github.com/onsi/gomega"
)

func Test_tokenize_text(t *testing.T) {

	tests := []struct {
		name     string
		text     string
		expected []string
		err      error
	}{
		{
			"Text white-space separated",
			"foo bar biz",
			[]string{"foo", "bar", "biz"},
			nil,
		},
		{
			"Test is lower cases w/o symbols",
			"Foo bar, biz!",
			[]string{"foo", "bar", "biz"},
			nil,
		},
	}

	var tokenizer Tokenizer = NewTextTokenizer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewGomegaWithT(t)
			out, err := tokenizer.Tokenize(tt.text)
			if tt.err != nil {
				g.Expect(err).To(gomega.MatchError(tt.err))
			} else {
				g.Expect(err).To(gomega.BeNil())
			}
			g.Expect(out).To(gomega.Equal(tt.expected))
		})
	}
}