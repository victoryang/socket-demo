package main

import "fmt"


func main() {
    var arr = []int{1, 2, 5, 1, 7, 2, 4, 2}

    var i int = 0
    var bucket = make(map[int]bool, 0)

    for i < len(arr) {
        if _, ok := bucket[arr[i]]; !ok {
            bucket[arr[i]] = true
        }
        i++
    }

    fmt.Println("res: ", bucket)
}