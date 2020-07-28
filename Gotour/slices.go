package main

import "fmt"

func main() {
  primes := [6]int{2, 3, 5, 7, 11, 13}
  var s []int = primes[1:4] // index 1 to 3
  fmt.Println(s)

  //slice literal
  p := []int{2, 3, 5, 7, 11, 13}
  fmt.Println("p ==", p)
  for i := 0; i < len(p); i++ {
    fmt.Printf("p[%d] == %d\n", i, p[i])
  }
}
