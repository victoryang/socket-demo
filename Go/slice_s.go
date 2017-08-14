package main

import (
	"fmt"
	//"log"
)

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
}
