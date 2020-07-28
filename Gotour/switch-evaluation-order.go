package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("When is Friday?")
  today := time.Now().Weekday()
  switch time.Friday {
  case today:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 1:
    fmt.Println("In two days.")
  default:
    fmt.Println("Too far away.")
  }
}
