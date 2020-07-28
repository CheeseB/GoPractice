package main

import "fmt"

var i, j int = 1, 2 // can initialize all at one line

func main() {
	var c, python, java = true, false, "no!" // type can be omitted
	fmt.Println(i, j, c, python, java)
}
