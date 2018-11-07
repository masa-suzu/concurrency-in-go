package chapter3

import (
	"sync"
	"testing"
)

func TestStrangeGoroutine(t *testing.T) {
	StrangeGoroutine()
}

func TestIntelligibleGoroutine(t *testing.T) {
	IntelligibleGoroutine()
}

func TestMeasureConsumptionOfMemory(t *testing.T) {
	MeasureConsumptionOfMemory()
}

func BenchmarkSwitchContext(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})

	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()

		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()

		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()

	b.StartTimer()
	close(begin)

	wg.Wait()
}
