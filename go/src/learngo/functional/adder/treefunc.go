package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left, Right *Node
}


func (node Node) Print() {
	fmt.Print(node.Value," ")
}

func CreateNode(value int) *Node {
	return &Node{Value:value}
}
//因为在此文件中没有定义函数类型为func(node *Node)的方法，要使用需通过匿名函数定义的方式进行使用
//在此函数中，将函数func(node *Node)以匿名函数方式进行定义，所以调用TraverseFunc里的f(node)时，使用的方法就是在Traverse定义的匿名方法
func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node){
		fmt.Println("11111111111111")
		node.Print()
	})
	fmt.Println("=========")
}

func (node *Node) TraverseFunc(f func(node *Node)) {
	if node == nil {
		return

	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node)TraverseWithChannel() chan *Node{
	c := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			c <- node
		})
		close(c)
	}()
	return c
}

func main() {
	var root Node

	root = Node{Value:3} //创建root，value为3，left和right都为nil
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = CreateNode(2)

	//root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *Node) {//在此函数中，TraverseFunc函数中f(node)使用的方法就是该匿名方法
		nodeCount++
	})
	fmt.Println("Node Count: ",nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Printf("Max node value:%d\n",maxNode)
}
