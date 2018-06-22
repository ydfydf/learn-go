package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

//闭包
//斐波那契数列
//1,1,2,3,5,8,13,21,34...
//  a,b
//    a,b
func fibonacci() func() int {
	a ,b := 0,1
	return func() int{
		a ,b = b, a+ b
		return a
	}
}
type intGen func() int//intGen是一个函数的类型(func() int)，在Go语言中，只要是类型就能实现接口，很灵活

func fibonacci1() intGen {
	a ,b := 0,1
	return func() int{
		a ,b = b, a+ b
		return a
	}
}



//为函数实现接口
func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0,io.EOF
	}
	s := fmt.Sprintf("==%d==\n",next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {//参数reader io.Reader是接口
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci1()

	printFileContents(f)
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())



}
