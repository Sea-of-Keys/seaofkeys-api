package main

// ######TODO######

import (
	"fmt"

	databae "github.com/Sea-of-Keys/seaofkeys-api/api/database"
	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

func main() {
	db := databae.Init()
	models.Setup(db)
	fmt.Printf("database: %v\n", db)
	fmt.Printf("gud: %v\n", "kronborg")
}
