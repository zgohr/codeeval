package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
)

type Stack struct {
  top *Element // Pointer to the current top element
  size int // current size of the Stack
}

type Element struct {
  value interface{} // dynamic type
  next *Element // Pointer to the next element
}

func (s *Stack) Len() int {
  return s.size
}

func (s *Stack) Push(value interface{}) {
  s.top = &Element{value, s.top}
  s.size++
}

// Remove top element, return its value. If empty, return nil
func (s *Stack) Pop() (value interface{}) {
  if s.size > 0 {
    value, s.top = s.top.value, s.top.next // Sweet assignment
    s.size--
    return
  }
  return nil
}


func main() {
    f, _ := os.Open(os.Args[1])
    r := bufio.NewReader(f)
    for {
        l, err := r.ReadString('\n')
        if err != nil { break }
        l = strings.TrimSpace(l)
        in_values := strings.Split(l, " ")

        stack := new(Stack)
        for i := 0; i < len(in_values); i++ {
          stack.Push(in_values[i])
        }

        for stack.Len() > 0 {
          fmt.Printf("%s ", stack.Pop().(string))
          stack.Pop()
        }
        fmt.Println()
    }
}
