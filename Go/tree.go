package main

import (
    "fmt"
)

type Node struct {
    val     int
    left    *Node
    right   *Node
}

func create_node(val int) *Node {
    node := new(Node)
    node.val = val
    node.left = nil
    node.right = nil

    return node
}

func search_and_insert(val int, node *Node) *Node {
    if node == nil {
        node = create_node(val)
        return node
    }

    if node.val > val {
        node.left = search_and_insert(val, node.left)
    } else {
        node.right = search_and_insert(val, node.right)
    }

    return node
}

func build_binary_tree(arr []int) {
    var root *Node = nil

    for _,v := range arr {
        root = search_and_insert(v, root)
    }

    traverse_tree(root)

    fmt.Println()
}

func traverse_tree(node *Node) {
    if node!= nil {
        fmt.Printf("%d ", node.val)

        traverse_tree(node.left)
        traverse_tree(node.right)
    }
}

func main() {
    var arr []int = []int{3,4,2,5,8,6,1,7,9}

    build_binary_tree(arr)

    fmt.Println("arr: ", arr)
}