package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("aaa", "bbb")
	fmt.Println("a: ", a, "b: ", b)
}
