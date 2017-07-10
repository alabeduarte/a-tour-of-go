package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  fields := strings.Fields(s)
  m := make(map[string]int)

  for _, key := range fields {
    elem, ok := m[key]

    if ok {
      elem += 1
    } else {
      elem = 1
    }

    m[key] = elem
  }

  return m
}

func main() {
  wc.Test(WordCount)
}
