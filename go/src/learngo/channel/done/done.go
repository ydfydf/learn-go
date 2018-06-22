package main

import (
	"fmt"
	"sync"
)

func doWorker(i int,c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker:%d received %c\n",i,n)
		//不要使用共享内存来通信，用通信来共享内存
		//协程通过channel来通知相应的协程数据处理完毕
		done <- true
	}
}

type worker struct {
	c chan int
	done chan bool
}

func createWorker(i int) worker {
	w := worker{
		c : make(chan int),
		done : make(chan bool),
	}
	go doWorker(i,w.c,w.done)
	return w
}

func channelDemo(){
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'a' + i
		//<-workers[i].done //按照这种方法，就相当于同步执行，一个workers往channel发数据，等待处理完后done为true，这样就是同步了，此法不可用
	}
	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'A' + i
		//<-workers[i].done
	}
	//等待所有协程全部处理完成
	for _, workeri := range workers{
		<-workeri.done
	}
}
//======================================================================================
func channelDemo1(){
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'a' + i
	}
	//等待上面10个协程全部处理完成
	for _, workeri := range workers{
		<-workeri.done
	}
	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'A' + i
	}
	//等待另外10个协程全部处理完成
	for _, workeri := range workers{
		<-workeri.done
	}
	//此发可行

	//下面这种方法就不行，会造成死锁
	//for _, workeri := range workers{
	//	<-workeri.done
	//	<-workeri.done
	//}
}

//======================================================================================
func doWorker2(i int,c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker:%d received %c\n",i,n)
		go func() {done <- true}() //也可以通过另起协程来异步实现将done置为true，此方法不好

	}
}

func createWorker2(i int) worker {
	w := worker{
		c : make(chan int),
		done : make(chan bool),
	}
	go doWorker2(i,w.c,w.done)
	return w
}

func channelDemo2(){
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'a' + i
	}
	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'A' + i
	}

}

//======================================================================================
//go语言提供WaitGroup来支持等待
type worker3 struct {
	c chan int
	wg *sync.WaitGroup
}

func doWorker3(i int,work worker3) {
	for n := range work.c {
		fmt.Printf("Worker:%d received %c\n",i,n)
		work.wg.Done()

	}
}

func createWorker3(i int,wg *sync.WaitGroup) worker3 {
	w := worker3{
		c : make(chan int),
		wg : wg,
	}
	go doWorker3(i,w)
	return w
}

func channelDemo3(){
	var workers [10]worker3
	var wg sync.WaitGroup
	wg.Add(20)//添加等待的任务数量
	for i := 0; i < 10; i++ {
		workers[i] = createWorker3(i,&wg)
	}

	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'a' + i
	}
	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'A' + i
	}
	wg.Wait()

}

//=============================================================
//上诉代码可以抽象成下列代码

type worker4 struct {
	c chan int
	done func() //将这里抽象成函数
}

func doWorker4(i int,work worker4) {
	for n := range work.c {
		fmt.Printf("Worker:%d received %c\n",i,n)
		work.done() //work.done()的实现wg.Done()，其中wg为函数createWorker4的参数wg *sync.WaitGroup

	}
}

func createWorker4(i int,wg *sync.WaitGroup) worker4 {
	w := worker4{
		c : make(chan int),
		done : func(){
			wg.Done() //注意，这里是结构体赋值，结构体内的函数done使用到了参数wg *sync.WaitGroup，使用这个赋值函数时，wg.Done()里的wg都会使用函数createWorker4的参数wg *sync.WaitGroup
		},
	}
	go doWorker4(i,w)
	return w
}

func channelDemo4(){
	var workers [10]worker4
	var wg sync.WaitGroup
	wg.Add(20)//添加等待的任务数量
	for i := 0; i < 10; i++ {
		workers[i] = createWorker4(i,&wg)
	}

	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'a' + i
	}
	for i := 0 ; i < 10 ; i++ {
		workers[i].c <- 'A' + i
	}
	wg.Wait()

}


func main() {
	channelDemo()
	channelDemo3()
	channelDemo4()
}