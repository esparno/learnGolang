package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for i, v:= range os.Args[0:]{
		fmt.Println(strconv.Itoa(i) + ": " + v)
	}
	fmt.Println(strings.Join(os.Args[0:], " "))
}
