package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
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
        arr := strings.Split(l, ",")
        var unique []string
        for i, _ := range arr {
            if !stringInSlice(arr[i], unique) {
                unique = append(unique, arr[i])
            }
        }
        fmt.Println(strings.Join(unique, ","))
    }
}
