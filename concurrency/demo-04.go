package main

import "fmt"

func print(s string /*  */) {
	for i := 0; i < 10; i++ {
		fmt.Println(s)
	}
}

func main() {
	print("Hello")
	print("World")
}
