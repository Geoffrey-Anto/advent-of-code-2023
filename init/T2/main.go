package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func Solve(s string) int {
	ans := 0

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

		inputString := string(text)

		res := Solve(inputString)

		ans += res
	}

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
