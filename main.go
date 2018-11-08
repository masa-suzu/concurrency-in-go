package main

import "github.com/masa-suzu/concurrency-in-go/chapter3"

func main() {
	chapter3.WaitGroup()
	chapter3.Mutex()
	chapter3.RWMutex()
	chapter3.Cond()
	chapter3.BroadCast()
}
