package main

// Subsets

import "fmt"
import "math"

func try(r *[]int, start int, k int, input []int, res *[][]int) {
    if len(*r) == k {
        *res = append(*res, *r)
        return
    }

    for i:=start+1; i<len(input); i++ {
        tmp := append(*r, input[i])
        try(&tmp, i, k, input, res)
    }
}

func sub_k(input []int, k int) [][]int {
    res := make([][]int, 0)

    for i:=0; i<len(input); i++ {
        v := input[i]
        tmp := make([]int, 0)
        tmp = append(tmp, v)
        try(&tmp, i, k, input, &res)
    }

    return res
}

func sub_set1(input []int){
    res := make([][]int, 0)

    for i:=1; i<=len(input); i++ {
        res = append(res, sub_k(input, i)...)
    }

    fmt.Println("res: ", res)
}

func sub_set2(input []int) {
    res := make([][]int, 0)

    n := math.Pow(2, float64(len(input)))
    for i:=1; i<int(n); i++ {
        tmp := i
        index := 0
        r := make([]int, 0)
        for tmp != 0 {
            if tmp&1 != 0 {
                r = append(r, input[index])
            }
            tmp = tmp >> 1
            index = index + 1
        }

        res = append(res, r)
    }

    fmt.Println("res: ", res)
}

func main() {
    fmt.Println("Please input the array, -1 to finish: ")

    var num int
    input := make([]int, 0)

    fmt.Scanf("%d", &num)
    for num != -1 {
        input = append(input, num)

        fmt.Scanf("%d", &num)
    }

    fmt.Println("input list: ", input)

    sub_set1(input)

    sub_set2(input)
}