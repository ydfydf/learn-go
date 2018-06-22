package queue

import (
	"fmt"
)

//此函数是一个特殊的test，其可以作为用户查看文档的Example（Go语言会自动生成此Example的文档）
//最后的//Output:的这个不是注释作用，其是一个固定格式(//Output:不能变)，用作此测试用例的输出，后面的输出值和上面的打印的值一致时，测试通过，所以不仅仅提供测试功能，
// 还会检查Output和实际输出的值是否一致
func ExampleUeue_Pop() { //直接打func Exam会自动补全，说明其是一段测试代码
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	//Output:
	//1
	//2
	//false
	//3
	//true
}