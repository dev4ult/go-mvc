package models

import (
	helper "praktikum/helpers"
	"strconv"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Id      int    `json:"id" form:"id" gorm:"type:int(11)"`
	UserId  int    `json:"user_id" form:"user_id" gorm:"type:int(11)"`
	Title   string `json:"title" form:"title" gorm:"type:varchar(100)"`
	Content string `json:"content" form:"content" gorm:"type:varchar(255)"`
}

type BlogModel struct {
	db *gorm.DB
}

func (bm *BlogModel) Init(db *gorm.DB) {
	bm.db = db
}

func (bm *BlogModel) InsertBlog(newBlog Blog) *Blog {
	if err := bm.db.Create(&newBlog).Error; err != nil {
		return nil
	}

	return &newBlog
}

func (bm *BlogModel) FindBlog(paramId string) map[string]interface{} {
	var blog Blog

	blogId, err := strconv.Atoi(paramId)

	if err != nil {
		return helper.Response(400, "Bad Request!")
	}

	result := bm.db.First(&blog, blogId)

	if result.RowsAffected < 1 {
		return helper.Response(404, "Not Found!")
	}

	return map[string]interface{}{
		"status": 200,
		"blog":   blog,
		"id":     blogId,
	}
}

func (bm *BlogModel) SelectAllblogs() []Blog {
	var blogs []Blog
	
	if err := bm.db.Find(&blogs).Error; err != nil {
		return nil
	}

	return blogs
}

func (bm *BlogModel) UpdateBlog(blogId int, newBlog Blog) bool {
	if err := bm.db.Table("blogs").Where("id", blogId).Updates(newBlog).Error; err != nil {
		return false
	}

	return true
}

func (bm *BlogModel) DeleteBlog(blogId int) bool {
	if err := bm.db.Delete(Blog{}, blogId).Error; err != nil {
		return false
	}

	return true
}