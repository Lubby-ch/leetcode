package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		Node := &TreeNode{
			Val:  val,
			Left: root,
		}
		root = Node
		return root
	}
	var List = []*TreeNode{root}
	var layer = 1
	var nodeNum = 1
	for layer < depth-1 {
		var tempNum int
		for i := 0; i < nodeNum; i++ {
			if List[0].Left != nil {
				List = append(List, List[0].Left)
				tempNum++
			}
			if List[0].Right != nil {
				List = append(List, List[0].Right)
				tempNum++
			}
			List = List[1:]
		}
		nodeNum = tempNum
		layer++
	}
	for i := 0; i < nodeNum; i++ {
		List[0].Left = &TreeNode{
			Val:  val,
			Left: List[0].Left,
		}
		List[0].Right = &TreeNode{
			Val:   val,
			Right: List[0].Right,
		}
		List = List[1:]
	}
	return root
}

// 深度优先
func addOneRow1(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		Node := &TreeNode{
			Val:  val,
			Left: root,
		}
		root = Node
		return root
	}
	var dfs func(node *TreeNode, layer int)
	dfs = func(node *TreeNode, layer int) {
		if layer == depth-1 {
			node.Left = &TreeNode{
				Val:  val,
				Left: node.Left,
			}
			node.Right = &TreeNode{
				Val:   val,
				Right: node.Right,
			}
			return
		}
		if node.Left != nil {
			dfs(node.Left, layer+1)
		}
		if node.Right != nil {
			dfs(node.Right, layer+1)
		}
	}
	dfs(root, 1)
	return root
}
