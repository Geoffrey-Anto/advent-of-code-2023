package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Bag struct {
	red   int
	blue  int
	green int
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

func RemoveLeftRightPadding(s string) string {
	i := 0
	n := len(s)

	for s[i] == ' ' && i < n {
		i++
	}

	j := n - 1

	for s[j] == ' ' && i >= 0 {
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
			res = append(res, RemoveLeftRightPadding(temp))
			temp = ""
		}
	}

	if temp != "" {
		res = append(res, RemoveLeftRightPadding(temp))
	}

	return res
}

func isGamePlayable(s string, b Bag) int {
	res := SplitByValue(s, ':')

	GameId := res[0]
	GameString := res[1]

	GameNoString := SplitByValue(GameId, ' ')[1]

	GameNo, _ := strconv.Atoi(GameNoString)

	GameRounds := SplitByValue(GameString, ';')

	for _, GameTurn := range GameRounds {
		GamePicks := SplitByValue(GameTurn, ',')

		for _, Pick := range GamePicks {
			count := SplitByValue(Pick, ' ')[0]
			BallCount, _ := strconv.Atoi(count)
			BallType := SplitByValue(Pick, ' ')[1]

			if BallType == "blue" {
				if BallCount > b.blue {
					return 0
				}
			} else if BallType == "green" {
				if BallCount > b.green {
					return 0
				}
			} else if BallType == "red" {
				if BallCount > b.red {
					return 0
				}
			} else {
				log.Fatal("Wrong Color provided!")
			}
		}
	}

	return GameNo
}

func main() {
	in := createFileReader("input.txt")
	out := createFileWriter("output.txt")

	bag := Bag{
		red:   12,
		blue:  14,
		green: 13,
	}

	ans := 0

	for {
		text, _, err := in.ReadLine()

		if err != nil {
			break
		}

		inputString := string(text)

		res := isGamePlayable(inputString, bag)

		ans += res
	}

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output\n", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
