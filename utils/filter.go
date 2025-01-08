package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

func lowercaseFlter(tokens []string) []string {
	// convert all the tokens to lowercase
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func stopWordFilter(tokens []string) []string {
	var stopWords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}

	r := make([]string, 0, len(tokens))
	for i, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			// append the token removing the stopwords to the result slice
			r = append(r, tokens[i])
		}
	}
	return r
}

func stemFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		// stem the token
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
