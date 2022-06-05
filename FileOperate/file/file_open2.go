package main

import (
	"fmt"
	"os"
)

func main() {
	// 这里因为制定了文件的打开方式(创建或追加),因此不会存在覆盖现象
	fp, err := os.OpenFile("fw.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fp)
	defer fp.Close() // 关闭文件, 释放资源

}
