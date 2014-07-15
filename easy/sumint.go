package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    f, _ := os.Open(os.Args[1])
    r := bufio.NewReader(f)
    s := 0
    for {
        l, err := r.ReadString('\n')
        if err != nil { break }
        l = strings.TrimSpace(l)
        i, _ := strconv.Atoi(l)
        s += i
    }
    fmt.Printf("%d\n", s)
}
