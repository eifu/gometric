package gometric

import (

"fmt"
	"testing"
)

func TestHash(t *testing.T) {

	n := InitNode([][4]uint{
						[4]uint{0, 1, 1, 1},
						[4]uint{0, 0, 0, 0},
						[4]uint{0, 0, 0, 0},
						[4]uint{0, 0, 0, 0},
						})

	if Hash(n) != Hash(Dehash(Hash(n))) {
		t.Errorf("%d should be the same but ", Hash(n))
	}

}

func TestHash2(t *testing.T){
	n := InitNode([][4]uint{
						[4]uint{0, 1, 1, 1},
						[4]uint{0, 0, 0, 0},
						[4]uint{0, 0, 0, 0},
						[4]uint{0, 0, 0, 0},
						})

	fmt.Println(Hash(n))
}

func TestSetX2Y2On(t *testing.T){

	n := InitNode([][4]uint{
					[4]uint{0, 1, 1, 0},
					[4]uint{0, 0, 0, 0},
					[4]uint{0, 0, 0, 0},
					[4]uint{0, 0, 0, 0},
					})
	n.setX2Y2On(3,0)

	if Hash(n) != 14 {
		t.Errorf("Hash should be 14, but %d",Hash(n))
	}
}

func TestSetX2Y2Off(t *testing.T){

	n := InitNode([][4]uint{
					[4]uint{0, 1, 1, 1},
					[4]uint{1, 0, 0, 0},
					[4]uint{0, 0, 0, 0},
					[4]uint{0, 0, 0, 0},
					})
	// fmt.Println(Hash(n))
	n.setX2Y2Off(0,1)
	if Hash(n) != 14 {
		t.Errorf("Hash should be 14, but %d", Hash(n))
	}
}

// func TestCase1(t *testing.T) {

// 	a := [][4]int{
// 		[4]int{0, 0, 0, 0},
// 		[4]int{0, 1, 0, 0},
// 		[4]int{0, 1, 0, 0},
// 		[4]int{0, 1, 0, 0},
// 	}

// 	if Solve(a) != 3 {
// 		t.Errorf("should be 3 but %d", a)
// 	}

// }

// func TestCase2(t *testing.T) {

// 	a := [][4]int{
// 		[4]int{0, 0, 0, 0},
// 		[4]int{0, 1, 0, 0},
// 		[4]int{0, 1, 0, 0},
// 		[4]int{0, 1, 1, 0},
// 	}

// 	if Solve(a) != 6 {
// 		t.Errorf("should be 6 but %d", a)
// 	}

// }

// func TestCase3(t *testing.T) {

// 	a := [][4]int{
// 		[4]int{0, 0, 0, 0},
// 		[4]int{0, 1, 0, 0},
// 		[4]int{0, 1, 0, 1},
// 		[4]int{0, 1, 1, 1},
// 	}

// 	if Solve(a) != 15 {
// 		t.Errorf("should be 15 but %d", Solve(a))
// 	}

// }
