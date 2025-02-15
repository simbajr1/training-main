package main

import (
	"fmt"
	"os"
)

const size = 9 // Sudoku grid size

func main() {
	// Check if exactly 9 arguments are provided
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	// Create a 9x9 Sudoku board
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}

	// Parse input arguments into the board
	for i, arg := range os.Args[1:] {
		// Check if each argument is exactly 9 characters long
		if len(arg) != size {
			fmt.Println("Error")
			return
		}
		for j, char := range arg {
			if char == '.' {
				board[i][j] = 0 // Use 0 to represent empty cells
			} else if char >= '1' && char <= '9' {
				board[i][j] = int(char - '0') // Convert character to integer
			} else {
				fmt.Println("Error") // Invalid character
				return
			}
		}
	}

	// Check if the initial board is valid
	if !isValidSudoku(board) {
		fmt.Println("Error")
		return
	}

	// Solve the Sudoku
	if !solveSudoku(board) {
		fmt.Println("Error") // If no solution exists
		return
	}

	// Print the solved Sudoku board
	printBoard(board)
}

// isValidSudoku checks if the Sudoku board is valid
func isValidSudoku(board [][]int) bool {
	// Check rows and columns for duplicates
	for i := 0; i < size; i++ {
		row := make([]bool, size+1)
		col := make([]bool, size+1)
		for j := 0; j < size; j++ {
			// Check row
			if board[i][j] != 0 {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}
			// Check column
			if board[j][i] != 0 {
				if col[board[j][i]] {
					return false
				}
				col[board[j][i]] = true
			}
		}
	}

	// Check 3x3 subgrids for duplicates
	for i := 0; i < size; i += 3 {
		for j := 0; j < size; j += 3 {
			box := make([]bool, size+1)
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					if board[x][y] != 0 {
						if box[board[x][y]] {
							return false
						}
						box[board[x][y]] = true
					}
				}
			}
		}
	}

	return true
}

// solveSudoku solves the Sudoku using backtracking
func solveSudoku(board [][]int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 { // Find an empty cell
				for num := 1; num <= size; num++ { // Try numbers 1-9
					if isSafe(board, i, j, num) { // Check if the number is safe
						board[i][j] = num // Place the number
						if solveSudoku(board) { // Recursively solve
							return true
						}
						board[i][j] = 0 // Backtrack if no solution
					}
				}
				return false // No valid number found
			}
		}
	}
	return true // Sudoku solved
}

// isSafe checks if placing a number in a cell is safe
func isSafe(board [][]int, row, col, num int) bool {
	// Check row and column
	for i := 0; i < size; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Check 3x3 subgrid
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

// printBoard prints the Sudoku board
func printBoard(board [][]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}