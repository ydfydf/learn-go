package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

const url = "http://www.imooc.com"

func main() {
	request,err := http.NewRequest(http.MethodGet,url,nil)
	//这个http的User-Agent头是手机端发送的，用在这里模拟手机端的请求
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//resp ,err := http.DefaultClient.Do(request)
	client := &http.Client{//和http.DefaultClient.Do(request)是等价的
		CheckRedirect: func(req *http.Request, via []*http.Request) error {//第一个参数是request的目标，第二个参数是将所有需要重定向的路径都放在里面
		fmt.Println("Redirec:",request)
		return nil
		},
	}
	resp ,err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()//使用http.Get，就必须要Body.Close

	s , err := httputil.DumpResponse(resp,true)//第二个参数是：是否将responce的body写到resp里去，如果为false，则只会写http头进去
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",s)
}
