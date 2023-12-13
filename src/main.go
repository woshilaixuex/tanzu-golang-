package main

import (
	"goweb/db"
	"goweb/router"
)

func main() {
	db.InitDb()
	router.InitRouter()
}
