package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// 因为 LisetenAndServe 在源码中启用的是Server对象
	// 那么我们手动创建一个
	server := &http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}

	// mux 也实现 Handler interface
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})

	server.Handler = mux

	log.Println("starting server 03 ...")
	// log.Fatal(http.ListenAndServe(":4000", mux))
	log.Fatal(server.ListenAndServe())
}

// 通过自行定义 mux 进行自定义路由转发
type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world, this is version 03."))
}
