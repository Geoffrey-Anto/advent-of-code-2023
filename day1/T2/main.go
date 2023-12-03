package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func wordsToInt(s string) int {
	var WORDS = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, word := range WORDS {
		if s == word {
			return i
		}
	}

	return 0
}

func isWord(i int, j int, line string) bool {
	var WORDS = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	substr := line[i:j]
	for _, word := range WORDS {
		if substr == word {
			return true
		}
	}
	return false
}

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

func isNumeric(ch byte) bool {
	if ch <= 57 && ch >= 48 {
		return true
	}

	return false
}

func getCalibrationValue(line string) int {
	n := len(line)

	first, last := 100000, -100000
	firstWord, lastWord := 100000, -100000
	firstWordValue, lastWordValue := -1, -1

	for i := 0; i < n; i++ {
		if isNumeric(line[i]) {
			first = i
			break
		}
	}

	shouldEnd := false

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isWord(i, j+1, line) {
				firstWord = i
				firstWordValue = wordsToInt(line[i:j+1]) + 1
				shouldEnd = true
				break
			}
		}
		if shouldEnd {
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		if isNumeric(line[i]) {
			last = i
			break
		}
	}

	shouldEnd = false

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if isWord(i, j+1, line) {
				lastWord = i
				lastWordValue = wordsToInt(line[i:j+1]) + 1
				shouldEnd = true
				break
			}
		}
		if shouldEnd {
			break
		}
	}

	left := ""
	right := ""

	// fmt.Println(first)
	// fmt.Println(firstWord)
	// fmt.Println(last)
	// fmt.Println(lastWord)

	if first < firstWord {
		left = fmt.Sprintf("%c", line[first])
	} else {
		left = fmt.Sprint(firstWordValue)
	}

	if last > lastWord {
		right = fmt.Sprintf("%c", line[last])
	} else {
		right = fmt.Sprint(lastWordValue)
	}

	ansres, err := strconv.Atoi(fmt.Sprintf("%s%s", left, right))

	// fmt.Println(ansres)

	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to convert digit")
	}

	return ansres
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

		// fmt.Println(string(text))

		res := getCalibrationValue(string(text))

		ans += res
	}

	res, err := out.WriteString(fmt.Sprintln(ans))

	fmt.Printf("%d bytes written to output\n\n", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
