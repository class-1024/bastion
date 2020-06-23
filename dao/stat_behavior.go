package dao

import (
	"bastion/database"
	"bastion/pkg/datasource"
)

func FindAllBehaviors(pagesize, page int, order string) (rows []database.StatBehavior, total int, e error) {
	offset := (page - 1) * pagesize

	var data []database.StatBehavior
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	err = datasource.GormPool.Model(&database.StatBehavior{}).Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&data).Error

	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func CreateBehaviors(info database.StatBehavior) error {
	err := datasource.GormPool.Create(&info).Error
	return err
}

func UpdateBehaviors(info database.StatBehavior) error {
	err := datasource.GormPool.Model(&database.StatBehavior{}).Updates(&info).Error
	return err
}

