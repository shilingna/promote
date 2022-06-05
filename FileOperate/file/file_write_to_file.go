package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileOut, err := os.Create("./fw_write.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileOut.Close()
	defer os.Remove("./fw_write.txt")
	for i := 0; i < 5; i++ {
		outStr := fmt.Sprintf("%s:%f\r\n", "Hello Go", i)
		// 写入文件
		fileOut.WriteString(outStr)              // string信息
		fileOut.Write([]byte("write to go\r\n")) // byte类型
		fileOut.WriteAt([]byte("插入一句话"), 5)      // 从指定位置写
	}

	// 打印内容
	content, err := ioutil.ReadFile("./fw_write.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", string(content))
}
