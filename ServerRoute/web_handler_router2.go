package main

import (
	"fmt"
	"log"
	"net/http"
)

// 使用Handle()方法
// 这种方式需要自定义处理函数和结构体

type LanguageHandler struct {
	Language string
}

func (h LanguageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", h.Language)
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/cn", LanguageHandler{Language: "中文！"})
	mux.Handle("/en", LanguageHandler{Language: "英语！"})

	server := &http.Server{
		Addr:    "192.168.31.134:8081",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
