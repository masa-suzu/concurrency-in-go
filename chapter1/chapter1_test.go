package chapter1_test

import (
	"testing"

	"github.com/masa-suzu/concurrency-in-go/chapter1"
)

func TestRaceData(t *testing.T) {

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
				got := chapter1.RaceData()

				if got != tt.want {
					t.Fatalf("chapter1.RaceData() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func DeadLock(t *testing.T) {

	tests := []struct {
		name string
		in   int
	}{
		{name: "trial=1", in: 1},
	}
	defer func() {
		t.Error("fatal error: all goroutines are asleep - deadlock!")
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.in; i++ {
				chapter1.DeadLock()
			}
		})
	}
}

func TestLiveLock(t *testing.T) {

	want := true
	got := chapter1.LiveLock()

	if got != want {
		t.Errorf("chapter1.LiveLock() = %v, want %v", got, want)
	}
}

func TestShortResources(t *testing.T) {
	chapter1.ShortResources()
}
