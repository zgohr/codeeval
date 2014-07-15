package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "math"
)

func fib(n int) (int) {
    r := (math.Pow(math.Phi, float64(n)) - (math.Pow(-1, float64(n)) /
        math.Pow(math.Phi, float64(n)))) / math.Sqrt(5)
    return int(math.Floor(r+0.5))
}

func main() {
    f, _ := os.Open(os.Args[1])
    r := bufio.NewReader(f)
    for {
        l, err := r.ReadString('\n')
        if err != nil { break }
        l = strings.TrimSpace(l)
        n, _ := strconv.Atoi(l)
        fmt.Println(fib(n))
    }
}
