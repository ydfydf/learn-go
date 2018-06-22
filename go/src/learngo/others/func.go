package others

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

//Go语言返回多个值
func div(a,b int) (int ,int ){
	return a / b, a % b
}

//给返回值起名字,可以使用形参直接进行赋值，但是这种方法不适用于过长的函数，因为不方便查看返回值到底是在哪里赋的值
func div2(a,b int) (q,r int ){
	q = a / b
	r = a % b
	return
}

//函数返回多个值，除了上面的算术用法外，一般用于返回错误error(最常用)
func div3(a,b int) (int,error){
	if a < 0  || b < 0 || b == 0{
		//返回error使用fmt.Errorf
		return -1,fmt.Errorf("a or b cannot be less than 0 and b cannot be equal 0")
	}else {
		return a / b,nil
	}
}

//函数式编程,Go语言是一门函数式编程语言，它的参数、返回值、甚至是函数体内都可以有函数(匿名函数)
func apply(op func(int,int) (int,error) ,a,b int ) (int, error) { //注意函数apply的返回值数量和类型要和op函数的返回值相同
	//获取调用的当前函数名与其参数
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args " +
		"(%d,%d)\n",opName,a,b)
	return op(a,b)
}

//重写pow函数
func pow(a,b int) int {
	return int(math.Pow(float64(a),float64(b)))
}

func apply1(op func(int ,int) int,a,b int) int{
	return op(a,b)
}

//Go语言没有默认参数，可选参数，函数重载等，但有可变参数列表
func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}

func main() {
	fmt.Println(div(3,4))

	//接收全部返回值
	q, r := div2(13,3)
	fmt.Println(q,r)

	//接收不完全的返回值，如果不想接收所有的返回值，则使用 _ 来表示跳过不想要接收的值
	a,_ := div2(13,3)
	fmt.Println(a)

	if b,c := div3(3,1);c != nil {
		fmt.Println(c)
	}else {
		fmt.Println(b)
	}

	if d,e := apply(div3,4,2);e != nil {
		fmt.Println(e)
	}else {
		fmt.Println(d)
	}

	fmt.Println(apply1(pow,3,4))

	//函数体内定义函数，即匿名函数
	fmt.Println(apply1(
		func(a,b int) int { //匿名函数不需要实际的函数名，直接func(参数) 返回值就可以了
			return int(math.Pow(float64(a),float64(b)))
		},3,4,//注意这里逗号，如果4后面换行后才加))，就需要逗号，如果4后面不换行直接加))，就不需要逗号
		))

	fmt.Println(sum(1,2,3,444,55,666,8889,990))

}
