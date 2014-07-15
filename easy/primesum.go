package main

import (
    "fmt"
    "math"
)

func main() {
    total := 0 
    count := 0
    i := 2
    for count < 1000 {
        broken := false
        for n := 2; n < i/2+1; n++ {
            if math.Mod(float64(i), float64(n)) == 0 {
                broken = true
                break
            }
        }
        if !broken {
            count += 1
            total += i
        }
        i += 1
    }
    fmt.Println(total)
}
