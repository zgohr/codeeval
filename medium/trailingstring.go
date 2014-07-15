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
        l = strings.TrimSpace(l)
        split := strings.Split(l, ",")
        idx := strings.Index(split[0], split[1])
        if idx != -1 && idx + len(split[1]) == len(split[0]) {
            fmt.Println("1")
        } else { 
            fmt.Println("0")
        }
        // Eat errors
        if err != nil { break }
    }
}
