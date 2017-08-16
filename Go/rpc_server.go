package main

import(
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Server struct {
	address string
	ch chan string
	l net.Listener
}

type Args struct {
	A int
}

type Repl struct {
	R bool
}

func (master *Server) DoHello(args *Args, rep *Repl) error {
	if args.A == 1 {
		fmt.Println("hello world")
		rep.R = true
	} else {
		rep.R = false
		fmt.Println("err")
	}
	return nil
}

func main(){
	srv := new(Server)
	srv.address = "Server"
	srv.ch = make(chan string)
	go newWorker("worker1", srv.ch)
	rpc.Register(srv)
	os.Remove(srv.address)
	l, e := net.Listen("unix", srv.address)
	if e != nil {
		fmt.Println("err!")
	}
	fmt.Println("listening...")
	srv.l = l
	srv.ch <- srv.address
	go func(){
			for {
				conn, err:= srv.l.Accept()
				if err == nil {
					go func(){
						rpc.ServeConn(conn)
						conn.Close()
					}()
				}
			}
		}()
	Reply := <- srv.ch
	if Reply == "good" {
		fmt.Println("exiting...")
	}
}