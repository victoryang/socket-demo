package main

import (
    "fmt"
)

const (
    INIT_STATE = iota
    WORD_STATE
    SPACE_STATE
)

func word_count(input string) int {
    var count int
    var state = INIT_STATE

    for i:=0; i<len(input); i++ {
        switch state {
        case INIT_STATE:
            if input[i] == ' '{
                state = SPACE_STATE
            } else {
                state = WORD_STATE
            }

        case WORD_STATE:
            if input[i] == ' '{
                count++
                state = SPACE_STATE
            } else {
                continue
            }

        case SPACE_STATE:
            if input[i] == ' '{
                continue
            } else {
                state = WORD_STATE
            }
        }
    }

    if state == WORD_STATE {
        count++
    }

    return count
}

func main() {
    var input string = "sa sdf dfd ss sss dafdfadf ewer asfadf assdfadf   asdfasdf sdfd     sdfafa  asdfaf"

    fmt.Println("words count: ", word_count(input))
}