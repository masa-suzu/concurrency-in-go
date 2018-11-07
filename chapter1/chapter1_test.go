package chapter1

import (
	"testing"
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
				got := RaceData()

				if got != tt.want {
					t.Fatalf("chapter1.RaceData() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func _DeadLock(t *testing.T) {

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
				DeadLock()
			}
		})
	}
}

func TestLiveLock(t *testing.T) {

	want := true
	got := LiveLock()

	if got != want {
		t.Errorf("chapter1.LiveLock() = %v, want %v", got, want)
	}
}

func TestShortResources(t *testing.T) {
	ShortResources()
}
