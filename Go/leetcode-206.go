package main

import "fmt"
import "strconv"

type Node struct {
    data int
    next *Node
}

func build_list(arr *[10]int, length int) *Node {
    var link *Node = new(Node)
    link.next = nil

    for i:=0; i<length; i++ {
        tmp := new(Node)

        tmp.data = (*arr)[i]
        tmp.next = link.next
        link.next = tmp
    }

    return link
}

func print_list(list *Node) {
    tmp := list.next
    i := 0

    for tmp != nil {
        fmt.Println("data["+strconv.Itoa(i)+"] = ", tmp.data)

        tmp = tmp.next
        i++
    }
}

func reverse_list(list *Node) *Node {
    t1 := list
    t2 := list.next
    t1.next = nil
    t1 = nil

    for t2 != nil {
        t3 := t2.next
        t2.next = t1

        t1 = t2
        t2 = t3
    }

    list.next = t1

    return list
}

func main() {
    var arr = [10]int{0,1,2,3,4,5,6,7,8,9}

    list := build_list(&arr, 10)

    print_list(list)

    rlist := reverse_list(list)

    print_list(rlist)
}