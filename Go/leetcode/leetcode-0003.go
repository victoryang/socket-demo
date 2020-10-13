package main

// Given a string s, find the length of the longest substring without repeating characters.

import "fmt"

/*
func lengthOfLongestSubstring(s string) int {
    length := len(s)
    t := 0
    maximumWindowLength := 0
    
    for t < length {
        j := t + 1
        for j < length {
                i := t
                window := 0
                for j < length && s[i] == s[j] && s[j-1] != s[j] {
                        window = window + 1
                        i = i + 1
                        j = j + 1
                }

                if window > maximumWindowLength {
                        maximumWindowLength = window
                }

                j = j + 1
        }
        
        t = t + 1
    }
    
    return maximumWindowLength
}
*/

func lengthOfLongestSubstring(str string){
	length := len(s)
    if length == 0 || length == 1{
            return length
    }

    t := 0
    maximumWindowLength := 0
    
    for t < length {
        subs := make(map[byte]bool)
        window := 0
        j := t
        for j < length {
                if _,ok := subs[s[j]]; !ok {
                        subs[s[j]] = true
                        window = window + 1
                } else {
                        break
                }

                j = j + 1
        }

        if window > maximumWindowLength {
                maximumWindowLength = window
        }
        
        t = t + 1
    }

    return maximumWindowLength
}

func main() {
	var a_str = "abcabcbb"
	var b_str = "bbbbb"
	var c_str = "pwwkew"

	lengthOfLongestSubstring(a_str)
	lengthOfLongestSubstring(b_str)
	lengthOfLongestSubstring(c_str)
}