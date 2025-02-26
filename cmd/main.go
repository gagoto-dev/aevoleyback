package main

import (
	"log"

	"github.com/gagoto-dev/aevoleyback/internal/api"
	"github.com/gagoto-dev/aevoleyback/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	sv := api.NewAPIServer(":8123", db)

	if err = sv.Run(); err != nil {
		log.Fatal(err)
	}
}
