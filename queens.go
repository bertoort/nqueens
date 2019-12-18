package main

import (
	"math"
)

func nQueens(n int) *[]int {
	board := 0
	for board < n {
		if solution := placeQueen(board, n, make([]int, n), 0); solution != nil {
			return &solution
		}
		board++
	}
	return nil
}

func placeQueen(col int, n int, solution []int, placed int) []int {
	var row int
	for row < n {
		var i int
		for i < n {
			if safe(col, row, solution) {
				solution[col] = row
				if placed == n {
					return solution
				}
				column := col + 1
				if column > n-1 {
					column = 0
				}
				if attempt := placeQueen(column, n, solution, placed+1); attempt != nil {
					return attempt
				}
			}
			i++
		}
		row++
	}
	return nil
}

func safe(col int, row int, solution []int) bool {
	var i int
	for i < col {
		if solution[i] == row || math.Abs(float64(solution[i]-row)) == math.Abs(float64(col-i)) {
			return false
		}
		i++
	}
	return true
}
