package main

import "fmt"

func main() {
  five_sos := [4]string {
    "Ashton",
    "Luke",
    "Michael",
    "Calum",
  }
  fmt.Println(five_sos)

  a := five_sos[0:2]
  b := five_sos[1:3]
  fmt.Println(a, b)

  b[0] = "Luke Hemmings"
  fmt.Println(a, b)
  fmt.Println(five_sos)
}
