package main

import (
    "fmt"
    "math"
)

func reverse(s string) string {
    b := make([]byte, len(s));
    var j int = len(s) - 1;
    for i := 0; i <= j; i++ {
        b[j-i] = s[i]
    }
    return string (b);
}

func main() {
    for i := int64(1000); i > 0; i-- {
        max := i / 2 + 1
        broke := false
        for j := int64(2); j < max; j++ {
            if math.Mod(float64(i), float64(j)) == 0 {
                broke = true
                break
            }
        }
        str := fmt.Sprintf("%d", i)
        if !broke && str == reverse(str) {
            fmt.Println(str)
            break
        }
    }
}
