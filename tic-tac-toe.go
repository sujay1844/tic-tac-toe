///usr/bin/true; exec /usr/bin/env go run "$0" "$@"

package main

import (
	"fmt"
	"os"
	"time"
)

func main(){

	// Variable initializing
	game := [9]int{0,0,0,0,0,0,0,0,0}
	x_turn := bool_randomizer()
	var mode_choice int
	number_of_moves_left := 0

	// Choice for single player mode
	fmt.Println("1 for player VS computer (Beta)")
	fmt.Println("2 for player VS player")
	fmt.Scanln(&mode_choice)

	player_is_x := bool_randomizer()
	level := 0
	if mode_choice == 2 {
		fmt.Println("Select level")
		fmt.Println("1 for easy")
		fmt.Println("2 for medium")
		fmt.Println("3 for hard")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&level)
	}

	for {
		clrscr()

		// Print the grid
		print_game(game)

		if mode_choice == 2 {
			x_turn, game  = pvp(x_turn, game)
		} else {
			x_turn, game = single_player(x_turn, game, player_is_x, level)
		}

		// Counting number of moves left
		number_of_moves_left = 0
		for i :=0 ; i<9 ; i++ {
			if game[i] == 0 { number_of_moves_left += 1 }
		}

		if check_win(game) == 1 {
			// X won
			clrscr()
			print_game(game)
			fmt.Println("X won the game")
			os.Exit(0)
		} else if check_win(game) == 2 {
			// O won
			clrscr()
			print_game(game)
			fmt.Println("O won the game")
			os.Exit(0)
		} else if number_of_moves_left == 0 {
			// Tie match
			clrscr()
			print_game(game)
			fmt.Println("It's a tie")
			os.Exit(0)
		}
	}
}

func print_game(g [9]int) {
	// Using g for grid to shorten code
	fmt.Println("TIC TAC TOE")
	fmt.Println("To choose your square press a number from 1 to 9 like a numpad")
	fmt.Println("Enter your chosen number and press enter")
	fmt.Println("					Reference")
	fmt.Println("╭─────┬─────┬─────╮		╭─────┬─────┬─────╮")
	fmt.Println("│  " + disp(g[0]) + "  │  " + disp(g[1]) + "  │  " + disp(g[2]) + "  │		│  1  │  2  │  3  │")
	fmt.Println("├─────────────────┤		├─────────────────┤")
	fmt.Println("│  " + disp(g[3]) + "  │  " + disp(g[4]) + "  │  " + disp(g[5]) + "  │		│  4  │  5  │  6  │")
	fmt.Println("├─────────────────┤		├─────────────────┤")
	fmt.Println("│  " + disp(g[6]) + "  │  " + disp(g[7]) + "  │  " + disp(g[8]) + "  │		│  7  │  8  │  9  │")
	fmt.Println("╰─────┴─────┴─────╯		╰─────┴─────┴─────╯")
}

func disp(n int) string{
	if n == 1 { return "X" }
	if n == 2 { return "O" }
	return " "
}

func clrscr() {
	// Clearing the terminal with ANSI Escape codes
	fmt.Println("\033[H\033[2J")
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
	// Using g for grid to shorten code

	//----------------------//
	// Checking if X has won//
	//----------------------//

	// Checking if horizontal rows are equal
	if g[0] == g[1] && g[1] == g[2] && g[2] != 0 && g[0] == 1 { return 1 } else
	if g[3] == g[4] && g[4] == g[5] && g[5] != 0 && g[3] == 1 { return 1 } else
	if g[6] == g[7] && g[7] == g[8] && g[2] != 0 && g[6] == 1 { return 1 } else
	// Checking if vertical columns are equal
	if g[0] == g[3] && g[3] == g[6] && g[6] != 0 && g[0] == 1 { return 1 } else
	if g[1] == g[4] && g[4] == g[7] && g[7] != 0 && g[1] == 1 { return 1 } else
	if g[2] == g[5] && g[5] == g[8] && g[8] != 0 && g[2] == 1 { return 1 } else
	// Checking if diagonals are equal
	if g[0] == g[4] && g[4] == g[8] && g[8] != 0 && g[0] == 1 { return 1 } else
	if g[2] == g[4] && g[4] == g[6] && g[6] != 0 && g[2] == 1 { return 1 } else

	//---------------------//
	//Checking if O has won//
	//---------------------//

	// Checking if horizontal rows are equal
	if g[0] == g[1] && g[1] == g[2] && g[2] != 0 && g[0] == 2 { return 2 } else
	if g[3] == g[4] && g[4] == g[5] && g[5] != 0 && g[3] == 2 { return 2 } else if g[6] == g[7] && g[7] == g[8] && g[2] != 0 && g[6] == 2 { return 2 } else
	// Checking if vertical columns are equal
	if g[0] == g[3] && g[3] == g[6] && g[6] != 0 && g[0] == 2 { return 2 } else
	if g[1] == g[4] && g[4] == g[7] && g[7] != 0 && g[1] == 2 { return 2 } else
	if g[2] == g[5] && g[5] == g[8] && g[8] != 0 && g[2] == 2 { return 2 } else
	// Checking if diagonals are equal
	if g[0] == g[4] && g[4] == g[8] && g[8] != 0 && g[0] == 2 { return 2 } else
	if g[2] == g[4] && g[4] == g[6] && g[6] != 0 && g[2] == 2 { return 2 }

	return 0
}

func bool_randomizer() bool{
	now := time.Now()
	x := now.Nanosecond()
	if x%2 == 0 { return true } else { return false }
}

func pvp(x_turn bool, game[9] int) (bool, [9]int){
	choice := 5
	if x_turn {
		// toggle x_turn to give to the other player
		x_turn = false

		//Taking user input
		fmt.Print("X's turn: ")
		fmt.Scan(&choice)
		if choice >= 1 && choice <=9 {

			// Making sure that the chosen square is empty
			if game[choice-1] == 0 {
				game = edit_grid(game, choice-1, "X")
			} else {
				fmt.Println("Spot already taken. Try again")
				time.Sleep(time.Duration(1000) * time.Millisecond)
				// Giving the player another chance if input was wrong
				x_turn = true
			}
		} else {
			fmt.Println("Invalid input")
			fmt.Println("There are only 9 squares, buddy")
			time.Sleep(time.Duration(1000) * time.Millisecond)
			// Giving the player another chance if input was wrong
			x_turn = true
		}
	} else {
		// toggle x_turn to give to the other player
		x_turn = true
		fmt.Print("O's turn: ")
		fmt.Scan(&choice)
		if choice >= 1 && choice <=9 {
			// Making sure that the chosen square is empty
			if game[choice-1] == 0 {
				game = edit_grid(game, choice-1, "O")
			} else {
				fmt.Println("Spot already taken. Try again")
				time.Sleep(time.Duration(1000) * time.Millisecond)
				// Giving the player another chance if input was wrong
				x_turn = false
			}
		} else {
			fmt.Println("Invalid input")
			fmt.Println("There are only 9 squares, buddy")
			time.Sleep(time.Duration(1000) * time.Millisecond)
			// Giving the player another chance if input was wrong
			x_turn = false
		}
	}
	return x_turn, game
}

func single_player(x_turn bool, game [9]int, player_is_x bool, level int) (bool, [9]int) {

	 

	return x_turn, game
}

func easy(game [9]int) int{
	// For easy mode, a random spot is chosen
	choice := 0
	for {
		random_slot := now.Nanosecond()%10
		if random_slot >=1 && random_slot <=9 {
			if game[random_slot] == 0 {
				choice = random_slot
			}
		}
	}
	return choice
}
