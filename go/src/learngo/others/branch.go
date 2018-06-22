package others

import (
	"fmt"
	"io/ioutil"
)

//if条件语句
func bounded(v int) int { //参数为参数值 参数类型，返回值类型在最后面
	if v > 100 { //if的条件里不需要小括号()
		return 100
	}else if v < 0 {
		return 0
	}else {
		return v
	}
}

//switch条件语句
//switch会自动break，不需要手动添加，除非使用fallthrough
func eval(a,b int,op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)//panic为报错函数，中断程序的执行
	}

	return result
}

//switch的另一种写法
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score:%d\n",score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	//default:
	//	panic(fmt.Sprintf("Wrong Score:%d\n",score))
	}
	return g
}

func main() {
	var a int;
	a = bounded(2)
	fmt.Println(a)

	const filename = "helloworld.go"
	/*contents , err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}*/
	//上面if的这种写法显得比较麻烦，Go语言支持if的条件里进行赋值，注意contents的作用域为if条件句完后
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}
	//fmt.Printf("%s\n",contents)    错误，contents的作用域完后就被系统回收了v


	//switch
	a = eval(10,90,"+")
	fmt.Println(a)

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		//grade(101),
	)
}
