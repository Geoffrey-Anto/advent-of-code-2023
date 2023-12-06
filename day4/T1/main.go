package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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

func RemoveLeftRightPadding(s string) string {
	i := 0
	n := len(s)

	for i < n && s[i] == ' ' {
		i++
	}

	j := n - 1

	for j >= 0 && s[j] == ' ' {
		j--
	}

	return s[i : j+1]
}

func SplitByValue(s string, ch byte) []string {
	res := []string{}

	temp := ""

	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] != ch {
			temp += string(s[i])
		} else {
			if temp != "" {
				res = append(res, RemoveLeftRightPadding(temp))
			}
			temp = ""
		}
	}

	if temp != "" {
		res = append(res, RemoveLeftRightPadding(temp))
	}

	return res
}

func GetPointsForGame(s string) int {
	ans := 0

	GameData := SplitByValue(s, ':')[1]

	WinningPart, CurrentCard := SplitByValue(GameData, '|')[0], SplitByValue(GameData, '|')[1]

	WinningNumbers := SplitByValue(WinningPart, ' ')
	CurrentReceivedNumbers := SplitByValue(CurrentCard, ' ')

	for _, w := range WinningNumbers {
		found := false
		pt := -1
		wr, _ := strconv.Atoi(w)

		for idx, c := range CurrentReceivedNumbers {
			cr, _ := strconv.Atoi(c)
			if wr == cr {
				found = true
				pt = idx
				break
			}

		}

		if found {
			ans++
			CurrentReceivedNumbers[pt] = "-1"
		}
	}

	if ans == 0 {
		return 0
	} else {
		return int(math.Pow(2, float64(ans-1)))
	}
}

func main() {
	in := createFileReader("input.txt")
	out := createFileWriter("output.txt")

	ans := 0

	for {
		text, _, err := in.ReadLine()

		if err != nil {
			break
		}

		inputString := string(text)

		res := GetPointsForGame(inputString)

		ans += res
	}

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output\n", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
