package T1

import (
	"bufio"
	"fmt"
	"log"
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

func isNumeric(ch byte) bool {
	if ch <= 57 && ch >= 48 {
		return true
	}

	return false
}

func getCalibrationValue(line string) int {
	n := len(line)

	first, last := 0, 0

	for i := 0; i < n; i++ {
		if isNumeric(line[i]) {
			first = i
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		if isNumeric(line[i]) {
			last = i
			break
		}
	}

	ans, err := strconv.Atoi(fmt.Sprintf("%c%c", line[first], line[last]))

	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to convert digit")
	}

	return ans
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

		res := getCalibrationValue(string(text))

		ans += res
	}

	res, err := out.WriteString(fmt.Sprintln(ans))

	fmt.Printf("%d bytes written to output", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
