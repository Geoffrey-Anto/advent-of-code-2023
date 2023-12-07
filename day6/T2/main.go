package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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

func GetNoOfWinningWaysForTime(t int, d int) int {
	ans := 0

	for timeHeld := 0; timeHeld <= t; timeHeld++ {
		speedAccumulated := timeHeld

		distanceTraveled := speedAccumulated * (t - timeHeld)

		if distanceTraveled > d {
			ans++
		}
	}

	return ans
}

func GetTotalWinningWays(t string, d string) int {
	ans := 1

	time := SplitByValue(t, ':')[1]
	distance := SplitByValue(d, ':')[1]

	times := []int{}
	distances := []int{}

	for _, v := range SplitByValue(time, ' ') {
		r, _ := strconv.Atoi(v)
		times = append(times, r)
	}

	for _, v := range SplitByValue(distance, ' ') {
		r, _ := strconv.Atoi(v)
		distances = append(distances, r)
	}

	n := len(times)

	for i := 0; i < n; i++ {
		r := GetNoOfWinningWaysForTime(times[i], distances[i])
		ans *= r
	}

	return ans
}

func main() {
	in := createFileReader("input.txt")
	out := createFileWriter("./output.txt")

	ans := 0

	text, _, err := in.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	t := string(text)

	text, _, err = in.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	distance := string(text)

	clockStart := time.Now()

	ans = GetTotalWinningWays(t, distance)

	fmt.Printf("Time Taken: %+v\n", time.Now().Sub(clockStart))

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
