package chapter4

import "fmt"

func generate(done <-chan interface{}, integers ...int) <-chan int {
	in := make(chan int, len(integers))

	go func() {
		defer close(in)
		for _, i := range integers {
			select {
			case <-done:
				return
			case in <- i:
			}
		}
	}()
	return in
}

func apply(done <-chan interface{}, in <-chan int, f func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			select {
			case <-done:
				return
			case out <- f(i):
			}
		}
	}()
	return out
}

func PipeLine() {
	done := make(chan interface{})
	defer close(done)

	in := generate(done, 1, 2, 3, 4)
	twice := func(v int) int { return 2 * v }
	plusOne := func(v int) int { return 1 + v }

	pipeline := apply(done, in, plusOne)
	pipeline = apply(done, pipeline, twice)

	for value := range pipeline {
		fmt.Println(value)
	}
}
