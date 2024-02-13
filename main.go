package main

import (
	"golang-jwt/database"
	"golang-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}