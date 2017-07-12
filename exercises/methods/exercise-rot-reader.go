// https://tour.golang.org/methods/23

package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (r13 *rot13Reader) Convert(b byte, alphabet []string) byte {
  for i := range alphabet {
    if alphabet[i] == strings.ToUpper(string(b)) {
      rot13 := 13

      if i > (len(alphabet) - 1)/2 {
        rot13 *= -1
      }

      return []byte(alphabet[i + rot13])[0]
    }
  }

  return b
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
  alphabet := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
  n, err := r13.r.Read(b)

  for i := 0; i < len(b); i++ {
    b[i] = r13.Convert(b[i], alphabet);
  }

  return n, err
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
