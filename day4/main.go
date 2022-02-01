package main

import "fmt"

type Square struct {
	number int
	hit    bool
}
type Board struct {
	Rows [][]Square
}

func (b Board) HasWon() (won bool) {
	//check horizontal
	for row := 0; row < len(b.Rows); row++ {
		isWinningRow := true
		for column := 0; column < len(b.Rows[row]); column++ {
			if b.Rows[row][column].hit == false {
				isWinningRow = false
			}
		}
		if isWinningRow == true {
			fmt.Println("Winning Row", row)
			return true
		}
	}

	//check vertical
	for column := 0; column < len(b.Rows); column++ {
		isWinningColumn := true
		for row := 0; row < len(b.Rows[row]); row++ {
			if b.Rows[row][column].hit == true {
				isWinningColumn = false
			}
		}
		if isWinningColumn == true {
			fmt.Println("Winning Column", column)
		}
	}

	return false
}

func main() {
	numbers := LineToIntArray("input.txt")
	fmt.Println(numbers)
}
