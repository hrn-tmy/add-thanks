package main

import (
	"add-thanks/internal/infra/database"
	"fmt"
	"log"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	// 一時的に出力
	fmt.Println("db: ", db)
}