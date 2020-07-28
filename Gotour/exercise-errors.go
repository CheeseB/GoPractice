package main

import (
  "fmt"
  "math"
)

// in the "fmt" package
/*
type error interface {
	Error() string
}
*/

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

const e = 1e-8

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }
  z := x
  for {
    new_z := z - ((z * z - x) / (2 * z))
    if math.Abs(new_z - z) < e {
      return new_z, nil
    }
    z = new_z
  }
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
