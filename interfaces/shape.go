package interfaces

type Shape interface {
	Area() int
}

type Hybrids []Shape

type Square struct {
	side int
}

func (square Square) Area() int {
	return square.side * square.side
}

type Rectangle struct {
	length  int
	breadth int
}

func (rectangle Rectangle) Area() int {
	return rectangle.length * rectangle.breadth
}

type Hybrid struct {
	square Shape
	rectangle Shape
}

func (hybrid Hybrid) Area() int {
	return hybrid.square.Area() + hybrid.rectangle.Area()
}

func (hybrids Hybrids) Area() int {
	area := 0
	for _,shape := range hybrids {
		area += shape.Area()
	}
	return area
}
