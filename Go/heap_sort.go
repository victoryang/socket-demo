package main

import "fmt"

const (
    SIZE = 10
)

func adjust(arr *[SIZE]int, length int) {
    var i int = 0
    var pos = 0

    for i<length {
        if 2*i+2 <length {
            if arr[2*i+2] > arr[2*i+1] {
                pos = 2*i+2
            } else {
                pos = 2*i+1
            }
        } else if 2*i+1 < length {
            pos = 2*i+1
        } else {
            break
        }

        if arr[i] > arr[pos] {
            break
        } else {
            arr[i], arr[pos] = arr[pos], arr[i]
            i = pos
        }
    }
}


func heapify(arr *[SIZE]int, root int) {
    var i int = root

    for i >= 0 {
        if arr[i] < arr[i*2+1] {
            arr[i], arr[i*2+1] = arr[i*2+1], arr[i]
        }

        if i*2+2 < SIZE {
            if arr[i] < arr[i*2+2] {
                arr[i], arr[i*2+2] = arr[i*2+2], arr[i]
            }
        }

        if i==0 {
            break
        }

        i = (i-1)/2
    }
}

func heap_sort(arr *[SIZE]int) {
    last := SIZE-1
    for i:=0;i<=(last-1)/2;i++ {
        heapify(arr, i)
    }

    fmt.Println("heap tree: ", arr)

    for i:=last;i>0;i-- {
        arr[0], arr[i] = arr[i], arr[0]

        adjust(arr, i)
    }
}

func main() {
    var arr = [SIZE]int{3,6,2,8,5,9,1,4,7,0}

    heap_sort(&arr)

    fmt.Println(arr)
}