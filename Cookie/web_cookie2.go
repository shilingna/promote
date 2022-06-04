package main

import (
	"fmt"
	"net/http"
)

func cookieGetHandle(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("fw")
	fmt.Printf("w,cookie:%#v,err:%v\n", c, err)
	w.Write([]byte("ok"))

}
func main() {
	http.HandleFunc("/", cookieGetHandle)
	http.ListenAndServe(":8085", nil)
}
