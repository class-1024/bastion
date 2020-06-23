package dao

import (
"bastion/database"
"bastion/pkg/datasource"
)

func FindAllProjects(pagesize, page int, order string) (rows []database.StatProject, total int, e error) {
	offset := (page - 1) * pagesize

	var data []database.StatProject
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	err = datasource.GormPool.Model(&database.StatProject{}).Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&data).Error

	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func CreateProjects(info database.StatProject) error {
	err := datasource.GormPool.Create(&info).Error
	return err
}

func UpdateProjects(info database.StatProject) error {
	err := datasource.GormPool.Model(&database.StatProject{}).Updates(&info).Error
	return err
}

