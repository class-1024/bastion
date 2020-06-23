package controller

import (
	"bastion/pkg/response"
	"bastion/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
)

type System struct {
}

var startTime = utils.NowTime()

// @Summary 系统信息
// @Produce json
// @Success 200 {object} response.Response
// @Router /go/sys/info [get]
func (s *System) Info(c *gin.Context) {
	var msg struct {
		Start   string `json:"start"`
		StartZh string `json:"startZh"`
		Env     string `json:"env"`
	}

	msg.Start = startTime

	str, _ := utils.PastFromNow(startTime)
	msg.StartZh = str

	msg.Env = viper.GetString("base.env")

	response.Success(c, msg)
	return
}

func (s *System) Error(c *gin.Context) {
	_, err := ioutil.ReadFile("filepath")
	panic(err)
	return
}
