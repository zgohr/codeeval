package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
)


func main() {
    f, err := os.Open(os.Args[1])

    if err != nil {
      fmt.Println(err)
      return
    }

    defer f.Close()

    r := bufio.NewReader(f)
    for {
        l, err := r.ReadString('\n')
        if err != nil { break }
        l = strings.TrimSpace(l)
        fmt.Println(l)
    }
}
