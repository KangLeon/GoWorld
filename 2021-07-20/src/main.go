package main

import (
	"fmt"
	"net/http"
)

func ServerHttp(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"通过HandleFunc启动一个http服务")
}

func main() {
	http.HandleFunc("/",ServerHttp)
	http.ListenAndServe("8080",nil)
}