package main

import (
	"fmt"
	"os"
	"bufio"
)

func tryDefer() {
	fmt.Println(1)
	fmt.Println(2)
}
//使用defer关键字修饰的语句，会在函数结束时调用
func tryDefer1() {
	defer fmt.Println(1)
	fmt.Println(2)
}

func tryDefer2() {
	defer fmt.Println(1)//defer里面相当于有一个栈，先进后出，1先进，2后进，所以打印出2，1
	defer fmt.Println(2)
	fmt.Println(3)
	return
}

//参数在defer语句时计算，意思是每次调用defer语句时，参数i会进行计算，然后defer将运算结果保存在栈中，最后函数结束时执行defer修饰的函数
func tryDeferFunc() {
	for i := 0; i < 100; i ++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many values")
		}
	}
}

func writeFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//str := "111111111111111"
	//buf := []byte(str)
	//f.Write(buf)

	//这里使用bufio.NewWriter函数来进行文件的写操作，因为其效率要比f.Write高
	w := bufio.NewWriter(f)
	defer w.Flush()
	fmt.Fprintln(w,"222222222222222222222222222")


}

func writeFileFunc(filename string) {
	//f, err := os.Create(filename)
	f , err := os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)
	if err != nil {
		//err = errors.New("define error by yourself") //自己定义想要的错误，定义后的error就不是os.PathError
		//fmt.Printf("error:%s\n",err.Error()) //打印出define error by yourself
		if pathError, ok := err.(*os.PathError); !ok {
			fmt.Println("unknown error",err)
		}else {
			//fmt.Printf("pathError.Op:%s\npathError.Path:%s\npathError.Err:%s\n",pathError.Op,pathError.Path,pathError.Err)
			fmt.Println(pathError.Err)
		}
		return
	}
	defer f.Close()

	//str := "111111111111111"
	//buf := []byte(str)
	//f.Write(buf)

	//这里使用bufio.NewWriter函数来进行文件的写操作，因为其效率要比f.Write高
	w := bufio.NewWriter(f)
	defer w.Flush()
	fmt.Fprintln(w,"222222222222222222222222222")


}

func writeFileFunc2(filename string) {
	f , err := os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("ERROR===>%s\n",err.Error())
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()
	fmt.Fprintln(w,"222222222222222222222222222")
}

func main() {
	tryDefer()
	fmt.Println("==============")
	tryDefer1()
	fmt.Println("==============")
	tryDefer2()
	fmt.Println("==============")
	tryDeferFunc()
	fmt.Println("==============")
	writeFileFunc("test.txt")


}
