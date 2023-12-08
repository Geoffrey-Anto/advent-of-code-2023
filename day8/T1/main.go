package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Route struct {
	left  string
	right string
}

type Graph map[string]Route

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

func AddRouteToGraph(graph *Graph, s string) {
	source, dest := strings.Split(s, " = ")[0], strings.Split(s, " = ")[1]

	dest = dest[1 : len(dest)-1]

	left := strings.Split(dest, ", ")[0]
	right := strings.Split(dest, ", ")[1]

	(*graph)[source] = Route{
		left:  left,
		right: right,
	}
}

func GetLengthOfTraversedPath(graph *Graph, s string) int {
	ans := 0

	pathIdx := 0
	currVal := "AAA"

	for {
		if pathIdx >= len(s) {
			pathIdx = 0
		}

		currSide := string(s[pathIdx])

		switch currSide {
		case "L":
			currVal = (*graph)[currVal].left
		case "R":
			currVal = (*graph)[currVal].right
		default:
			log.Fatal("ERROR")
		}

		pathIdx++
		ans++

		if currVal == "ZZZ" {
			break
		}
	}

	return ans
}

func main() {
	in := createFileReader("input.txt")
	out := createFileWriter("./output.txt")

	var graph Graph = Graph{}

	text, _, err := in.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	path := string(text)

	in.ReadLine()

	ans := 0

	for {
		text, _, err := in.ReadLine()

		if err != nil {
			break
		}

		inputString := string(text)

		AddRouteToGraph(&graph, inputString)
	}

	ans = GetLengthOfTraversedPath(&graph, path)

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
