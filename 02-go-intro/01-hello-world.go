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

	//var char rune = 'a'
	/*
		var char int32 = 97
		fmt.Printf("%c\n", char)
	*/
	fmt.Println(username, age, city)
	/*
		fmt.Printf("%s (%d) %s\n", username, age, city)
		fmt.Printf("%T (%T) %T\n", username, age, city)
	*/
	/*
		v1 := fmt.Sprintf("%s (%d) %s\n", username, age, city)
		v2 := fmt.Sprintf("%T (%T) %T\n", username, age, city)
		fmt.Printf("%s%s", v1, v2)
	*/

	i := -50
	fmt.Printf("i = %v (%T)\n", i, i)
	f := float32(i)
	fmt.Printf("f = %v (%T)\n", f, f)
	ui := uint(i)
	fmt.Printf("ui = %v (%T)\n", ui, ui)
	x := f * float32(i)
	fmt.Println(x)
}
