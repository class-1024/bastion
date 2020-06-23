package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"log"
	"xorm.io/core"
)

var XormEngine *xorm.Engine

func InitXorm() error {
	driverName := viper.GetString("mysql.driver_name")
	dataSourceName := viper.GetString("mysql.data_source_name")

	var (
		orm *xorm.Engine
		err error
	)
	if orm, err = xorm.NewEngine(driverName, dataSourceName);err!=nil{
		return err
	}

	orm.SetMaxIdleConns(10)
	orm.SetMaxOpenConns(20)

	logger := xorm.NewSimpleLogger(log.Writer())
	logger.SetLevel(core.LOG_INFO)
	logger.ShowSQL(true)
	orm.SetLogger(logger)
	orm.ShowExecTime(true)

	XormEngine = orm

	fmt.Println("初始化 xorm")
	return nil
}
