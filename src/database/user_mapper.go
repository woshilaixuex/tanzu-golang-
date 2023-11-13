package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/tanzu_db?charset=utf8mb4&parseTime=True&loc=Local"

func CreateDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
func SelectAllUser(db *gorm.DB) {

}
