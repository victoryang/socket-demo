package main

import "fmt"
import "unsafe"

func add(a int, b int) (c int) {
	c = a + b
	return c
}

var e = "avc"

func main() {
	a := 1
	var b, c = 2, 0
	c = add(a, b)
	fmt.Println("a:", a, " b:", b)
	a, b = b, a
	fmt.Println("after change:", "a:", a, " b:", b)
	fmt.Println("unsafe.Sizeof e: ", unsafe.Sizeof(e))
	fmt.Println(c)
}
