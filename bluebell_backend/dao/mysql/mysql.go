package mysql

import (
	"bluebell_backend/settings"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init 初始化MySQL连接
func Init(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		DisableAutomaticPing: false,
	})
	if err != nil {
		fmt.Println("数据库配置失败", err)
		return err
	}
	DB = db
	fmt.Println("数据库开启成功")
	return
}

// Close 关闭MySQL连接
func Close() {
	sql, _ := DB.DB()
	err := sql.Close()
	if err != nil {
		fmt.Println("数据库关闭失败", err)
		return
	}
}
