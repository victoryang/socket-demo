package main

import (
	"fmt"
	//"log"
)

func printCapIncreament() {
	fmt.Println("test reslice...")
	s := make([]int, 0, 1)
	c := cap(s)

	for i:=0; i<50; i++ {
		s = append(s,1)
		if c != cap(s) {
			fmt.Println("cap:", c, "->", cap(s))
			c = cap(s)
		}
	}
}

func testCopySlice(){
	fmt.Println("test copy slice...")

	data1 := []int{0,1,2,3,4,5,6,7,8,9}
	data2 := data1[5:8]

	copy(data1, data2)
	fmt.Println("copy data2 to data1:", data1)
}

func main(){
		a :=[]int{1,2,3,4}
		for i, j:= range a {
			fmt.Println(i,j)		
		}

		b := a[1:]
		for i,j := range b {
			fmt.Println(i,j)
		}

		kvs := map[string]int{"a":1,"b":2}
		for k,v := range kvs {
			fmt.Println(k,v);
		}

		data := [...]int{0,1,2,3,4,5,6,7,8,9}
		s1 := data[1:4:8]
		fmt.Println("data[1:4:8]: ", s1)

		s2 := s1[:5]
		fmt.Println("s1[:5]:", s2)

		printCapIncreament()

		testCopySlice()
}
