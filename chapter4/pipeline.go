package chapter4

import (
	"fmt"
)

type maps func(int) int

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

func apply(done <-chan interface{}, in <-chan int, m maps) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			select {
			case <-done:
				return
			case out <- m(i):
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

type stream struct {
	done  <-chan interface{}
	value <-chan interface{}
}

func Repeat(done <-chan interface{}, values ...interface{}) *stream {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case ch <- v:
				}
			}
		}
	}()
	return &stream{done: done, value: ch}
}

func (s *stream) Take(num int) *stream {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for i := 0; i < num; i++ {
			select {
			case <-s.done:
				return
			case ch <- <-s.value:
			}
		}
	}()
	return &stream{done: s.done, value: ch}
}

func (s *stream) Plus(v int) *stream {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for value := range s.value {
			i, ok := value.(int)
			if !ok {
				ch <- fmt.Sprintf("%#v is not integer", value)
				continue
			}
			select {
			case <-s.done:
				return
			case ch <- v + i:
			}
		}
	}()
	return &stream{done: s.done, value: ch}
}

func TakeTenOnes() {
	done := make(chan interface{})
	defer close(done)

	for one := range Repeat(done, 0, "string").Take(10).Plus(1).value {
		fmt.Printf("%v ", one)
	}
	fmt.Println("")
}
