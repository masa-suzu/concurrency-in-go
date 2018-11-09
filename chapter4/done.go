package chapter4

import (
	"fmt"
	"math/rand"
	"time"
)

func Done1() {
	do := func(strings <-chan string, done <-chan interface{}) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("Done do")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := do(nil, done)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling...")
		close(done)
	}()
	<-terminated
	fmt.Println("Done Done1")
}

func Done2() {
	do := func(done <-chan interface{}) <-chan int {
		i := make(chan int)

		go func() {
			defer fmt.Println("Done do")
			defer close(i)

			for {
				select {
				case i <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return i
	}
	done := make(chan interface{})
	got := do(done)

	for i := 0; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-got)
	}
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("Done Done1")
}
