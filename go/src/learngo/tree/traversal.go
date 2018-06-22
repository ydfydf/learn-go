package tree

//树的中序遍历
func (node *Node) Traverse() {
	if node == nil {
		return

	}
	node.Left.Traverse()//递归,合理不要需要判断if node.left != nil，因为nil也可以调用函数
	node.Print()
	node.Right.Traverse()
}
