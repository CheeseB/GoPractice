package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
  fmt.Printf("And i have %g problems.\n", math.Nextafter(2, 3))
  fmt.Printf("But we have %g problems.", math.Nextafter(2, 1))
}
