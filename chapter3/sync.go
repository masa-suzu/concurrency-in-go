package chapter3

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup

	wg.Add(5)
	for i := range [5]struct{}{} {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%vth goroutine sleeping.\n", i+1)
			time.Sleep(1)
		}(i)
	}

	wg.Wait()
	fmt.Println("all goroutines complete.")

}

func Mutex() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()

		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()

		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("arithmetic complete")
}

func RWMutex() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			l.Lock()
			l.Unlock()

			time.Sleep(1)
		}
	}
	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")

	for i := 0; i < 22; i += 2 {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m))
	}
}

func Cond() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")

		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

func BroadCast() {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing windows.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	clickRegistered.Wait()
}

func Once() {
	var count int
	increment := func() {
		count++
	}
	var once sync.Once
	var increments sync.WaitGroup

	increments.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer increments.Done()
			once.Do(increment)
		}(i)
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

func Pool() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new instance")
			return struct {
			}{}
		},
	}

	pool.New()
	instance := pool.Get()
	pool.Put(instance)
	pool.Get()
	pool.Get()

	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			fmt.Println("Allocating 1kb.")

			return &mem
		},
	}
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
