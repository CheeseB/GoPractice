package main

import "fmt"

func a() {
  i := 1
  //execute after return, but printed i is 1
  defer fmt.Println(i)
  i++
  return
}

func c() (i int) {
  //execute after i(1) return, and finally i++(2) return
  defer func() { i++ }()
  return 1
}

func main() {
  //execute after main function return.
  defer fmt.Println("world")
  fmt.Println("hello")
  a()
  fmt.Println(c())
}
