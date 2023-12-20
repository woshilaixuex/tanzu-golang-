package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"goweb/data-models"
	"goweb/db"
	"log"
	"strings"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/tanzu_db?charset=utf8mb4&parseTime=True&loc=Local"

func InsertUser(user *data_models.User) (*data_models.User, error) {
	user.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	result := db.Db.FirstOrCreate(&user, data_models.User{})
	if result.Error != nil && result.RowsAffected != 1 {
		return nil, result.Error
	}
	log.Println("用户" + user.ID + ":" + user.UserName + "保存成功")
	user.PassWord = ""
	return user, nil
}

func SelectUserByID(username string) (*data_models.User, error) {
	var user *data_models.User
	result := db.Db.Where("user_name = ?", username).Take(&user)
	if result.Error != nil {

	}
	return user, nil
}

func SelectAllUser(db *gorm.DB) {

}
