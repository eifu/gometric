package main

import (
	"github.com/eifu/gometric/gometric"
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

func main() {

	var height, width int 

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("height?")
	guess, err := reader.ReadString('\n')
	if err!=nil{
		fmt.Fprintf(os.Stderr, "Read error: %v\n", guess, err)
		os.Exit(1)
	}

	guess = strings.Replace(guess, "\n", "", -1)

	if height, err = strconv.Atoi(guess); err!=nil{
		fmt.Fprintf(os.Stderr, "Conv error: %v\n", guess, err)
		os.Exit(1)
	}

	fmt.Println("width?")
	guess, err = reader.ReadString('\n')
	if err!=nil{
		fmt.Fprintf(os.Stderr, "Read error: %v\n", guess, err)
		os.Exit(1)
	}

	guess = strings.Replace(guess, "\n", "", -1)

	if width, err = strconv.Atoi(guess); err!=nil{
		fmt.Fprintf(os.Stderr, "Conv error: %v\n", guess, err)
		os.Exit(1)
	}


	var input [][]uint
	var row []uint

	for y:= 0; y < height; y++{
		guess, err = reader.ReadString('\n')
		if err!=nil{
			fmt.Fprintf(os.Stderr, "Read error: %v\n", guess, err)
			os.Exit(1)
		}

		row = []uint{}
		for x := 0; x < width; x ++{
			if guess[x] == '1'{
				row = append(row, 1)
			} else{
				row = append(row, 0)
			}
		}
		input = append(input, row)
	}

	c := gometric.Count(input)

	i :=gometric.RemoveDuplicates(c)

	fmt.Println(i)
}
