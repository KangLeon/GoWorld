package main

import (
	"fmt"
	"net/http"
)

type DefineServerMux struct {

}

func (dsm *DefineServerMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer,"通过自定义类型启动一个http")
}

func ServerHttp(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"通过HandleFunc启动一个http服务")

	fmt.Fprintln(w,r.URL.RawQuery)
	fmt.Fprintln(w,r.URL.Host)
	fmt.Fprintln(w,r.URL.Path)
}

func main() {
	http.HandleFunc("/geturl",ServerHttp)

	defineServerMux := DefineServerMux{}
	http.Handle("/getUserInfo",&defineServerMux)

	http.ListenAndServe(":8000",nil)
}