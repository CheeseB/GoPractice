package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{
  "Bell Labs": {40.7, -74.4},
  "Google": {37.4, -122.1}, // "Google": Vertex{37.4, -122.1}
}

func main() {
  fmt.Println(m)
}
