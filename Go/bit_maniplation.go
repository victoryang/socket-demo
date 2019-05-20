package main

import (
    "fmt"
)

func last_zero_nums(num int) int {
    var i int = 0

    for num&1 == 0{
        num = num >> 1
        i++
    }

    return i
}

func find_one_nums(num int) int {
    var i int = 0

    var n int = num
    for n != 0 {
        n = n&(n-1)
        i++
    }

    return i
}

func set_bit(num *int, index int) {
    *num |= 1 << uint(index)
}

func unset_bit(num *int, index int) {
    *num &^= 1 << uint(index)
}

func main() {
    var num int = 8

    fmt.Println("last zero's numbers: ", last_zero_nums(num))
    fmt.Println("one's numbers: ", find_one_nums(num))

    set_bit(&num, 2)
    fmt.Println("set bit", 2, ": ", num)

    unset_bit(&num, 2)
    fmt.Println("unset bit", 2, ": ", num)
}