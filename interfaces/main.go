package main

import "fmt"

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

type Shape interface {
	Area() float64
}

func info(z Shape) {
	fmt.Println(z.Area())
}

func main() {
	var sq = Square{4}
	info(sq)
}
