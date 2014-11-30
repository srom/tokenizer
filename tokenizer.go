package tokenizer

type Tokenizer interface {
	Tokenize(text string) []string
}
