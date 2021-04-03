package main

import (
	"fmt"
	"time"
)

func print1(s string, c1 chan string, c2 chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		<-c1
		fmt.Println(s)
		c2 <- "Done"
	}
	fmt.Println(s + " - Done")
}

func print2(s string, c1 chan string, c2 chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		<-c1
		fmt.Println(s)
		c2 <- "Done"
	}
	fmt.Println(s + " - Done")
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go print1("Hello", c1, c2)
	go print2("World", c2, c1)
	c1 <- "start"
	var s string
	fmt.Scanln(&s)
}
