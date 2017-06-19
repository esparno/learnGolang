package main

import "fmt"

func main() {
	var a int = 4
	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(a)
	fmt.Println(&a)
	*b = 42

	fmt.Println(*b)
}
