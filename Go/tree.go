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

func search_and_insert(val int, pNode **Node) {
    if *pNode == nil {
        *pNode = create_node(val)
        return
    }

    node := *pNode
    if node.val > val {
        search_and_insert(val, &node.left)
    } else {
        search_and_insert(val, &node.right)
    }

    return
}

func build_binary_tree(arr []int) *Node {
    var root *Node = nil

    for _,v := range arr {
        search_and_insert(v, &root)
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

func find_item_in_tree(val int, pos *Node, node *Node) bool {
    if node==nil {
        return false
    }

    if node.val == val {
        pos = node
        return true
    }

    return find_item_in_tree(val, pos, node.left) || find_item_in_tree(val, pos, node.right)
}

func remove_from_tree(val int, root **Node) bool {
    var pos *Node = nil

    if root == nil || *root == nil {
        return false
    }

    if !find_item_in_tree(val, pos, *root) {
        return false
    }

    return true
}

func main() {
    var arr []int = []int{3,4,2,5,6,1,7,9}
    var other int = 8

    fmt.Println("arr: ", arr)

    root := build_binary_tree(arr)

    traverse_tree(root)

    search_and_insert(other, &root)

    traverse_tree(root)

    ret := remove_from_tree(other, &root)
    fmt.Println("try to find ", other, ": ", ret)
}