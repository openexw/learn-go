package tree

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

/**
工厂函数：创建节点
@return 返回局部变量地址
*/
func createTreeNode(value int) *Node {
	return &Node{value: value}
}

/**
遍历，值传递
Note: print在后面 后面使用root.print() 调用
*/
func (node Node) Print() {
	fmt.Print(node.value, " ")
}

/**
设置 value
*/
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Errorf("%s", "Error， Node is not nill")
		return
	}
	node.value = value
}

/**
遍历
*/
func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.left.Traverse()
	node.Print()
	node.right.Traverse()
}
