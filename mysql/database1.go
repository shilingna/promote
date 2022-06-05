package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1)/fw_go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
