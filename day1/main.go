package main

import "fmt"

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
	input := FileToIntArray("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
