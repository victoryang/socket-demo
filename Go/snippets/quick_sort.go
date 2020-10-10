package main

import (
    "fmt"
)

func find_pos(arr *[10]int, start int, end int) int {
    var i int = start
    var j int = end

    for i<j {
        if arr[i] > arr[j] {
            arr[i], arr[j] = arr[j], arr[i]
            i++
        } else {
            j--
        }
    }

    return i
}

func quick_sort(arr *[10]int, start int, end int) {
    var pos int

    if start >= end {
        return
    }

    pos = find_pos(arr, start, end)
    fmt.Println(*arr)
    quick_sort(arr, 0, pos-1)
    quick_sort(arr, pos+1, end)
}

func main() {
    var arr = [10]int{3,6,2,8,5,9,1,4,7,0}

    quick_sort(&arr, 0, 9)

    fmt.Println(arr)
}