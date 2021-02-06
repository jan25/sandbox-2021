package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hi!")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		j := i
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(j)
		}()
	}

	wg.Wait()
}
