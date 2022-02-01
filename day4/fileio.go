package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LineToIntArray(filename string) (data []int) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	rawData := strings.Split(line, ",")
	for i := 0; i < len(rawData); i++ {
		num, err := strconv.Atoi(rawData[i])
		check(err)
		data = append(data, num)
	}
	return
}

func BoardTo2DStringArray(filename string) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// first line is the numbers
	scanner.Scan()

}
