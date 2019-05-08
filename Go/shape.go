package main

import (
    "fmt"
)

const (
    SIZE = 10
)

func star_1() {
    for i:=1; i<=SIZE;i++ {
        for j:=1; j<=i;j++ {
            fmt.Print("* ")
        }
        fmt.Print("\n")
    }
}

func star_2() {
    for i:=1; i<=SIZE;i++ {
        for j:=i; j<=SIZE;j++ {
            fmt.Print("* ")
        }
        fmt.Print("\n")
    }
}

func star_3() {
    for i:=1; i<=SIZE;i++ {
        for j:=1; j<=i;j++ {
            fmt.Printf("%d ", i*j)
        }
        fmt.Print("\n")
    }
}

func main() {
    fmt.Println("Shape 1:")
    star_1()
    fmt.Println("Shape 2:")
    star_2()
    fmt.Println("Shape 3:")
    star_3()
}