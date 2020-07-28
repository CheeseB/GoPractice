package main

import "fmt"

func main() {
  s := []int{2, 3, 5, 7, 11, 13}

  s = s[:] // from first to end
  fmt.Println(s)

  s = s[1:4] // from 1 to 3
  fmt.Println(s)

  s = s[:2] // from first to 1
  fmt.Println(s)

  s = s[1:] // from 1 to end
  fmt.Println(s)
}
