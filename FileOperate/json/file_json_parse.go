package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 解析json

// 定义配置文件解析后的结构
type (
	OrderDetail struct {
		OrderId       string `json:"order_id"`
		ReceiveName   string `json:"receiveName"`
		ReceiveMobile string `json:"receiveMobile"`
		ReceiveAddr   string `json:"receiveAddr"`
	}
	Order struct {
		OrderNo     string      `json:"orderNo"`
		OrderDetail OrderDetail `json:"orderDetail"`
	}
)

func main() {
	v := Order{}
	date, err := ioutil.ReadFile("./json_parse.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(date, &v)
	json, _ := json.Marshal(v)
	fmt.Println()
}
