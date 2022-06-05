package main

import (
	"fmt"
	"log"
	"net/http"
)

// 启动Server 端的主要代码就是 http.ListenAndServe()，这个方法提供了带参数的方法和无参的方法

//下面是http.HandleFunc() 方式启动server 端，
//http.ListenAndServe(":8081", nil) 第二个参数是nil ，实际会走DefaultServeMux()来进行处理
func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func main() {
	http.HandleFunc("/", hi)
	if err := http.ListenAndServe("192.168.31.134:8081", nil); err != nil {
		log.Fatal(err)
	}
}
