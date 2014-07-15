package main

import (
    "os"
    "bufio"
    "fmt"
)

var codes = []byte {
  ' ',
  'E', 'T',
  'I', 'A', 'N', 'M',
  'S', 'U', 'R', 'W', 'D', 'K', 'G', 'O',
  'H', 'V', 'F', '?', 'L', '?', 'P', 'J', 'B', 'X', 'C', 'Y', 'Z', 'Q', '?', '?',
  '5', '4', '?', '3', '?', '?', '?', '2', '?', '?', '+', '?', '?', '?', '?', '1', '6', '=', '/', '?', '?', '?', '?', '?', '7', '?', '?', '?', '8', '?', '9', '0',
}

func main() {
    f, _ := os.Open(os.Args[1])
    r := bufio.NewReader(f)
    p := 0
    for {
        c, err := r.ReadByte()
        if err != nil { break }
        switch c {
        case '\n':
          fmt.Printf("%c\n", codes[p])
          p = 0
        case ' ':
          fmt.Printf("%c", codes[p])
          p = 0
        case '.':
          p = 2 * p + 1
        case '-':
          p = 2 * p + 2
        }
    }
}
