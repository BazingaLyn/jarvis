package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, "连接成功")

	buf := make([]byte, 1024)
	for {
		//阻塞等待用户发送的数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		//切片截取，只截取有效数据
		result := buf[:n]
		fmt.Printf("接收到数据来自[%s]==>[%d]:%s\n", ipAddr, n, string(result))
		if "exit" == string(result) {
			fmt.Println(ipAddr, "退出连接")
			return
		}
		conn.Write([]byte(strings.ToUpper(string(result))))
	}
}
