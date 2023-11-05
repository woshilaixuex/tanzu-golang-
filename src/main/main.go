package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"goweb/database"
	"goweb/utils"
	"log"
	"net/http"
	"strings"
)

type UserLogin struct {
	Username string
	Password string
	Vercode  string
}

var vcode utils.VerCode

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tanzu_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.SelectAllUser(db)
	//注册gin驱动
	r := gin.Default()
	store := cookie.NewStore([]byte("relink"))
	r.Use(sessions.Sessions("session", store))
	r.GET("/vercode", func(c *gin.Context) {
		session := sessions.Default(c)
		vcode.CreatVerifyCode()
		session.Set("vercode", vcode.Vcode)
		session.Save()
		c.Data(http.StatusOK, "image/png", vcode.OutInput())
	})
	r.POST("/login", func(c *gin.Context) {
		var login_user UserLogin
		session := sessions.Default(c)
		svcode := session.Get("vercode").(string)
		if err := c.ShouldBindJSON(&login_user); err == nil {
			if strings.ToLower(login_user.Vercode) == strings.ToLower(svcode) {
				c.JSON(http.StatusOK, gin.H{"user": login_user})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code", "user": login_user})
			}
		}
		return
	})
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	if ginErr := r.Run(); ginErr != nil {
		log.Panic(ginErr)
	}
}
