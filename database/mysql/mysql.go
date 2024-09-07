package mysql

import (
	"fmt"
	"time"

	"github.com/newdee/aipaper-util/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetGlobalDBIns() *gorm.DB {
	return DB
}

func InitDatabase(dbConfig database.DatabaseConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.GetUser(),
		dbConfig.GetPassword(), dbConfig.GetHost(), dbConfig.GetPort(), dbConfig.GetDataBase())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("连接数据库失败, error=" + err.Error())
	}
	DB = db

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	return nil
}
