package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("fw.txt")

	// 如果路径不对将出现以下错误
	// open fw.txt: The system cannot find the file specified.
	// <nil>
	// invalid argument
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}
