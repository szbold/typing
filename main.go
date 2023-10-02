package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
	"time"
  "math"
)

const url string = "http://random-word-api.herokuapp.com/word"

func typingGame(words []string) (int, float32) {
  var start time.Time
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

    if i == 0 {
      start = time.Now()
    }

		if char == inputted_char[0] {
			correctCount++
			correct.Print(inputted_char)
		} else {
			incorrectCount++
			incorrect.Print(inputted_char)
		}
	}

  end := time.Since(start).Minutes()

  wpm := int(math.Floor(float64(len(words)) * (1 / end)))
  accuracy := float32(correctCount)/float32(correctCount+incorrectCount)*100

	fmt.Printf("\nSpeed: %v WPM\n", wpm)

	fmt.Printf("Accuracy: %.2f\n", accuracy)

  return wpm, accuracy
}

func main() {
	var err error
	numberOfWords, wordLength, err := parseArgs()

	if err != nil {
		panic(err)
	}

	words := fetchWords(numberOfWords, wordLength)

  wpm, accuracy := typingGame(words)

  err = save(wpm, accuracy)

  if err != nil {
    fmt.Println("Could not save file", err)
  }

  avgWpm, avgAccuracy, err := getAverage()

  if err != nil {
    fmt.Println("Error retrieving data", err)
  }

  fmt.Printf("Average wpm: %v\nAverage accuracy: %v\n", avgWpm, avgAccuracy)
}
