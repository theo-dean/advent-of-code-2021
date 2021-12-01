package main

import (
	"bufio"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FileToIntArray(filename string) (data []int) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineInt, err := strconv.Atoi(line)
		check(err)
		data = append(data, lineInt)
	}

	return
}
