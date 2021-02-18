package model

import (
	. "server/pkg/db"
	"time"
)

// 创建User接口参数
type CreateUserParam struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}

// User - 对应数据库表
type UserModel struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}

func (UserModel) Paginate(page int, size int) []UserModel {
	limit := size
	offset := (page - 1) * size
	users := []UserModel{}
	DB.Limit(limit).Offset(offset).Find(&users)
	return users
}

func (UserModel) Find(id int) *UserModel {
	user := UserModel{}
	DB.First(&user, id)
	if user.ID == 0 {
		return nil
	}
	return &user
}

func (UserModel) Create(params *CreateUserParam) *UserModel {
	user := UserModel{Name: params.Name}
	DB.Create(&user)
	return &user
	return nil
}

var User = UserModel{}
