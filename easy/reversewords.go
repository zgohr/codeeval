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
        if l == "" {
            continue
        }
        arr := strings.Split(l, " ")
        for i := len(arr) - 1; i > -1; i-- {
            if i == 0 {
                fmt.Printf("%s", arr[i])
            } else {
                fmt.Printf("%s ", arr[i])
            }
        }
        fmt.Println()
    }
}
