package main

import "fmt"

func main() {
  s := []int{2, 3, 5, 7 ,11, 13}
  printSlice(s)

  s = s[:0] // zero length
  printSlice(s)

  s = s[:4] // extend length
  printSlice(s)

  s = s[2:] // drop first two values
            // its len, cap will decrease
  printSlice(s)

  // s = s[1:7] //out of bound error
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
  // len: length, cap: capacity
}
