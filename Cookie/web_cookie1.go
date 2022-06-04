package main

// 设置cookie
import "net/http"

func cookieHandle(w http.ResponseWriter, _ *http.Request) {
	cookie := &http.Cookie{
		Name:   "token",
		Value:  "AAC8A721-FF21-4B04-A350-AE1A243A92A0",
		MaxAge: 3600,
		Domain: "192.168.31.134",
		Path:   "/",
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", cookieHandle)
	http.ListenAndServe(":8085", nil)
}
