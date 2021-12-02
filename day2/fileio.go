package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func FileToVectors(filename string) (directions []vector) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		usefulLine := strings.Fields(line)
		if magnitude, err := strconv.Atoi(usefulLine[1]); err == nil {
			directions = append(directions, vector{usefulLine[0], magnitude})
		}
	}
	return
}
