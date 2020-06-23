package dao

import (
	"bastion/database"
	"bastion/pkg/datasource"
	"bastion/utils"
	"fmt"
	"testing"
)

func init() {
}

func TestGorm(t *testing.T) {
	e := datasource.GormPool.Delete(database.MUser{}, "nick_name LIKE ?", "test").Error
	utils.Must(e)

	//uid := time.Now().Unix()
	//u := db.MUser{Openid: strconv.Itoa(int(uid)), NickName: "test"}
	//
	//err := datasource.GormPool.Create(&u).Error
	//utils.Must(err)
	//
	//fmt.Printf("%v \n", u)
}

// 所有用户
func TestFindAllUsers(t *testing.T) {
	rows, total, e := FindAllUsers(10, 1, "")

	utils.Must(e)
	s, e := utils.PrintJson(rows)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Printf("%v \n", total)
	fmt.Printf("%s \n", s)
}

func TestFindUserByOpenid(t *testing.T) {
	//
	//res, e := FindUserByOpenid("otneZ5fVdpIoF0LzKhoYqd0-bEt0")
	//
	//utils.Must(e)
	//s, e := utils.PrintJson(res)
	//if e != nil {
	//	t.Fatal(e)
	//}
	//fmt.Printf("%s \n", s)
}

//
func TestCreateUser(t *testing.T) {
	e := CreateUser("ddd")
	fmt.Printf("%v  \n", e)
}
