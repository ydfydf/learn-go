package main

import (
	"fmt"
	"time"
	"sync"
)

//系统自带atomic库进行atomic操作(原子化操作)，作用为多个goroutine之间进行操作是安全的,
//比如atomic.AddInt32就是再多个goroutine之间进行Add时是安全的，在实际的使用中
//应该使用系统的atomic进行操作，此文件仅为测试：使用mutes实现atomicInt操作

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment(){
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value ++
}
//假如上面的increment函数很长，使用defer的话会执行完所有代码才会unlock，所以需要将defer控制在一块代码区内，
//让主要的加锁逻辑完后就解锁，需要使用匿名函数，代码改进如下
func (a *atomicInt) increment1(){
	//defer a.lock.Unlock()会在匿名函数完时执行
	func(){
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value ++
	}()
}


func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
