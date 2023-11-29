package servce

import (
	"errors"
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

func GetVcode(c *gin.Context) (ResResult, error) {
	var login_user UserLogin
	session := sessions.Default(c)
	svcode := session.Get("vercode").(string)
	if svcode == "" {
		return ResErr("验证码为空", "ERROR"), errors.New("验证码为空")
	}
	if err := c.ShouldBindJSON(&login_user); err != nil {
		return ResErr("数据格式错误", "ERROR"), errors.New("数据格式错误")
	}
	if strings.ToLower(login_user.Vercode) != strings.ToLower(svcode) {
		return ResErr("验证码错误", "ERROR"), errors.New("验证码错误")
	}
	return ResSucceed(login_user), nil
}
func CreateUserServe(r *gin.Engine, db *gorm.DB) error {
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
		response, err := GetVcode(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, response)
		}
		//查询数据库
		c.JSON(http.StatusOK, response)
		return
	})
	r.POST("/register", func(c *gin.Context) {
		response, err := GetVcode(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, response)
		}
		
	})
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	if ginErr := r.Run(); ginErr != nil {
		log.Panic(ginErr)
	}
	return nil
}
