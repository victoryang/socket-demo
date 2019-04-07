package main

import "fmt"

type Node struct {
	val 	int
	next 	*Node
}

func make_list(arr []int) *Node {
	var list = new(Node)
	list.val = 0
	list.next = nil

	for _,v := range arr {
		tmp := new(Node)
		tmp.val = v
		tmp.next = list.next
		list.next = tmp
	}

	return list
}

func print_list(list *Node){
	var t *Node = list.next

	fmt.Print("List:")
	for t!=nil {
		fmt.Print(" ", t.val)
		t = t.next
	}
	fmt.Print("\n")
}

func add_two_numbers(a,b *Node) *Node {
	var pa = a.next
	var pb = b.next

	var pc = new(Node)
	pc.val = 0
	pc.next = nil
	var c *Node = pc

	var carry = 0
	for pa != nil || pb != nil {
		var x int
		if pa != nil {
			x = pa.val
		}

		var y int
		if pb != nil {
			y = pb.val
		}

		sum := carry + x + y
		carry = sum / 10

		tmp := new(Node)
		tmp.val = sum % 10
		tmp.next = c.next
		c.next = tmp
		c = tmp

		pa = pa.next
		pb = pb.next
	}

	return pc
}

func main() {
	a := [3]int{3,4,2}
	b := [3]int{4,6,5}

	c := make_list(a[:])
	d := make_list(b[:])

	print_list(c)
	print_list(d)

	e := add_two_numbers(c, d)
	print_list(e)
}