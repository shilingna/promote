package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// NewServeMux目前是不能模式匹配路由的，比如/user/*info
// HttpRouter 则弥补了这个不足，HttpRouter 是一个高性能、可拓展的第三方Http 路由的包
// go get github.com/julienschmidt/httprouter

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Index"))
}
func Hello3(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello World"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello", Hello3)
	/*log.Fatal(http.ListenAndServe("192.168.31.134:8081", router))*/
	router.GET("/getUser", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("get 获取用户"))
	})

	router.POST("/user", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("post 用户信息"))
	})

	// 精确匹配
	router.GET("/user/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("用户 name" + p.ByName("name")))
	})

	/*// 匹配所有
	router.GET("/user/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("接口info name" + p.ByName("name")))
	})*/
	http.ListenAndServe("192.168.31.134:8081", router)

}
