package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
)

func main() {
    f, _ := os.Open(os.Args[1])
    r := bufio.NewReader(f)
    for {
        l, err := r.ReadString('\n')
        if err != nil { break }
        l = strings.TrimSpace(l)
        fmt.Println(strings.ToLower(l))
    }
}
