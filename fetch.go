package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func fetchWords(numberOfWords int, wordLength int) []string {
	var err error

	if numberOfWords == 0 {
		numberOfWords = 10
	}

	if wordLength == 0 {
		wordLength = 5
	}

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	query := req.URL.Query()

	query.Add("number", strconv.Itoa(numberOfWords))
	query.Add("length", strconv.Itoa(wordLength))

	req.URL.RawQuery = query.Encode()

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	words := make([]string, numberOfWords)

	err = json.Unmarshal(body, &words)

	if err != nil {
		panic(err)
	}

	return words
}


