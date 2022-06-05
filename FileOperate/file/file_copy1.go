package main

import (
	"fmt"
	"io"
	"os"
)

// func Copy(dst Writer, src Reader) (written int64, err error)
// dst: 源文件指针
// src: 目标文件指针

func main() {
	// 先创建一个名为: fw.copy.txt文件
	file, err := os.Create("./fw_copy.txt")
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString("Hello Go")
	// 打开文件fw_copy.txt, 获取文件指针
	srcFile, err := os.Open("./fw_copy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()
	// 打开文件要复制的新文件名fw_copy2.txt, 获取文件指针
	dstFile, err := os.OpenFile("./fw_copy2.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("opne file err = %v\n", err)
		return
	}
	defer dstFile.Close()
	// 通过Copy方法
	result, err := io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println("复制完成, 共复制字节数为: ", result)
	}
	defer os.Remove("fw_copy.txt")
	defer os.Remove("fw_copy2.txt")
}
