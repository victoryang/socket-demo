package main

/* Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

 * Follow up: The overall run time complexity should be O(log (m+n)).

 */

import "fmt"

func binarySearchMedian(num1 []int, num2 []int, min int, max int) float64 {
	i := (min+max)/2
	j := (len(num1) + len(num2))/2 - i

	var left_a
	if i == 0 {
		
	}

	if num1[i] < 
}

func findMedianSortedArrays(num1 []int, num2 []int) float64 {
	var result float64

	len1 := len(num1)
	len2 := len(num2)

	if len1 > len2 {
		num1, num2 := num2, num1
	}

	return binarySearchMedian(num1, num2, 0, len(num1))
}

func main() {
	num1 := []int{1,2,3,4,5,10}
	num2 := []int{6,7,8,9}

	fmt.Println(findMedianSortedArrays(num1, num2))
}