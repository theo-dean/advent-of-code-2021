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
	count = 0
	for i, current := range data {
		if i != 0 && current > data[i-1] {
			count++
		}
	}
	return
}

func main() {
	data := FileToIntArray("input.txt")
	triplets := triplets(data)
	count := increaseCount(triplets)
	fmt.Print(count)
}
