package main

import "fmt"

func main() {

	// traditional for loop
	for i:=1; i<=10; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j<=10 {
		fmt.Println(j)
		j+=2
	}

	// do while loop
	k := 0
	for {
		fmt.Println(k)
		if k >=12 {
			break
		}
		k+= 3
	}

	l := 2

	switch l {
	case 1 :
		fmt.Println("switch case: 1")
	case 2 :
		fmt.Println("switch case: 2")
	case 3 :
		fmt.Println("switch case: 3")
	case 4 :
		fmt.Println("switch case: 4")
	default:
		fmt.Println("switch case: default")
	}
}
