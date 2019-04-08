package main

import "fmt"

func longest_substring_without_repeating(str string){
	var length = len(str)

	i := 0
	j := i + 1
	res := make(map[rune]bool, 0)
	for i<length && j<length {
		if str[j] == str[j-1] {
			for _,v := range str[i:j] {
				if _,ok := res[v];!ok {
					res[v] = true
				}
			}
			i = i + 1
			j = i + 1
		} else {
			j = j + 1
		}
	}
	
	for _,v := range str[i:j] {
		if _,ok := res[v];!ok {
			res[v] = true
		}
	}

	var r string
	for k,_ :=range res {
		r = r + string(k)
	}
	fmt.Println("r: ", r)
}

func main() {
	var a_str = "abcabcbb"
	var b_str = "bbbbb"
	var c_str = "pwwkew"

	longest_substring_without_repeating(a_str)
	longest_substring_without_repeating(b_str)
	longest_substring_without_repeating(c_str)
}