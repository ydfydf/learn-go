package others

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func basicFor() {//无需返回值时，最后就不要返回值类型
	sum := 0
	for i:= 0; i <= 100; i ++ {//for的条件里不需要小括号
		sum += i
	}
	fmt.Println(sum)
}

//for的条件里可以省略初始条件，结束条件，递增表达式

//省略初始条件的for循环
func convertToBin(n int) string {//函数作用：将数字转换为二进制字符串
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//省略递增表达式
func printFile(filename string){
	file ,err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//省略结束条件
func forever() {
	for { //死循环，Go语言中没有while关键字，所以不能使用while(1)来表示死循环，用此方法代替
		fmt.Println("ydf")
	}
}

func main() {
	basicFor()
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)
	printFile("helloworld.go")
	forever()
}
