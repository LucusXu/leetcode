package main

type TNode struct {
	left *TNode
	right *TNode
	val int
}

func twoNodeDistance(root, p, q *TNode) int {
	// 最近公共祖先
	parent := getFirstCommonNode(root, p, q)
	dp := distance(parent, p)
	dq := distance(parent, q)
	if dp == -1 || dq == -1 {
		return -1
	}
	return dp + dq
}

func getFirstCommonNode(root, p, q *TNode) *TNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	left := getFirstCommonNode(root.left, p, q)
	right := getFirstCommonNode(root.right, p, q)
	if left == nil && right == nil {
		return nil
	} else if right == nil {
		return left
	} else if left == nil {
		return right
	}
	return root
}

func distance(root, p *TNode) int {
	if root == p {
		return 1
	}
	// 计算层
	var level = 0
	// 队列
	var queue []*TNode
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {
			tmp := queue[0]
			queue = queue[1:]
			if tmp.left != nil {
				queue = append(queue, tmp.left)
			}
			if tmp.right != nil {
				queue = append(queue, tmp.right)
			}
		}
	}
	return -1
}

func main() {
	root := &TreeNode{left:nil, right:nil, val: 1}
	convert(root)
}
