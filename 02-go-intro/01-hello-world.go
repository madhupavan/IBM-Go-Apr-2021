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
var username, message string = "Magesh", "Hello!"

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

	fmt.Println(username, message)
}
