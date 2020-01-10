package main

import (
	"awesomeProject1/tree"
	"fmt"
)

// 如何扩展系统类型或别人的类型
// 1.定义别名  2.使用组合

// 使用组合的方式扩展已有的类型或结构
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	fmt.Println(root)

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{} // 因为left是指针类型 所以要取地址赋值给left
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value: 12},
		{},
		{33,nil,&root},
		{15, nil,nil},
	}

	fmt.Println(root)
	fmt.Println(nodes)

	root.Print()

	root.Right.Left.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount: ", nodeCount)

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()


	// 使用goroutine和channel遍历tree
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value: ", maxNode)

}














