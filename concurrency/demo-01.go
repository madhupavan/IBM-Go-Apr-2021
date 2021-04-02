package main

import (
	"fmt"
	"time"
)

func print(s string) {
	fmt.Println(s)
}

func main() {
	go print("Hello")
	print("World")
	/* fmt.Println("end of main") */
	time.Sleep(3 * time.Second)
}
