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
        max_sum := int64(0)
        positive_sum := int64(0)
        for _, v := range arr {
            v = strings.TrimSpace(v)
            vi, _ := strconv.ParseInt(v, 0, 0)
            if vi + positive_sum > 0 {
                positive_sum = vi + positive_sum
            } else {
                positive_sum = 0
            }
            if max_sum < positive_sum {
                max_sum = positive_sum
            }
        }
        fmt.Println(max_sum)
    }
}
