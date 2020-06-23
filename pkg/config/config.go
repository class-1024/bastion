package config

import (
	"bastion/utils"
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type configFileInfo struct {
	Type    string
	Name    string
	Path    string
	AbsPath string
}

func Load(path string) {
	file := utils.GetAbsFileWithEnv(path)

	dir := filepath.Dir(file)
	base := filepath.Base(file)
	b := strings.Split(base, ".")

	if len(b) < 2 {
		panic("config file ext split error")
	}

	viperRead(configFileInfo{
		Name:    b[0],
		Type:    b[1],
		Path:    dir,
		AbsPath: file,
	})
}

func viperRead(conf configFileInfo) {
	viper.SetConfigType(conf.Type)
	viper.SetConfigName(conf.Name) // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(conf.Path) // 第一个搜索路径

	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Printf("%v %v\n", "viperRead success: ", conf.AbsPath)
}
