package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"os"
	"fmt"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error{
	panic(123)
}

func testUserErr(writer http.ResponseWriter, request *http.Request) error{
	return userError("user error")
}

func noErr(writer http.ResponseWriter, request *http.Request) error{
	fmt.Fprintln(writer,"no error")
	return nil
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error{
	return os.ErrNotExist
}

func errInternal(writer http.ResponseWriter, request *http.Request) error{
	return os.ErrInvalid
}

//表格驱动测试
var tests = []struct {
h HandlerHttp
code int
message string
}{
{errPanic,500,"Internal Server Error"},
{noErr,200,"no error"},
{testUserErr,400,"user error"},
{errNotFound,404,"Not Found"},
{errInternal,500,"Internal Server Error"},
}

//只是测试函数ErrorWraper,使用了两个假参数response，request去调用
func TestErrWrapper(t *testing.T){


	for _, tt := range tests {
		f := ErrorWraper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com",
			nil,
		)
		f(response,request)
		all ,_ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(all),"\n")

		if response.Code != tt.code || body != tt.message {
			t.Errorf("Expected (%d,%s); Got (%d,%s)",tt.code,tt.message,response.Code,body)
		}

	}
}

//上面的测试用例是模拟一个reques，这个测试用例会真正的启一个server来进行测试，测试的是服务器
func TestErrWrapperInServer(t *testing.T) {
	for _ , tt := range tests{
		f := ErrorWraper(tt.h) //f的类型为func(writer http.ResponseWriter, request *http.Request)
		//type HandlerFunc func(ResponseWriter, *Request),HandlerFunc是一个类型，http.HandlerFunc(f)就是为这个函数的类型赋值，这里即为接口赋值
		server := httptest.NewServer(http.HandlerFunc(f))
		resp ,_ := http.Get(server.URL)

		all ,_ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(all),"\n")

		if resp.StatusCode != tt.code || body != tt.message {
			t.Errorf("Expected (%d,%s); Got (%d,%s)",tt.code,tt.message,resp.StatusCode,body)
		}
	}
}
