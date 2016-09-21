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

func tostringHelper(data int) string{
	var s string
	for i:=0; i<4; i++{
		if 0x1&data == 1{
			s += "1"
		}else{
			s += "0"
		}
		data = data >> 1
	}
	return s 
}

func ToString(i int) string{
	var s string
	s += tostringHelper(0xF&i) + "\n"
	s += tostringHelper(0xF&(i>>4)) + "\n"
	s += tostringHelper(0xF&(i>>8)) + "\n"
	s += tostringHelper(0xF&(i>>12)) + "\n"
	return s
}

func (n *node) setX2Y2On(x2, y2 int){
	n.body |= 1 << uint(y2*4 + x2)
	n.size += 1
}

func (n *node) setX2Y2Off(x2, y2 int){
	n.body ^= 1 << uint(y2*4 + x2)
	n.size -= 1
}

func (n *node) NextTo(x2, y2 int, input[][]uint) bool {

	if 0x1&(n.body>>uint(4*y2+x2))==1 {
		return false
	}

	var i, j int
	i = x2 
	j = y2 -1 
	if 0 <= j{
		if 0x1&(n.body>>uint(4*j+i))==1{
			return true
		}	
	}

	i = x2 
	j = y2 +1 
	if j < 4 {
		if  0x1&(n.body>>uint(4*j+i))==1{
			return true
		}	
	}

	i = x2 -1
	j = y2  
	if 0 <= i {
		if  0x1&(n.body>>uint(4*j+i))==1{
			return true
		}	
	}

	i = x2 +1
	j = y2  
	if i < 4{
		if  0x1&(n.body>>uint(4*j+i))==1{
			return true
		}	
	}
	return false
}

func InitNode(input [][]uint) *node {

	var body uint = 0
	var size uint = 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			body |= input[y][x]<<uint(y*4 + x)
			size += input[y][x]
		}
	}
	return &node{
		size: size,
		body: body,
	}
}

func Count(input [][]uint) []int {

	var n *node
	var counted []int

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
	return counted
}



func helper(n *node, x1, y1 int, input [][]uint, counted []int) []int{

	for y2 := 0; y2 < 4; y2 ++{
		for x2 := 0; x2 < 4; x2 ++{	
			if input[y2][x2] == 1 && n.NextTo(x2, y2, input) {
				n.setX2Y2On(x2, y2)
				counted = append(counted, Hash(n))
				counted = helper(n, x2, y2, input, counted)
				n.setX2Y2Off(x2,y2)
			}
		}
	}

	return counted
}

func RemoveDuplicates(c []int) int{
	t := make(map[int]int)
	result := 0
	for _, e := range c{
		if t[e] != 1{
			t[e] = 1
			// fmt.Printf("%#v \n%s\n",Dehash(e), ToString(e))
			result += 1
		}
	}
	return result

}

