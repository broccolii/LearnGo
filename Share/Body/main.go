package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age int
}

func main () {
	router := mux.NewRouter()

	router.HandleFunc("/query", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("进入了请求")
		buf, _ := ioutil.ReadAll(request.Body)
		user := User{}
		_ = json.Unmarshal(buf, &user)

		if user.Name == "kk" {
			user.Name = "金刚"
		}

		responseBuf, _ := json.Marshal(user)
		writer.Write(responseBuf)

	})

	err := http.ListenAndServe(":4000", router)
	if err != nil {
		fmt.Printf("服务器启动发生错误 error : %v", err)
	} else {
		fmt.Println("服务已启动")
	}
}
