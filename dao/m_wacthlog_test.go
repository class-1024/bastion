package dao

import (
	"bastion/utils"
	"fmt"
	"testing"
)

func init() {
}

// 观看记录
func TestFindAllUserMovieLogs(t *testing.T) {

	//
	rows, total, e := FindAllMovieLogs(10, 1, "")
	//
	utils.Must(e)
	s, e := utils.PrintJson(rows)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Printf("%v \n", total)
	fmt.Printf("%s \n", s)
}

func TestCrateMovieLog(t *testing.T) {

	e := CrateMovieLog(6, 111, "88:88:88")
	fmt.Printf("%v  \n", e)
}

func TestFindMovieLogByMovieIdAndUserId(t *testing.T) {
	//
	//log, e := FindMovieLogByMovieIdAndUserId(32, 4)
	//utils.Must(e)
	//
	//s, e := utils.PrintJson(log)
	//utils.Must(e)
	//
	//fmt.Printf("%s \n", s)
}

func TestUpdateMovieLog(t *testing.T) {

	UpdateMovieLog(120, "99:99:99")
}

func TestFindMovieLog(t *testing.T) {
	//
	//movie, e := FindMovieLog(98)
	//utils.Must(e)
	//
	//s, e := utils.PrintJson(movie)
	//utils.Must(e)
	//
	//fmt.Printf("%s \n", s)
}
