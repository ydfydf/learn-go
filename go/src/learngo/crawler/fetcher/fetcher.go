package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"log"
	"golang.org/x/text/encoding/unicode"
	"time"
)

func DetermineEncoding(r *bufio.Reader) encoding.Encoding{
	//由于charset.DetermineEncoding()是截取前1024个字节来做判断，所以只需传1024个字节就够了,但是Peek完后，r *io.Reader里的r的当前指针就指向了1025，下次再去读的话就会从1025开始读
	bytes , err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v",err)
		return unicode.UTF8 //如果出错，则返回默认UTF-8编码
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

var rateLimit = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte,error){
	<- rateLimit //每10毫秒去进行一次fetch，不然获取速度太快，超过珍爱网的所允许的速度
	resp , err := http.Get(url)
	if err != nil {
		return nil ,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("Wrong status code:%d",resp.StatusCode)//生成error的方法有两种：errors.New(),fmt.Errorf()
	}

	//将检测网站上的编码转UTF-8
	r := bufio.NewReader(resp.Body)
	e := DetermineEncoding(r)
	utf8Reader := transform.NewReader(r,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

