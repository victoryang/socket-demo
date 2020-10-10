package main

import "fmt"
import "container/list"

type Queue struct {
    list    []int
}

func NewQueue() *Queue {
    return &queue{
        list: make([]int, 0),
    }
}

func (q *Queue) Enqueue(val int) {
    q.list = append(q.list, val)
}

func (q *Queue) Dequeue() int {
    ret := q.list[0]
    q.list = q.list[1:]
    return ret
}

func (q *Queue) Front() int {
    return q.list[0]
}

func (q *Queue) Back() int {
    return q.list[len(q.list)-1]
}

func main() {
    var arr = [...]{3,5,4,7,2,6,9,1,8}

    for _,v :=range arr {

    }
}