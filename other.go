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

type ListNode struct{
	val int
	next *ListNode
}

func listToBst(head *ListNode) *TreeNode {
	var size int = 0
	var curr = head
	// 统计长度
	for curr != nil {
		size++
		curr = curr.next
	}
	curr = head
	return buildTree(curr, size)
}

func buildTree(list *ListNode, size int) *TreeNode {
	if size == 0 {
		return nil
	}
	var root = &TreeNode{}
	root.left = buildTree(list, size / 2)
	root.val = list.val
	list = list.next
	root.right = buildTree(list, size - size / 2)
	return root
}

func main() {
	root := &TreeNode{left:nil, right:nil, val: 1}
	convert(root)
}
