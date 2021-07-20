package main

import (
	"fmt"
	"net/http"
)

func ServerHttp(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"通过HandleFunc启动一个http服务")
}

type DefineServerMux struct {

}

func (dsm *DefineServerMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer,"通过自定义类型启动一个http")
}

func main() {
	http.HandleFunc("/",ServerHttp)
	defineServerMux := DefineServerMux{}
	http.Handle("getUserInfo",&defineServerMux)

	http.ListenAndServe(":8000",nil)
}