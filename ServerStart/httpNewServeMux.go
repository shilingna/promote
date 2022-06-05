package main

import (
	"fmt"
	"log"
	"net/http"
)

// http. NewServeMux ()的作用是注册网络访问的多路路由。
// 因为它采用的是自定义的多路由分发任务方式，所以称之为自定义多路由分发服务，
// Mux是multiplexer的缩写，意思是多路由转换器（多路路由器）

func hi2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world2")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi2)

	server := &http.Server{
		Addr:    "192.168.31.134:8081",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
