package main

import (
	"fmt"
	"time"
)

func triplets(data []int) (triplets []int) {
	for i, current := range data[:len(data)-2] {
		triplet := current + data[i+1] + data[i+2]
		triplets = append(triplets, triplet)
	}
	return
}

func increaseCount(data []int) (count int) {
	for i, current := range data {
		if i != 0 && current > data[i-1] {
			count++
		}
	}
	return
}

func part1(input []int) (count int) {
	return increaseCount(input)
}

func part2(input []int) (count int) {
	triplets := triplets(input)
	return increaseCount(triplets)
}

func main() {
	start := time.Now()
	input := FileToIntArray("input.txt")
	part1 := part1(input)
	part2 := part2(input)
	time := time.Since(start)

	fmt.Println(part1)
	fmt.Println(part2)
	fmt.Println("Time: ", time)
}
