package main

import "fmt"
import "time"

func init(){
	fmt.Println("init1")
}

func init(){
	fmt.Println("init2")
}

func init(){
	fmt.Println("init3")
}

func testChanBuffer(){
	data := make(chan int)
	exit := make(chan bool)

	fmt.Println("start...")
	go func() {
		for v := range data {
			fmt.Println("data:", v)
		}

		fmt.Println("stops")
		exit <- true
	}()

	data <- 1
	data <- 2
	data <- 3
	close(data)

	<-exit
	fmt.Println("quit...")
}

func testSingleDirection(){
	c := make(chan int, 1)

	var send chan<- int = c
	var rcv <-chan int = c

	send <- 1
	d := <-rcv
	fmt.Println("receive:", d)
}

func testRandomRead(){
	a, b := make(chan int, 3), make(chan int, 1)

	go func() {
		v, ok, s := 0, false, ""

		for {
			select {
			case v,ok = <-a:
				s = "a"
			case v,ok = <-b:
				s = "b"
			}

			if ok {
				fmt.Println(s,v)
			} else {
				return
			}
		}
	}()

	for i:=0; i<5;i++ {
		select {
		case a<-i:
		case b<-i:
		}
	}

	close(b)
	close(a)
}

func factoryNewChan() chan int {
	c := make(chan int)

	go func(){
		c <- 1
	}()

	return c
}

func testFactoryMode(){
	c := factoryNewChan()
	dc := <-c
	fmt.Println(dc)
}

func testTimeout() {
	w := make(chan bool)
	c := make(chan int, 2)

	go func() {
		for {
			select {
			case v:= <-c:
				fmt.Println("get ", v)
			case <-time.After(time.Second*3):
				fmt.Println("timeout")
				goto QUIT
			}
		}

	QUIT:
		w<-true
	}()

	c<-1
	<-w

	fmt.Println("timeout caught")
}

func main() {
	testChanBuffer()

	testSingleDirection()

	testRandomRead()

	testFactoryMode()

	testTimeout()
}