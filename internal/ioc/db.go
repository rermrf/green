package ioc

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"green/internal/repository/dao"
)

func InitDB() *gorm.DB {
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	var cfg Config = Config{
		// 这只默认值
		DSN: "root:root@tcp(localhost:3308)/webook?charset=utf8mb4&parseTime=True&loc=Local",
	}
	err := viper.UnmarshalKey("db.mysql", &cfg)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
		//Logger: glogger.New(gormLoggerFunc(l.Debug), glogger.Config{
		//	// 满查询阈值，只有执行时间超过这个阈值，才会使用
		//	// 50ms, 100ms
		//	// SQL 查询必然要求命中索引，最好就是走一次磁盘 IO
		//	// 一次磁盘 IO 是不到 10ms
		//	SlowThreshold:             time.Millisecond * 10,
		//	IgnoreRecordNotFoundError: true,
		//	ParameterizedQueries:      false,
		//	LogLevel:                  glogger.Info,
		//}),
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
