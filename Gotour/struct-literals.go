package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1, 2} // X:1, Y:2
  v2 = Vertex{X: 1} // X:1, Y:0
  v3 = Vertex{} // X:0, Y:0
  p = &Vertex{1, 2}
)

func main() {
  fmt.Println(v1, p, *p, v2, v3)
}
