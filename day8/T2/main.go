package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func Iterative(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(Iterative(a, b)))
}

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

func GetNextTraversedPath(graph *Graph, s string, startingNode string, pathIdx int) (string, int) {
	currVal := startingNode

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

	return currVal, pathIdx + 1
}

func GetNoOfPathsTraversed(graph Graph, path string) int {
	ans := 0
	startingNodes := []string{}

	for s := range graph {
		if s[len(s)-1] == 'A' {
			startingNodes = append(startingNodes, s)
		}
	}

	c := make([]string, len(startingNodes))
	p := make([]int, len(startingNodes))

	for i, v := range startingNodes {
		c[i] = v
	}

	steps := make([]int, len(startingNodes))

	for {
		for i := 0; i < len(startingNodes); i++ {
			c[i], p[i] = GetNextTraversedPath(&graph, path, c[i], p[i])
		}

		ans++

		for i := 0; i < len(startingNodes); i++ {
			if c[i][len(c[i])-1] == 'Z' {
				steps[i] = ans
			}
		}

		if !slices.Contains(steps, 0) {
			break
		}
	}

	ans = int(Lcm(int64(steps[0]), int64(steps[1])))

	for i := 2; i < len(startingNodes); i++ {
		ans = int(Lcm(int64(ans), int64(steps[i])))
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

	ans = GetNoOfPathsTraversed(graph, path)

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
