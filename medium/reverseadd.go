package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "sync"
)

type PotentiallyPalindromal int

func (p PotentiallyPalindromal) isPalindrome() bool {
  return p == p.reverse()
}

func (p PotentiallyPalindromal) reverse() (r PotentiallyPalindromal) {
  for p != 0 {
    r = r*10 + p%10
    p /= 10
  }
  return
}

func (p *PotentiallyPalindromal) increment() {
  *p += p.reverse()
}

func (p *PotentiallyPalindromal) makePalindromal() (i int) {
  for !p.isPalindrome() {
    p.increment()
    i++
  }
  return
}

func (p PotentiallyPalindromal) String() string {
  count := p.makePalindromal()
  return fmt.Sprintf("%d %d", count, p)
}

func main() {
    f, err := os.Open(os.Args[1])
    if err != nil {
      fmt.Println(err)
      return
    }
    defer f.Close()
    r := bufio.NewReader(f)

    wg := sync.WaitGroup{}

    for {
        l, err := r.ReadString('\n')
        if err != nil { break }

        wg.Add(1)
        go func() {
          n, _ := strconv.Atoi(strings.TrimSpace(l))
          fmt.Println(PotentiallyPalindromal(n))
          wg.Done()
        }()
    }
    wg.Wait()
}
