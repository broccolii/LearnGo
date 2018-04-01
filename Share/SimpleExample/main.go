package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":4000", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(([]byte)("hello world"))
		}))
}
