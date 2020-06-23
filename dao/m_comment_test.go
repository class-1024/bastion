package dao

import (
	"bastion/utils"
	"testing"
)

func init() {
}

func TestFindAllComments(t *testing.T) {
	rows, total, e := FindAllComments(10, 1, "")

	utils.Must(e)
	utils.PrintJsonString(rows)
	utils.Print(total)
}
