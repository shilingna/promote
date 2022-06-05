package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// ①使用net.Dial建立tcp协议,并连接服务端提供的ip和端口
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	// 客户端可以发送单行数据, 然后就推出
	// ②使用os.Stdin从客户端读取需要发送的数据, 如果客户端输入"exit"指定, 将退出客户端
	reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准输入[终端]
	for {
		// 从终端读取一行用户输入, 并准备发送个服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("用户推出客户端")
			break
		}
		// 再将line发送给服务器
		// ③调用conn.Write写入数据传给服务端, 并输入传输的字节数
		content, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("客户端发送了%d字节的数据到服务端\n", content)
	}
}
