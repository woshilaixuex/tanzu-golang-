package api

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
		log.Println(response)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//查询数据库
	if response.Data == nil {
		c.JSON(http.StatusBadRequest, net_models.ResErr(err.Error(), "ERROR"))
		return
	} else {
		user, err := database.UserLoginDataServer(response.Data.(net_models.UserLogin))
		if err != nil {
			c.JSON(http.StatusBadRequest, net_models.ResErr(err.Error(), "ERROR"))
			return
		}
		token, errr := utils.CreatToken(*user)
		if errr != nil {
			c.JSON(http.StatusBadRequest, net_models.ResErr(errr.Error(), "ERROR"))
			return
		}
		c.JSON(http.StatusOK, net_models.ResSucceed(token))
	}
}

func Register(c *gin.Context) {
	var login_user net_models.UserLogin
	if err := c.ShouldBindJSON(&login_user); err != nil {
		c.JSON(http.StatusBadRequest, net_models.ResErr(err.Error(), "ERROR"))
	}
	user, err := database.UserRegisterDataServer(login_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, net_models.ResErr(err.Error(), "ERROR"))
		return
	}
	token, _ := utils.CreatToken(*user)
	c.JSON(http.StatusOK, net_models.ResSucceed(token))

}
