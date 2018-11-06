package chapter1

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func LiveLock() bool {
	cadence := sync.NewCond(&sync.Mutex{})

	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}

		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32

	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) bool {
		var out bytes.Buffer
		defer func() {
			fmt.Println(out.String())
		}()

		defer walking.Done()

		fmt.Fprintf(&out, "%v is tring to scoot", name)
		for _, _ = range [5]struct{}{} {
			if tryLeft(&out) || tryRight(&out) {
				return true
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation", name)
		return false
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	var ret1, ret2 bool

	go func() { ret1 = walk(&peopleInHallway, "Alice") }()
	go func() { ret2 = walk(&peopleInHallway, "Barbara") }()
	peopleInHallway.Wait()

	return ret1 && ret2
}
