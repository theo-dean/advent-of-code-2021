package main

import (
	"fmt"
	"time"
)

type vector struct {
	direction string
	magnitude int
}

func sumVector1(vectors []vector) (horizontal int, depth int) {
	for _, vec := range vectors {
		switch vec.direction {
		case "forward":
			horizontal += vec.magnitude
		case "down":
			depth += vec.magnitude
		case "up":
			depth -= vec.magnitude
		default:
			panic("Ahhhh! What is this Vector?!?! Ahhhhh!")
		}
	}
	return
}

func sumVector2(vectors []vector) (horizontal int, depth int) {
	aim := 0
	for _, vec := range vectors {
		switch vec.direction {
		case "forward":
			horizontal += vec.magnitude
			depth += aim * vec.magnitude
		case "down":
			aim += vec.magnitude
		case "up":
			aim -= vec.magnitude
		default:
			panic("Ahhhh! What is this Vector?!?! Ahhhhh!")
		}
	}
	return
}

func part1(input []vector) (answer int) {
	horizontal, depth := sumVector1(input)
	return horizontal * depth
}

func part2(input []vector) (answer int) {
	horizontal, depth := sumVector2(input)
	return horizontal * depth
}

func main() {
	start := time.Now()
	vectors := FileToVectors("input.txt")
	solution1 := part1(vectors)
	solution2 := part2(vectors)
	time := time.Since(start)

	fmt.Println(solution1)
	fmt.Println(solution2)
	fmt.Println("Time: ", time)
}
