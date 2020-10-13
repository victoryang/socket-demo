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
    var res *ListNode = nil
    
    tail := new(ListNode)
    tail.Val = 0
    tail.Next = nil
    res = tail

    carry := 0
    var t1Val int
    var t2Val int
    for t1!=nil || t2!=nil {
        if t1 == nil {
            t1Val = 0
        } else {
            t1Val = t1.Val
            t1 = t1.Next
        }
        
        if t2 == nil {
            t2Val = 0
        } else {
            t2Val = t2.Val
            t2 = t2.Next
        }

        val := t1Val + t2Val + carry
        if val >= 10 {
            val = val % 10
            carry = 1
        } else {
            carry = 0
        }
        
        tmp := new(ListNode)
        tmp.Val = val
        tmp.Next = tail.Next        
        tail.Next = tmp
        tail = tmp
    }
    
    if carry == 1 {
        tmp := new(ListNode)
        tmp.Val = 1
        tmp.Next = tail.Next        
        tail.Next = tmp
        tail = tmp
    }

    return res.Next
}

/*
 func make_list(num []int) *ListNode {
    var list *ListNode = nil
    
    // first value
    tail := new(ListNode)
    tail.Val = num[0]
    tail.Next = nil
    list = tail

    i := 1
    for i < len(num) {
        tmp := new(ListNode)
        tmp.Val = num[i]
        tmp.Next = nil
        tail.Next = tmp
        tail = tmp
        i = i + 1
    }

    return list
 }

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    t1 := l1
    t2 := l2
    
    res := make([]int, 0)
    addOne := 0
    for t1!=nil || t2!=nil {
        var t1Val int
        var t2Val int
        if t1 == nil {
            t1Val = 0
        } else {
            t1Val = t1.Val
            t1 = t1.Next
        }
        if t2 == nil {
            t2Val = 0
        } else {
            t2Val = t2.Val
            t2 = t2.Next
        }
        val := t1Val + t2Val + addOne
        if val >= 10 {
            addOne = 1
            val = val % 10
        } else {
            addOne = 0
        }
        
        res = append(res, val)
    }

    if addOne == 1 {
        res = append(res, 1)
    }

    return make_list(res)
}

 */

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