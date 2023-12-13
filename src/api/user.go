package api

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	database "goweb/database"
	net_models "goweb/net-models"
	"goweb/utils"
	"log"
	"net/http"
	"strings"
)

var vcode utils.VerCode

func ApiGetVcode(c *gin.Context) (net_models.ResResult, error) {
	var login_user net_models.UserLogin
	//检查验证码
	session := sessions.Default(c)
	svcode := session.Get("vercode").(string)
	if svcode == "" {
		return net_models.ResErr("验证码为空", "ERROR"), errors.New("验证码为空")
	}
	if err := c.ShouldBindJSON(&login_user); err != nil {
		return net_models.ResErr("数据格式错误", "ERROR"), errors.New("数据格式错误")
	}
	if strings.ToLower(login_user.Vercode) != strings.ToLower(svcode) {
		return net_models.ResErr("验证码错误", "ERROR"), errors.New("验证码错误")
	}
	return net_models.Pending(login_user), nil
}

func Vcode(c *gin.Context) {
	session := sessions.Default(c)
	vcode.CreatVerifyCode()
	session.Set("vercode", vcode.Vcode)
	session.Save()
	c.Data(http.StatusOK, "image/png", vcode.OutInput())
}

func Login(c *gin.Context) {
	response, err := ApiGetVcode(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
	}
	//查询数据库
	user, err := database.UserLoginDataServer(response.Data.(net_models.UserLogin))
	if err != nil {
		c.JSON(http.StatusBadRequest, net_models.ResErr(err.Error(), "ERROR"))
	}
	c.JSON(http.StatusOK, net_models.ResSucceed(user))
	return
}

func Register(c *gin.Context) {
	response, err := ApiGetVcode(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
	}
}

func CreateUserServe(r *gin.Engine, db *gorm.DB) error {
	store := cookie.NewStore([]byte("relink"))
	r.Use(sessions.Sessions("session", store))
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	if ginErr := r.Run(); ginErr != nil {
		log.Panic(ginErr)
	}
	return nil
}
