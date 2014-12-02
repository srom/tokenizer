package tokenizer

import (
	"regexp"
	"strings"
)

// TreebankWordTokenizer is an English specific tokenizer which uses regular expressions to
// tokenize text as in Penn Treebank.
//
// Ported from NLTK's implementation: http://www.nltk.org/_modules/nltk/tokenize/treebank.html
//
// Regexp initialization happens at first call of Tokenize(). You can initialize in advance by
// creating the Tokenizer via NewTreebankWordTokenizer method.
type TreebankWordTokenizer struct {
	initialized bool
	steps       []step
}

type step struct {
	Regexp      *regexp.Regexp
	Replacement string
	Literal     bool
}

func (t *TreebankWordTokenizer) init() {
	steps := make([]step, 0, 25)

	// Starting quotes.
	steps = append(
		steps,
		// Starting quotes.
		step{regexp.MustCompile(`^"`), "``", true},
		step{regexp.MustCompile("(``)"), " ${1} ", false},
		step{regexp.MustCompile(`([ (\[{<])"`), "${1} `` ", false},

		// Punctuation.
		step{regexp.MustCompile(`([:,])([^\d])`), " ${1} ${2}", false},
		step{regexp.MustCompile(`\.\.\.`), " ... ", true},
		step{regexp.MustCompile(`[;@#$%&]`), " ${0} ", false},
		step{regexp.MustCompile(`([^\.])(\.)([\]\)}>"']*)\s*$`), "${1} ${2}${3} ", false},
		step{regexp.MustCompile(`[?!]`), " ${0} ", false},
		step{regexp.MustCompile(`([^'])' `), "${1} ' ", false},

		// Brackets and stuff.
		step{regexp.MustCompile(`[\]\[\(\)\{\}\<\>]`), " ${0} ", false},
		step{regexp.MustCompile(`--`), " -- ", true},

		// Ending quotes.
		step{regexp.MustCompile(`"`), " '' ", true},
		step{regexp.MustCompile(`(\S)('')`), "${1} ${2} ", false},

		// Contractions 1.
		step{regexp.MustCompile(`([^' ])('[sS]|'[mM]|'[dD]|') `), "${1} ${2} ", false},
		step{regexp.MustCompile(`([^' ])('ll|'LL|'re|'RE|'ve|'VE|n't|N'T) `), "${1} ${2} ", false},

		// Contractions 2.
		step{regexp.MustCompile(`(?i)\b(can)(not)\b`), " ${1} ${2} ", false},
		step{regexp.MustCompile(`(?i)\b(d)('ye)\b`), " ${1} ${2} ", false},
		step{regexp.MustCompile(`(?i)\b(gim)(me)\b`), " ${1} ${2} ", false},
		step{regexp.MustCompile(`(?i)\b(gon)(na)\b`), " ${1} ${2} ", false},
		step{regexp.MustCompile(`(?i)\b(got)(ta)\b`), " ${1} ${2} ", false},
		step{regexp.MustCompile(`(?i)\b(lem)(me)\b`), " ${1} ${2} ", false},
		step{regexp.MustCompile(`(?i)\b(mor)('n)\b`), " ${1} ${2} ", false},
		step{regexp.MustCompile(`(?i)\b(wan)(na)\s`), " ${1} ${2} ", false},

		// Contractions 3.
		step{regexp.MustCompile(`(?i) ('t)(is)\b`), " ${1} ${2} ", false},
		step{regexp.MustCompile(`(?i) ('t)(was)\b`), " ${1} ${2} ", false},

		// Clean up whitespaces.
		step{regexp.MustCompile(`\s`), " ", false},
		step{regexp.MustCompile(`\s+`), " ", false},
	)

	(*t).steps = steps
	(*t).initialized = true
}

func (t *TreebankWordTokenizer) Tokenize(text string) []string {
	if !(*t).initialized {
		t.init()
	}

	_text := text
	for _, step := range (*t).steps {
		if step.Literal {
			_text = step.Regexp.ReplaceAllLiteralString(_text, step.Replacement)
		} else {
			_text = step.Regexp.ReplaceAllString(_text, step.Replacement)
		}
	}
	_text = strings.TrimSpace(_text)
	return strings.Split(_text, " ")
}

func NewTreebankWordTokenizer() *TreebankWordTokenizer {
	t := new(TreebankWordTokenizer)
	t.init()
	return t
}
