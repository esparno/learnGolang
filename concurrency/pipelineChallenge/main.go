package main

import (
	"fmt"
)

func main() {
	cn := createNumbers()
	c := factorial(cn)
	for n := range c {
		fmt.Println(n)
	}
}
func createNumbers() chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			for j := 1; j <= 10; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}
func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- fact(n)
		}

		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
