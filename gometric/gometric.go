package gometric



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

func Hash(n *node) uint {
	return n.body
}

func Dehash(i uint) *node {
	var s uint = 0
	tmp := i
	for tmp > 0 {
		if int(tmp&0x1) == 1 {
			s++
		}
		tmp = tmp >> 1
	}
	return &node{
		size: s,
		body: i,
	}
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

// func Solve(input [][4]int) int {

// 	var n *node
// 	counted := make(map[int]int)

// 	for y1 := 0; y1 < 4; y1++ {
// 		for x1 := 0; x1 < 4; x1++ {
// 			if input[y1][x1] == 1 {
// 				n = &node{
// 					size: 1,
// 					body: uint(y1*4 + x1),
// 				}
// 				counted = helper(n, x1, y1, input, counted)
// 			}
// 		}
// 	}
// 	fmt.Println(counted)

// 	result := 0

// 	return result
// }

// func helper(n *node, x1, y1 int, input [][4]int, counted map[int]int) map[int]int {
// 	var x2, y2 int

// 	// right
// 	x2 = x1 + 1
// 	y2 = y1
// 	if x2 < 4 && input[y2][x2] == 1 {
// 		n.setX2Y2(x2, y2)
// 		counted[Hash(n)] = 1
// 		//fmt.Printf("%#v\n", *n)
// 		counted = helper(n, x2, y2, input, counted)
// 	}

// 	// bottom
// 	x2 = x1
// 	y2 = y1 + 1
// 	if y2 < 4 && input[y2][x2] == 1 {
// 		n.setX2Y2(x2, y2)
// 		counted[Hash(n)] = 1
// 		//fmt.Printf("%#v\n", *n)
// 		counted = helper(n, x2, y2, input, counted)
// 	}
// 	return counted
// }
