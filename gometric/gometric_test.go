package gometric

import (
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

func TestCount(t *testing.T){
	var n *node
	var counted []int

	input := [][4]uint{
		[4]uint{1, 0, 0 ,0},
		[4]uint{1, 0, 0, 0},
		[4]uint{1, 0, 0, 0},
		[4]uint{0, 0, 0, 0},
	}

	for y1 := 0; y1 < 4; y1++ {
		for x1 := 0; x1 < 4; x1++ {
			if input[y1][x1] == 1 {
				n = &node{
					size: 1,
					body: 1 << uint(y1*4+x1),
				}				
				counted = helper(n, x1, y1, input, counted)
			}
		}
	}

	table := make(map[int]int)
	result := 0
	for _, e := range counted{
		if table[e] != 1{
			table[e] = 1
			result += 1
		}
	}
	if result != 3{
		t.Errorf("should be 3 but %d", result)
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

func TestCase4(t *testing.T) {

	a := [][4]uint{
		[4]uint{0, 0, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{1, 1, 0, 0},
		[4]uint{0, 1, 0, 0},
	}

	if RemoveDuplicates(Count(a)) != 7 {
		t.Errorf("should be 7 but %d", RemoveDuplicates(Count(a)))
	}
}

func TestCase5(t *testing.T) {

	a := [][4]uint{
		[4]uint{0, 0, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{1, 1, 1, 0},
		[4]uint{0, 1, 0, 0},
	}

	if RemoveDuplicates(Count(a)) != 15 {
		t.Errorf("should be 15 but %d", RemoveDuplicates(Count(a)))
	}
}


func TestCase6(t *testing.T) {

	a := [][4]uint{
		[4]uint{0, 1, 0, 0},
		[4]uint{1, 1, 0, 0},
		[4]uint{1, 1, 1, 0},
		[4]uint{0, 0, 0, 0},
	}

	if RemoveDuplicates(Count(a)) != 27{
		t.Errorf("should be 27 but %d", RemoveDuplicates(Count(a)))
	}
}
func TestCase7(t *testing.T) {

	a := [][4]uint{
		[4]uint{0, 0, 0, 0},
		[4]uint{1, 1, 0, 0},
		[4]uint{1, 1, 0, 0},
		[4]uint{0, 0, 0, 0},
	}

	if RemoveDuplicates(Count(a)) != 9{
		t.Errorf("should be 22 but %d", RemoveDuplicates(Count(a)))
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
	if !strings.EqualFold(s, s_tostring){
		t.Errorf("%s and %s should be the same", s, s_tostring)
	}
}

func TestNextTo(t *testing.T){

	a := [][4]uint{
		[4]uint{0, 0, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 0, 0},
		[4]uint{0, 1, 1, 0},
	}

	n := InitNode(a)

	if n.NextTo(0,0,a) != false{
		t.Errorf("should be false")
	}

	if n.NextTo(1,0,a) != true{
		t.Errorf("should be true")
	}

	if n.NextTo(1,1,a) != false{
		t.Errorf("should be false")
	}

}
