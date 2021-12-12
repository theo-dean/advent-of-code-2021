package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func zeroCount(data []string) (zeroCount []int) {
	zeroCount = make([]int, len(data[0]))

	for _, line := range data {
		chars := strings.Split(line, "")
		for i, char := range chars {
			if char == "0" {
				zeroCount[i]++
			}
		}
	}
	return
}

func GammaAndEpsilon(data []string, zeroCount []int) (gamma string, epsilon string) {
	var gammaBuffer bytes.Buffer
	var epsilonBuffer bytes.Buffer
	var dataLength = len(data)
	for _, count := range zeroCount {
		if count > dataLength/2 {
			gammaBuffer.WriteString("0")
			epsilonBuffer.WriteString("1")
		} else {
			gammaBuffer.WriteString("1")
			epsilonBuffer.WriteString("0")
		}
	}
	gamma = gammaBuffer.String()
	epsilon = epsilonBuffer.String()
	return
}

func part1(data []string) (answer int64) {
	zeroCount := zeroCount(data)
	gamma, epsilon := GammaAndEpsilon(data, zeroCount)
	gammaInt, err1 := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, err2 := strconv.ParseInt(epsilon, 2, 64)
	if err1 != nil || err2 != nil {
		panic("Parsing int from gamma and epsilon binary strings failed!")
	}
	return gammaInt * epsilonInt
}

func Oxygen(data []string, pos int) (answer string) {
	zeroCount := zeroCount(data)
	filteredData := []string{}
	for _, str := range data {
		if str[pos] == '0' && zeroCount[pos] > len(data)/2 {
			filteredData = append(filteredData, str)
		} else if (str[pos] == '1' && zeroCount[pos] < len(data)/2) || (str[pos] == '1' && zeroCount[pos] == len(data)/2) {
			filteredData = append(filteredData, str)
		}
	}
	if len(filteredData) == 1 {
		return filteredData[0]
	}
	newPos := pos + 1
	return Oxygen(filteredData, newPos)
}

func CO2(data []string, pos int) (answer string) {
	zeroCount := zeroCount(data)
	filteredData := []string{}
	for _, str := range data {
		if str[pos] == '1' && zeroCount[pos] > len(data)/2 {
			filteredData = append(filteredData, str)
		} else if (str[pos] == '0' && zeroCount[pos] < len(data)/2) || (str[pos] == '0' && zeroCount[pos] == len(data)/2) {
			filteredData = append(filteredData, str)
		}
	}
	if len(filteredData) == 1 {
		return filteredData[0]
	}
	newPos := pos + 1
	return CO2(filteredData, newPos)
}

func part2(data []string) (answer int64) {
	o2String := Oxygen(data, 0)
	co2String := CO2(data, 0)
	o2, err1 := strconv.ParseInt(o2String, 2, 64)
	if err1 != nil {
		panic("Oxygen String to Binary Conversion Failed!")
	}
	co2, err2 := strconv.ParseInt(co2String, 2, 64)
	if err2 != nil {
		panic("CO2 String to Binary Conversion Failed!")
	}
	return o2 * co2
}

func main() {
	start := time.Now()
	data := FileToStringArray("input.txt")
	answer1 := part1(data)
	answer2 := part2(data)
	time := time.Since(start)
	fmt.Println(answer1)
	fmt.Println(answer2)
	fmt.Printf("Time: %v\n", time)
}
