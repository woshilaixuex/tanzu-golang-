package data_models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID                    string `gorm:"primaryKey"`
	UserName              string
	PassWord              string
	AccountNonExpired     bool       `gorm:"default:true"`
	AccountNonLocked      bool       `gorm:"default:true"`
	CredentialsNonExpired bool       `gorm:"default:true"`
	Enabled               bool       `gorm:"default:true"`
	Role                  []Role     `gorm:"many2many:user_role;"`
	UserInform            UserInform `gorm:"foreignKey:ID;references:ID"`
}

type Role struct {
	ID     int64  `gorm:"primaryKey"`
	Name   string `gorm:"default:'ROLE_user'"`
	NameZh string
}

type UserInform struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Message  string
	AvaImage string
	Email    string
}

type History struct {
	Time     string
	Emission int64
}

// CreatUserModels 创建用户模型
func CreatUserModels(db *gorm.DB) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		}
	}()
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserInform{})
}
