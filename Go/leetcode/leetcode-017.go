package main

// Letter Combination of a Phone Number

import "fmt"

var table = map[rune]string{
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
}

func try(layer int, src []string, res string, r *[]string) {
    if layer >= len(src) {
        *r = append(*r, res)
        return
    }

    for _,v :=range src[layer] {
        try(layer+1, src, res+string(v), r)
    }
}

func count(input string){
    // 确定解空间
    src := make([]string, 0)
    r := make([]string, 0)

    // 开始回溯
    for _, v :=range input {
        src = append(src, table[v])
    }

    try(0, src, "", &r)

    fmt.Println("res: ", r)
}

func main() {
    var input string
    fmt.Println("Please input the conbination: ")
    fmt.Scanf("%s", &input)

    count(input)
}