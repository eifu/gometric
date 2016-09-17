package gometric

import (
	"testing"
)

func TestHash(t *testing.T) {

	n := &node{
		X1: 0, Y1: 0, X2: 3, Y2: 2}

	if Hash(n) != Hash(Dehash(Hash(n))) {
		t.Errorf("%d should be the same but ", Hash(n))
	}

}

func TestCase1(t *testing.T) {

	a := [][4]int{
		[4]int{0, 0, 0, 0},
		[4]int{0, 1, 0, 0},
		[4]int{0, 1, 0, 0},
		[4]int{0, 1, 0, 0},
	}

	if Solve(a) != 3 {
		t.Errorf("should be 3 but %d", a)
	}

}

func TestCase2(t *testing.T) {

	a := [][4]int{
		[4]int{0, 0, 0, 0},
		[4]int{0, 1, 0, 0},
		[4]int{0, 1, 0, 0},
		[4]int{0, 1, 1, 0},
	}

	if Solve(a) != 6 {
		t.Errorf("should be 6 but %d", a)
	}

}

func TestCase3(t *testing.T) {

	a := [][4]int{
		[4]int{0, 0, 0, 0},
		[4]int{0, 1, 0, 0},
		[4]int{0, 1, 0, 1},
		[4]int{0, 1, 1, 1},
	}

	if Solve(a) != 15 {
		t.Errorf("should be 15 but %d", Solve(a))
	}

}
