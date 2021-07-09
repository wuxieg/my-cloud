package model

import db2 "wangy1.top/my_cloud/pkg/db"

type Role struct {
	BaseModel
	Name     string `gorm:"column:name;type:varchar(20);not null"`
	Describe string `gorm:"column:describe;type:varchar(20);not null"`
}

// TableName 表名
func (Role) TableName() string {
	return "t_role"
}
func (u *Role) Create() (err error) {
	if err = db2.DB.Create(&u).Error; err != nil {
		return err
	}
	return
}

// Update 更新非空字段 /*
func (u *Role) Update() (err error) {
	if err = db2.DB.Model(&User{}).Updates(&u).Error; err != nil {
		return err
	}
	return
}

func (u *Role) Delete() (err error) {
	if err = db2.DB.Model(&User{}).Delete(&u).Error; err != nil {
		return err
	}
	return
}
