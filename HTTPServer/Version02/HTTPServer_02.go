package main

import (
	"log"
	"net/http"
)

func main() {

	// mux 也实现 Handler interface
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})

	log.Println("starting server 02 ...")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

// 通过自行定义 mux 进行自定义路由转发
type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world, this is version 02."))
}
