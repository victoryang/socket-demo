package main

import "fmt"

func remove_duplicated_from_sorted(arr [10]int) []int{
	var j int = 0

	for i:=1; i<10; i++ {
		if arr[i] != arr[j] {
			j += 1
			arr[j] = arr[i]
		}
	}

	return arr[:j+1]
}

func main() {
	var arr = [10]int{0,0,1,1,1,2,2,3,3,4}

	fmt.Println("arr: ", arr)
	res := remove_duplicated_from_sorted(arr)
	fmt.Println("res: ", res)
}