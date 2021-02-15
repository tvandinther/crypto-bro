package main

import (
	"encoding/json"
	"fmt"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type analyser struct {
	positiveWordsGraph *map[string]interface{}
	negativeWordsGraph *map[string]interface{}
	cryptoCurrencies *map[string]string
	tickers *map[string]string
}

func main() {
	positiveWords := make(map[string]interface{})
	negativeWords := make(map[string]interface{})
	cryptoCurrencies := make(map[string]string)
	tickers := make(map[string]string)
	PopulateMapsFromFile("positive_words_graph.json", &positiveWords)
	PopulateMapsFromFile("negative_words_graph.json", &negativeWords)
	PopulateMapsFromFile("cryptocurrencies.json", &cryptoCurrencies)
	PopulateMapsFromFile("tickers.json", &tickers)

	apiHandle, _ := reddit.NewScript("My agent", 5 * time.Second)

	subreddits := []string{
		"cryptocurrency",
	}

	cfg := graw.Config{
		SubredditComments: subreddits,
	}

	analyser := &analyser{
		positiveWordsGraph: &positiveWords,
		negativeWordsGraph: &negativeWords,
		cryptoCurrencies: &cryptoCurrencies,
		tickers: &tickers,
	}

	stop, wait, err := graw.Scan(analyser, apiHandle, cfg)
	defer stop()

	if err != nil {
		fmt.Printf("graw scan encountered an initialisation error: %v\n", err)
	}

	if err := wait(); err != nil {
		fmt.Printf("graw scan encountered a runtime error: %v\n", err)
	}
}

func (a *analyser) Comment(comment *reddit.Comment) error {
	crypto, sentiment, err := a.processText(comment.Body)
	if err == nil {
		formattedText := fmt.Sprintf("[%s] %d\n%s\n", strings.Join(crypto, ", "), sentiment, comment.Body)
		fmt.Println(formattedText)
	}

	return nil
}


func PopulateMapsFromFile(filename string, mapRef interface{}) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, mapRef)
	if err != nil {
		panic(err)
	}
}

