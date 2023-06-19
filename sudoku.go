// File name: sudoku.go
// Author: Jingyu Zhang, Lishan Li
// VUnetid: zhanj56, lil38
// Email: jingyu.zhang@vanderbilt.edu, lishan.li@vanderbilt.edu
// Class: CS3270
// Date: 12/05/2022
// Honor statement: I pledge that I have neither providing or receving help when working on this
// project Assignment Number: project 4
// Description: This is a sudoku solver write in GoLang

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
* Global variables that indicates the number of rows and columns.
 */
var ROW_NUM int = 9
var COL_NUM int = 9

/**
* The main function that executes the program.
*
* @author Jingyu Zhang
 */
func main() {

	var sudokuBoard [9][9]int

	fmt.Println("Enter the file path of the sudoku: ")
	var filePath string
	fmt.Scanln(&filePath)
	readFile(filePath, &sudokuBoard)

	fmt.Println()
	fmt.Println("here is the Sudoku you entered: ")
	printBoard(&sudokuBoard)
	fmt.Println()

	if solveSudoku(&sudokuBoard) {
		fmt.Println("The puzzle is solved! Here is the solution: ")
		printBoard(&sudokuBoard)
		fmt.Println()
	} else {
		fmt.Println("This Sudoku is unsolvable!")
		fmt.Println()
	}

}

/**
* A function that reads a file and load the sudoku into a sudokuBoard
*
* @author Jingyu Zhang
* @param  filePath  A path to the location of the file containing the sudoku.
* @param  sudokuBoard  A 9*9 array represents the sudoku board in our program.
 */
func readFile(filePath string, sudokuBoard *[9][9]int) {
	f, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < ROW_NUM; i++ {
		for j := 0; j < COL_NUM; j++ {
			scanner.Scan()
			intVar, _ := strconv.Atoi(scanner.Text())
			sudokuBoard[i][j] = intVar
		}
	}

}

/**
* A function that prints the sudokuBoard
*
* @author Lishan Li
* @param  sudokuBoard  A 9*9 array represents the sudoku board in our program.
 */
func printBoard(sudokuBoard *[9][9]int) {
	for row := 0; row < ROW_NUM; row++ {
		for col := 0; col < COL_NUM; col++ {
			fmt.Print(" ")
			fmt.Print(sudokuBoard[row][col])
			fmt.Print(" ")
			if col == 2 || col == 5 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 {
			fmt.Println("\n+--------+---------+---------+")
		} else {
			fmt.Println()
		}
	}
}

/**
* A function that solves the Sudoku, if the sudoku is insolvable, return false.
*
* @author Lishan Li
* @param  sudokuBoard  A 9*9 array represents the sudoku board in our program.
* @return a boolean value indicating whether the sudoku is solvable.
 */
func solveSudoku(sudokuBoard *[9][9]int) bool {
	var row, col int
	if isSolved(&row, &col, sudokuBoard) {
		return true
	} else {
		for option := 1; option <= 9; option++ {
			if isDoable(row, col, option, sudokuBoard) {
				sudokuBoard[row][col] = option
				if solveSudoku(sudokuBoard) {
					return true
				}
				sudokuBoard[row][col] = 0
			}
		}
	}
	return false

}

/**
* A function that checks whether the sudoku is solved, if not, set the row and column into
* the location to be solved and return false.
*
* @author Lishan Li
* @param  row  The reference to the row number where a digit waited to be solved.
* @param  col  The reference to the column number where a digit waited to be solved.
* @param  sudokuBoard  A 9*9 array represents the sudoku board in our program.
* @param  a boolean value indicating whether the puzzle is solved or not.
 */
func isSolved(row *int, col *int, sudokuBoard *[9][9]int) bool {
	for i := 0; i < ROW_NUM; i++ {
		for j := 0; j < COL_NUM; j++ {
			if sudokuBoard[i][j] == 0 {
				*row = i
				*col = j
				return false
			}
		}
	}
	return true
}

/**
* A function that checks whether the number can be placed in the current location.
*
* @author Lishan Li
* @param  row  The reference to the row number where a digit waited to be placed.
* @param  col  The reference to the column number where a digit waited to be placed.
* @param  option  The value to be checked for placing in current location.
* @param  sudokuBoard  A 9*9 array represents the sudoku board in our program.
* @param  a boolean value indicating whether the value can be placed here.
 */
func isDoable(row int, col int, option int, sudokuBoard *[9][9]int) bool {
	return checkRow(row, option, sudokuBoard) && checkCol(col, option, sudokuBoard) && checkGrid(row, col, option, sudokuBoard)
}

/**
* A function that checks whether the number can be placed in the current row.
*
* @author Jingyu Zhang
* @param  row  The reference to the row number where a digit waited to be placed.
* @param  option  The value to be checked for placing in current location.
* @param  sudokuBoard  A 9*9 array represents the sudoku board in our program.
* @param  a boolean value indicating whether the value can be placed in current row.
 */
func checkRow(row int, option int, sudokuBoard *[9][9]int) bool {
	for i := 0; i < COL_NUM; i++ {
		if sudokuBoard[row][i] == option {
			return false
		}
	}
	return true
}

/**
* A function that checks whether the number can be placed in the current column.
*
* @author Jingyu Zhang
* @param  col  The reference to the column number where a digit waited to be placed.
* @param  option  The value to be checked for placing in current location.
* @param  sudokuBoard  A 9*9 array represents the sudoku board in our program.
* @param  a boolean value indicating whether the value can be placed in current column.
 */
func checkCol(col int, option int, sudokuBoard *[9][9]int) bool {
	for i := 0; i < ROW_NUM; i++ {
		if sudokuBoard[i][col] == option {
			return false
		}
	}
	return true
}

/*
*
* A function that checks whether the number can be placed in the current grid.
*
* @author Lishan Li
* @param  row  The reference to the row number where a digit waited to be placed.
* @param  col  The reference to the column number where a digit waited to be placed.
* @param  option  The value to be checked for placing in current location.
* @param  sudokuBoard  A 9*9 array represents the sudoku board in our program.
* @param  a boolean value indicating whether the value can be placed in current grid.
 */
func checkGrid(row int, col int, option int, sudokuBoard *[9][9]int) bool {
	var startRow int = row - row%3
	var startCol int = col - col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if sudokuBoard[i][j] == option {
				return false
			}
		}
	}
	return true
}
