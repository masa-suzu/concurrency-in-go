package chapter4

import "testing"

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