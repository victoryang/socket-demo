package main

import (
    "fmt"
)

// 1. Matrix
var graph_matrix [5][5]int

// 2. Link
type Line struct {
    end        int
    weight     int
    next       *Line
}

type Vectex struct {
    start       int
    number      int
    neighbor    *Line
    next        *Vectex
}

type Graph struct {
    count       int
    head        *Vectex
}

func create_new_line(end int, weight int) *Line {
    return &Line {
        end: end,
        weight: weight,
        next: nil,
    }
}

func create_new_vectex(start int) *Vectex {
    return &Vectex {
        start: start,
        number: 0,
        neighbor: nil,
        next: nil,
    }
}

func create_new_vectex_for_graph(start int, end int, weight int) *Vectex {
    v := create_new_vectex(start)

    v.neighbor := create_new_line(end, weight)
}

func create_new_graph(start int, end int, weight int) *Graph {
    graph := new(Graph)

    graph.count = 1
    graph.head = create_new_vectex_for_graph(start, end, weight)
}

func find_vectex_in_graph(v *Vectex, start int) *Vectex {
    if v == nil {
        return nil
    }

    for v!=nil {
        if start == v.start {
            return v
        }

        v = v.next
    }

    return nil
}

func find_line_in_graph(l *Line, end int) *Line {
    if l == nil {
        return nil
    }

    for l!=nil {
        if end == l.end {
            return l
        }

        l = l.next
    }

    return nil
}

func insert_vectex_into_graph(graph *Graph, start int, end int, weight int) bool {
    if graph == nil {
        return false
    }

    if graph.head == nil {
        graph.head = create_new_vectex_for_graph(start, end, weight)
        graph.head.number++
        graph.count++
        return true
    }

    // Try to find out if start already exists
    vec := find_vectex_in_graph(graph.head, start)
    if  vec==nil {
        vec = create_new_vectex_for_graph(start, end, weight)
        vec.next = graph.head
        graph.head = vec
        graph.head.number++
        graph.count++
        return true
    }

    line := find_line_in_graph(vec.neighbor, end, weight)
    if line!=nil {
        return false
    }

    line = create_new_line(end, weight)
    line.next = vec.neighbor
    vec.neighbor = line

    return true
}

func delete_old_vectex(start int, pv **Vectex) bool {
    if pv == nil || *pv == nil {
        return false
    }

    v := find_vectex_in_graph(*pv, start)
    if v == nil {
        return false
    }

    if v == *pv {
        *pv = v.next
        return true
    }

    pre := *pv
    for pre.next!=v {
        pre = pre.next
    }

    pre.next = v.next
    return true
}

func delete_old_line(end int, pline **Line) bool {
    if pline == nil || *pline == nil {
        return false
    }

    l := find_line_in_graph(*pline, end)
    if l == nil {
        return false
    }

    if l == *pline {
        *pline = l.next
        return true
    }

    pre := *pline
    for pre.next!=l {
        pre = pre.next
    }

    pre.next = l.next
    return true
}

func delete_vectex_from_graph(graph *Graph, start int, end int, weight int) bool {
    if graph==nil || graph.head==nil {
        return false
    }

    v := find_vectex_in_graph(graph.head, start)
    if v==nil {
        return false
    }

    line := find_line_in_graph(v.neighbor, end)
    if line==nil {
        return false
    }

    res := delete_old_line(end, &v.neighbor)
    v.number--
    if len(v.neighbor) == 0 {
        res = delete_old_vectex(start, &graph.head)
    }

    graph.count--
    return res
}

func main() {
    fmt.Println(graph_matrix)
}