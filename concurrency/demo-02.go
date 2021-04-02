package main

import (
	"fmt"
	"time"
)

func writeData(ch chan int) {
	fmt.Println("Writing data into the ch")
	time.Sleep(5 * time.Second)
	ch <- 100
}

func main() {
	ch := make(chan int)
	go writeData(ch)
	result := <-ch
	fmt.Printf("Result from goroutine = %d\n", result)
}
