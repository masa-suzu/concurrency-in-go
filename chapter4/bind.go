package chapter4

func Bind() <-chan int {
	ch := make(chan int, 4)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	return ch
}
