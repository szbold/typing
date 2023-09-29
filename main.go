package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const url string = "http://random-word-api.herokuapp.com/word"

func typingGame(words []string, start time.Time) {
	input := strings.Join(words, " ")
	color.Yellow(input)

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	b := make([]byte, 1)

	correct := color.New(color.FgGreen)
	incorrect := color.New(color.FgRed)

	correctCount := 0
	incorrectCount := 0

	for i := 0; i < len(input); i++ {
		char := input[i]

		os.Stdin.Read(b)

		inputted_char := string(b)

		if inputted_char[0] == 127 {
			os.Stdout.Write([]byte("\b"))

			if i > 0 {
				i -= 2
			} else {
				i = -1
			}

			continue
		}

		if char == inputted_char[0] {
			correctCount++
			correct.Print(inputted_char)
		} else {
			incorrectCount++
			incorrect.Print(inputted_char)
		}
	}

	fmt.Printf("Speed: %v WPM", time.Since(start))

	fmt.Printf("\nScore: %.2f\n", float32(correctCount)/float32(correctCount+incorrectCount)*100)
}

func main() {
	var err error
	numberOfWords, wordLength, err := parseArgs()

	if err != nil {
		panic(err)
	}

	words := fetchWords(numberOfWords, wordLength)

	typingGame(words, time.Now())
}
