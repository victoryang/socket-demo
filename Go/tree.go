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

func build_binary_tree(arr []int) *Node {
    var root *Node = nil

    for _,v := range arr {
        root = search_and_insert(v, root)
    }

    return root
}

func traverse_tree_pre(node *Node) {
    if node!= nil {
        fmt.Printf("%d ", node.val)

        traverse_tree_pre(node.left)
        traverse_tree_pre(node.right)
    }
}

func traverse_tree_suf(node *Node) {
    if node!= nil {
        traverse_tree_suf(node.left)
        fmt.Printf("%d ", node.val)
        traverse_tree_suf(node.right)
    }
}

func traverse_tree(root *Node) {
    fmt.Printf("traverse_tree_pre: ")
    traverse_tree_pre(root)
    fmt.Println()

    fmt.Printf("traverse_tree_suf: ")
    traverse_tree_suf(root)
    fmt.Println()
}

func main() {
    var arr []int = []int{3,4,2,5,6,1,7,9}
    var other int = 8

    fmt.Println("arr: ", arr)

    root := build_binary_tree(arr)

    traverse_tree(root)

    search_and_insert(other, root)

    traverse_tree(root)
}