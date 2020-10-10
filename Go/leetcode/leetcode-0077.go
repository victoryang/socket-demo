package main

// Combination

import "fmt"

func try(start int, r *[]int, k int, n int, res *[][]int){
    if len(*r) == k {
        *res = append(*res, *r)
        return
    }

    for i:=start+1; i<=n; i++ {
        tmp := append(*r, i)
        try(i, &tmp, k, n, res)
    }
}

func combination(n, k int) {
    res := make([][]int, 0)

    for i:=1; i<=n; i++ {
        r := make([]int, 0)
        r = append(r, i)
        try(i, &r, k, n, &res)
    }

    fmt.Println("res: ", res)
}

func main() {
    fmt.Println("Please input n: ")
    var n int
    fmt.Scanf("%d", &n)

    fmt.Println("Please input k: ")
    var k int
    fmt.Scanf("%d", &k)

    combination(n, k)
}