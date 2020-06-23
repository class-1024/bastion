package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	GormPool *gorm.DB
)

func InitGorm() error {
	driverName := viper.GetString("mysql.driver_name")
	dataSourceName := viper.GetString("mysql.data_source_name")

	Gorm, err := gorm.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	isLog := viper.GetString("base.env") == "dev"

	Gorm.LogMode(isLog)
	Gorm.DB().SetMaxIdleConns(10)
	Gorm.DB().SetMaxOpenConns(20)
	Gorm.SingularTable(true)

	GormPool = Gorm

	fmt.Printf("%v \n", "gorm ready!")
	return nil
}
