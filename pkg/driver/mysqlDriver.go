package driver

import (
	"context"
	"fmt"
	"gitee.com/chunanyong/zorm"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"jwk/test/pkg/config"
	"jwk/test/utils"
)

var ctx = context.Background()
var Db *zorm.DBDao

func init() {
	InitMysql()
}

func InitMysql() {
	var err error
	mysqlConfig := config.Config.Mysql

	dbConfig := &zorm.DataSourceConfig{
		DSN:                   fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database),
		DriverName:            "mysql", //数据库驱动名称
		Dialect:               "mysql", //数据库类型
		SlowSQLMillis:         0,       //慢sql的时间阈值,单位毫秒.小于0是禁用SQL语句输出;等于0是只输出SQL语句,不计算执行时间;大于0是计算SQL执行时间,并且>=SlowSQLMillis值
		MaxOpenConns:          50,      //数据库最大连接数,默认50
		MaxIdleConns:          50,      //数据库最大空闲连接数,默认50
		ConnMaxLifetimeSecond: 600,     //连接存活秒时间. 默认600(10分钟)后连接被销毁重建.
		//避免数据库主动断开连接,造成死连接.MySQL默认wait_timeout 28800秒(8小时)
		DefaultTxOptions: nil, //事务隔离级别的默认配置,默认为nil
	}

	Db, err = zorm.NewDBDao(dbConfig)

	if err != nil {
		utils.Logger.Error("数据库连接异常 ", zap.Error(err))
		panic(err)
	}
}
