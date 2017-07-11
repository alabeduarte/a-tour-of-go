// https://tour.golang.org/methods/20

package main

import (
  "fmt"
  "math"
)

const delta = 1e-6

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }

  z := x
  n := 0.0
  for math.Abs(n-z) > delta {
    n, z = z, z-(z*z-x)/(2*z)
  }

  return z, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
