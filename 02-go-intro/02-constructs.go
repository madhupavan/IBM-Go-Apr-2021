package main

import "fmt"

func main() {
	//if else construct
	/*
		no := 42
		if no%2 == 0 {
			fmt.Printf("%v is even\n", no)
		} else {
			fmt.Printf("%v is odd\n", no)
		}
	*/

	if no := 42; no%2 == 0 {
		fmt.Printf("%v is even\n", no)
	} else {
		fmt.Printf("%v is odd\n", no)
	}

	//for construct
	/*
		sum := 0
		for i := 0; i < 9; i++ {
			sum += i
		}
		fmt.Println(sum)
	*/

	//without pre/post statements (while alternative)
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	//infinite for
	/*
		for {

		}
	*/

	//switch case

	/*
		no := 51
		remainder := no % 2
		switch remainder {
		case 0:
			fmt.Printf("%v is even \n", no)
		case 3 - 2:
			fmt.Printf("%v is odd\n", no)
		}
	*/
	/*
		0 - 3 = Terrible
		4 - 5 = Mediocre
		6 - 7 = Not bad
		8 - 9 = Good
		10 = Perfect
	*/
	/*
		score := 11
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5:
			fmt.Println("Mediocre")
		case 6, 7:
			fmt.Println("Not bad!")
		case 8, 9:
			fmt.Println("Good!!")
		case 10:
			fmt.Println("Perfect!!!")
		default:
			fmt.Printf("%v is out of range\n", score)
		}
	*/

	//using 'fallthrough' construct
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	default:
		fmt.Println("Try again!")
	}
}
