package main

import (
	"./gometric"
	"fmt"
)

func main() {
	a := [][4]int{
		[4]int{0, 1, 1, 1},
		[4]int{1, 0, 0, 0},
		[4]int{1, 0, 0, 1},
		[4]int{1, 1, 0, 1},
	}

	fmt.Println(gometric.Solve(a))

}
