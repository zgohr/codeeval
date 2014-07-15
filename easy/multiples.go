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
    for {
        l, err := r.ReadString('\n')
        if err != nil { break }
        l = strings.TrimSpace(l)
        arr := strings.Split(l, ",")
        x, _ := strconv.Atoi(arr[0])
        n, _ := strconv.Atoi(arr[1])
        i := 0
        count := 0
        for i < x {
            count++
            i = n * count
        }
        fmt.Printf("%d\n", i)
    }
}
