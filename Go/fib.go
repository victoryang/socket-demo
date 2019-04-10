package main

import "fmt"

func main() {
    var arr = [10]int{0}

    arr[0] = 1
    arr[1] = 1

    for i:=2; i<10;i++ {
        arr[i] = arr[i-1] + arr[i-2]
    }

    fmt.Println("arr: ", arr)
}