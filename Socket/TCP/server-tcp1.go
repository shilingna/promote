package main

import (
	"fmt"
	"log"
	"net"
)

func handler(c net.Conn) {
	defer c.Close()
	for {
		// 1.等待客户端通过conn发送消息
		// 2.如果客户端没有write【发送】，那么协程就阻塞在这里
		fmt.Printf("等待客户端%s发送消息\n", c.RemoteAddr().String())
		// ⑤建议一个1024字节的缓存用于读取客户端传来的消息并输出, 调用c.Read()读取数据到缓存中
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}

		// 3.显示客户端发送的内容到服务端的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	// ①使用net.Listen建立一个tcp协议的IP:port监听的客户端
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	// ②声明listener的关闭
	defer listener.Close()

	// 循环等待客户端访问,并输出客户端的信息
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("客户端ip=%v\n", conn.RemoteAddr().String())
		// ④开启一个线程用于处理客户端发送的内容
		go handler(conn)
	}
}
