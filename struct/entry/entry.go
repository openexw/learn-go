package main

/**
扩展已有的包
*/
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	myTreeNode{myNode.node.left}.postOrder()
	myTreeNode{myNode.node.right}.postOrder()
	myNode.node.Print()
}
func main() {
	var root tree.Node

	root = TreeNode{value: 3}
	root.left = &TreeNode{}
	root.right = &TreeNode{5, nil, nil}
	root.right.left = new(TreeNode)
	root.left.right = createTreeNode(2)
	//fmt.Println(root)
	root.left.right.SetValue(12)
	//root.left.right.print()
	root.Traverse()
	/**
	   3
	/    \
	0     5
	\     /
	 2   0
	*/
}
