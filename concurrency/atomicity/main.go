package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var count int32
func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", count)
}
func incrementor(s string) {
	for i:=0; i<31; i++ {
		atomic.AddInt32(&count, 1)
		fmt.Println(s, i, "Counter: ", count)
	}
	wg.Done()
}