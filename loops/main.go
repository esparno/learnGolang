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
}
