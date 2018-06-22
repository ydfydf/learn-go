package main

import (
	"regexp"
	"fmt"
)
//go语言中换行字符串使用``符号
const text = `My email is ydf@qq.com@ABC.com
email is abc@efg.org
email2 is kkk@qwe.cn
email3 is ddd@123.com.cn
`

//正则表达式提取功能，使用()标记需要提取的内容
func regexpSubString(expr string) [][]string{
	re := regexp.MustCompile(expr)
	return re.FindAllStringSubmatch(text,-1)
}

func main() {
	//*表示匹配前面的子表达式零次或多次，.表示匹配前面的子表达式一次或多次
	//注意.*和.+的区别，.代表任意一个字符，.*代表0个或多个字符，.+则一个或多个字符
	//regexp  := regexp.MustCompile(`.......`)//regexp.MustCompile的参数必须要符合正则表达式
	regexp , err := regexp.Compile(`[A-Za-z0-9]+@[A-Za-z0-9.]+\.[A-Za-z0-9]+`)//[]里的.可以不用转义字符\
	if err != nil {
		panic(err)
	}

	str := regexp.FindAllString(text,-1)
	fmt.Println(str)
	str1 := regexpSubString(`([A-Za-z0-9]+)@([A-Za-z0-9.]+)(\.[A-Za-z0-9]+)`)
	fmt.Println(str1)
	for _,m := range str1{
		for _,v := range m{
			fmt.Println(v)
		}
	}
}
