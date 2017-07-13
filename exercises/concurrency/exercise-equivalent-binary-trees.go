// https://tour.golang.org/concurrency/7
// https://tour.golang.org/concurrency/8

package main

import (
  "golang.org/x/tour/tree"
  "fmt"
)

func Walk(t *tree.Tree, ch chan int) {
  _walk(t, ch)
  close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
  if t.Left != nil {
    _walk(t.Left, ch)
  }

  ch <- t.Value

  if t.Right != nil {
    _walk(t.Right, ch)
  }
}

func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)

  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for i := range ch1 {
    if i == <- ch2 {
      return true
    }
  }

  return false
}

func main() {
  fmt.Println(Same(tree.New(1), tree.New(1)), "true")
  fmt.Println(Same(tree.New(1), tree.New(2)), "false")
}
