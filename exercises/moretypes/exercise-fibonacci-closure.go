// https://tour.golang.org/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  prev := 0
  curr := 0
  next := 0

  return func() int {
    switch {
    case curr == 0, curr == 1:
      prev = curr
      curr++
      next++

      return curr
    default:
      next = prev + curr
      prev = curr
      curr = next

      return curr
    }
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}

