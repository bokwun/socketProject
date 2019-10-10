package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func dealConn(conn net.Conn) {
	defer conn.Close()
	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, "连接成功！")

	buf := make([]byte, 1024)
	for {
		//接收数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		//截取有效数据
		result := buf[:n]
		fmt.Printf("接收到的数据来自[%s]===>[%d]:%s\n", ipAddr, n, string(result)) //转化为字符串格式
		if "exit" == string(result) {
			fmt.Println(ipAddr, "退出连接")
			return
		}

		//将接收到客户端的数据变成大写格式，然后发送回给客户端
		conn.Write([]byte(strings.ToUpper(string(result)))) //转换成[]byte 类型格式
	}

}

func main() {
	listener, err := net.Listen("tcp", "192.168.1.108:8082")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go dealConn(conn)
	}

}
