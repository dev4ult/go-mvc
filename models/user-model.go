package models

import (
	"strconv"

	"gorm.io/gorm"

	helper "praktikum/helpers"
)

type User struct {
	gorm.Model
	Id    int    `json:"id" form:"id" gorm:"type:int(11)"`
	Name  string `json:"name" form:"name" gorm:"type:varchar(255)"`
	Email string `json:"email" form:"email" gorm:"type:varchar(255)"`
	Blog []Blog
}

type UserModel struct {
	db *gorm.DB
}

func (um *UserModel) Init(db *gorm.DB) {
	um.db = db
}

func (um *UserModel) InsertUser(userData interface{}) error {
	err := um.db.Create(userData).Error

	return err
}


func (um *UserModel) FindUser(paramId string) map[string]interface{} {
	var user = User{}

	userId, err := strconv.Atoi(paramId)

	if err != nil {
		return helper.Response(400, "Bad Request!")
	}

	result := um.db.First(&user, userId)

	if result.RowsAffected < 1 {
		return helper.Response(404, "Not Found!")
	}

	return map[string]interface{} {
		"status": 200,
		"user": user,
		"id": userId,
	}
}

func (um *UserModel) SelectAllUser(users *([]User)) error {
	err := um.db.Find(users).Error

	return err
}


func (um *UserModel) UpdateUser(userId int, userData interface{}) error {
	err := um.db.Table("users").Where("id", userId).Updates(userData).Error

	return err
}

func (um *UserModel) DeleteUser(userId int) error {
	err := um.db.Delete(&User{}, userId).Error

	return err
}
