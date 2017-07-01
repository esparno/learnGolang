package main

import (
	"fmt"
	//"time"
	"sync"
)

func main() {
	//c:= make(chan int)
	//go func() {
	//	for i:=1; i<=10; i++ {
	//		c <- i
	//	}
	//}()
	//go func() {
	//	for {
	//		fmt.Println(<-c)
	//	}
	//}()

	//ch:= make(chan int)
	//go func() {
	//	for i:=1; i<=10; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//go func() {
	//	for n := range ch {
	//		fmt.Println(n)
	//	}
	//}()

	//time.Sleep(time.Second)

	var wg sync.WaitGroup
	wg.Add(2)
	ch3 := make(chan int)
	go func() {
		for i := 11; i <= 20; i++ {
			ch3 <- i
		}
		wg.Done()
	}()
	go func() {
		for i := 11; i <= 20; i++ {
			ch3 <- i
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(ch3)
	}()

	for n := range ch3 {
		fmt.Println(n)
	}

}
