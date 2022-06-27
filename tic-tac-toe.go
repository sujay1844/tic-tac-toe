package main

import (
	//"bufio"
	"fmt"
	"os"
	"time"
)

func main(){
	fmt.Println("TIC TAC TOE")
	fmt.Println("To choose your square press a number from 1 to 9 like a numpad")
	time.Sleep(time.Duration(2000) * time.Millisecond)
	game := [9]int{0,0,0,0,0,0,0,0,0}
	fmt.Println(game)
	x_turn := true
	var choice int
	number_of_moves_left := 0
	for {
		// Clear screen
		fmt.Println("\033[H\033[2J")
		// Print the grid
		print_game(game)
		if x_turn {
			x_turn = false
			fmt.Print("X's turn: ")
			fmt.Scan(&choice)
			if choice >= 1 && choice <=9 {
				if game[choice-1] == 0 {
					game = edit_grid(game, choice-1, "X")
				} else {
					fmt.Println("Spot already taken. Try again")
					time.Sleep(time.Duration(1000) * time.Millisecond)
					x_turn = true
				}
			} else {
				fmt.Println("Invalid input")
				fmt.Println("There are only 9 squares, buddy")
				time.Sleep(time.Duration(1000) * time.Millisecond)
				x_turn = true
			}
		} else {
			x_turn = true
			fmt.Print("O's turn: ")
			fmt.Scan(&choice)
			if choice >= 1 && choice <=9 {
				if game[choice-1] == 0 {
					game = edit_grid(game, choice-1, "O")
				} else {
					fmt.Println("Spot already taken. Try again")
					time.Sleep(time.Duration(1000) * time.Millisecond)
					time.Sleep(1)
					x_turn = false
				}
			} else {
				fmt.Println("Invalid input")
				fmt.Println("There are only 9 squares, buddy")
				time.Sleep(time.Duration(1000) * time.Millisecond)
				x_turn = false
			}
		}

		number_of_moves_left = 0
		for i :=0 ; i<9 ; i++ {
			if game[i] == 0 { number_of_moves_left += 1 }
		}
		if number_of_moves_left == 0 {
			fmt.Println("\033[H\033[2J")
			print_game(game)
			fmt.Println("It's a tie")
			os.Exit(0)
		}
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
	if n == 1 { return "X" }
	if n == 2 { return "O" }
	return " "
}

func edit_grid(grid [9]int, pos int, player string) [9]int{
	n := 0
	if player == "X" { n = 1 }
	if player == "O" { n = 2 }
	if player == " " { n = 0 }
	grid[pos] = n
	return grid
}

func check_win(g[9]int) int{
	// Checking if O has won

	// Checking if horizontal rows are equal
	if g[0] == g[1] && g[1] == g[2] && g[2] != 0 && g[0] == 1 {
		return 1
	}
	if g[3] == g[4] && g[4] == g[5] && g[5] != 0 && g[3] == 1 {
		return 1
	}
	if g[6] == g[7] && g[7] == g[8] && g[2] != 0 && g[6] == 1 {
		return 1
	}
	// Checking if vertical columns are equal
	if g[0] == g[3] && g[3] == g[6] && g[6] != 0 && g[0] == 1 {
		return 1
	}
	if g[1] == g[4] && g[4] == g[7] && g[7] != 0 && g[1] == 1 {
		return 1
	}
	if g[2] == g[5] && g[5] == g[8] && g[8] != 0 && g[2] == 1 {
		return 1
	}
	// Checking if diagonals are equal
	if g[0] == g[4] && g[4] == g[8] && g[8] != 0 && g[0] == 1 {
		return 1
	}
	if g[2] == g[4] && g[4] == g[6] && g[6] != 0 && g[2] == 1 {
		return 1
	}

	//Checking if X has won

	// Checking if horizontal rows are equal
	if g[0] == g[1] && g[1] == g[2] && g[2] != 0 && g[0] == 1 {
		return 2
	}
	if g[3] == g[4] && g[4] == g[5] && g[5] != 0 && g[3] == 1 {
		return 2
	}
	if g[6] == g[7] && g[7] == g[8] && g[2] != 0 && g[6] == 1 {
		return 2
	}
	// Checking if vertical columns are equal
	if g[0] == g[3] && g[3] == g[6] && g[6] != 0 && g[0] == 1 {
		return 2
	}
	if g[1] == g[4] && g[4] == g[7] && g[7] != 0 && g[1] == 1 {
		return 2
	}
	if g[2] == g[5] && g[5] == g[8] && g[8] != 0 && g[2] == 1 {
		return 2
	}
	// Checking if diagonals are equal
	if g[0] == g[4] && g[4] == g[8] && g[8] != 0 && g[0] == 1 {
		return 2
	}
	if g[2] == g[4] && g[4] == g[6] && g[6] != 0 && g[2] == 1 {
		return 2
	}
	return 0
}
