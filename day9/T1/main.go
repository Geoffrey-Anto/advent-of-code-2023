package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func createFileReader(fileName string) *bufio.Reader {
	stream, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Failed to read file")
	}

	in := bufio.NewReader(stream)

	return in
}

func createFileWriter(fileName string) *os.File {
	stream, err := os.Create(fileName)

	if err != nil {
		log.Fatal("Failed to write file")
	}

	return stream
}

func GetDiff(history []int) []int {
	res := []int{}

	for i := 0; i < len(history)-1; i++ {
		res = append(res, history[i+1]-history[i])
	}

	return res
}

func isValid(history []int) bool {
	if len(history) == 0 {
		return false
	}

	cnt := 0

	for _, v := range history {
		if v == 0 {
			cnt++
		}
	}

	return cnt != len(history)
}

func GetPrediction(s string) int {
	currHistory := []int{}

	for _, v := range strings.Split(s, " ") {
		val, _ := strconv.Atoi(v)

		currHistory = append(currHistory, val)
	}

	history := [][]int{}

	history = append(history, currHistory)

	for isValid(history[len(history)-1]) {
		diffs := GetDiff(history[len(history)-1])
		history = append(history, diffs)
	}

	history[len(history)-1] = append(history[len(history)-1], 0)

	for i := len(history) - 2; i >= 0; i-- {
		history[i] = append(history[i], history[i+1][len(history[i+1])-1]+history[i][len(history[i])-1])
	}

	return history[0][len(history[0])-1]
}

func main() {
	in := createFileReader("input.txt")
	out := createFileWriter("./output.txt")

	ans := 0

	for {
		text, _, err := in.ReadLine()

		if err != nil {
			break
		}

		inputString := string(text)

		res := GetPrediction(inputString)

		ans += res
	}

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
