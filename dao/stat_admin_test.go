package dao

import (
	"bastion/database"
	"bastion/utils"
	"testing"
)

// 所有用户
func TestFindAllAdminUsers(t *testing.T) {
	rows, total, e := FindAllAdminUsers(10, 1, "")
	if e != nil {
		t.Fatal(e)
	}

	utils.PrintJsonString(rows)
	utils.Print(total)
}

// 创建
func TestCreateAdminUser(t *testing.T) {
	type args struct {
		info database.StatAdmin
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"测试1", args{info: database.StatAdmin{Nickname: "用户1"}}, false},
		{"测试2", args{info: database.StatAdmin{Nickname: "用户2"}}, false},
		{"测试3", args{info: database.StatAdmin{Nickname: "用户3"}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateAdminUser(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("CreateAdminUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateAdminUser(t *testing.T) {
	type args struct {
		info database.StatAdmin
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"UpdateAdminUser测试1", args{info: database.StatAdmin{Model: database.Model{ID: 12}, Nickname: "用户1", Avatar: "", Username: "nancode"}}, false},
		{"UpdateAdminUser测试2", args{info: database.StatAdmin{Model: database.Model{ID: 13}, Nickname: "mick", Email: "122"}}, false},
		{"UpdateAdminUser测试3", args{info: database.StatAdmin{Model: database.Model{ID: 14}, Nickname: "用户1", Phone: 13185010167}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateAdminUser(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdminUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFindAdminUserById(t *testing.T) {
	user, err := FindAdminUserById(2)
	if err != nil {
		t.Fatal(err)
	}
	utils.PrintJsonString(user)

}
