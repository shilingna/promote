package main

// 创建单个目录

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 根据path创建目录，例如fw_dir
	err := os.Mkdir("fw_dir", 0777)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mkdir success")
}
