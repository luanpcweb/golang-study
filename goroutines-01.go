package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	go coroutine()

	for i := 0; i < 5; i++ {
		fmt.Println("Main ", i)
		time.Sleep(1 * time.Millisecond)
	}

	wg.Wait()

}

func coroutine() {
	for i := 0; i < 5; i++ {
		fmt.Println("Coroutine ", i)
		time.Sleep(1 * time.Millisecond)
	}

	wg.Done()
}
