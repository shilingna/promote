package main

import (
	"fmt"
	"os"
)

// 创建多级目录

func main() {
	err := os.MkdirAll("fw_dir/1/2/3", 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("mkdir success")
}
