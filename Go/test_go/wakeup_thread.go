package main

import "fmt"
import "sync"
import "time"

func main(){
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	a := 2
	isChanged := false

	go func(){
		for {
			time.Sleep(1000 * time.Millisecond)
			cond.L.Lock()
			
			a = 1
			isChanged = true
			cond.Broadcast()
			cond.L.Unlock()
		}
	}()

	cond.L.Lock()
	fmt.Println("waiting...")
	if !isChanged {
		cond.Wait()
	}
	fmt.Println("times up: ", a)
	cond.L.Unlock()
	//cond.Wait()
}