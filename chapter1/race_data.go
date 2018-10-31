package chapter1

import "time"

func SleepAfterCreatingGoroutine() int {
	var data int

	go func() { data++ }()
	time.Sleep(5)

	if data == 0 {
		return data - 1
	}

	return 2 * data

}
