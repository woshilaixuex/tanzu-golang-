package database

import (
	"errors"
	"goweb/data-models"
	net_models "goweb/net-models"
	"log"
)

func UserRegisterDataServer() {

}
func UserLoginDataServer(userLogin net_models.UserLogin) (*data_models.User, error) {
	user, err := SelectUserByID(userLogin.Username)
	if err != nil {
		log.Println("用户不存在")
		return nil, errors.New("用户不存在")
	}
	if user.PassWord != userLogin.Password {
		log.Println("账户或密码错误")
		return nil, errors.New("账户或密码错误")
	}
	user.PassWord = ""
	return user, nil
}
