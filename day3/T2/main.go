package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func Contains(a []Point, x Point) bool {
	for _, n := range a {
		if x.x == n.x && x.y == n.y {
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
		if s[row][col] == '*' {
			return true
		}
	}

	return false
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
				ans[pt] = []int{}
			}
		}
	}

	return ans
}

func isNumberPartOfEngineSchematic(s []string, num string, tempStart int, tempEnd int, row int) []Point {
	n := len(s)

	var arr []Point = []Point{}

	for i := tempStart; i <= tempEnd; i++ {
		for _, pos := range POSITIONS {
			if isValid(s, row+pos.y, i+pos.x, n) {
				pt := Point{
					x: i + pos.x,
					y: row + pos.y,
				}

				if Contains(arr, pt) {
					continue
				}

				arr = append(arr, pt)
			}
		}
	}

	return arr
}

func GetPartsForEngineSchematics(s []string, starPos map[Point][]int) int {
	n := len(s)
	ans := 0

	for i := 0; i < n; i++ {
		temp := ""
		for j := 0; j < n; j++ {
			if isNumeric(s[i][j]) {
				temp += string(s[i][j])

				if j != n-1 {
					continue
				}
			}
			if temp == "" {
				continue
			} else {
				num := temp

				tempStart := j - len(temp)
				tempEnd := j - 1
				pts := isNumberPartOfEngineSchematic(s, num, tempStart, tempEnd, i)

				if len(pts) > 0 {
					res, _ := strconv.Atoi(num)

					for _, pt := range pts {
						starPos[pt] = append(starPos[pt], res)
					}
				}

				temp = ""
			}
		}
	}

	for pt, nums := range starPos {
		fmt.Printf("%+v ---> %+v\n", pt, nums)
		if len(nums) == 2 {
			n1 := nums[0]
			n2 := nums[1]

			ans += (n1 * n2)
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
