package dao

import (
	"bastion/database"
	"bastion/pkg/datasource"
)

func FindAllMovies(pagesize, page int, order string) (rows []*database.MMovie, total int, e error) {
	offset := (page - 1) * pagesize

	var movies []*database.MMovie
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	err = datasource.GormPool.Model(&database.MMovie{}).Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&movies).Error

	if err != nil {
		return nil, 0, err
	}

	return movies, count, nil
}

func FindMovieById(id int64) (*database.MMovie, error) {
	movie := database.MMovie{}

	e := datasource.GormPool.Model(&database.MMovie{}).Where("id = ?", id).First(&movie).Error

	if e != nil {
		return nil, e
	}
	return &movie, nil
}
