package others

import (
	"math/cmplx"
	"fmt"
	"math"
)

func complexTest(){
	c := 3 + 4i//3 + 4i这种表示方式，编译器就会知道它是复数，这里不能写成3 + 4*i，4*i的话i就为变量
	fmt.Println(cmplx.Abs(c))//复数的库为cmplx
}

func euler1(){
	//欧拉公式里面的虚数i要表示应该为1i，不能写成1*i，math.E代表数学中的e，math.Pi代表数学中的Π，函数Pow作用为求以e为底，iΠ为指数的值
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	//求出值为(0+1.2246467991473515e-16i)，实部为0，虚部为1.2246467991473515e-16i，因为complex128（复数类型）是由实部为float64和虚部为float64组成，遵循浮点数标准
}

func eular2(){
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)//函数Exp作用为 求默认e为底，参数为指数的值
	fmt.Printf("%.3f\n",cmplx.Exp(1i * math.Pi) + 1)
}

//强制类型转换
func triangle(){
	var a,b int = 3,4
	var c int
	//c = math.Sqrt(a * a + b * b),因为Go语言不存在隐式类型转换，所以这种方式错误
	c = int(math.Sqrt(float64(a * a + b * b)))//sqrt函数作用为开根号
	fmt.Println(c)
}

//常量
const filename1 = "aaa.txt" //包内常量
func consts(){
	//const filename = "ydf.txt"
	//const a, b = 3, 4
	//简便写法：
	const (
		filename = "ydf.txt"
		a,b = 3,4
	)
	var c int
	c = int(math.Sqrt((a * a + b * b)))
	//注意这里的a，b是没有定义类型的，数值可以作为各种类型使用，使用时就相当于做了文本替换工作，所以这里参数不需要强制转换成float64
	//如果定义const a, b int = 3, 4，c = int(math.Sqrt(float64(a * a + b * b)))，这里必须进行类型强制转换
	fmt.Println(filename1,filename,c)
}

//枚举
//Go语言没有专门的枚举类型，使用const块来进行定义使用
func enums(){
	//普通枚举类型
	const(
		c       = 0
		cpp     = 1
		java    = 2
		golang  = 3
	)
	fmt.Println(c,cpp,java,golang)

	//自增枚举类型
	const (
		c1  = iota   //元素iota代表这组const是自增值的
		_            //_代表跳过这个值，这里跳过1
		cpp1
		java1
		golang1
	)
	fmt.Println(c1,cpp1,java1,golang1)

	//iota可以参与运算
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	complexTest()
	euler1()
	eular2()
	triangle()
	consts()
	enums()
}
