package main

import (
	"github.com/BazingaLyn/jarvis/trie"
	"log"
	"net/http"
)

var urlRouter trie.UriTrie = trie.NewTrie()

func main() {

	urlRouter.AddNode(":/get/user/{id}", getUserById)

	http.HandleFunc("/", router)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenerAndServe:", err)
		return
	}
}

func getUserById(writer http.ResponseWriter, request *http.Request) {

}

func router(writer http.ResponseWriter, request *http.Request) {
	path := request.RequestURI
	method := request.Method
	node, ok := urlRouter.Search(path)
	if ok {
		if len(node.Params) > 0 {
			if method == "GET" {
			}

		}
	}
}
