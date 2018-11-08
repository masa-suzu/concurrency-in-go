package chapter3

import (
	"bytes"
	"fmt"
	"os"
)

func Channel() {
	stream := make(chan string)

	go func() {
		stream <- "in1"
		close(stream)
	}()

	for input := range stream {
		fmt.Printf("%v\n", input)
	}
}

func DeadLockChannel() {
	stream := make(chan string)

	go func() {
		if 0 != 1 {
			return
		}
		stream <- "in"
	}()
	fmt.Println(<-stream)
}

func BufferedChannels() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	stream := make(chan int, 3)

	go func() {
		defer close(stream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done")
		for i := 0; i < 5; i++ {
			stream <- i
			fmt.Fprintf(&stdoutBuff, "Sendind: %d\n", i)
		}
	}()

	for i := range stream {
		fmt.Fprintf(&stdoutBuff, "Received %v\n", i)
	}

}

func ReadOnlyChannel() {
	createReadOnlyChannel := func() <-chan int {
		buf := make(chan int, 5)
		go func() {
			defer close(buf)
			for i := 0; i <= 5; i++ {
				fmt.Printf("Sent: %d\n", i)
				buf <- i
			}
		}()
		return buf
	}
	for ret := range createReadOnlyChannel() {
		fmt.Printf("Recieved: %d\n", ret)
	}
	fmt.Println("Done")
}
