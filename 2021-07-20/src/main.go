package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type DefineServerMux struct {
}

func (dsm *DefineServerMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "通过自定义类型启动一个http")
}

func ServerHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过HandleFunc启动一个http服务")

	fmt.Fprintln(w, r.URL.RawQuery)
	fmt.Fprintln(w, r.URL.Host)
	fmt.Fprintln(w, r.URL.Path)
	rawquery := r.URL.RawQuery
	va, _ := url.ParseQuery(rawquery)

	fmt.Fprintln(w, "获取url中key为name的值", va.Get("name"))
	fmt.Fprintln(w, "=================================")

	for key := range r.Header {
		fmt.Fprintf(w, "%s:%s\n", key, r.Header[key])
	}

}

func main() {
	http.HandleFunc("/geturl", ServerHttp)

	defineServerMux := DefineServerMux{}
	http.Handle("/getUserInfo", &defineServerMux)

	http.ListenAndServe(":8000", nil)
}
