package servce

import (
	"github.com/gin-gonic/gin"
	"goweb/database"
)

func RunServe() {
	db, err := database.CreateDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	CreateUserServe(r, db)
}
