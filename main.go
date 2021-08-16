package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

const (
	height = 6                    // Atleast 3
	width  = height * 4           // Atleast 12
	delay  = time.Second / height // Dynamic Delay According To Height
)

func main() {
	x := 0
	y := 0
	x_velocity := 1
	y_velocity := 1
	first := true
	_ = x + y + x_velocity + y_velocity
	for {
		board := [height][width]bool{}
		board[y][x] = true
		screen.Clear()
		screen.MoveTopLeft()
		for i, row := range board {
			for j, col := range row {
				if ((i == 0 && j == 0) || (i == height-1 && j == 0) || (i == height-1 && j == width-1) || (i == 0 && j == width-1)) && !col {
					fmt.Printf("*") // Corner
				} else if (j == 0 || j == width-1) && !col {
					fmt.Printf("|") // Vertical Side
				} else if (i == 0 || i == height-1) && !col {
					fmt.Printf("-") // Horizontal Side
				} else if !col {
					fmt.Printf(" ")
				} else {
					fmt.Printf("O")
				}
			}
			fmt.Println()
		}
		fmt.Printf("\nX: %d, Y: %d\n", x, y)
		fmt.Printf("\nX Velocity: %d, Y Velocity: %d\n", x_velocity, y_velocity)
		if y == height-1 || (y == 0 && !first) {
			y_velocity *= -1
		} else if x == width-1 || (x == 0 && !first) {
			x_velocity *= -1
		}
		if first {
			first = !first
		}
		x += x_velocity
		y += y_velocity
		time.Sleep(delay)
	}
}
