package main

import "fmt"
import "reflect"

type test struct {
	i int
}

func (a *test) read() {
	fmt.Println(a.i)
}

func read2(a interface{}) {
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).Type().Name())
}

func main(){
	a := &test{3}
	a.read()
	read2(a)
}