package main

import (
	"fmt"
	"github.com/strongerlin666/msweb/msgo"
	"net/http"
)

func main() {
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	// TODO: error处理
	//	fmt.Fprintf(writer, "%s 欢迎来到码神之路goweb教程", "mszlu.com")
	//})
	//err := http.ListenAndServe(":8111", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	engine := msgo.New()
	engine.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		// TODO: error处理
		fmt.Fprintf(w, "%s 欢迎来到码神之路goweb教程", "mszlu.com")
	})
	engine.Run()
}
