package main

import (
	"github.com/eifu/gometric/gometric"
	"fmt"
)

func main() {
	a := [][4]uint{
		[4]uint{0, 1, 1, 1},
		[4]uint{0, 0, 0, 0},
		[4]uint{0, 0, 0, 0},
		[4]uint{0, 0, 0, 0},
	}

	n := gometric.InitNode(a)
	fmt.Println(n.GetSize())
	fmt.Println(n.GetBody())

}
