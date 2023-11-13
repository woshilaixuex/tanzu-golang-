package servce

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"goweb/utils"
	"log"
	"net/http"
	"strings"
)

var vcode utils.VerCode

type UserLogin struct {
	Username string
	Password string
	Vercode  string
}

func CreateUserServe(r *gin.Engine, db *gorm.DB) {
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
