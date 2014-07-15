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
        sum := 0
        for i := 0; i < len(l); i++ {
            in, _ := strconv.Atoi(string(l[i]))
            sum += in
        }
        fmt.Println(sum)
    }
}
