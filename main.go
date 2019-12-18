package main

import (
	"fmt"
	"os"
	"strconv"
)

const queen = '♕'
const box = '☐'

var simpleSolution = []int{0, 4, 7, 5, 2, 6, 1, 3}
var demoQueens = [][]int{{0, 1}, {1, 3}, {2, 0}, {3, 2}}

func main() {
	queens := getQueens(8)
	solution := nQueens(queens)
	if solution == nil {
		fmt.Println("No solution")
		return
	}
	board := generateBoard(convertToMatrix(*solution))

	fmt.Println(board)
}

func getQueens(def int) int {
	queens := def
	args := os.Args
	if len(args) < 2 {
		return queens
	}
	flag := os.Args[1]
	queens, err := strconv.Atoi(flag)
	if err != nil {
		queens = def
	}
	return queens
}

func convertToMatrix(solution []int) (matrix [][]int) {
	for x, y := range solution {
		matrix = append(matrix, []int{x, y})
	}
	return
}

func generateBoard(queens [][]int) (board string) {
	for i, loc := range queens {
		for j := range queens {
			if loc[0] == i && loc[1] == j {
				board += string(queen)
			} else {
				board += string(box)
			}
			board += " "
		}
		board += "\n"
	}
	return
}
