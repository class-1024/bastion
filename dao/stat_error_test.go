package dao

import (
	"bastion/entry"
	"bastion/pkg/setup"
	"bastion/utils"
	"testing"
)

func init() {
	setup.InitTest()
}

func TestCreateErrors(t *testing.T) {

}

func TestFindErrorsByAppId(t *testing.T) {


	query := entry.ErrorsQuery{}
	query.AppId = "1"

	data, total, err := FindErrors(query, 10, 1, "")

	utils.Must(err)
	utils.PrintJsonString(total)
	utils.PrintJsonString(data)
}
