// https://tour.golang.org/methods/18

package main

import "fmt"

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
  size := len(ipAddr)
  str := ""

  for i := 0; i < size; i++ {
    decimal := fmt.Sprintf("%d", ipAddr[i])

    if i < (size-1) {
      str += decimal + "."
    } else {
      str += decimal
    }
  }

  return str
}

func main() {
  hosts := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
  }
}
