package dao

import (
	"bastion/database"
	"bastion/pkg/datasource"
	"github.com/medivhzhan/weapp/v2"
	"time"
)

func FindAllUsers(pagesize, page int, order string) (rows []*database.MUser, total int, e error) {
	offset := (page - 1) * pagesize

	var users []*database.MUser
	var count int

	if order == "" {
		order = " last_login desc, created_at desc"
	}

	var err error
	err = datasource.GormPool.Model(&database.MUser{}).Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func FindUserByOpenid(openid string) (*database.MUser, error) {
	var user database.MUser
	err := datasource.GormPool.Model(&database.MUser{}).Where("openid = ?", openid).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(openid string) error {
	user := database.MUser{
		Openid:    openid,
		LastLogin: time.Now(),
	}
	err := datasource.GormPool.Model(&database.MUser{}).Create(&user).Error
	return err
}

func UpdateUser(info *weapp.UserInfo) error {
	err := datasource.GormPool.Model(&database.MUser{}).Updates(&info).Error
	return err
}

func UpdateUserLoginTime(userId int) {
	user := database.MUser{
		Model: database.Model{ID: uint(userId)},
	}
	datasource.GormPool.Model(&user).Update("last_login", time.Now())
}
