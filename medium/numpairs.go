package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func strToIntSlice(orig []string) (asInt []int64) {
    for _, v := range orig {
        v, _ := strconv.ParseInt(v, 0, 0)
        asInt = append(asInt, v)
    }
    return
}

func main() {
    f, _ := os.Open(os.Args[1])
    r := bufio.NewReader(f)
    for {
        l, err := r.ReadString('\n')
        // Eat errors
        if err != nil { break }
        // Trim newline character
        l = strings.TrimSpace(l)
        // ["1,2,3", "4"]
        split := strings.Split(l, ";")
        // [1,2,3]
        nums := strToIntSlice(strings.Split(split[0], ","))
        // 4
        sum, _ := strconv.ParseInt(split[1], 0, 0)

        pairs := []string{}
        for k1, v1 := range nums {
            for _, v2 := range nums[k1+1:] {
                if v1 + v2 == sum {
                    pairs = append(pairs, fmt.Sprintf("%d,%d", v1, v2))
                }
            }
        }

        if len(pairs) == 0 {
            fmt.Println("NULL")
        } else {
            fmt.Println(strings.Join(pairs, ";"))
        }
    }
}
