package others

import "fmt"

//Go语言一般通过Slice切片的方式对数组进行操作、修改
//Slice本身没有数据，是对底层array的一个view
func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s := arr[2:6]//s即为切片，范围为半开半闭区间，即包含左边不包含右边
	fmt.Println("arr[2:6] = ",s)

	fmt.Println("arr[0:6] = ",arr[0:6])
	fmt.Println("arr[:6] = ",arr[:6])
	fmt.Println("arr[2:] = ",arr[2:])
	fmt.Println("arr[:] = ",arr[:])
	fmt.Println("========================================")

	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("s1 = ",s1)
	fmt.Println("s2 = ",s2)
	updateSlice(s1)
	updateSlice(s2)
	fmt.Println("After updateSlice")
	fmt.Println("s1 = ",s1)
	fmt.Println("s2 = ",s2)
	fmt.Println("arr = ",arr)
	fmt.Println("========================================")

	//Reslice,对切片进行再切片
	fmt.Println("Reslice")
	fmt.Println("s2 = ",s2)
	s2 = s2[:5]
	fmt.Println("s2 = ",s2)
	s2 = s2[2:]
	fmt.Println("s2 = ",s2)
	fmt.Println("arr = ",arr)
	fmt.Println("========================================")

	//Slice的扩展
	fmt.Println("Extending Slice")
	fmt.Printf("arr = %v,len(arr) = %d,cap(arr) = %d\n",arr,len(arr),cap(arr))
	s1 = arr[2:6]//注意这里s1只有4个元素
	fmt.Printf("s1 = %v,len(s1) = %d,cap(s1) = %d\n",s1,len(s1),cap(s1))
	s2 = s1[3:5]//按理说是取不到s1[4]的(不包含5)，但是Go语言slice可以根据原底层数组进行向后扩展(s[i]不可以超过len(s)，向后扩展不可以超越原底层数组cap(s))；不能向前扩展，也就是说切片的前面的切掉了就不能使用了，后面的能提供切片扩展
	//简而言之，slice后面，超过的区间会有原底层数组数据，但是不能使用slice[i](i>=len(slice))，也就是不能直接使用超过区间的数据，报错越界，但是超过的区间数据可以为slice提供切片扩展
	fmt.Printf("s2 = %v,len(s2) = %d,cap(s2) = %d\n",s2,len(s2),cap(s2))
	fmt.Println("========================================")

	fmt.Println("arr = ",arr)
	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)
	fmt.Println("s3, s4, s5 = ",s3, s4, s5)
	//s4和s5不再是arr的view，而是新分配的一个长于arr的view
	fmt.Println("arr = ",arr)

}
