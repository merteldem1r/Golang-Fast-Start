package main

import (
	"fmt"
)

func main() {
	fmt.Println("Goroutines and Channels")

	ch := make(chan string)

	go func() {
		ch <- "data"
	}()

	// main goroutine waits for the message, that's because main function waits to either closed or to message to be received (this is the join point)
	msg := <-ch // this is blocking line of code
	fmt.Println(msg)
}
