package main

import(
	"fmt"
	"net/rpc"
)

type client struct {
	name string
	channel chan string
	srv_address string
}

func newWorker(name string, c chan string) {
	worker := new(client)
	worker.name = name
	worker.channel = c
	worker.srv_address = <- c
	fmt.Println(worker.srv_address)
	caller, err := rpc.Dial("unix", worker.srv_address)
	if err != nil {
		fmt.Println("err")
	}
	var reply Repl
	fmt.Println("before calling....")
	err = caller.Call("Server.DoHello", &Args{1}, &reply)
	fmt.Println("calling....")
	if err == nil {
		fmt.Println("successed!")
		c <- "good"
	} else {
		fmt.Println("error during rpc call: ",err)
	}
	caller.Close()
}