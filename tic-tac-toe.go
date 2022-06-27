package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("TIC TAC TOE")
	fmt.Println("To choose your square press a number from 1 to 9 like a numpad")
	//game := [3][3]int{
		//{0,0,0},
		//{0,0,0},
		//{0,0,0},
	//}
	game := [9]int{0,0,0,0,0,0,0,0,0}
	fmt.Println(game)
	x_turn := true
	for {
		// Delay 500ms between changes
		time.Sleep(time.Duration(500) * time.Millisecond)
		if x_turn {
			x_turn = false
			fmt.Println("\033[H\033[2J")
			fmt.Println("X's turn")
		} else {
			x_turn = true
			fmt.Println("\033[H\033[2J")
			fmt.Println("O's turn")
		}
	}
}

func check_win(grid []int) int{
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
