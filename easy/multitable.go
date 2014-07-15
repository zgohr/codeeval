package main

import (
    "fmt"
)

func main() {
    count := 0
    for x := 1; x < 13; x++ {
        count++
        for n := 1; n < 13; n++ {
            fmt.Printf("%4d", n * count)
        }
        fmt.Println("")
    }
}
