package main

// Remove Element

import "fmt"

func remove_item_from_unsorted(arr [10]int, num int) []int {
	var j int = 0
	for i:=0; i<10; i++ {
		if arr[i] != num && j!=i {
			j = j + 1
			arr[j] = arr[i]
		}
	}

	return arr[:j+1]
}

func main() {
	var arr = [10]int{0,0,1,1,1,2,2,3,3,4}

	fmt.Println("arr: ", arr)
	fmt.Println("Please input the element you want: ")

	var num int
	fmt.Scanf("%d", &num)

	res := remove_item_from_unsorted(arr, num)

	fmt.Println("res: ", res)
}