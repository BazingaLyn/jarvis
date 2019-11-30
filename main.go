package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var users = map[int]User{
	1: User{
		Id:   1,
		Name: "小明",
		Age:  10,
	},
}

type User struct {
	Id   int
	Name string
	Age  int
}

//GET localhost:9090/save/user?id=3&name=bazinga&age=21
//POST localhost:9090/save/user
//{
////	"Id": 4,
////	"Name": "jarvis",
////	"Age": 22
////}
func main() {

	http.HandleFunc("/user", getUser)
	http.HandleFunc("/save/user", saveUser)
	http.HandleFunc("/update/user", updateUser)
	http.HandleFunc("/delete/user", deleteUser)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenerAndServe:", err)
		return
	}
}

func deleteUser(writer http.ResponseWriter, request *http.Request) {

	params := request.URL.Query()
	idStr := params["id"][0]
	id, _ := strconv.Atoi(idStr)

	delete(users, id)

	fmt.Fprintln(writer, "delete success")
	return

}

func updateUser(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		bytes, _ := ioutil.ReadAll(request.Body)
		updateUser := User{}
		err := json.Unmarshal(bytes, &updateUser)

		if err != nil {
			panic("param is err")
		}
		_, ok := users[updateUser.Id]
		if ok {
			users[updateUser.Id] = updateUser
			fmt.Fprintln(writer, "update success")
		} else {
			fmt.Fprintln(writer, "user not exist not need update")
		}
		return
	}

}

func saveUser(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		bytes, _ := ioutil.ReadAll(request.Body)
		saveUser := User{}
		err := json.Unmarshal(bytes, &saveUser)

		if err != nil {
			panic("param is err")
		}
		users[saveUser.Id] = saveUser
		fmt.Fprintln(writer, "success")
		return
	}

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
		fmt.Fprintln(writer, "success")
		return
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
