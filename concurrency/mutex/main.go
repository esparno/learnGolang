package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var wg sync.WaitGroup
var count int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", count)
}
func incrementor(s string) {
	for i := 0; i < 31; i++ {
		m.Lock()
		count++
		m.Unlock()
		fmt.Println(s, i, "Counter: ", count)
	}
	wg.Done()
}
