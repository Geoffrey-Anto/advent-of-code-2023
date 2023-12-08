package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand int

const (
	FIVE_OF_A_KIND  Hand = 6
	FOUR_OF_A_KIND  Hand = 5
	FULL_HOUSE      Hand = 4
	THREE_OF_A_KIND Hand = 3
	TWO_PAIR        Hand = 2
	ONE_PAIR        Hand = 1
	HIGH_CARD       Hand = 0
)

func GetEnumFromString(s string) Hand {
	ans := HIGH_CARD

	mp := map[string]int{}

	jokerCount := 0

	for _, v := range s {
		if v == 'J' {
			jokerCount++
			continue
		}
		mp[string(v)] += 1
	}

	mx, mxVal := 0, ""

	for k, v := range mp {
		if v > mx {
			mx = v
			mxVal = k
		}
	}

	mp[mxVal] += jokerCount

	if len(mp) == 1 {
		ans = FIVE_OF_A_KIND
	} else if len(mp) == 2 {
		isFullHouse := false
		for _, v := range mp {
			if v == 2 {
				isFullHouse = true
			}
		}

		if isFullHouse {
			ans = FULL_HOUSE
		} else {
			ans = FOUR_OF_A_KIND
		}
	} else if len(mp) == 3 {
		isThree := false
		for _, v := range mp {
			if v == 3 {
				isThree = true
			}
		}

		if isThree {
			ans = THREE_OF_A_KIND
		} else {
			ans = TWO_PAIR
		}
	} else if len(mp) == 4 {
		ans = ONE_PAIR
	} else {
		ans = HIGH_CARD
	}

	return ans
}

type Card struct {
	handValue string
	hand      Hand
	bid       int
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

func isHandBigger(a string, b string) bool {
	hands := map[string]int{"A": 12, "K": 11, "Q": 10, "T": 8, "9": 7, "8": 6, "7": 5, "6": 4, "5": 3, "4": 2, "3": 1, "2": 0, "J": -1}

	for i := 0; i < 5; i++ {
		if a[i] == b[i] {
			continue
		} else {
			return hands[string(a[i])] > hands[string(b[i])]
		}
	}

	return false
}

func compare(a Card, b Card) int {
	if a.hand != b.hand {
		if a.hand > b.hand {
			return 1
		} else {
			return -1
		}
	} else {
		handA := a.handValue
		handB := b.handValue

		for i := 0; i < 5; i++ {
			if isHandBigger(handA, handB) {
				return 1
			} else {
				return -1
			}
		}
	}

	return 1
}

func GetTotalWinnings(s []Card) int {
	ans := 0

	slices.SortFunc(s, func(a, b Card) int {
		res := compare(a, b)
		fmt.Println(a.handValue, b.handValue, res)
		return res
	})

	rank := 1

	for _, val := range s {
		ans += (rank * val.bid)
		rank++
	}

	return ans
}

func main() {
	in := createFileReader("input.txt")
	out := createFileWriter("./output.txt")

	ans := 0

	input := []Card{}

	for {
		text, _, err := in.ReadLine()

		if err != nil {
			break
		}

		inputString := string(text)

		card := strings.Split(inputString, " ")

		bid, _ := strconv.Atoi(card[1])

		input = append(input, Card{
			handValue: card[0],
			hand:      GetEnumFromString(card[0]),
			bid:       bid,
		})
	}

	ans = GetTotalWinnings(input)

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
