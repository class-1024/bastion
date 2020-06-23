package dao

import (
	"bastion/database"
	"bastion/pkg/datasource"
)

func CreateAdminUser(info database.StatAdmin) error {
	err := datasource.GormPool.Create(&info).Error
	return err
}

func FindAllAdminUsers(pagesize, page int, order string) (rows []*database.StatAdmin, total int, e error) {
	offset := (page - 1) * pagesize

	var users []*database.StatAdmin
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	err = datasource.GormPool.Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func FindAdminUserByUserName(username string) (*database.StatAdmin, error) {
	user := &database.StatAdmin{}
	err := datasource.GormPool.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindAdminUserById(id int) (*database.StatAdmin, error) {
	user := &database.StatAdmin{}
	err := datasource.GormPool.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateAdminUser(info database.StatAdmin) error {
	err := datasource.GormPool.Model(&database.StatAdmin{}).Updates(&info).Error
	return err
}
