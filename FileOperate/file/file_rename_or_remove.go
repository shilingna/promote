package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./fw_rename.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	err = os.Rename("./fw_rename.txt", "./fw_rename_new.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
