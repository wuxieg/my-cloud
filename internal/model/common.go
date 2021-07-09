package model

import (
	"github.com/lexkong/log"
	"gorm.io/gorm"
	db2 "wangy1.top/my_cloud/pkg/db"
)

func SetupDb() *gorm.DB {
	db := db2.InitMysql()

	err := db.AutoMigrate(
		&User{}, &Role{})

	if err != nil {
		log.Fatal("db autoMigrate error", err)
		return nil
	}

	return db
}
