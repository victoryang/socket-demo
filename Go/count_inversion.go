package main

import(
    "fmt"
)

const (
    ArrSize = 5
)

type Tuple struct { x, y int}

func count_inversion_version1(arr *[ArrSize]int) []Tuple {
    var res = make([]Tuple, 0)

    for i:=0; i<ArrSize; i++ {
        for j:=i; j<ArrSize ;j++ {
            if arr[i] > arr[j] {
                res = append(res, Tuple{arr[i], arr[j]})
            }
        }
    }

    return res
}

func main() {
    var arr = [ArrSize]int{2,4,1,3,5}

    res := count_inversion(&arr)

    fmt.Println("res: ", res)
}