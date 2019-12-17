package main

import "fmt"

const queen = '♕'
const box = '☐'

var demoQueens = [][]int{{0, 1}, {1, 3}, {2, 0}, {3, 2}}

func main() {
	board := generateBoard(demoQueens)
	fmt.Println(board)
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
