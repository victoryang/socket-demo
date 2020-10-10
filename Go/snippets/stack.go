package main

import "fmt"
import "container/list"

type Stack struct {
    v   *list.List
}

func NewStack() *Stack{
    v := list.New()
    return &Stack{
        v: v,
    }
}

func (s *Stack) Push(value int) {
    s.v.PushBack(value)
}

func (s *Stack) Pop() int {
    e := s.v.Back()
    if e != nil {
        s.v.Remove(e)
        return int((e.Value).(int))
    }
    
    return -1
}

func (s *Stack) Top() int {
    e := s.v.Back()
    if e != nil {
        return int((e.Value).(int))
    }
    
    return -1
}

func main() {
    var arr = [...]int{9,8,7,6,5,4,3,2,1}

    s := NewStack()

    for _,v := range arr {
        s.Push(v)
    }

    fmt.Printf("res ")
    e := s.Top()
    for e!=-1 {
        e = s.Pop()
        fmt.Printf(" %d", e)
    }
    fmt.Printf("\n")
}