package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world, this is version 01."))
	})

	log.Println("starting server 01 ...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
