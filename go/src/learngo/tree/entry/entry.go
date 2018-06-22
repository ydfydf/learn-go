package main

import (
	"fmt"
	"learngo/tree"
)

//通过组合的方式拓展类型
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder(){
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
	//创建结构体的几种方法
	root = tree.Node{Value:3} //创建root，value为3，left和right都为nil
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //无论地址还是结构本身，一律使用 . 来访问成员
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value:3},
		{},//在slice里，可以省略treeNode这几个字（结构体名字），{}表示类型为结构体的数都为nil
		{6,nil,&root},
	}
	fmt.Println(nodes)//打印出 [{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc0420023e0}]

	fmt.Println("========================================")
	root.Print()
	root.Right.Left.SetValue(4)//原值为0
	root.Right.Left.Print()
	//注意setValue和print函数的参数一个为指针，一个为值，但是使用方法都是相同的，是因为Go语言编译器进行了自动解析，root.right.left是地址，
	// 调用print时编译器会将地址自动转换成指针所指向的值然后进行调用

	root.SetValue(10)
	root.Print()
	pRoot := &root
	pRoot.SetValue(20)
	pRoot.Print()

	var p1Root *tree.Node
	//nil指针也可以调用方法
	p1Root.SetValue(100)
	p1Root = &root
	p1Root.SetValue(200)
	p1Root.Print()
	fmt.Println("========================================")

	root.Traverse()
	fmt.Println()

	p2Root := myTreeNode{&root}
	p2Root.postOrder()
	fmt.Println()


}

