package main

import (
	"fmt"
	"os"
)

// solveSudoku attempts to solve the given Sudoku board using a backtracking algorithm.
// It returns true if the board is successfully solved, otherwise false.
func solveSudoku(board [][]byte) bool {
	// Define a recursive function to attempt solving the Sudoku.
	var solve func() bool
	solve = func() bool {
		// Iterate through each cell in the 9x9 board.
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				// Check if the current cell is empty (denoted by '.').
				if board[r][c] == '.' {
					// Try placing each number from '1' to '9' in the empty cell.
					for num := byte('1'); num <= '9'; num++ {
						// Check if placing the number is valid according to Sudoku rules.
						if isValid(board, r, c, num) {
							board[r][c] = num // Place the number on the board.
							if solve() { // Recursively attempt to solve with this number placed.
								return true // If successful, return true.
							}
							board[r][c] = '.' // Reset the cell if it leads to no solution.
						}
					}
					return false // Return false if no number can be placed in this cell.
				}
			}
		}
		return true // Return true if all cells are filled correctly.
	}
	return solve() // Start the solving process.
}

// isValid checks if placing a number in a specific cell is valid according to Sudoku rules.
// It ensures that the number does not already exist in the same row, column, or 3x3 subgrid.
func isValid(board [][]byte, r, c int, num byte) bool {
	for i := 0; i < 9; i++ {
		// Check current row and column for duplicates.
		if board[r][i] == num || board[i][c] == num ||
			board[(r/3)*3+i/3][(c/3)*3+i%3] == num { // Check corresponding 3x3 subgrid.
			return false // Return false if a duplicate is found.
		}
	}
	return true // Return true if no duplicates are found.
}

// isValidInput checks if the input array conforms to expected Sudoku format.
// It verifies that there are exactly 9 lines and each line contains exactly 9 characters,
// with valid characters being '1'-'9' or '.' for empty cells.
func isValidInput(input []string) bool {
	if len(input) != 9 { // Ensure there are exactly 9 rows.
		return false
	}

	for _, line := range input {
		if len(line) != 9 { // Ensure each row has exactly 9 characters.
			return false
		}
		for _, ch := range line {
			if ch != '.' && (ch < '1' || ch > '9') { // Validate characters in each row.
				return false
			}
		}
	}

	return true // Return true if all checks pass.
}

// main function serves as the entry point of the program.
// It reads input from command line arguments and processes it to solve a Sudoku puzzle.
func main() {
	if len(os.Args) != 10 { // Check for correct number of command line arguments (1 for program name + 9 for board).
		fmt.Println("Error") // Print error message if incorrect number of arguments provided.
		return
	}

	input := os.Args[1:] // Capture input lines from command line arguments.

	if !isValidInput(input) { // Validate the input format before processing.
		fmt.Println("Error") // Print error message for invalid input format.
		return
	}

	board := make([][]byte, 9) // Create a new 2D slice to represent the Sudoku board.
	for i, line := range input {
		board[i] = []byte(line) // Convert each input string into a byte slice for processing.
	}

	if solveSudoku(board) { // Attempt to solve the Sudoku puzzle using the defined function.
		for _, row := range board { // Iterate over each row of the solved board.
			for i, num := range row {
				if i != len(row)-1 {
					fmt.Printf("%c ", num) // Print numbers with space between them except for last number in row.
				} else {
					fmt.Printf("%c", num) // Print last number without trailing space.
				}
			}
			fmt.Println() // Move to next line after printing a row of numbers.
		}
	} else {
		fmt.Println("Error") // Print error message if no solution exists for the given Sudoku puzzle.
	}
}
