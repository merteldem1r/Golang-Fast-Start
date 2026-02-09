package main

import (
	"fmt"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	fmt.Println("Goroutines and Channels")

	ch := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch <- "data"
	}()

	go func() {
		ch2 <- "data2"
	}()

	select {
	case msgFromCh := <-ch:
		fmt.Println(msgFromCh)
	case msgFromCh2 := <-ch2:
		fmt.Println(msgFromCh2)
	}
}
