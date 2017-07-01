package main

import "fmt"

func main() {
	//function expression
	sayhi := func() {
		fmt.Println("Hello world!")
	}

	defer greet("Erin")
	defer fmt.Println(fullname("Erin", "Sparno"))
	fmt.Println(bondStyleName("Erin", "Sparno"))
	fmt.Println(average(5, 4, 3, 2, 1, 6))
	data := []float64{4, 5, 6, 7, 8, 9, 10}
	fmt.Println(average(data...))
	fmt.Println(average2(data))
	sayhi()
	fc := createFunction()
	fc()

	callbackExample([]int{3, 4, 5}, func(x int) {
		fmt.Println(x)
	})
	fmt.Println(factorial(4))

}

// void
func greet(name string) {
	fmt.Println("Hi " + name)
}

// return string
func fullname(fname, lname string) string {
	return fname + " " + lname
}

// return multiple values
func bondStyleName(fname, lname string) (string, string) {
	return lname, fname + " " + lname
}

// variadic
func average(sf ...float64) float64 {
	fmt.Println(sf)
	//fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func average2(sf []float64) float64 {
	fmt.Println(sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// returning a function
func createFunction() func() {
	return func() {
		fmt.Println("Function Created!")
	}
}

// callbacks
func callbackExample(nums []int, callback func(int)) {
	for _, n := range nums {
		callback(n)
	}
}

// recursion
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
