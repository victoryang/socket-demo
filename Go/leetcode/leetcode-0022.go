package main

// Generate parentheses

import "fmt"

var target = map[string]int{"(": 1, ")": -1}

func try(r string, left_total int, num int, res *[]string) {
    if len(r)==num*2 {
        *res = append(*res, r)
        return
    }

    if left_total < len(r) - left_total {
        return
    }

    add_left := left_total+1
    if add_left <= num {
        try(r+"(", add_left, num, res)
    }

    right := len(r)-left_total
    if right + 1 <= num {
        try(r+")", left_total, num, res)
    }
}

func generate(n int) {
    res := make([]string, 0)

    try("", 0, n, &res)

    fmt.Println("res: ", res)
}

func main() {
    fmt.Println("Please input number: ")
    var num int

    fmt.Scanf("%d", &num)
    generate(num)
}