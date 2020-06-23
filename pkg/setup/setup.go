package setup

import (
	"bastion/pkg/config"
	"bastion/pkg/datasource"
	"bastion/utils"
	"flag"
	"github.com/georgehao/log"
)

// 测试
func InitTest() {
	var relPath string
	relPath = "./secret.dev.toml"
	Ready(relPath)
}

func Boot() {
	var relPath string
	flag.StringVar(&relPath, "conf", "./secret_dev.toml", "请输入配置文件路径")
	flag.Parse()
	Ready(relPath)
}

func Ready(relPath string)  {

	// 配置文件加载
	config.Load(relPath)

	// 日志
	s := utils.GetAbsFileWithEnv("./logs/bastion.log")
	log.Init(s, log.DebugLevel, true, log.SetCaller(true))

	var err error
	err = datasource.NewRedisClient()
	utils.Must(err)

	err = datasource.InitGorm()
	utils.Must(err)
}