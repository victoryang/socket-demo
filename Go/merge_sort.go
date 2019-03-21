package main

import "fmt"

const (
    ArrSize = 10
)

func merge(arr_x []int, arr_y []int) []int {
    var tmp = make([]int, 0)
    var i int = 0
    var j int = 0

    for i<len(arr_x) && j<len(arr_y) {

        var m int = 0
        if arr_x[i] < arr_y[j] {
            m = arr_x[i]
            i += 1
        } else {
            m = arr_y[j]
            j += 1
        }

        tmp = append(tmp, m)
    }

    for i<len(arr_x) {
        tmp = append(tmp, arr_x[i])
        i += 1
    }

    for j<len(arr_y) {
        tmp = append(tmp, arr_y[j])
        j += 1
    }

    return tmp
}

func merge_sort(arr *[ArrSize]int, start int, end int) []int {
    if start >= end {
        tmp := make([]int, 0)
        tmp = append(tmp, arr[start])
        return tmp
    }

    var mid int = (start+end)/2
    tmp1 := merge_sort(arr, start, mid)
    tmp2 := merge_sort(arr, mid+1, end)

    tmp := merge(tmp1, tmp2)

    return tmp
}

func main() {
    var arr = [ArrSize]int{3,6,2,8,5,9,1,4,7,0}

    res := merge_sort(&arr, 0, ArrSize-1)

    fmt.Println(res)
}