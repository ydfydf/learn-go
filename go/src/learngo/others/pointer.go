package others

import "fmt"

//Go语言的指针相当于C语言的指针，但是Go语言的指针不能进行运算，比如不能进行指针加减，
// 只能进行值传递（把参数进行副本拷贝），不能进行引用传递（C++能进行引用传递，使用&符号，没有进行副本拷贝，直接使用原值）

func swap(a,b int)  {
	b,a = a,b
}

func swap1(a,b *int){
	*b,*a = *a,*b
}

func swap2(a,b int ) (int ,int) {
	return b,a
}
func main() {
	a,b := 3,4
	swap(a,b)
	fmt.Println(a,b) //值没变

	swap1(&a,&b)
	fmt.Println(a,b) //值交换

	b,a = swap2(a,b)
	fmt.Println(a,b) //值交换,此方法更好
}
