package gometric


import (
"fmt")
type node struct {
	size uint
	body uint
}

func (n *node)GetSize() uint{
	return n.size
}

func (n *node)GetBody() uint{
	return n.body
}

func Hash(n *node) int {
	return int(n.body)
}

func Dehash(i int) *node {
	var s uint = 0
	tmp := uint(i)
	for tmp > 0 {
		if int(tmp&0x1) == 1 {
			s++
		}
		tmp = tmp >> 1
	}
	return &node{
		size: s,
		body: uint(i),
	}
}

func pow(a, b int) int{
	if b == 0{
		return 1
	} else if b == 1 {
		return a
	} else if b%2 == 0{
		return pow(a,b/2)*pow(a,b/2)
	} else{
		return pow(a,b/2)*pow(a,b/2) *a
	}

}

func (n *node) setX2Y2On(x2, y2 int){
	n.body |= 1 << uint(y2*4 + x2)
	n.size += 1
}

func (n *node) setX2Y2Off(x2, y2 int){
	n.body ^= 1 << uint(y2*4 + x2)
	n.size -= 1
}

func (n *node) getX1() {
	return
}
func (n *node) getY1() {
	return
}

func InitNode(input [][4]uint) *node {

	var body uint = 0
	var size uint = 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			body |= (input[y][x]<<uint(y*4 + x))
			size += input[y][x]
		}
	}
	return &node{
		size: size,
		body: body,
	}
}

func Count(input [][4]uint) []int {

	var n *node
	var counted []int

	for y1 := 0; y1 < 4; y1++ {
		for x1 := 0; x1 < 4; x1++ {
			if input[y1][x1] == 1 {
				n = &node{
					size: 1,
					body: uint(y1*4 + x1),
				}
				counted = helper(n, x1, y1, input, counted)
			}
		}
	}
	fmt.Println(counted)

	return counted
}

func helper(n *node, x1, y1 int, input [][4]uint, counted []int) []int{
	var x2,y2 int
	// right
	x2 = x1 + 1
	y2 = y1
	if x2 < 4 && input[y2][x2] == 1 {
		n.setX2Y2On(x2, y2)
		counted = append(counted, Hash(n))
		counted = helper(n, x2, y2, input, counted)
		n.setX2Y2Off(x2,y2)
	}

	// bottom
	x2 = x1
	y2 = y1 + 1
	if y2 < 4 && input[y2][x2] == 1 {
		n.setX2Y2On(x2, y2)
		counted = append(counted, Hash(n))
		counted = helper(n, x2, y2, input, counted)
		n.setX2Y2Off(x2, y2)
	}

	return counted
}

func RemoveDuplicates(c []int) int{
	t := make(map[int]int)
	result := 0
	for _, e := range c{
		if t[e] != 1{
			t[e] = 1
			result += 1
		}
	}
	return result

}

