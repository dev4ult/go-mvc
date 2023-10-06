package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	helper "praktikum/helpers"
	model "praktikum/models"
)

type BookController struct {
	model model.BookModel
}

func (bc *BookController) InitBookController(model model.BookModel) {
	bc.model = model
}

func (bc *BookController) GetBooks(ctx echo.Context) error {
	books := bc.model.SelectAllBooks()

	if books == nil {
		return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something went wrong!"))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "Books Listed!",
		"books": books,
	})
}

func (bc *BookController) CreateBook(ctx echo.Context) error {
	var book model.Book
	ctx.Bind(&book)

	result := bc.model.InsertBook(book)

	if result == nil {
		ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something went wrong"))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "Book Created!",
		"book": book,
	})
}

func (bc *BookController) GetBook(ctx echo.Context) error {
	book := bc.model.FindBook(ctx.Param("id"))

	if book["status"] != http.StatusOK {
		return ctx.JSON(book["status"].(int), book)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "Book Found!",
		"book": book["book"],
	})
}

func (bc *BookController) EditBook(ctx echo.Context) error {
	book := bc.model.FindBook(ctx.Param("id"))

	if book["status"] != http.StatusOK {
		return ctx.JSON(book["status"].(int), book)
	}

	var newBook model.Book

	ctx.Bind(&newBook)

	update := bc.model.UpdateBook(book["id"].(int), newBook)

	if !update {
		return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, helper.Response(200, "Book Updated!"))
}

func (bc *BookController) RemoveBook(ctx echo.Context) error {
	book := bc.model.FindBook(ctx.Param("id"))

	if book["status"] != http.StatusOK {
		return ctx.JSON(book["status"].(int), book)
	}

	delete := bc.model.DeleteBook(book["id"].(int))

	if !delete {
		return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, helper.Response(200, "book Deleted!"))
}