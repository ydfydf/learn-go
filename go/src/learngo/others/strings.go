package others

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱你"
	fmt.Printf("%s,%d\n",[]byte(s),len(s))
	for _, b := range []byte(s){
		fmt.Printf("%X ",b)
	}
	fmt.Println()
	fmt.Printf("len(s)=%d\n",len(s))//utf-8编码，一个中文等于三个字节

	for i, ch := range s {//ch is a rune
		//fmt.Printf("(%d %X) ",i ,ch)
		fmt.Printf("(%d %c) ",i, ch) //rune是4字节，
	}
	fmt.Println()
	fmt.Printf("utf8.RuneCountInString(s)=%d\n",utf8.RuneCountInString(s))//获取s中字符个数

	bytes := []byte(s) //定义切片bytes，使用[]byte可以获得字节
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //DecodeRune作用为将byte切片进行解码，返回rune字符与该字符大小
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
		//打印出(0 59) (1 65) (2 73) (3 6211) (6 7231) (9 4F60) ，这里0，1，2，3，6，9代表字节数
	}
	fmt.Println()

	//使用range遍历position，rune对
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ",i, ch)
		//打印出(0 Y) (1 e) (2 s) (3 我) (4 爱) (5 你) ，这里0，1，2，3，4，5不是代表字节数，只是代表个数
	}

	fmt.Println()
	s1 := "qwee1234一二三四"
	fmt.Println(len(s1))
	s2 := []rune(s1)
	fmt.Println(len(s2))
	s3 := []byte(s1)
	fmt.Println(len(s3))
	s4 := [...]rune{'q','w','e','e',1,2,3,4,'一','二','三','四'}
	fmt.Println(len(s4))
}
