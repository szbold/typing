package main

import (
	"os"
	"strconv"
)

func parseArgs() (int, int, error) {
	var err error
	args := os.Args[1:]

	numberOfWords := 0
	wordLength := 0

	for i, arg := range args {
		switch arg {
		case "-n":
			numberOfWords, err = strconv.Atoi(args[i+1])

			if err != nil {
				return 0, 0, err
			}
		case "-l":
			wordLength, err = strconv.Atoi(args[i+1])

			if err != nil {
				return 0, 0, err
			}
		}
	}

	return numberOfWords, wordLength, nil
}


