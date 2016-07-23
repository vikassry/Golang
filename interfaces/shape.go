package interfaces

type Shape interface {
	Area() int
}

type Square struct {
	side int
}

func (s Square) Area() int {
	return s.side * s.side
}

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) Area() int {
	return r.length * r.breadth
}

type Hybrid struct {
	s Shape
	r Shape
}

func (h Hybrid) Area() int {
	return h.s.Area() + h.r.Area()
}
