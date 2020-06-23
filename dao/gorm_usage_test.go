package dao

import (
	"bastion/database"
	"bastion/pkg/datasource"
	"bastion/utils"
	"testing"
)

func init() {
}

func TestLearnGormCreate(t *testing.T) {


	//uid := time.Now().Unix()
	//u := db.MUser{Openid: strconv.Itoa(int(uid)), NickName: "test"}
	//
	//err := datasource.GormPool.Create(&u).Error
	//utils.Must(err)
	//
	//fmt.Printf("%v \n", u)
}

func TestLearnGormQuery(t *testing.T) {


	//uid := time.Now().Unix()
	//u := db.MUser{Openid: strconv.Itoa(int(uid)), NickName: "test"}
	//
	//err := datasource.GormPool.Create(&u).Error
	//utils.Must(err)
	//
	//fmt.Printf("%v \n", u)
}

func TestLearnGormUpdate(t *testing.T) {


	//uid := time.Now().Unix()
	//u := db.MUser{Openid: strconv.Itoa(int(uid)), NickName: "test"}
	//
	//err := datasource.GormPool.Create(&u).Error
	//utils.Must(err)
	//
	//fmt.Printf("%v \n", u)
}

func TestLearnGormDelete(t *testing.T) {


	e := datasource.GormPool.Unscoped().Delete(database.MUser{}, "nick_name LIKE ?", "test").Error
	utils.Must(e)

	//uid := time.Now().Unix()
	//u := db.MUser{Openid: strconv.Itoa(int(uid)), NickName: "test"}
	//
	//err := datasource.GormPool.Create(&u).Error
	//utils.Must(err)
	//
	//fmt.Printf("%v \n", u)
}
