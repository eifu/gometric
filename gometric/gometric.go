package gometric

type node struct {
	X1 int // X1, Y1 are top left coordinates
	Y1 int
	X2 int // X2, Y2 are bottum right coordinates
	Y2 int
}

func Hash(n *node) int {
	return n.X1*64 + n.Y1*16 + n.X2*4 + n.Y2

}

func Dehash(i int) *node {

	x1 := i / 64
	i = i % 64
	y1 := i / 16
	i = i % 16
	x2 := i / 4
	i = i % 4
	y2 := i
	return &node{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}

}

func (n *node) setX2Y2(x2, y2 int) {
	n.X2 = x2
	n.Y2 = y2
}

func Solve(input [][4]int) int {

	var n *node
	counted := make(map[int]int)

	for y1 := 0; y1 < 4; y1++ {
		for x1 := 0; x1 < 4; x1++ {
			if input[y1][x1] == 1 {
				n = &node{
					X1: x1,
					Y1: y1,
				}
				counted = helper(n, x1, y1, input, counted)
			}
		}
	}

	result := 0
	var x1, x2, y1, y2 int
	for y1 = 0; y1 < 4; y1++ {
		for x1 = 0; x1 < 4; x1++ {
			for y2 = y1; y2 < 4; y2++ {
				for x2 = x1; x2 < 4; x2++ {
					if y2 == y1 && x2 == x1 {
						continue
					}
					n = &node{
						X1: x1,
						Y1: y1,
						X2: x2,
						Y2: y2,
					}
					if counted[Hash(n)] == 1 {
						result++
					}
				}
			}

		}
	}
	return result
}

func helper(n *node, x1, y1 int, input [][4]int, counted map[int]int) map[int]int {
	var x2, y2 int

	// left
	x2 = x1 + 1
	y2 = y1
	if x2 < 4 && input[y2][x2] == 1 {
		n.setX2Y2(x2, y2)
		counted[Hash(n)] = 1
		//fmt.Printf("%#v\n", *n)
		counted = helper(n, x2, y2, input, counted)
	}

	// bottom
	x2 = x1
	y2 = y1 + 1
	if y2 < 4 && input[y2][x2] == 1 {
		n.setX2Y2(x2, y2)
		counted[Hash(n)] = 1
		//fmt.Printf("%#v\n", *n)
		counted = helper(n, x2, y2, input, counted)
	}
	return counted
}
