package main

//使用者文件
import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
)

type Retriever123 interface { //接口名字可以与结构名不一样，叫Retriever123也行，最好相同，为了后面说明，特别把名字取不同
	Get(url string) string//不能加func关键字，因为接口里面全是函数，省略了
	//Get(url string) string函数的定义要与实现者定义的函数结构一致
}

func download(r Retriever123) string{
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever123
	//这里是关键，将结构与接口相关联，mock.Retriever返回接收者(r Retriever)；假如GET函数的接收者是指针接收者，func (r *Retriever) Get(url string) string
	//r = mock.Retriever{}会报错mock.Retriever does not implement Retriever123 (Get method has pointer receiver)，因为mock.Retriever定义的接收者是指针接收者
	//应该用r = &mock.Retriever{"this is a fake imooc.com"}，但是一般接口很少用指针接收者
	r = &mock.Retriever{"this is a fake imooc.com"} //这个时候r的接口类型就变为*mock.Retriever
	//r = mock.Retriever{"this is a fake imooc.com"}
	fmt.Println(download(r))
	fmt.Println(r.Get("www.imooc.com"))
	inspect(r)

	//获取接口里的值，方法二type assertion
	mR := r.(*mock.Retriever)
	fmt.Println(mR.Contents)

	q := mock.Retriever{"ydf"}//另一种写法,编译器会自动将r转换成指针接收者
	fmt.Println(q.Get("www.imooc.com"))
	//r = new(mock.Retriever)
	//fmt.Println(download(r))


	//指针接收者实现只能以指针方式使用；值接收者都可以
	r = &real2.Retriever{}//另一种写法，编译器会自动将r转换成指针接收者
	//fmt.Println(download(r))
	inspect(r)
	inspect(&q) //q是指针

	//获取接口里的值，方法二type assertion
	rR := r.(*real2.Retriever)
	fmt.Println(rR.TimeOut)

	//mR := q.(mock.Retriever)
}

//获取接口里的值，方法一
func inspect(r Retriever123) {
	fmt.Printf("%T %v\n",r,r)//%T打印类型type，%v打印值
	switch v := r.(type) { //接口r的类型
	case *mock.Retriever:
		fmt.Println("Conetents:",v.Contents)
	case real2.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)

	}
}
