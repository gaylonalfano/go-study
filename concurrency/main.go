package main

import "fmt"

// Goroutines (similar to a thread) and Channels (similar to a queue)
// To start a Go routine, you put 'go' in front of a function
// -- Reading from or writing to a Channel blocks until it's been read.
// -- Read/write to multiple channels at once using Select statements

// Concurrency - Select to r/w to multiple channels
func writer(num int, ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i * num
	}
}

func concurrencySelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go writer(1, ch1) // routine (~thread)
	go writer(2, ch2) // another routine (~thread)
	for i := 0; i < 20; i++ {
		// 'select' picks a channel to be read/written at RANDOM to prevent deadlock!
		select {
		case v := <-ch1:
			fmt.Println("from 1:", v)
		case v := <-ch2:
			fmt.Println("from 2:", v)
		}
	}
}

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		// Using a closure that pass 'i' as the argument, and writes
		// the value to the channel 'ch'
		go func(val int) {
			ch <- val * 2
		}(i)
	}

	// Here we're reading the channel 10 times
	for i := 0; i < 10; i++ {
		result := <-ch
		fmt.Println(result)
	}

	concurrencySelect()
}
