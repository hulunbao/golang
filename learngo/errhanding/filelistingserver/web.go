package main

import (
	"fmt"
	"golang/learngo/errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func(){
			r := recover()
			if r != nil{
				log.Printf("Panic: %v", r)
				http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
		}()
		err := handler(writer,request)
		if err != nil{
			log.Printf("Error handling request: %s",err.Error())

			if userErr, ok := err.(userErr); ok {
				http.Error(writer,userErr.Message(),http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

type userErr interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888",nil)
	if err != nil{
		panic(err)
	}
	type demo string
	d := demo("hello")
	fmt.Println(d)

}
