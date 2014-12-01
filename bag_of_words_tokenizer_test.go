package tokenizer

import (
	"fmt"
)

func ExampleBagOfWordsTokenizer() {
	tokenizer := NewBagOfWordsTokenizer("stop_words.txt")

	tokens := tokenizer.Tokenize(
		// Example string from http://nlp.stanford.edu/software/tokenizer.shtml
		`"Oh, no," she's saying, "our $400 blender can't handle something this hard!"`)

	fmt.Println(tokens)
	// Output: [`` Oh , no , '' she 's saying , `` our $ 400 blender ca n't handle something this hard ! '']
}
