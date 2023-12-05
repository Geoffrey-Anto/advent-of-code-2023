package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x int
	y int
}

var POSITIONS = []Point{
	{x: -1, y: -1},
	{x: -1, y: 0},
	{x: -1, y: 1},
	{x: 0, y: -1},
	{x: 0, y: 1},
	{x: 1, y: -1},
	{x: 1, y: 0},
	{x: 1, y: 1},
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

func Replace(s string, start int, end int, ch byte) string {
	bytes := []byte(s)

	for i := range bytes {
		if i >= start && i <= end {
			bytes[i] = ch
		}
	}

	return string(bytes)
}

func isNumeric(ch byte) bool {
	if ch <= 57 && ch >= 48 {
		return true
	}

	return false
}

func isValid(s []string, row int, col int, n int) bool {
	if row < n && row >= 0 && col < n && col >= 0 {
		if s[row][col] != '.' && !isNumeric(s[row][col]) {
			return true
		}
	}

	return false
}

func isNumberPartOfEngineSchematic(s []string, num string, tempStart int, tempEnd int, row int) bool {
	n := len(s)

	flag := false

	for i := tempStart; i <= tempEnd; i++ {
		for _, pos := range POSITIONS {
			if isValid(s, row+pos.y, i+pos.x, n) {
				flag = true
			}
		}
	}

	return flag
}

func GetAllStarsMap(s []string) map[Point][]int {
	ans := make(map[Point][]int)

	for i, row := range s {
		for j, ch := range row {
			if ch == '*' {
				pt := Point{
					x: i,
					y: j,
				}
				ans[pt] = []int{0, 0}
			}
		}
	}

	return ans
}

func GetPartsForEngineSchematics(s []string, starPos map[Point][]int) int {
	n := len(s)
	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

		}
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

	starPos := GetAllStarsMap(input)

	fmt.Println(len(starPos))

	ans = GetPartsForEngineSchematics(input, starPos)

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output\n", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
