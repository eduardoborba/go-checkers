package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Square struct {
	usable bool
	value  string
}

type Checkers struct {
	board [8][8]Square
	turn  string
}

func main() {
	var checkers Checkers = Checkers{
		board: [8][8]Square{
			{
				Square{false, ""}, Square{true, "Black"}, Square{false, ""},
				Square{true, "Black"}, Square{false, ""}, Square{true, "Black"},
				Square{false, ""}, Square{true, "Black"},
			},
			{
				Square{true, "Black"}, Square{false, ""}, Square{true, "Black"},
				Square{false, ""}, Square{true, "Black"}, Square{false, ""},
				Square{true, "Black"}, Square{false, ""},
			},
			{
				Square{false, ""}, Square{true, "Black"}, Square{false, ""},
				Square{true, "Black"}, Square{false, ""}, Square{true, "Black"},
				Square{false, ""}, Square{true, "Black"},
			},
			{
				Square{true, ""}, Square{false, ""}, Square{true, ""},
				Square{false, ""}, Square{true, ""}, Square{false, ""},
				Square{true, ""}, Square{false, ""},
			},
			{
				Square{false, ""}, Square{true, ""}, Square{false, ""},
				Square{true, ""}, Square{false, ""}, Square{true, ""},
				Square{false, ""}, Square{true, ""},
			},
			{
				Square{true, "White"}, Square{false, ""}, Square{true, "White"},
				Square{false, ""}, Square{true, "White"}, Square{false, ""},
				Square{true, "White"}, Square{false, ""},
			},
			{
				Square{false, ""}, Square{true, "White"}, Square{false, ""},
				Square{true, "White"}, Square{false, ""}, Square{true, "White"},
				Square{false, ""}, Square{true, "White"},
			},
			{
				Square{true, "White"}, Square{false, ""}, Square{true, "White"},
				Square{false, ""}, Square{true, "White"}, Square{false, ""},
				Square{true, "White"}, Square{false, ""},
			},
		},
		turn: "White",
	}

	for checkers.gameNotFinished() {
		checkers.printBoard()

		readMove(checkers)
	}
}

func (c Checkers) printBoard() {
	clearTerminal()

	fmt.Println(". _ . _ . _ . _ . _ . _ . _ . _ .")
	for i := 0; i < 8; i += 1 {
		for j := 0; j < 8; j += 1 {
			fmt.Print("| ", valueToPrint(c.board[i][j]), " ")
		}

		fmt.Println("|")
		fmt.Println(". _ . _ . _ . _ . _ . _ . _ . _ .")
	}
}

func clearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func valueToPrint(square Square) string {
	value_to_print := square.value

	if value_to_print == "" {
		value_to_print = " "
	} else {
		value_to_print = value_to_print[0:1]
	}

	return value_to_print
}

func (c Checkers) gameNotFinished() bool {
	has_black := false
	has_white := false

	for i := 0; i < 8; i += 1 {
		for j := 0; j < 8; j += 1 {
			if c.board[i][j].value == "Black" {
				has_black = true
			}

			if c.board[i][j].value == "White" {
				has_white = true
			}
		}
	}

	return has_black && has_white
}

func readMove(c Checkers) {
	var from_row, from_column, to_row, to_column int

	fmt.Print("From row: ")
	fmt.Scanf("%d", &from_row)
	fmt.Print("From column: ")
	fmt.Scanf("%d", &from_column)

	fmt.Print("To row: ")
	fmt.Scanf("%d", &to_row)
	fmt.Print("To column: ")
	fmt.Scanf("%d", &to_column)

	from := c.board[from_row][from_column]
	to := c.board[to_row][to_column]

	fmt.Println(from_row, from_column, from.value)
	fmt.Println(to_row, to_column, to.value)
	if !from.usable || !to.usable {
		print("Not valid")
	}
}
