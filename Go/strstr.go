package main

import "fmt"

func strstr1(target string, sub string) {
    if len(target) == 0 || len(sub) == 0 {
        fmt.Println(false)
        return
    }

    var length = len(sub)
    for len(target)>=length {
        var i int
        for i=0; i<length; i++ {
            if target[i]!=sub[i]{
                break
            }
        }

        if i == length {
            fmt.Println(true)
            return
        }

        target = target[i+1:]
    }

    fmt.Println(false)
}

/*----------------------------------------------------*/

func strstr(target string, sub string) {
    if len(target) == 0 || len(sub) == 0 {
        fmt.Println(false)
        return
    }

    var length = len(sub)
    for len(target)>=length {
        var i int
        for i=0; i<length; i++ {
            if target[i]!=sub[i]{
                break
            }
        }

        if i == length {
            fmt.Println(true)
            return
        }

        target = target[i+1:]
    }

    fmt.Println(false)
}

func main() {
    var target string = "hello world"
    var sub1 string = "world"
    var sub2 string = "china"

    strstr(target, sub1)
    strstr(target, sub2)
}