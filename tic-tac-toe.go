package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("TIC TAC TOE")
	fmt.Println("To choose your square press a number from 1 to 9 like a numpad")
	game := [9]int{0,0,0,0,0,0,0,0,0}
	fmt.Println(game)
	x_turn := true
	for {
		// Delay 500ms between changes
		time.Sleep(time.Duration(500) * time.Millisecond)
		if x_turn {
			x_turn = false
			fmt.Println("X's turn")
		} else {
			x_turn = true
			fmt.Println("O's turn")
		}
		print_game(game)
		//fmt.Println("\033[H\033[2J")
	}
}

func print_game(grid [9]int) {
	fmt.Println("+-----+-----+-----+")
	fmt.Println("|  " + disp(grid[0]) + "  |  " + disp(grid[1]) + "  |  " + disp(grid[2]) + "  |")
	fmt.Println("+-----+-----+-----+")
	fmt.Println("|  " + disp(grid[3]) + "  |  " + disp(grid[4]) + "  |  " + disp(grid[5]) + "  |")
	fmt.Println("+-----+-----+-----+")
	fmt.Println("|  " + disp(grid[6]) + "  |  " + disp(grid[7]) + "  |  " + disp(grid[8]) + "  |")
	fmt.Println("+-----+-----+-----+")
}

func disp(n int) string{
	if n == 0 { return " " }
	if n == 1 { return "X" }
	if n == 2 { return "O" }
	return "error"
}

func check_win(grid [9]int) int{
	// Checking if O has won

	// Checking if horizontal rows are equal
	if (grid[0] == grid[1] && grid[1] == grid[2] && grid[2] != 0 && grid[0] == 1) {
		return 1
	}
	if (grid[3] == grid[4] && grid[4] == grid[5] && grid[5] != 0 && grid[3] == 1) {
		return 1
	}
	if (grid[6] == grid[7] && grid[7] == grid[8] && grid[2] != 0 && grid[6] == 1) {
		return 1
	}
	// Checking if vertical columns are equal
	if (grid[0] == grid[3] && grid[3] == grid[6] && grid[6] != 0 && grid[0] == 1) {
		return 1
	}
	if (grid[1] == grid[4] && grid[4] == grid[7] && grid[7] != 0 && grid[1] == 1) {
		return 1
	}
	if (grid[2] == grid[5] && grid[5] == grid[8] && grid[8] != 0 && grid[2] == 1) {
		return 1
	}
	// Checking if diagonals are equal
	if (grid[0] == grid[4] && grid[4] == grid[8] && grid[8] != 0 && grid[0] == 1) {
		return 1
	}
	if (grid[2] == grid[4] && grid[4] == grid[6] && grid[6] != 0 && grid[2] == 1) {
		return 1
	}

	//Checking if X has won

	// Checking if horizontal rows are equal
	if (grid[0] == grid[1] && grid[1] == grid[2] && grid[2] != 0 && grid[0] == 1) {
		return 2
	}
	if (grid[3] == grid[4] && grid[4] == grid[5] && grid[5] != 0 && grid[3] == 1) {
		return 2
	}
	if (grid[6] == grid[7] && grid[7] == grid[8] && grid[2] != 0 && grid[6] == 1) {
		return 2
	}
	// Checking if vertical columns are equal
	if (grid[0] == grid[3] && grid[3] == grid[6] && grid[6] != 0 && grid[0] == 1) {
		return 2
	}
	if (grid[1] == grid[4] && grid[4] == grid[7] && grid[7] != 0 && grid[1] == 1) {
		return 2
	}
	if (grid[2] == grid[5] && grid[5] == grid[8] && grid[8] != 0 && grid[2] == 1) {
		return 2
	}
	// Checking if diagonals are equal
	if (grid[0] == grid[4] && grid[4] == grid[8] && grid[8] != 0 && grid[0] == 1) {
		return 2
	}
	if (grid[2] == grid[4] && grid[4] == grid[6] && grid[6] != 0 && grid[2] == 1) {
		return 2
	}
	return 0
}
