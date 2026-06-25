package main

import "fmt"

func main() {
	var chess_size int
	chess_size = 8
	for i := 1; i <= chess_size; i++ {
		for j := 1; j <= chess_size; j++ {
			if j%2 == 0 && i%2 == 1 {
				fmt.Print("⬛")
			}
			if j%2 == 1 && i%2 == 1 {
				fmt.Print("⬜")
			}
			if j%2 == 1 && i%2 == 0 {
				fmt.Print("⬛")
			}
			if j%2 == 0 && i%2 == 0 {
				fmt.Print("⬜")
			}
		}
		fmt.Println("")
	}
}
