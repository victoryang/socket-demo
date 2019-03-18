package main

import (
    "fmt"
    "sync"
    "strconv"
)

const (
    thread_num = 3

    Task = 100
)

type ThreadPool struct{
    Lock        *sync.Mutex
    Cond        *sync.Cond
}

func main() {
    var lock = new(sync.Mutex)
    var p = ThreadPool{
        Lock: lock,
        Cond: sync.NewCond(lock),
    }

    var buf = 0

    var wg = new(sync.WaitGroup)
    for i:=0; i<thread_num; i++ {
        wg.Add(1)

        tag := i
        go func(tag int){
            defer wg.Done()
            for true {
                p.Lock.Lock()

                for buf%thread_num != tag && buf != Task{
                    p.Cond.Wait()
                }

                if buf == Task {
                    p.Lock.Unlock()
                    break
                }

                fmt.Println("print ", strconv.Itoa(buf), " from thread", strconv.Itoa(tag))
                buf++

                p.Lock.Unlock()

                p.Cond.Broadcast()
            }

            fmt.Println("Thread finish...")
        }(tag)
    }
    wg.Wait()

    fmt.Println("Leaving...")
}