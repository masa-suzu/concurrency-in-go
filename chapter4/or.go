package chapter4

import (
	"fmt"
	"time"
)

func OR(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-OR(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
func ORSample() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-OR(
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(1*time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
