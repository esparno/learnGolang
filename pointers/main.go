package main

import "fmt"

func main() {
	var a int = 4
	var b *int = &a

	fmt.Println("b", b)
	fmt.Println("*b", *b)
	fmt.Println("a", a)
	fmt.Println("&a", &a)
	*b = 42

	fmt.Println(*b)
}
