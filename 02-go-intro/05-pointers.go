package main

import "fmt"

func main() {
	var no int = 10
	var addressOfno *int = &no
	fmt.Println(no, addressOfno, *addressOfno)
	fmt.Printf("Before Incrementing, no = %d\n", no)
	increment(&no)
	fmt.Printf("After Incrementing, no = %d\n", no)

	nos := []int{10, 20}
	fmt.Printf("Before inserting, nos = %v\n", nos)
	insertValue(&nos, 30)
	fmt.Printf("After inserting, nos = %v\n", nos)

	var x, y int = 10, 20
	fmt.Printf("Before swapping x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping x = %d, y = %d\n", x, y)
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func increment(x *int) {
	*x += 1
}

func insertValue(nos *[]int, no int) {
	*nos = append(*nos, no)
}
