package main

import (
	"errors"
	"strings"
)

func (a *analyser) IdentifyCrypto(text string) ([]string, error) {
	cryptos := *a.cryptoCurrencies
	tickers := *a.tickers
	var matches = make(map[string]bool)
	for _, word := range strings.Split(text, " ") {
		lowercaseWord := strings.ToLower(word)
		elem, ok := cryptos[lowercaseWord]
		if ok {
			matches[elem] = true
		}

		elem, ok = tickers[word]
		if ok {
			matches[elem] = true
		}
	}

	keys := make([]string, 0, len(matches))
	for k := range matches {
		keys = append(keys, k)
	}

	if len(keys) > 0 {
		return keys, nil
	}
	return keys, errors.New("no matches")
}
