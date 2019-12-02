package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BazingaLyn/jarvis/trie"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var urlRouter trie.UriTrie = trie.NewTrie()

var userMap = map[int]User{
	1: User{
		Id:   1,
		Name: "bazinga",
		Age:  21,
	},
}

func main() {

	urlRouter.AddNode("/get/user/:id", getUserById)

	http.HandleFunc("/", router)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenerAndServe:", err)
		return
	}
}

type User struct {
	Id   int
	Name string
	Age  int
}

func getUserById(writer http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()
	idStr := query["id"][0]
	id, _ := strconv.Atoi(idStr)
	user := userMap[id]
	bytes, err := json.Marshal(user)
	if err == nil {
		fmt.Fprintln(writer, string(bytes))
	}
}

func router(writer http.ResponseWriter, request *http.Request) {
	path := request.RequestURI
	method := request.Method
	node, ok := urlRouter.Search(path)
	if ok {
		if len(node.Params) > 0 {
			if method == "GET" {
				query := request.URL.Query()
				for k, v := range node.Params {
					query.Set(k, v)
				}
				var nativeUrl url.URL
				nativeUrl.Path = request.URL.Path
				request.URL, _ = request.URL.Parse(nativeUrl.String() + "?" + query.Encode())
			} else {
				decoder := json.NewDecoder(request.Body)
				var params map[string]string
				decoder.Decode(&params)
				for k, v := range node.Params {
					params[k] = v
				}
				wholeParams, _ := json.Marshal(params)
				request.Body = ioutil.NopCloser(bytes.NewBuffer(wholeParams))
			}
			handle := node.Handle
			handle(writer, request)
		}
	} else {
		fmt.Fprintln(writer, "not found")
	}
}
