package main

import (
	"strings"
)

func (a *analyser) CalculateSentiment(text string) int {
	text = strings.ToLower(text)
	sentiment := 0
	for _, word := range strings.Split(text, " ") {
		sentiment += a.checkWord(word)
	}

	return sentiment
}

func (a *analyser) checkWord(word string) int {
	//Positivity Check
	node := *a.positiveWordsGraph
	for _, letter := range strings.Split(word, "") {
		newNode, ok := node[letter].(map[string]interface{})
		if ok {
			node = newNode
		} else {
			break
		}
	}
	if node["."] == "." {
		return 1
	}

	//Negativity Check
	node = *a.negativeWordsGraph
	for _, letter := range strings.Split(word, "") {
		newNode, ok := node[letter].(map[string]interface{})
		if ok {
			node = newNode
		} else {
			break
		}
	}
	if node["."] == "." {
		return -1
	}

	return 0
}

