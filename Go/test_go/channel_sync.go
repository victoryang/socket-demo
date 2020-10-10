package main

import(
	"fmt"
	"time"
)

func hello(c chan string){
	data := <-c
	if(data == "hello") {
		fmt.Println(data)
	}
}

func main() {
	c := make(chan string)

	go hello(c)

	c <- "hello"
	fmt.Println("send")
	time.Sleep(1000)
}
