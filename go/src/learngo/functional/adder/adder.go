package main

import "fmt"

//闭包
func adder() func(int) int{
	sum := 0//sum是自由变量
	fmt.Println("==========")
	return func(v int) int {
		fmt.Println(">>>>>>",sum)
		sum += v //v可以看作是局部变量
		return sum
	}
}

func main() {
	//根据打印结果可以理解代码的运行过程为：a := adder()，a调用adder()函数，返回类型为func(int) int的函数，sum被初始化为0并保存
	//然后调用a(1),a(2)等其实只是调用函数func(v int) int，当然自由变量sum会被保持
	a := adder()
	fmt.Printf("%T\n",a)//打印func(int) int
	for i := 0 ;i < 10 ;i ++ {
		fmt.Printf("0 + ... + %d = %d\n",i,a(i))
	}
	fmt.Println(a(100))
}
