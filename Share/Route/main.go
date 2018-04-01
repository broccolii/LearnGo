package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Write([]byte("user is " + name))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{name}", userHandler)

	http.ListenAndServe(":4000", router)
}