package dao

import (
	"bastion/database"
	"bastion/pkg/datasource"
)

func FindAllComments(pagesize, page int, order string) (rows []database.CommentBiz, total int, e error) {
	offset := (page - 1) * pagesize

	var comments []database.CommentBiz
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	err = datasource.GormPool.Model(&database.CommentBiz{}).Preload("User").Count(&count).
		Order(order).Offset(offset).Limit(pagesize).Find(&comments).Error

	if err != nil {
		return nil, 0, err
	}

	return comments, count, nil
}

func CrateComment(userId int, content string) error {
	com := database.MComment{
		UserID:  userId,
		Content: content,
	}
	err := datasource.GormPool.Create(&com).Error
	return err
}
