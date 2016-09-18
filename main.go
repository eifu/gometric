package main

import (
	"github.com/eifu/gometric/gometric"
	"fmt"
)

func main() {
	a := [][4]uint{
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 1, 0},
	}

	c := gometric.Count(a)
	// fmt.Println(c)
	i :=gometric.RemoveDuplicates(c)
	fmt.Println(i)
}
