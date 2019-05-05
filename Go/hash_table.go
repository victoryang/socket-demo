package main

import (
    "fmt"
)

type Node struct {
    val     int
    next    *Node
}

var table []*Node

func create_table(num int) {
    table = make([]*Node, num)
    for i,_ := range table {
        table[i] = nil
    }
}

func find_element_in_table(n *Node, v int) bool {
    if n.val == v {
        return true
    }

    if n.next == nil {
        return false
    }

    return find_element_in_table(n.next, v)
}

func insert_into_hash(val int) bool {
    if table[val%10] == nil {
        n := new(Node)
        n.val = val
        n.next = nil
        table[val%10] = n
        return true
    }

    if find_element_in_table(table[val%10], val) {
        return false
    }

    n := table[val%10]
    for n.next!=nil {
        n = n.next
    }

    n.next = new(Node)
    n.next.val = val
    n.next.next = nil
    return true
}

func remove_from_hash(val int) bool {
    if table[val%10] == nil {
        return false
    }

    if find_element_in_table(table[val%10], val) == false {
        return false
    }

    n := table[val%10]
    if n.val == val {
        table[val%10] = n.next
        return true
    }

    pn := n
    n = n.next
    for n != nil {
        if n.val == val {
            pn.next = n.next
            return true
        }
        n = n.next
        pn = pn.next
    }

    return false
}

func print_table() {
    for i,v := range table {
        fmt.Print("table[",i,"] = ")
        if v == nil {
            fmt.Print("nil")
        } else {
            fmt.Print(v.val)
            n := v.next
            for n!=nil {
                fmt.Print(" ")
                fmt.Print(n.val)
                n = n.next
            }
        }
        fmt.Print("\n")
    }
}

func main() {
    var arr = []int{0,1,2,3,4,5,6,7,8,9,10,11}

    create_table(10)

    for _,v := range arr {
        insert_into_hash(v)
    }

    print_table()

    var rm = 6
    fmt.Println("remove", rm, "from table")
    remove_from_hash(rm)

    print_table()
}