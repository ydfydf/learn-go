package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) //随机睡眠1500毫秒以内的时间
			out <- i
			i ++
		}
	}()
	return out
}

func worker(id int,c chan int) {
	for n:= range c{
		time.Sleep(time.Second)
		fmt.Printf("Worker:%d received %d\n",id,n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id,c)
	return c
}

func main() {
	var c1, c2 = generator(),generator()
	var worker = createWorker(0)

	var values []int//将channel c1、c2中的数据存起来，因为如果不存起来，c1、c2产生数据速度太快，worker消耗速度不够，会造成数据被冲刷掉
	tm := time.After(10 * time.Second) //定时器，10秒后返回一个channel time
	tick := time.Tick(time.Second) //定时器，每1s会返回一个tick channel
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		//select运行机制是同时等待多个channel返回，当有channel返回时，select就获取其值做出相应处理，然后退出select语句，所以一般配合for循环进行使用，当
		//select收到一个channel的值返回后，由于外面嵌套了for循环，又会进入select语句中去
		select {//从多个channel里接收数据
		//case语句会阻塞住来等待channel的返回
		case n := <-c1:
			//w <- n //希望将n再发送到w，但是这样的话这里又变成阻塞式运行，所以需要将w <- n进行case，让select来管理
			values = append(values, n)
		//另外一个select重要特性，当case的channel为nil时，也就是只定义而未创建时（var n chan int，没有make），case语句会一直阻塞不会执行
		case n := <-c2:
			//w <- n
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond)://当select前一次从channel获取数据，800毫秒后没有channel的数据返回，就打印Time out
			fmt.Println("Time out")
		case <-tick:
			fmt.Println("queue len = ",len(values))
		case <-tm://当从channel tm中接收到数据时，说明定时器已返回，然后进行函数退出
			fmt.Println("bye ydf")
			return

		//default://defaut是当case的channel都没数据来时，它就会非阻塞式的执行defalut下的语句
			//time.Sleep(time.Microsecond)
			//fmt.Println("No value received")

		}
	}
}
