package main

import "fmt"

func main() {

	// create slice with 3 values
	var mySlice1 []int = []int{2, 7, 4}
	mySlice1 = append(mySlice1, 5, 6, 7)
	fmt.Println(mySlice1)

	// slice created with 5 empty values
	mySlice2 := make([]int, 5)
	for i := 0; i < len(mySlice2); i++ {
		mySlice2[i] = i + 2
	}
	fmt.Println(mySlice2)

	// slice w/ 3 values
	mySlice3 := []int{2, 4, 6}
	fmt.Println(mySlice3)

	// array with length 0 and capacity 10
	mySlice4 := make([]int, 0, 10)
	mySlice4 = append(mySlice4, 4, 8, 10, 12, 14, 16)
	fmt.Println(mySlice4)

	//create a slice from an array
	var arr = [5]int{1, 2, 3, 4, 5}
	var mySlice5 = arr[:]
	fmt.Println(mySlice5)
	fmt.Println(arr)
}
