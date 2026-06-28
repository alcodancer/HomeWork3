package main

import (
	"fmt"
)

var chess_size = 8

func main() {
	draw_desk()
}

func draw_desk() {
	for i := 1; i <= chess_size; i++ {
		fmt.Printf("%d ", i)
		for j := 1; j <= chess_size; j++ {
			if j%2 == 0 && i%2 == 1 {
				fmt.Print("▒")
				fmt.Print("|")
			}
			if j%2 == 1 && i%2 == 1 {
				fmt.Print(" ")
				fmt.Print("|")
			}
			if j%2 == 1 && i%2 == 0 {
				fmt.Print("▒")
				fmt.Print("|")
			}
			if j%2 == 0 && i%2 == 0 {
				fmt.Print(" ")
				fmt.Print("|")
			}
		}
		fmt.Println("")
	}
}
