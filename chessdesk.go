package main

import (
	"fmt"
)

var chess_size = 7

func main() {
	draw_desk()
}
func draw_desk() {
	k := " |▒|"
	for i := 1; i <= chess_size; i++ {
		if k == "▒| |" {
			k = " |▒|"
		} else {
			k = "▒| |"
		}
		for j := 1; j <= chess_size/2; j++ {
			fmt.Print(k)
		}
		fmt.Println("")
	}
}
