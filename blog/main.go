package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%s 欢迎来到码神之路goweb教程", "mszlu.com")
	})
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
