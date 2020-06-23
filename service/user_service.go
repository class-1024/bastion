package service

import (
	"bastion/dao"
	"bastion/database"
)

func FindOrCreateUserByOpenid(openid string) (*database.MUser, error) {

	// 通过openid 检查用户是否存在
	isExist, user, err := UserIsExist(openid)
	if err != nil {
		return nil, err
	}

	// 存在-->登录
	if isExist {
		dao.UpdateUserLoginTime(int(user.ID))
		return user, nil
	}

	// 不存在-->创建用户-->并登录
	err = dao.CreateUser(openid)
	if err != nil {
		return nil, err
	}

	u, err := dao.FindUserByOpenid(openid)
	return u, err
}

func UserIsExist(openid string) (isExist bool, user *database.MUser, error error) {
	user, e := dao.FindUserByOpenid(openid)
	if e != nil {
		return false, nil, e
	}
	if user != nil {
		return true, user, nil
	}
	return false, nil, nil
}

func LogIsExist(userId int, movieId int) (isExist bool, log *database.MWatchLog, error error) {
	movieLog, err := dao.FindMovieLogByMovieIdAndUserId(userId, movieId)
	if err != nil {
		return false, nil, err
	}
	if movieLog != nil {
		return true, movieLog, nil
	}
	return false, nil, nil
}
