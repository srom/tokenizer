package tokenizer

import (
	"fmt"
	"testing"
)

func ExampleTreebankWordTokenizer() {
	tokenizer := NewTreebankWordTokenizer()

	tokens := tokenizer.Tokenize(
		// Example string from http://nlp.stanford.edu/software/tokenizer.shtml
		`"Oh, no," she's saying, "our $400 blender can't handle something this hard!"`)

	fmt.Println(tokens)
	// Output: [`` Oh , no , '' she 's saying , `` our $ 400 blender ca n't handle something this hard ! '']
}

type testCase struct {
	Input          string
	ExpectedOutput []string
}

func TestTreebankWordTokenizer(t *testing.T) {
	cases := []testCase{
		testCase{
			"They'll save and invest more.",
			[]string{"They", "'ll", "save", "and", "invest", "more", "."},
		},
		testCase{
			"Good muffins cost $3.88\nin New York.  Please buy me\ntwo of them.\nThanks.",
			[]string{
				"Good", "muffins", "cost", "$", "3.88", "in", "New", "York.",
				"Please", "buy", "me", "two", "of", "them.", "Thanks", ".",
			},
		},
		testCase{
			"Hello, there!\n I'm Romain.",
			[]string{"Hello", ",", "there", "!", "I", "'m", "Romain", "."},
		},
	}
	for _, test_case := range cases {
		tokenizer := NewTreebankWordTokenizer()
		words := tokenizer.Tokenize(test_case.Input)
		for i, word := range words {
			if word != test_case.ExpectedOutput[i] {
				t.Errorf("\nGot:      %s\nExpected: %s", word, test_case.ExpectedOutput[i])
			}
		}
	}
}
