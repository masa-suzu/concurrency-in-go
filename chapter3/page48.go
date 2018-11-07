package chapter3

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup

	wg.Add(5)
	for i, _ := range [5]struct{}{} {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%vth goroutine sleeping.\n", i+1)
			time.Sleep(1)
		}(i)
	}

	wg.Wait()
	fmt.Println("all goroutines complete.")

}
