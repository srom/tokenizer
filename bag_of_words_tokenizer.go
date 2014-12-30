package tokenizer

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

type contraction struct {
	Regexp      *regexp.Regexp
	Replacement string
}

var wordWithoutLastPunct *regexp.Regexp = nil
var commonContractions []contraction = nil

// BagOfWordsTokenizer output a list of words suitable to be used to build a bag of words matrix.
//
// It uses a Treebank tokenizer internally but removes puncutation, and stopwords.
type BagOfWordsTokenizer struct {
	wordTokenizer Tokenizer
	stopWords     map[string]struct{}
	initalized    bool
}

func (t *BagOfWordsTokenizer) Tokenize(text string) []string {
	if !(*t).initalized {
		panic("BagOfWordsTokenizer: Initalize first by calling NewBagOfWordsTokenizer")
	}
	_text := t.preprocessing(text)

	words := (*t).wordTokenizer.Tokenize(_text)

	tokens := make([]string, 0, len(words))
	for _, word := range words {
		cleanWord := t.cleanupWord(word)
		if cleanWord != "" {
			tokens = append(tokens, cleanWord)
		}
	}
	return tokens
}

func (t *BagOfWordsTokenizer) init(pathToStopWords string) {
	f, err := os.Open(pathToStopWords)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wordWithoutLastPunct = regexp.MustCompile(`^(\pL+)[^\pL]?$`)
	commonContractions = []contraction{
		contraction{regexp.MustCompile("can't"), "can not"},
		contraction{regexp.MustCompile("won't"), "will not"},
		contraction{regexp.MustCompile(`([^' ])('[sS]) `), "${1} is "},
		contraction{regexp.MustCompile(`([^' ])('[mM]) `), "${1} am "},
		contraction{regexp.MustCompile(`([^' ])('[dD]) `), "${1} had "},
	}
	(*t).stopWords = loadStopWords(f)
	(*t).wordTokenizer = NewTreebankWordTokenizer()
	(*t).initalized = true
}

func (t *BagOfWordsTokenizer) preprocessing(text string) string {
	_text := text
	for _, commonContraction := range commonContractions {
		_text = commonContraction.Regexp.ReplaceAllString(_text, commonContraction.Replacement)
	}
	return _text
}

func (t *BagOfWordsTokenizer) cleanupWord(word string) string {
	if !wordWithoutLastPunct.MatchString(word) {
		return ""
	}
	_word := wordWithoutLastPunct.ReplaceAllString(word, "${1}")
	_word = strings.ToLower(_word)
	if _, isStopWord := (*t).stopWords[_word]; isStopWord {
		return ""
	}
	return _word
}

func NewBagOfWordsTokenizer(pathToStopWords string) *BagOfWordsTokenizer {
	t := new(BagOfWordsTokenizer)
	t.init(pathToStopWords)
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
