package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	api "goweb/api"
)

func InitRouter() {
	r := gin.Default()
	store := cookie.NewStore([]byte("relink"))
	r.Use(sessions.Sessions("session", store))
	/* @ShuCoding
	basic service：
		register、login、vcode
	*/
	basicAPI := r.Group("/tanzu")
	basicAPI.GET("/vercode", api.Vcode)
	basicAPI.POST("/user/login", api.Login)
	basicAPI.POST("/user/register", api.Register)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
