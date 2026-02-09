package main

import (
	"fmt"
	"time"
)

// doWork runs in a loop until it receives a signal on the done channel.
// The parameter type <-chan bool is a receive-only channel — this function
// can only read from it, not send to it, enforcing clear ownership at compile time.
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			// When a value is received (or the channel is closed), we return
			// and the goroutine exits cleanly. This is the "done channel" pattern —
			// a common way to signal a goroutine to stop.
			return
		default:
			// default makes this select non-blocking: if done has no signal yet,
			// execution falls through here immediately instead of waiting.
			// This creates a tight loop that keeps printing until done is signaled.
			fmt.Println("Doing Work...")
		}
	}
}

func main() {
	fmt.Println("Goroutines and channels")

	// Create an unbuffered channel used purely as a signal (the bool value doesn't matter).
	done := make(chan bool)

	// Launch doWork in a separate goroutine. It will keep printing "Doing Work..."
	// in a tight loop until it receives on the done channel.
	go doWork(done)

	// Sleep for 3 seconds, during which doWork keeps running and printing.
	time.Sleep(3 * time.Second)

	// NOTE: The program exits here WITHOUT ever sending to or closing the done channel.
	// This means doWork never receives the stop signal — it's terminated abruptly
	// when main() returns and the process exits. To cleanly stop doWork, you would
	// send a signal (done <- true) or close the channel (close(done)) before exiting.
	fmt.Println("Exiting from main")
}
