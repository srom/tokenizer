package tokenizer

import (
	"fmt"
	"testing"
)

func ExampleBagOfWordsTokenizer() {
	tokenizer := NewBagOfWordsTokenizer("fixtures/stop_words.txt")

	tokens := tokenizer.Tokenize(
		// Example string from http://nlp.stanford.edu/software/tokenizer.shtml
		`"Oh, no," she's saying, "our $400 blender can't handle something this hard!"`)

	fmt.Println(tokens)
	// Output: [oh saying blender handle something hard]
}

func TestBagOfWordsTokenizer(t *testing.T) {
	cases := []testCase{
		testCase{
			"They'll save and invest more.",
			[]string{"save", "invest", "more"},
		},
		testCase{
			"Good muffins cost $3.88\nin New York.  Please buy me\ntwo of them.\nThanks.",
			[]string{
				"good", "muffins", "cost", "new", "york",
				"please", "buy", "two", "thanks",
			},
		},
		testCase{
			"Hello, there!\n I'm Romain.",
			[]string{"hello", "romain"},
		},
	}
	for _, test_case := range cases {
		tokenizer := NewBagOfWordsTokenizer("fixtures/stop_words.txt")
		words := tokenizer.Tokenize(test_case.Input)
		for i, word := range words {
			if word != test_case.ExpectedOutput[i] {
				t.Errorf("\nGot:      %s\nExpected: %s", word, test_case.ExpectedOutput[i])
			}
		}
	}
}
