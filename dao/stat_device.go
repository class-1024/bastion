package dao

import (
	"bastion/database"
	"bastion/pkg/datasource"
	"github.com/jinzhu/gorm"
)

func CreateIfNotExistDevice(data database.StatDevice) error {
	d := database.StatDevice{
		UId: data.UId,
	}
	err := datasource.GormPool.Where("uid = ?", d.UId).Find(&d).Error
	// 未找到 创建
	if err != nil && gorm.IsRecordNotFoundError(err) {
		err := datasource.GormPool.Create(&data).Error
		return err
	}
	return nil
}

func FindByUid(uid string) (*database.StatDevice, error) {
	res := database.StatDevice{}
	err := datasource.GormPool.Where("uid = ?", uid).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func FindAllDevice(pagesize, page int, order string) (rows []database.StatDevice, total int, e error) {
	offset := (page - 1) * pagesize

	var res []database.StatDevice
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	err = datasource.GormPool.Model(&database.StatDevice{}).Count(&count).
		Order(order).Offset(offset).Limit(pagesize).Find(&res).Error

	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}
