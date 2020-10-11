package main

import(
	"sort"
)

type ArrayInt []int

func (a ArrayInt) Len() int {
	return len(a)
}

func (a ArrayInt) Swap(i,j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ArrayInt) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	arr1 := ArrayInt{4,3,2,1}
	sort.Sort(arr1)
	fmt.Println(arr1)

	arr2 := []int{4,3,2,1}
	sort.Ints(arr2)
	fmt.Println(arr2)

	arr3 := []float64{4.0,3.0,2.0,1.0}
	sort.Float64s(arr3)
	fmt.Println(arr3)
}