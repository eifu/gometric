package gometric

import (
	"fmt"
	"testing"
	"strings"
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

	if Hash(n) != 14{
		t.Errorf("Hash should be 14 but %d", Hash(n))
	}
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

func TestCase1(t *testing.T) {

	a := [][4]uint{
		[4]uint{0, 0, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 0, 0},
	}

	if RemoveDuplicates(Count(a)) != 3 {
		t.Errorf("should be 3 but %d", RemoveDuplicates(Count(a)))
	}

}

func TestCase2(t *testing.T) {

	a := [][4]uint{
		[4]uint{0, 0, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 1, 0},
	}

	if RemoveDuplicates(Count(a)) != 6 {
		t.Errorf("should be 6 but %d", RemoveDuplicates(Count(a)))
	}

}

func TestCase3(t *testing.T) {

	a := [][4]uint{
		[4]uint{0, 0, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 0, 1},
		[4]uint{0, 1, 1, 1},
	}

	if RemoveDuplicates(Count(a)) != 15 {
		t.Errorf("should be 15 but %d", RemoveDuplicates(Count(a)))
	}
}

func TestToString(t *testing.T){

	s := "1000\n"+"1000\n"+"1000\n"+"0000\n"
	a := [][4]uint{
		[4]uint{1, 0, 0, 0},
		[4]uint{1, 0, 0, 0},
		[4]uint{1, 0, 0, 0},
		[4]uint{0, 0, 0, 0},
	}
	s_tostring := ToString(Hash(InitNode(a)))
	fmt.Println(s)
	fmt.Println(s_tostring)
	if !strings.EqualFold(s, s_tostring){
		t.Errorf("%s and %s should be the same", s, s_tostring)
	}

}
