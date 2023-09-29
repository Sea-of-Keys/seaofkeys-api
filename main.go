package main

// ######TODO######

import (
	"fmt"
	"log"

	databae "github.com/Sea-of-Keys/seaofkeys-api/api/database"
	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

func main() {
	db, err := databae.Init("mysql")
	// db, err := databae.Init("postgres")
	if err != nil {
		log.Panic(err)
	}
	models.Setup(db)
	fmt.Printf("database: %v\n", db)
	fmt.Printf("gud: %v\n", "kronborg")
}
