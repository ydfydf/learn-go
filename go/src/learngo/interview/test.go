package main

import (
	"fmt"
	"sync"
	"runtime"
	"time"
)

type WorkerPool struct {
	tasks    <-chan *string //任务队列长度
	poolSize int            //启动goroutine的数目
}

func (p *WorkerPool) Run() {
	var wg sync.WaitGroup
	for i := 0; i < p.poolSize; i++ {
		wg.Add(1)
		fmt.Printf("i=%d\n",i)
		go func() {
			for task := range p.tasks {
				fmt.Println("Consume Task", *task)
			}
			wg.Done()

		}()

	}

	wg.Wait()
	fmt.Println("WorkerPool | Pool exit.")
}

func main() {
	taskNum := 100
	projects := make(chan *string, taskNum)
	//启动生产任务goroutine
	go func() {
		for i := 0; ; i++ {
			s := "Project " + fmt.Sprintf("%d", i)
			fmt.Println("Produce task", s)
			projects <- &s
			time.Sleep(5 * time.Nanosecond)
			fmt.Printf("NumGoroutine=%d\n",runtime.NumGoroutine())
		}
	}()

	p := WorkerPool{
		tasks:    projects,
		poolSize: 100,
	}
	p.Run()
}