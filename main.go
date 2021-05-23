package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var Board = [8][8]string{
	{"-", "B", "-", "B", "-", "B", "-", "B"},
	{"B", "-", "B", "-", "B", "-", "B", "-"},
	{"-", "B", "-", "B", "-", "B", "-", "B"},
	{"0", "-", "0", "-", "0", "-", "0", "-"},
	{"-", "0", "-", "0", "-", "0", "-", "0"},
	{"W", "-", "W", "-", "W", "-", "W", "-"},
	{"-", "W", "-", "W", "-", "W", "-", "W"},
	{"W", "-", "W", "-", "W", "-", "W", "-"},
}

var Turn = "W"

func main() {
	for gameNotFinished() {
		printBoard()

		readMove()
	}
}

func printBoard() {
	clearTerminal()

	fmt.Println(". _ . _ . _ . _ . _ . _ . _ . _ .")
	for _, row := range Board {
		for _, value := range row {
			fmt.Print("| ", valueToPrint(value), " ")
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

func valueToPrint(value string) string {
	value_to_print := value

	if value == "-" || value == "0" {
		value_to_print = " "
	}

	return value_to_print
}

func gameNotFinished() bool {
	has_black := false
	has_white := false

	for _, row := range Board {
		for _, value := range row {
			if value == "B" {
				has_black = true
			}

			if value == "W" {
				has_white = true
			}
		}
	}

	return has_black && has_white
}

func readMove() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Piece: ")
	piece, _ := reader.ReadString('\n')
	// piece = strings.Split(piece, " ")
	// if Board[x][y] != Turn {
	// 	return
	// }

	// fmt.Print("Move to: ")
	// move_to, _ := reader.ReadString('\n')
}
