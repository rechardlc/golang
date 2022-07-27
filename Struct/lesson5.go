/*
	1.go只支持封装，不支持继承与多态
	2.go没有class，只有struct
*/
package main

import (
	"fmt"
)

// 二叉树结构体
type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.right = &treeNode{}
	root.left = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) // 无论指针还是非指针都使用(.)来访问
	root.right.right = createNode(10)
	fmt.Println("修改前：", root.right.left.value)
	root.right.left.setValue(100)
	fmt.Println("修改后", root.right.left.value)
}
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}
