package main

import (
	"fmt"

	databae "github.com/Sea-of-Keys/seaofkeys-api/api/database"
)

func main() {
	db := databae.Init()
	fmt.Printf("database: %v\n", db)
	fmt.Printf("lol %v\n", "kronborg")
}
