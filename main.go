package main

import "fmt"

const queen = '♕'
const box = '☐'

var simpleSolution = []int{0, 4, 7, 5, 2, 6, 1, 3}
var demoQueens = [][]int{{0, 1}, {1, 3}, {2, 0}, {3, 2}}

func main() {
	eightBoard := generateBoard(convertToMatrix(simpleSolution))
	fourBoard := generateBoard(demoQueens)

	fmt.Println(eightBoard)
	fmt.Println(fourBoard)
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
