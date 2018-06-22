package others

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s = %v, len = %d, cap = %d\n",s,len(s),cap(s))
}

func main() {
	//定义s为一个slice
	var s []int //Go语言一个特性，Zero Value for slice is nil
	for i := 0;i < 100;i++ {
		printSlice(s) //首次打印出len = 0, cap = 0，表现出Zero Value for slice is nil，即切片为空时值为0
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	//slice的另一种创建方式
	s1 := []int{2,4,6,8}
	//slice的另一种创建方式,slice长度为16
	s2 := make([]int,16)
	//slice的另一种创建方式,slice长度为10，容量为32
	s3 := make([]int,10,32)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("========================================")
	fmt.Println("Copy Slice")
	copy(s2,s1)
	printSlice(s2)

	//通过append方式，变向将元素从slice中删除
	fmt.Println("Delete elements from Slice")
	s2 = append(s2[:3],s2[4:]...) //把元素8从s2中删除了
	printSlice(s2)

	fmt.Println("Poping from front")//删除首位
	front := s2[0]
	s2 = append(s2[1:])
	fmt.Printf("front = %d, s2 = %v\n",front, s2)

	fmt.Println("Poping from tail")//删除尾部
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) -1]
	fmt.Printf("tail = %d, s2 = %v\n",tail, s2)



}
