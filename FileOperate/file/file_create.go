package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件
	// 默认权限是0666,如果文件已经已存在,则将文件清空,即覆盖
	fp, err := os.Create("./text.txt")
	if err != nil {
		fmt.Println("文件创建失败.")
		return
	}

	//　关闭文件,释放资源
	defer fp.Close()
}
