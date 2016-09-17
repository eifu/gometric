package gometric

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {

	n := &node{
		X1: 0, Y1: 0, X2: 3, Y2: 2}

	fmt.Println(n)
	fmt.Println(Dehash(Hash(n)))

}
