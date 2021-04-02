package main

import "fmt"

func main() {
	var x interface{}
	x = true
	fmt.Println(x)

	/*
		value, ok := x.(int)
		if ok {
			fmt.Printf("%v is of type int\n", value)
		}
	*/

	switch v := x.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	/* x = "abc"
	fmt.Println(x)

	x = true
	fmt.Println(x) */
}
