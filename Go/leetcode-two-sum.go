package main

import "fmt"

func try(i int, target int, arr [4]int, res *[][]int) {
	for j:=i +1; j<4; j++ {
		if arr[i] + arr[j] == target {
			r := make([]int, 0)
			r = append(r, arr[i], arr[j])
			*res = append(*res, r)
		}
	}
}

func two_sum(arr [4]int, target int) {
	res := make([][]int, 0)
	for i:=0; i<3;i++ {
		try(i, target, arr, &res)
	}

	fmt.Println("res: ", res)
}

func main() {
	arr := [4]int{2, 7, 11, 15}
	target := 9

	two_sum(arr, target)
}