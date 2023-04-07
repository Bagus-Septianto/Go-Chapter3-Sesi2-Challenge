package main

import (
	"challenge294/database"
	"challenge294/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}