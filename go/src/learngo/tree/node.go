package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

//给结构体定义函数
func (node Node) Print() { //(node treeNode)是接收者，准确的说是值接收者，调用方法为root.print()；也可以写成func print()(node treeNode)，调用方法就为print(root)
	fmt.Print(node.Value," ")
}

//Go语言所有参数都是值传递，结构体定义函数也是，所有需要使用指针
//(node *treeNode)是接收者，准确的说是指针接收者
//一致性：如果有指针接收者，最好都用指针接收者
func (node *Node) SetValue(value int) {
	//nil指针也可以调用方法
	if node == nil {
		fmt.Println("Setting value to nil node.Ignored")
		return
	}
	node.Value = value//直接使用node.value，而不是node->value
}


//自定义工厂函数
func CreateNode(value int) *Node {
	return &Node{Value:value}
	//注意这里的创建的treeNode{value:value}是局部变量，返回局部变量的地址，在C/C++是不允许的，但是在Go语言中可以这样做，可以返回局部变量
	//那么Go语言中结构创建在堆上还是栈上？
	//答案是不需要知道，它是由Go语言的编译器和编译环境来共同决定的，比如你创建了一个函数，没有返回局部变量，那么就会在栈上分配内存，而像此例中返回了局部变量，那么就会在堆上分配内存，会有垃圾回收机制来进行资源回收
}
