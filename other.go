package main

type TreeNode struct {
	left *TreeNode
	right *TreeNode
	val int
}

func convert(root *TreeNode) *TreeNode {
	var node *TreeNode
	Bst2List(root, &node)
	var header = node
	for header != nil && header.left != nil {
		header = header.left
	}
	return header
}

func Bst2List(root *TreeNode, prev **TreeNode) {
	if root == nil {
		return
	}
	// 当前节点
	var curr= root
	if curr.left != nil {
		Bst2List(curr.left, prev)
	}
	curr.left = *prev
	if *prev != nil {
		prev.right = curr
	}
	*prev = curr
	if curr.right != nil {
		Bst2List(curr.right, prev)
	}
}

func main() {
	root := &TreeNode{left:nil, right:nil, val: 1}
	convert(root)
}
