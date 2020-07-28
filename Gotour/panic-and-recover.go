package main

import "fmt"

func main() {
  f()
  fmt.Println("Returned normally from f.") // execute
}

func f() {
  defer func() {
    //recover is only useful inside deferred functions.
    if r := recover(); r != nil {
      fmt.Println("Recovered in f", r) // r == i == 4
    }
  }()
  fmt.Println("Calling g.")
  g(0)
  fmt.Println("Returned normally from g.") //not execute
}

func g(i int) {
  if i > 3 {
    fmt.Println("Panicking!")
    panic(fmt.Sprintf("%v", i)) // i == 4
  }
  defer fmt.Println("Defer in g", i)
  fmt.Println("Printing in g", i)
  g(i + 1)
}
// Calling g
// Printing in g 0 1 2 3
// Panicking
// Defer in g 3 2 1 0
// Recovered in f 4
// Returned normally from f
