package tokenizer

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

var WORD_WITHOUT_LAST_PUNCT = regexp.MustCompile(`^(\w+)[^\w]?$`)

type BagOfWordsTokenizer struct {
	wordTokenizer *TreebankWordTokenizer
	stopWords     map[string]struct{}
}

func (t *BagOfWordsTokenizer) Tokenize(text string) []string {
	words := (*t).wordTokenizer.Tokenize(text)
	tokens := make([]string, 0, len(words))
	for _, word := range words {
		cleanWord := t.cleanupWord(word)
		if cleanWord != "" {
			tokens = append(tokens, cleanWord)
		}
	}
	return tokens
}

func (t *BagOfWordsTokenizer) cleanupWord(word string) string {
	if !WORD_WITHOUT_LAST_PUNCT.MatchString(word) {
		return ""
	}
	_word := WORD_WITHOUT_LAST_PUNCT.ReplaceAllString(word, "${1}")
	if _, isStopWord := (*t).stopWords[_word]; isStopWord {
		return ""
	}
	return _word
}

func NewBagOfWordsTokenizer(pathToStopWords string) *BagOfWordsTokenizer {
	f, err := os.Open("stop_words.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := new(BagOfWordsTokenizer)
	(*t).stopWords = loadStopWords(f)
	(*t).wordTokenizer = NewTreebankWordTokenizer()
	return t
}

func loadStopWords(reader io.Reader) map[string]struct{} {
	stopWords := make(map[string]struct{})
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		stopWords[scanner.Text()] = struct{}{}
	}
	return stopWords
}
