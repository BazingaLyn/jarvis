package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

//ToString方法
func (u User) String() string {
	return "User[ Id " + string(u.Id) + "]"
}

//设置Name方法
func (u *User) SetName(name string) string {
	oldName := u.Name
	u.Name = name
	return oldName
}

//年龄数+1
func (u *User) AddAge() bool {
	u.Age++
	return true
}

//测试方法
func (u User) TestUser() {
	fmt.Println("我只是输出某些内容而已....")
}

func main() {
	//通过反射的方式调用结构体类型的方法
	var setNameStr string = "SetName"
	//var addAgeStr string = "AddAge"
	user := User{
		Id:   1,
		Name: "sherlock",
		Age:  18,
	}
	//1.获取到结构体类型变量的反射类型
	refUser := reflect.ValueOf(&user) //需要传入指针，后面再解析
	fmt.Println(refUser)
	//2.获取确切的方法名
	//带参数调用方式
	setNameMethod := refUser.MethodByName(setNameStr)
	args := []reflect.Value{reflect.ValueOf("Mike")}

	call := setNameMethod.Call(args)

	fmt.Println(call[0].Interface().(string))

}
