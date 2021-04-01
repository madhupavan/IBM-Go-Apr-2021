package main

import "fmt"

/*
func main() {
	fmt.Println("Hello world!!")
}
*/

/* func main() {
	var message string
	message = "Hello World!"
	fmt.Println(message)
} */

/* func main() {
	var message string = "Hello World!"
	fmt.Println(message)
} */

/* func main() {
	message := "Hello World!"
	fmt.Println(message)
} */

//package level members
/*
	var username, message string = "Magesh", "Hello!"
*/

func main() {
	/*
		username := "Magesh"
		message := "Hello!"
	*/
	/*
		username, message := "Magesh", "Hello!"
	*/
	/*
		var username string
		var message string
		username = "Magesh"
		message = "Hello!"
	*/
	/*
		var username, message string
		username = "Magesh"
		message = "Hello!"
	*/
	/*
		var username, message string
		username, message = "Magesh", "Hello!"
	*/

	/*
		const username, message = "Magesh", "Hello!"
		fmt.Println(username, message)
	*/

	//variables of different types
	/*
		var username string
		var age int
		var city string

		username = "Magesh"
		age = 40
		city = "Bangalore"
	*/

	/*
		var (
			username string
			age      int
			city     string
		)

		username = "Magesh"
		age = 40
		city = "Bangalore"
	*/

	/*
		var (
			username string = "Magesh"
			age      int    = 40
			city     string = "Bangalore"
		)
	*/

	/*
		var (
			username, city string = "Magesh", "Bangalore"
			age            int    = 40
		)
	*/

	const (
		username, city string = "Magesh", "Bangalore"
		age            int    = 40
	)

	fmt.Println(username, age, city)
}
