package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.108:8082")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		fmt.Println("请输入发送内容：")
		fmt.Scan(&buf)
		fmt.Println("发送的内容为:", string(buf))
		conn.Write(buf)

		//接收数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		//截取有效数据
		result := buf[:n]
		fmt.Println("接收到的数据[%d]：%s", n, string(result)) //转化为字符串格式

	}

}
