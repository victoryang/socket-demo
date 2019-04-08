package main

import "fmt"

func main() {
    var num uint = 140

    fmt.Printf("num is %b\n", num)

    var total int
    /*for num != 0 {
        if num&1 == 1 {
            total += 1
        }
        num = num >> 1
    }*/

    for num != 0 {
        total += 1
        // n & (n-1) will make last "1" be 0
        num = num&(num-1)
    }

    fmt.Println("res: ", total)
}