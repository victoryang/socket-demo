package main

/**
 * Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

 * You may assume that each input would have exactly one solution, and you may not use the same element twice.

 * You can return the answer in any order.

 */

import "fmt"

func twoSum(nums []int, target int) []int {
    i := 0
    n := len(nums)
    t := make(map[int]int)
    for i < n {
        k := target - nums[i]
        t[k] = i
        i = i + 1
    }

    j := 0
    var res []int
    for j < n {
        k := nums[j]
        if v,ok := t[k]; ok {
        	if v != j {
                res = []int{v,j}
            }
        }
        j = j + 1
    }

    return res
}

func main() {
    arr1 := []int{2, 7, 11, 15}
    target1 := 9

    res := twoSum(arr1, target1)
    fmt.Println(res)
}