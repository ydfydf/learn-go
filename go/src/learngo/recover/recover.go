package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() { //recover仅在defer中调用
		r := recover()
		fmt.Printf("r.(error)=%v,%T\n",r.(error),r.(error))
		if err, ok := r.(error); ok {
			fmt.Println("Error Occurred:",err)
		}else {
			panic(r)
		}
	}()//此匿名函数写法等价于 defer handlePanic(){}

	panic(errors.New("this is an error"))
}//此函数最终打印出Error Occurred: this is an error，而不是直接挂掉并且报一堆难看红色错误代码出来

func tryRecover1() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error Occurred:",err)
		}else {
			panic(r)
		}
	}()

	b := 0
	a := 5 / b
	fmt.Println(a)
}//此函数最终打印出Error Occurred: runtime error: integer divide by zero，而不是直接挂掉并且报一堆难看红色错误代码出来

//repanic
func tryRecover2() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error Occurred:",err)
		}else {
			panic(fmt.Sprintf("I dont't know what to de:%v",r))

		}
	}()

	panic(123)
}//打印出panic: 123 [recovered]
//panic: I dont't know what to de:123，直接挂掉并且报一堆难看红色错误代码出来,
// 因为panic(123)在recover中是不知道什么类型的错误，上次的panic进入了reover中，所以需要再次panic

func main() {
	tryRecover()
	//tryRecover1()
	//tryRecover2()
}
