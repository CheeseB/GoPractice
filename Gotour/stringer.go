package main

import "fmt"

// in the "fmt" package
/*
type Stringer interface {
	String() string
}
*/

type Person struct {
	Name string
	Age int
}

// Person type implements fmt.Stringer
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Sonia", 24}
	z := Person{"Jimon", 26}
	fmt.Println(a, z)
}
