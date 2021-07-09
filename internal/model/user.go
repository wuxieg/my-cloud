package model

import (
	"time"
	"wangy1.top/my_cloud/internal/handler/user/request"
	db2 "wangy1.top/my_cloud/pkg/db"
)

type User struct {
	BaseModel
	Username      string    `gorm:"column:username;type:varchar(50);not null"`
	Password      string    `gorm:"column:password;type:varchar(20);not null"`
	LastLoginTime time.Time `gorm:"column:last_login_time"`
}

// TableName 表名
func (User) TableName() string {
	return "t_user"
}

func (u *User) Create() (err error) {
	if err = db2.DB.Create(&u).Error; err != nil {
		return err
	}
	return
}

// Update 更新非空字段 /*
func (u *User) Update() (err error) {
	if err = db2.DB.Model(&User{}).Updates(&u).Error; err != nil {
		return err
	}
	return
}

func (u *User) Delete() (err error) {
	if err = db2.DB.Model(&User{}).Delete(&u).Error; err != nil {
		return err
	}
	return
}

type UserQuery struct {
}

func (uq *UserQuery) GetById(id string) (*User, error) {
	u := &User{}
	if err := db2.DB.First(&u, id).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// List 获取用户列表
func (uq *UserQuery) List(pageParams *request.UserListParams) (users []*User, err error) {
	users = make([]*User, 0)
	var offset, limit = 0, 0
	if pageParams.Page > 0 {
		offset = (pageParams.Page - 1) * pageParams.Size
		limit = pageParams.Size
	}
	if err := db2.DB.Offset(offset).Limit(limit).Order(pageParams.OrderBy).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
