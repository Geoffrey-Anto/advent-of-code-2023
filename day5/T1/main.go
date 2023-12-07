package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type AlmanacMap struct {
	source int
	dest   int
	length int
}

type Almanac struct {
	seed2Soil            []AlmanacMap
	soil2Fertilizer      []AlmanacMap
	fertilizer2Water     []AlmanacMap
	water2Light          []AlmanacMap
	light2Temperature    []AlmanacMap
	temperature2Humidity []AlmanacMap
	humidity2Location    []AlmanacMap
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

func BuildMap(fileName string) []AlmanacMap {
	in := createFileReader(fileName)

	str := []string{}

	for {
		text, _, err := in.ReadLine()

		if err != nil {
			break
		}

		str = append(str, string(text))
	}

	ans := []AlmanacMap{}

	for _, s := range str {
		mapInfo := SplitByValue(s, ' ')

		dest, _ := strconv.Atoi(mapInfo[0])
		start, _ := strconv.Atoi(mapInfo[1])
		length, _ := strconv.Atoi(mapInfo[2])

		temp := AlmanacMap{}
		temp.dest = dest
		temp.source = start
		temp.length = length

		ans = append(ans, temp)
	}

	return ans
}

func GetMappedValueForAlmanac(mp []AlmanacMap, source int) int {
	ans := 0

	found := false

	for _, m := range mp {
		if m.source <= source && m.source+m.length > source {
			found = true

			diff := source - m.source

			ans = m.dest + diff
		}
	}

	if !found {
		return source
	}

	return ans
}

func GetLocationForSeed(almanac *Almanac, seed int) int {
	soil := GetMappedValueForAlmanac(almanac.seed2Soil, seed)
	fertilizer := GetMappedValueForAlmanac(almanac.soil2Fertilizer, soil)
	water := GetMappedValueForAlmanac(almanac.fertilizer2Water, fertilizer)
	light := GetMappedValueForAlmanac(almanac.water2Light, water)
	temperature := GetMappedValueForAlmanac(almanac.light2Temperature, light)
	humidity := GetMappedValueForAlmanac(almanac.temperature2Humidity, temperature)
	location := GetMappedValueForAlmanac(almanac.humidity2Location, humidity)

	return location
}

func main() {
	isTest := false

	env := ""

	if isTest {
		env = "test"
	} else {
		env = "prod"
	}

	in := createFileReader(fmt.Sprintf("%s/input.txt", env))
	out := createFileWriter("./output.txt")

	seeds := []int{}

	seedBytes, _, err := in.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	seeds_ := SplitByValue(string(seedBytes), ' ')

	maxSeed := -1

	for _, s := range seeds_ {
		temp, _ := strconv.Atoi(s)
		maxSeed = max(maxSeed, temp)
		seeds = append(seeds, temp)
	}

	ans := 99999999999

	almanac := Almanac{}

	almanac.seed2Soil = BuildMap(fmt.Sprintf("%s/s2s.txt", env))
	almanac.soil2Fertilizer = BuildMap(fmt.Sprintf("%s/s2f.txt", env))
	almanac.fertilizer2Water = BuildMap(fmt.Sprintf("%s/f2w.txt", env))
	almanac.water2Light = BuildMap(fmt.Sprintf("%s/w2l.txt", env))
	almanac.light2Temperature = BuildMap(fmt.Sprintf("%s/l2t.txt", env))
	almanac.temperature2Humidity = BuildMap(fmt.Sprintf("%s/t2h.txt", env))
	almanac.humidity2Location = BuildMap(fmt.Sprintf("%s/h2l.txt", env))

	for _, seed := range seeds {
		res := GetLocationForSeed(&almanac, seed)
		ans = min(ans, res)
	}

	res, err := out.WriteString(fmt.Sprint(ans))

	fmt.Printf("%d bytes written to output\n", res)

	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
