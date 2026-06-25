package main

import "fmt"

func main() {
	var chess_size int
	chess_size = 8
	for i := 1; i <= chess_size; i++ {
		for j := 1; j <= chess_size/2; j++ {
			if i%2 == 0 {
				fmt.Print("# ")
			}
			if i%2 == 1 {
				fmt.Print(" #")
			}
		}
		fmt.Println("")
	}
}
