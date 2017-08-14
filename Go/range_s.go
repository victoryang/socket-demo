package main

import(
	"fmt"
	"os"
	"log"
)


func main() {
	a := []int{1,2,3,4}
	for i,j := range a {
		log.Printf("array[%d]=%d",i,j);
	}
	fmt.Println(os.Getenv("GOROOT"))
	b:="asdfdf"
	fmt.Println(b)
	var c *int
	d:=1
	c = &d
	fmt.Println(c)
}
