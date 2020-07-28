package main

import "fmt"
import "reflect"

func main() {
	var i, j int = 1, 2 // original declaration
	k := 3 // short declaration
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(j), reflect.TypeOf(k))
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(python))
	fmt.Println(reflect.TypeOf(java))
}
