package dao

import (
	"bastion/database"
	"bastion/pkg/datasource"
)

func FindAllMovieLogs(pagesize, page int, order string) (rows []database.UserMovieLogBiz, total int, e error) {
	offset := (page - 1) * pagesize

	var logs []database.UserMovieLogBiz
	var count int

	if order == "" {
		order = "updated_at desc, created_at desc"
	}

	var err error
	err = datasource.GormPool.Model(&database.UserMovieLogBiz{}).
		Preload("User").Preload("Movie").Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&logs).Error

	if err != nil {
		return nil, 0, err
	}

	return logs, count, nil
}

func FindMovieLogByMovieIdAndUserId(userId int, movieId int) (*database.MWatchLog, error, ) {
	log := database.MWatchLog{}

	err := datasource.GormPool.Model(&database.UserMovieLogBiz{}).
		Where("user_id = ? AND movie_id >= ?", userId, movieId).First(&log).Error

	if err != nil {
		return nil, err
	}

	return &log, nil
}

func CrateMovieLog(userId int, movieId int, progress string) error {
	log := database.MWatchLog{
		UserID:   userId,
		MovieID:  movieId,
		Progress: progress,
	}
	err := datasource.GormPool.Create(&log).Error
	return err
}

func UpdateMovieLog(id int, progress string) {
	log := database.MWatchLog{
		Model: database.Model{ID: uint(id)},
	}
	datasource.GormPool.Model(&log).Update("progress", progress, )
}

func FindMovieLog(id int64) (*database.UserMovieLogBiz, error) {
	log := database.UserMovieLogBiz{}

	e := datasource.GormPool.Model(&database.UserMovieLogBiz{}).
		Preload("User").Preload("Movie").
		Where("id = ?", id).First(&log).Error

	if e != nil {
		return nil, e
	}
	return &log, nil
}
