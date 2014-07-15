package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
)

func contains(a []string, b string) bool {
    for _, c := range a {
        if c == b {
            return true
        }
    }
    return false
}

func main() {
    f, _ := os.Open(os.Args[1])
    r := bufio.NewReader(f)
    for {
        l, err := r.ReadString('\n')
        if err != nil { break }
        l = strings.TrimSpace(l)
        arr := strings.Split(l, ";")
        x := strings.Split(arr[0], ",")
        y := strings.Split(arr[1], ",")
        var match []string
        for _, a := range x {
            if contains(y, a) {
                match = append(match, a)
            }
        }
        fmt.Println(strings.Join(match, ","))
    }
}
