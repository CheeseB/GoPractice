package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "Woooooorrrrrld"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	//const NotTrue := false //error
	fmt.Println("Go rules?", Truth)
}
