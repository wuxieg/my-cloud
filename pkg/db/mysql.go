package db

import (
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitMysql() *gorm.DB {
	url := createdbURL(viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"),
		viper.GetInt("db.port"), viper.GetString("db.database"))
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Database connection failed. Database url: "+url+" error: ", err)
	} else {
		fmt.Print("\n\n------------------------------------------ GORM OPEN SUCCESS! -----------------------------------------------\n\n")
	}

	DB = db
	return db
}

func createdbURL(uname string, pwd string, host string, port int, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s",
		uname, pwd,
		host, port,
		name, true, "Local")
}
