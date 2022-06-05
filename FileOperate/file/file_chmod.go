package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件
	// 默认权限是0666
	fp, err := os.Create("./fw.txt")
	defer fp.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileInfo, err := os.Stat("./fw.txt")
	fileMode := fileInfo.Mode()
	fmt.Println(fileMode)
	os.Chmod("./fw.txt", 0777)
	fileInfo, err = os.Stat("./fw.txt")
	fileMode = fileInfo.Mode()
	fmt.Println(fileMode)
}
