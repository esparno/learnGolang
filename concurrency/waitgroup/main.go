package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup
func main() {
	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()
}
func f1() {
	for i:=0; i<30; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
func f2() {
	for i:=0; i<30; i++ {
		fmt.Println(i)
	}
	wg.Done()
}