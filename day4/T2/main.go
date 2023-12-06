package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var cache map[int]int = map[int]int{}

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

func GetPointsForGame(s string, id int) int {
	if cache[id] != -1 {
		return cache[id]
	}

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

	return ans
}

func PrefillCache(s []string) {
	for i := 0; i < len(s); i++ {
		cache[i] = -1
	}

	for i := 0; i < len(s); i++ {
		cache[i] = GetPointsForGame(s[i], i)
	}
}

func GetTotalCardsWon(s []string) int {
	ans := 0

	mem := map[int]int{}

	n := len(s)

	mem[0] = 0

	for i := 0; i < n; i++ {
		mem[i]++

		times := mem[i]

		res := GetPointsForGame(s[i], i)

		for j := 0; j < times; j++ {
			for k := i + 1; k <= i+res; k++ {
				mem[k] += 1
			}
		}
	}

	for _, v := range mem {
		ans += v
	}

	return ans
}

func main() {
	in := createFileReader("input.txt")
	out := createFileWriter("output.txt")

	ans := 0

	input := []string{}

	for {
		text, _, err := in.ReadLine()

		if err != nil {
			break
		}

		inputString := string(text)

		input = append(input, inputString)
	}

	PrefillCache(input)

	ans = GetTotalCardsWon(input)

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output\n", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
