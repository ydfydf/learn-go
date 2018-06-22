package others

import "fmt"

//数组是值类型的，也就是进行了副本拷贝，会拷贝数组
func printArray(arr [5]int){
	arr[0] = 100//外面arr[0]的值未改变
	for _,v := range arr {//range返回第一个值为下表，第二个值为下标对应的值
		fmt.Println(v)
	}
}

func printArray1(arr *[5]int){
	arr[0] = 100//外面arr[0]的值改变，注意这里的写法，不是（*arr）[0]，直接使用arr[0]即可
	for _,v := range arr {//range返回第一个值为下表，第二个值为下标对应的值
		fmt.Println(v)
	}
}

//综上所诉，Go语言中[10]int和[20]int是不同类型，调用func f(arr [10]int)会拷贝数组，很不方便，所以
//在Go语言中一般不直接使用数组，而是用切片
func main() {
	//定义数组，数量写在类型前面
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 5, 6, 8, 9, 0} //编译器根据后面数组内容决定数组长度

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历数组
	for i := 0;i < len(arr3);i++{
		fmt.Println(arr3[i])
	}
	fmt.Println("================================")
	//遍历更好的方法
	for i := range arr3 {//range可以获得下标
		fmt.Println(arr3[i])
	}
	fmt.Println("================================")
	for i,v := range arr3 {//range返回第一个值为下表，第二个值为下表对应的值
		fmt.Println(i,v)
	}
	fmt.Println("================================")
	printArray(arr1)
	//printArray(arr3) 错误，因为arr3之有7个元素的数组，这里需要5个元素的数组，Go语言认为他们是不同的类型
	fmt.Println("================================")
	fmt.Println(arr1[0])
	fmt.Println("================================")
	printArray1(&arr1)
	fmt.Println(arr1[0])

}
