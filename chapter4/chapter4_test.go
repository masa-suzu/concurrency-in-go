package chapter4

import (
	"math/rand"
	"runtime"
	"testing"
)

func TestBind(t *testing.T) {

	tests := []struct {
		name string
		want []int
	}{
		{name: "ReadOnlyChannel", want: []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := 0
			for x := range Bind() {
				if x != tt.want[i] {
					t.Errorf("got = %v, want %v", x, tt.want[i])
				}
				i++
			}

			if i != len(tt.want) {
				t.Errorf("len(got) = %v, want %v", i, len(tt.want))
			}
		})
	}
}

func BenchmarkTypedStream(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range RepeatString(done, "string").Take(b.N).value {
	}
}

func BenchmarkGenericStream(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range Repeat(done, "string").Take(b.N).ToString().value {
	}
}

func BenchmarkFinin(b *testing.B) {
	numFinders := runtime.NumCPU()
	done := make(chan interface{})
	b.StartTimer()

	primeStream := RepeatFrom(done, func() int { return rand.Intn(50000000) }).FindPrime()
	finders := make([]*intStream, numFinders)

	for i := 0; i < numFinders; i++ {
		finders[i] = primeStream
	}

	for _ = range FanIn(done, finders...).Take(b.N).value {
	}
}
