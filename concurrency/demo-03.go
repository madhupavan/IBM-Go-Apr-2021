package main

import (
	"fmt"
	"time"
)

func print(s string, ch chan string, delay time.Duration) {
	fmt.Println(s)
	time.Sleep(delay * time.Second)
	ch <- s + " - done"
}

func main() {
	ch := make(chan string)
	go print("Magesh", ch, 7)
	go print("Suresh", ch, 5)
	go print("Ramesh", ch, 3)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("End of main")
}
