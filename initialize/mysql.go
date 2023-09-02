package initialize

import (
	"go-template/global"
	"go-template/models"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
)

/*
Created by 斑斑砖 on 2023/8/13.
Description：
	初始化mysql，

*/

func InitMysql() {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       global.Config.Mysql.Dsn(),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}), &gorm.Config{
		Logger: getGormLogger(), // 打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: global.Config.Mysql.Singular, // 表明不加s
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConn) // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConn) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	global.MysqlDB = db
	global.Logger.Debug("mysql初始化成功")

	//	进行表迁移
	models.Migrate()

}
func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.Config.Mysql.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                   // 慢 SQL 阈值
		LogLevel:                  logMode,                                  // 日志级别
		IgnoreRecordNotFoundError: false,                                    // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.Config.Mysql.EnableFileLogWriter, // 禁用彩色打印
	})
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.Config.Mysql.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.Config.Log.RootDir + "/" + global.Config.Mysql.LogFilename,
			MaxSize:    global.Config.Log.MaxSize,
			MaxBackups: global.Config.Log.MaxBackups,
			MaxAge:     global.Config.Log.MaxAge,
			Compress:   global.Config.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}
