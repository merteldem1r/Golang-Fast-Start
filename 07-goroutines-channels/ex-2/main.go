package main

import (
	"fmt"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	fmt.Println("Goroutines and Channels")

	// Create two unbuffered channels that carry string values.
	// Unbuffered means a send blocks until another goroutine receives, and vice versa.
	ch := make(chan string)
	ch2 := make(chan string)

	// Launch a goroutine (lightweight concurrent thread) using an anonymous function.
	// It sends "data" into ch. Because ch is unbuffered, this send will block
	// until the main goroutine (or another goroutine) reads from ch.
	go func() {
		ch <- "data"
	}()

	// Launch a second goroutine that sends "data2" into ch2.
	// Both goroutines are now running concurrently, racing to deliver their values.
	go func() {
		ch2 <- "data2"
	}()

	// select waits on multiple channel operations simultaneously.
	// It blocks until ONE of the cases is ready, then executes that case.
	// Key behaviors:
	//   - If both channels have data ready at the same time, select picks one at RANDOM.
	//   - If neither is ready yet, select blocks until one becomes ready.
	//   - Only ONE case is executed per select (not both).
	// This means the output is non-deterministic — you'll see either "data" or "data2".
	select {
	case msgFromCh := <-ch:
		// This case runs if ch delivers its value first.
		fmt.Println(msgFromCh)
	case msgFromCh2 := <-ch2:
		// This case runs if ch2 delivers its value first.
		fmt.Println(msgFromCh2)
	}
	// Note: the goroutine whose channel was NOT selected will be left blocked on its send.
	// Since main() exits here, the program terminates and that goroutine is cleaned up.
}
