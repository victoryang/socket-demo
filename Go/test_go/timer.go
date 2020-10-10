package main

import "fmt"
import rand "math/rand"
import "time"


func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	i:=0
	for {
		n := rand.Int63n(1000)
		time.Sleep(time.Duration(n) * time.Millisecond)
		fmt.Println("1000ms")
		i++
		if i == 5 {
			break
		}
	}
	fmt.Println("quit...")

	fmt.Println("start a new test")
	t := time.NewTimer(time.Duration(3)*time.Second)
	<-t.C
	fmt.Println("3 second")
}