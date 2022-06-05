package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("fw.txt", os.O_CREATE|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	// 关闭文件句柄
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		// 读到一个换行就结束
		line, err := reader.ReadString('\n')
		// io.EOF表示文件的结尾
		if err == io.EOF {
			break
		}
		// 输出内容
		fmt.Print(line)
	}
}
