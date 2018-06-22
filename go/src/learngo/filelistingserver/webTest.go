package main

import (
	"net/http"
	"strings"
	"log"
	"os"
	"io/ioutil"
	_"net/http/pprof"//程序中没有用到net/http/pprof包里的东西，加下划线是为了告诉编译器我们会load这个包里的一些帮助程序进来
)
const prefix = "/ydf/"

type userError string


type userErr interface {
	error
	Message() string
}

func (e userError) Error() string{
	return e.Message()
}

func (e userError) Message() string{
	return string(e)
}

type HandlerHttp func(writer http.ResponseWriter, request *http.Request) error

func ErrorWraper(handler HandlerHttp) func(writer http.ResponseWriter, request *http.Request){
	return func(writer http.ResponseWriter, request *http.Request){
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recover:[%v]", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer,request)
		if err != nil {
			if userErr ,ok := err.(userErr);ok {
				http.Error(writer,userErr.Message(),http.StatusBadRequest)

				return
			}
		}
		log.Printf("")
		code := http.StatusOK
		if err != nil {
			log.Println(err)
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

func HandleFile(writer http.ResponseWriter, request *http.Request) error {

	if position := strings.Index(request.URL.Path,prefix);position == -1 {
		//return errors.New("Path must start with " + prefix)
		return userError("Path must start with " + prefix)
	}

	filename := request.URL.Path[len("/ydf/"):]
	file , err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	contents ,err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if _,err := writer.Write(contents); err != nil {
		return err
	}
	return nil
}


func main() {
	http.HandleFunc("/", ErrorWraper(HandleFile))
	err := http.ListenAndServe(":8887",nil)
	if err != nil {
		panic(err)
	}
}
