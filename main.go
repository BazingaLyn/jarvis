package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var users = make(map[int]User)

func init() {
	user := User{
		Id:   1,
		Name: "小明",
		Age:  10,
	}
	users[user.Id] = user
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/user", getUser)
	http.HandleFunc("/save/user", saveUser)
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

func saveUser(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		params := request.URL.Query()
		idStr := params["id"][0]
		id, _ := strconv.Atoi(idStr)
		name := params["name"][0]
		AgeStr := params["age"][0]
		age, _ := strconv.Atoi(AgeStr)

		newUser := User{
			Id:   id,
			Name: name,
			Age:  age,
		}
		users[id] = newUser

	}

}

func getUser(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	idStr := params["id"][0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintln(writer, "id format is error is not number type")
		return
	}
	user := users[id]
	bytes, _ := json.Marshal(user)
	fmt.Fprintln(writer, string(bytes))

}

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	fmt.Println(request.RequestURI)
	fmt.Fprintf(writer, "hello web golang")
}
