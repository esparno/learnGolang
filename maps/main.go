package main

import "fmt"

func main() {

	var dictionary = make(map[string]string)
	dictionary["Alister"] = "dog"
	dictionary["Hopsy"] = "rabbit"
	dictionary["Bartholomew"] = "cat"
	fmt.Println(dictionary )
	delete(dictionary, "Hopsy")
	for key, value := range dictionary {
		fmt.Println(key, " - ", value)
	}

	dictionary2 := map[int]int {}
	fmt.Println(dictionary2	== nil)

}
