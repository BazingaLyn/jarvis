package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenerAndServe:", err)
		return
	}
}

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello web golang")
}
