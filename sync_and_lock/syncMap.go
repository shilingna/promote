package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("name", "xuluo")
	m.Store("gender", "Male")
	v, ok := m.LoadOrStore("name1", "Jim")
	fmt.Println(ok, v)
	v, ok = m.LoadOrStore("name", "aaa")
	fmt.Println(ok, v)
	v, ok = m.Load("name")
	if ok {
		fmt.Println("key存在，值是： ", v)
	} else {
		fmt.Println("key不存在")
	}
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, ",", value)
		return true
	})
	m.Delete("name1")
	fmt.Println(m.Load("name1"))
}
