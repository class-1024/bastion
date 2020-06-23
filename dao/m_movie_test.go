package dao

import (
	"bastion/utils"
	"fmt"
	"testing"
)

func init() {
}

// 电影 列表
func TestFindAllMovies(t *testing.T) {

	rows, total, e := FindAllMovies(10, 1, "")

	utils.Must(e)
	s, e := utils.PrintJson(rows)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Printf("%v \n", total)
	fmt.Printf("%s \n", s)
}

func TestFindMovieById(t *testing.T) {

	//movie, e := FindMovieById(12312312312312)
	//utils.Must(e)
	//
	//s, e := utils.PrintJson(movie)
	//utils.Must(e)
	//
	//fmt.Printf("%s \n", s)
}
