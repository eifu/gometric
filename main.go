package main

import (
	"github.com/eifu/gometric/gometric"
	"fmt"
	"os"
	"bufio"
)

func main() {
	// a := [][4]uint{
	// 	[4]uint{1, 1, 1, 0},
	// 	[4]uint{0, 1, 0, 0},
	// 	[4]uint{0, 0, 0, 0},
	// 	[4]uint{0, 0, 0, 0},
	// }

	var a [][4]uint
	var aa [4]uint
	reader := bufio.NewReader(os.Stdin)
	for i:= 0;i < 4; i++{
		guess, err := reader.ReadString('\n')
		if err!=nil{
			fmt.Println("aa")
			return 
		}

		aa = [4]uint{0, 0, 0, 0}
		for j := 0; j < 4;j ++{
			if guess[j] == '1'{
				aa[j] = 1
			}
		}
		a = append(a, aa)
	}

	c := gometric.Count(a)

	i :=gometric.RemoveDuplicates(c)

	fmt.Println(i)
}
