package chapter1_test

import (
	"testing"

	"github.com/masa-suzu/concurrency-in-go/chapter1"
)

func TestSleepAfterCreatingGoroutine(t *testing.T) {

	tests := []struct {
		name string
		in   int
		want int
	}{
		{name: "trial=10", in: 10, want: 2},
		{name: "trial=10000", in: 10000, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.in; i++ {
				got := chapter1.SleepAfterCreatingGoroutine()

				if got != tt.want {
					t.Fatalf("chapter1.SleepAfterCreatingGoroutine() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
