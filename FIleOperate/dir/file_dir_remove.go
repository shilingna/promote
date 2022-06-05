package main

import (
	"fmt"
	"os"
)

// os.Remove删除单个目录，如果有该目录有子目录或文件则会报错
// 目录不存在也是删除成功，不会报错

func main() {
	os.Mkdir("fw_dir", 0777)
	err := os.Remove("fw_dir")
	if err != nil {
		fmt.Printf("remove fw_dir err : %v\n", err)
	}
	fmt.Println("remove success")
}
