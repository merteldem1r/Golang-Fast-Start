package main

import (
	"fmt"
	"sync"
	"time"
)

// count sends messages into the channel. It takes:
// - name: label for this goroutine
// - ch: channel to send messages into (chan<- means send-only, this function can't receive from it)
// - wg: pointer to WaitGroup so we can signal when this goroutine is done (*sync.WaitGroup = pointer, not a copy)
func count(name string, ch chan<- string, wg *sync.WaitGroup) {
	// defer = "run this when the function exits, no matter how it exits"
	// wg.Done() subtracts 1 from the WaitGroup counter
	defer wg.Done()

	for i := range 3 {
		// Send a formatted string INTO the channel
		// This line BLOCKS if nobody is receiving on the other end yet
		ch <- fmt.Sprintf("%v : %v", name, i)

		// Simulate slow work (like an API call)
		time.Sleep(500 * time.Millisecond)
	}
	// When this function returns, defer runs → wg.Done() → counter decrements by 1
}

func main() {
	// Create a channel that carries strings — this is the "pipe" between goroutines and main
	ch := make(chan string)

	// Create a WaitGroup — a counter that tracks how many goroutines are still running
	var wg sync.WaitGroup

	// Tell the WaitGroup: "2 goroutines will be working" and 2 of them will call Done()
	wg.Add(2)

	// Launch goroutine A — starts immediately in the background, doesn't block main
	// &wg passes the POINTER (memory address) so the goroutine modifies the original counter
	go count("A", ch, &wg)

	// Launch goroutine B — also starts immediately, both A and B now run concurrently
	go count("B", ch, &wg)

	// Launch a third goroutine whose only job is to close the channel when all senders are done
	// Why a goroutine? Because wg.Wait() blocks — if we put it here directly,
	// main would freeze waiting, and nobody would receive from ch, causing a deadlock
	go func() {
		wg.Wait() // Block until WaitGroup counter reaches 0 (both A and B called Done)
		close(ch) // Signal to the range loop below: "no more messages are coming"
	}() // call the function immediately

	// Range over the channel — receives messages one by one until the channel is closed
	// Without close(ch), this loop would wait forever after the last message (deadlock)
	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("Done")
}
