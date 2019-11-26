package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//客户端主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:10086")
	if err != nil {
		log.Fatal(err) //log.Fatal()会产生panic
		return
	}

	defer conn.Close() //关闭

	buf := make([]byte, 1024) //缓冲区
	for {
		fmt.Printf("请输入发送的内容：")
		fmt.Scan(&buf) //相当于Python中的input-->>buf = input("请输入发送的内容")
		fmt.Printf("发送的内容：%s\n", string(buf))

		conn.Write(buf)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := buf[:n]
		fmt.Printf("接收到数据[%d]:%s\n", n, string(result))
	}

}
