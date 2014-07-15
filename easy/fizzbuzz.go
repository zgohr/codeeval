package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "math"
)

func main() {
    f, _ := os.Open(os.Args[1])
    r := bufio.NewReader(f)
    for {
        l, err := r.ReadString('\n')
        if err != nil { break }
        l = strings.TrimSpace(l)
        nums := strings.Split(l, " ")
        a, _ := strconv.ParseInt(strings.TrimSpace(nums[0]), 0, 0)
        b, _ := strconv.ParseInt(strings.TrimSpace(nums[1]), 0, 0)
        n, _ := strconv.ParseInt(strings.TrimSpace(nums[2]), 0, 0)
        str := ""
        for j := int64(0); j < n; j++ {
            modA := math.Mod(float64(j+1), float64(a))
            modB := math.Mod(float64(j+1), float64(b))
            if modA == 0 && modB == 0 {
                str += "FB "
            } else if modA == 0 {
                str += "F "
            } else if modB == 0 {
                str += "B "
            } else {
                str += fmt.Sprintf("%d ", j+1)
            }
        }
        fmt.Println(strings.TrimSpace(str))
    }
}
