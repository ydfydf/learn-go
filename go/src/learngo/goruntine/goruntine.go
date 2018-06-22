package main

import (
	"fmt"
	"time"
)

//func printHello(){}
func main() {
	for i := 0; i < 1000 ; i ++ {
		//go后面的func (){}是匿名函数，也可以写成go printHello()，不过go后面的函数一般写成匿名函数的样子
		go func(i int) {//如果不加go，这个函数func(i int)是一个死循环函数（里面有没有退出条件的for循环），加了go就不是去调这个函数，而是并发的执行这个函数，并发数为i
			for {
				fmt.Printf("Hello from " + "goruntine %d\n",i)//这里的i引用的是go func(i int)里面的i，如果用外面for循环的i，会不安全
			}
		}(i)
	}
	time.Sleep(time.Microsecond)//因为main函数和go func(i int) 是并发的执行，go func(i int) 都还没来得及执行，main函数就先退出了，所以加个睡眠，让main不要先退出
	//最后在time.Microsecond时间内十个协程不停的执行函数，直到sleep时间到，main退出
}



