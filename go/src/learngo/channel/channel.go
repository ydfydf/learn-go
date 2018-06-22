package main

import (
	"fmt"
	"time"
)

//channel创建的时候，一般要先指定一个任务去阻塞的从该channel收数据，这样后面往该channel发数据时才不会造成dead lock
//比如channelDemo函数，生成的channel没人去接收，反而是往该channel发数据后才去收数据，这样就造成了dead lock
//但是有一种方法也可以避免channelDemo造成dead lock，就是再生成channel时制定缓冲区，这样就算该channel刚开始没有指定接收者，
//在往该channel发送不超该缓冲区大小的数据也不会造成dead lock

func channelDemo(){
	var c chan int//定义一个channel，channle里面的类型为int，此时的c == nil，channel还未被创建出来
	//c := make(chan int)
	c = make(chan int) //创建channle
	c <- 1 //往channel c发数据
	c <- 2
	n := <- c //从channel接收数据
	fmt.Println(n)
}
//函数channelDemo运行会造成deadlock，因为channel是goroutine之间的通道，一个协程发送了数据，必须要有一个协程来接收数据，不然会造成死锁

func channelDemo1(){
	c := make(chan int)
	go func() { //生成另外一个协程来收数据
		for {
			n := <- c
			fmt.Println(n) //可能只打印出一个1，2还来不及打印函数就返回了
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Microsecond)
}

//=============================================================================

//channel做参数
func worker(i int,c chan int) {
	for {
		fmt.Printf("Worker:%d received %c\n",i,<- c)
	}
}

func channelDemo2(){
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i,channels[i])
	}

	for i := 0 ; i < 10 ; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0 ; i < 10 ; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Microsecond)
}
//=============================================================================
//channel作返回值
func createWorker(i int) chan int {
	c := make(chan int)
	go func() {//创建channel后再创建相应的工作协程，常用方法
		for {
			fmt.Printf("Worker:%d received %c\n",i,<- c)
		}
	}()
	return c
}

func channelDemo3(){
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0 ; i < 10 ; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0 ; i < 10 ; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Microsecond)
}
//=============================================================================
//channel作返回值时，为了让使用者清晰的知道这个channel的作用，所以需要进一步修饰
func createWorker1(i int) chan<- int { //chan<-代表往channel发数据，<-chan代表从channel收数据
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker:%d received %c\n",i,<- c)
		}
	}()
	return c
}

func channelDemo4(){
	var channels [10]chan<- int //明确定义的channel作用，收还是发，因为createWorker1返回的是往channel发数据，所以定义channels必须明确作用，不然会报错
	for i := 0; i < 10; i++ {
		channels[i] = createWorker1(i)
	}

	for i := 0 ; i < 10 ; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0 ; i < 10 ; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Microsecond)
}
//=============================================================================
//channel的缓冲区
//像上面的例子，往channel发数据，必须要有协程从这个channel里收数据，会造成一些资源损耗，所以加入缓存区
//加入缓冲区后，往channel发数据，即使没有协程从channel里收（前提发送的数据要小于等于缓冲区大小），运行也不会deadlock
func bufferedChannel(){
	c := make(chan int, 3)
	go worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	//再没有协程接收数据时，往缓冲区大小为3的channel c发送第四个数据时，会造成deadlock
	//c <- 4
}

func bufferedChannel1(){
	c := createWorker2(0,3)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	//再没有协程接收数据时，往缓冲区大小为3的channel c发送第四个数据时，会造成deadlock
	c <- 'd'
	time.Sleep(time.Microsecond)
}

func worker1(i int,c chan int) {
	for {
		fmt.Printf("Worker:%d received %c\n",i,<- c)
	}
}
func createWorker2(i int,bufferSize int) chan int {
	c := make(chan int,bufferSize)
	go worker1(i,c)
	return c
}

//=============================================================================
//通知接收方数据发完了
func channelClose(){
	c := createWorker3(0,3)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) //只是channel的发送方关闭了发送功能，接收方还是能一直接收，只不过接收的是空值
	time.Sleep(time.Second)
}

func worker2(i int,c chan int) {
	for {
		n ,ok := <- c
		if  !ok  {
			break
		}
			fmt.Printf("Worker:%d received %c\n", i,n)
		}

	}
//等同于worker3的另一种写法
func worker3(i int,c chan int) {
	for n := range c {
		fmt.Printf("Worker:%d received %c\n", i,n)
	}

}
func createWorker3(i int,bufferSize int) chan int {
	c := make(chan int,bufferSize)
	go worker2(i,c)
	return c
}


func main() {
	channelDemo()
	//channelDemo1()
	//channelDemo2()
	//channelDemo3()
	//channelDemo4()
	//bufferedChannel()
	//bufferedChannel1()
	channelClose()
}
