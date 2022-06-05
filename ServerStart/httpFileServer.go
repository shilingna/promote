package main

import "net/http"

// http.ListenAndServer()函数用来启动Web服务，绑定并监听http端口。
// 其中第一个参数为监听地址，第二个参数表示提供文件访问服务的HTTP处理器Handler

func main() {
	http.ListenAndServe("192.168.31.134:8081", http.FileServer(http.Dir("./ServerStart/public/")))
}
