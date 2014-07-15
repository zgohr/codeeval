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
        n, _ := strconv.Atoi(arr[0])
        p1, _ := strconv.Atoi(arr[1])
        p2, _ := strconv.Atoi(arr[2])
        b := strconv.FormatInt(int64(n), 2)
        if b[len(b)-p1] == b[len(b)-p2] {
            fmt.Println("true")
        } else {
            fmt.Println("false")
        }
    }
}
