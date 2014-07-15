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
        str_vals := strings.Split(l, " ")

        // Init the map
        values := make([]int64, len(str_vals))
        next_map := make(map[int64]int)
        for i := 0; i < len(str_vals); i++ { 
            values[i], _ = strconv.ParseInt(str_vals[i], 10, 0)
            next_map[values[i]] = -1
        }

        for index, value := range values {
            // next_map maps a value to its next value
            if next_map[value] < 0 {
                // If this value's next is unknown, set it
                next_map[value] = int(values[index+1])
            } else {
                // We already know what the next value is going to be
                fmt.Print(value)
                n := next_map[value]
                for value != int64(n) {
                    fmt.Printf(" %d", n)
                    n = next_map[int64(n)]
                }
                fmt.Println()
                break
            }
        }
    }
}
