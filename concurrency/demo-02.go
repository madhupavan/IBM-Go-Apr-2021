package main

import (
	"fmt"
	"time"
)

func writeData(ch chan int) {
	fmt.Println("Writing data into the ch after 5 seconds")
	time.Sleep(5 * time.Second)
	ch <- 100
}

func main() {
	ch := make(chan int)
	go writeData(ch)
	fmt.Println("Attempting to read data from the channel after 7 seconds")
	time.Sleep(7 * time.Second)
	result := <-ch
	fmt.Printf("Result from goroutine = %d\n", result)
}
