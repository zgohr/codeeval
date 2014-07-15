package main

import (
    "os"
    "fmt"
)

func main() {
    fi, _ := os.Stat(os.Args[1])
    fmt.Println(fi.Size())
}
