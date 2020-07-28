package main

import "fmt"

type Vertex struct {
  X, Y int
  S string
}

func main() {
  v := new(Vertex)
  fmt.Println(v)
  v.X, v.Y = 11, 9
  v.S = "omg"
  fmt.Println(v)
}
