package database

import (
	"errors"
	"goweb/data-models"
	net_models "goweb/net-models"
	"log"
)

func UserRegisterDataServer(userLogin net_models.UserLogin) (*data_models.User, error) {
	user, _ := SelectUserByID(userLogin.Username)
	if user.UserName != "" {
		return nil, errors.New("用户已存在")
	}
	user, _ = InsertUser(&data_models.User{
		UserName: userLogin.Username,
		PassWord: userLogin.Password,
	})
	if user == nil {
		log.Println("用户创建失败")
	}
	user.PassWord = ""
	return user, nil
}
func UserLoginDataServer(userLogin net_models.UserLogin) (*data_models.User, error) {
	user, err := SelectUserByID(userLogin.Username)
	if err != nil || user == nil {
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
