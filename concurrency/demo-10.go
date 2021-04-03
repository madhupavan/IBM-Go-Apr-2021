package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int, quit chan string) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(1 * time.Second)
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quitting")
			return
		}
	}
}
func main() {
	fibChan := make(chan int)
	quitChan := make(chan string)
	go func() {
		for {
			fmt.Println(<-fibChan)
		}
	}()
	go fibonacci(fibChan, quitChan)
	var input string
	fmt.Println("Hit ENTER to stop!")
	fmt.Scanln(&input)
	quitChan <- "done"
	fmt.Println("Job Done!")
}
