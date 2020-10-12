package main

/**
 * You are given two non-empty linked lists representing two non-negative integers. 

 * The digits are stored in reverse order, and each of their nodes contains a single digit. 

 * Add the two numbers and return the sum as a linked list.

 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.

 */

import "fmt"

/** Definition for singly-linked list.
 *  type ListNode struct {
 *      Val int
 *      Next *ListNode
 *  }
 */

type ListNode struct {
	Val 	int
	Next 	*ListNode
}

func make_list(num int) *ListNode {
    var list *ListNode = nil

    n := num
    for n>0 {
        t := n % 10
        tmp := new(ListNode)
        tmp.Val = t
        tmp.Next = list
        list = tmp
        n = n / 10
    }

    return list
}

func print_list(list *ListNode){
    var t *ListNode = list

    for t!=nil {
        fmt.Print(" ", t.Val)
        t = t.Next
    }
    fmt.Print("\n")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    t1 := l1
    t2 := l2
    sum := 0
    unit := 1
    for t1!=nil && t2!=nil {
        sum = sum + (t1.Val + t2.Val)*unit
        unit = unit * 10

        t1 = t1.Next
        t2 = t2.Next
    }

    for t1!=nil {
        sum = sum + t1.Val*unit
        unit = unit * 10

        t1 = t1.Next
    }

    for t2!=nil {
        sum = sum + t2.Val*unit
        unit = unit * 10

        t2 = t2.Next
    } 

    return make_list(sum)
}

func main() {
    a := 9999999
    b := 9999
    c := make_list(a)
    d := make_list(b)

    print_list(c)
    print_list(d)

    e := addTwoNumbers(c, d)
    print_list(e)
}