package chapter3

import (
	"fmt"
	"sync"
)

func StrangeGoroutine() {
	var wg sync.WaitGroup
	for _, value := range []string{"hello", "greeting", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(value)
		}()
	}
	wg.Wait()
}
