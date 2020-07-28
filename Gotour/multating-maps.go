package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"]) //42

  m["Answer"] = 48
  fmt.Println("The value:", m["Answer"]) //48

  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"]) //zero value(0)

  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok) //0, false
}
