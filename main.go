package main

import (
	"github.com/eifu/gometric/gometric"
	"fmt"
	"os"
	"bufio"
)

func main() {

	var input [][4]uint
	var row [4]uint
	reader := bufio.NewReader(os.Stdin)
	for i:= 0;i < 4; i++{
		guess, err := reader.ReadString('\n')
		if err!=nil{
			fmt.Fprintf(os.Stderr, "Read error: %v\n", guess, err)
			os.Exit(1)
		}

		row = [4]uint{0, 0, 0, 0}
		for j := 0; j < 4;j ++{
			if guess[j] == '1'{
				row[j] = 1
			}
		}
		input = append(input, row)
	}

	c := gometric.Count(input)

	i :=gometric.RemoveDuplicates(c)

	fmt.Println(i)
}
