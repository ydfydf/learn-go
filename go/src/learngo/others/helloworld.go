package others

//**************************************
//Go语言学习之定义变量，2018/5/13
//**************************************

import "fmt"
//import . "fmt"            作用为省略fmt前缀，直接使用函数，如Println
//import f "fmt"            作用是为包fmt取一个别名f，使用函数时就为f.xxx，如f.Println

//在函数体外的包内（在go语言里没有全局变量这种说法）定义变量（此变量的作用域为当前的包），只能使用var进行定义，不能使用 := 这种方式
var aa = 1
var bb = true
var ss = "young"

//上面的方式频繁使用var显得繁琐，go语言定义包内变量一般使用如下方式集中定义变量
var (
	aaa = 11
	bbb = false
	sss = "ydf-young"
)

func variableInitialValue()  {
	var a int = 100
	var s string = "ydf"
	fmt.Println(a,s)
	//Println(a,s)
	//f.Println(a,s)
}

func variableZeroValue(){
	var a int
	var s string

	fmt.Printf("%d %q\n",a,s)
	//Printf("%d %q\n",a,s)
	//f.Printf("%d %q\n",a,s)
}

func variableTypeDeduction(){
	var a ,b ,c,s = 100,10,true,"ydf"  //编译器会根据后面的值自动为变量决定类型
	fmt.Println(a,b,c,s)
}

func variableShorter(){
	a ,b ,c,s := 100,10,true,"ydf" //在函数体内定义变量，一般用此方式，省略var
	b = 1000
	fmt.Println(a,b,c,s)
}

func main() {
	fmt.Println("hello world")
	//Println("hello world")
	//f.Println("hello world")

	variableInitialValue()
	variableZeroValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,ss,aaa,bbb,sss)
}
