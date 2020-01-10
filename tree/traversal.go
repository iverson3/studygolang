package tree

import "fmt"

// 遍历Node
func (node *Node) Traverse()  {
	if node == nil {
		return
	}
	// 中序遍历 (先左 再中间 再右)
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// 重写遍历Node的方法
func (node *Node) Traverse2()  {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node))  {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		// node遍历结束之后关闭channel
		close(out)
	}()
	return out
}









