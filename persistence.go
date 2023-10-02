package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const saveFileName string = "save.txt"
const delimiter string = ","

func save(wpm int, accuracy float32) error {
  var err error
	wpmStr := strconv.Itoa(wpm)
	accuStr := fmt.Sprintf("%.2f", accuracy)

	file, err := os.OpenFile(saveFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

  _, err = file.WriteString(wpmStr + delimiter + accuStr + "\n")

	if err != nil {
		return err
	}

	return nil
}

func getAverage() (float32, float32, error) {
	file, err := os.Open(saveFileName)

	if err != nil {
		return 0, 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wpmSum float64 = 0
	var accuSum float64 = 0
	var count float64 = 0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")

		wpm, err := strconv.ParseFloat(split[0], 32)
		accuracy, err := strconv.ParseFloat(split[1], 32)

		if err != nil {
			return 0, 0, err
		}

		wpmSum += wpm
		accuSum += accuracy
		count += 1
	}

	return float32(wpmSum) / float32(count), float32(accuSum) / float32(count), nil
}
