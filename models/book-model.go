package models

import (
	helper "praktikum/helpers"
	"strconv"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id        int    `json:"id" form:"id" gorm:"type:int(11)"`
	Title     string `json:"title" form:"title" gorm:"type:varchar(100)"`
	Author    string `json:"author" form:"author" gorm:"type:varchar(255)"`
	Publisher string `json:"publisher" form:"publisher" gorm:"type:varchar(255)"`
}

type BookModel struct {
	db *gorm.DB
}

func (bm *BookModel) Init(db *gorm.DB) {
	bm.db = db
}

func (bm *BookModel) InsertBook(newBook Book) *Book {
	if err := bm.db.Create(&newBook).Error; err != nil {
		return nil
	}

	return &newBook
}

func (bm *BookModel) FindBook(paramId string) map[string]interface{} {
	var book Book

	bookId, err := strconv.Atoi(paramId)

	if err != nil {
		return helper.Response(400, "Bad Request!")
	}

	result := bm.db.First(&book, bookId)

	if result.RowsAffected < 1 {
		return helper.Response(404, "Not Found!")
	}

	return map[string]interface{}{
		"status": 200,
		"book":   book,
		"id":     bookId,
	}
}

func (bm *BookModel) SelectAllBooks() []Book {
	var books []Book
	
	if err := bm.db.Find(&books).Error; err != nil {
		return nil
	}

	return books
}

func (bm *BookModel) UpdateBook(bookId int, newBook Book) bool {
	if err := bm.db.Table("books").Where("id", bookId).Updates(newBook).Error; err != nil {
		return false
	}

	return true
}

func (bm *BookModel) DeleteBook(bookId int) bool {
	if err := bm.db.Delete(Book{}, bookId).Error; err != nil {
		return false
	}

	return true
}