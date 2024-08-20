package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const HEIGHT = 10
const WIDTH = 30

// This function clears the terminal screen.
func ClearScreen() {

	args := []string{"clear"}
	if runtime.GOOS == "windows" {
		args = []string{"clear"}
		args = []string{"cmd", "/c", "cls"}
	}

	cmd := exec.Command("./", args...)
	if cmd.Err != nil {
		return
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}

func CleanThisshit() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func InitBoard(board *[HEIGHT][WIDTH]rune) {
	for j := 0; j < HEIGHT; j++ {
		for i := 0; i < WIDTH; i++ {
			board[j][i] = '.'
		}
	}
}

func main() {
	board := [HEIGHT][WIDTH]rune{}

	InitBoard(&board)

	for true {
		for j := 0; j < HEIGHT; j++ {
			for i := 0; i < WIDTH; i++ {
				fmt.Printf("%c", board[j][i])
			}
			fmt.Println()
		}
		CleanThisshit()
		time.Sleep(500)
	}
}
