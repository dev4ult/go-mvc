package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	helper "praktikum/helpers"
	model "praktikum/models"
)

type BlogController struct {
	model model.BlogModel
}

func (bc *BlogController) InitBlogController(m model.BlogModel) {
	bc.model = m
}

func (bc *BlogController) GetBlogs(ctx echo.Context) error {
	blogs := bc.model.SelectAllblogs()

	if blogs == nil {
		return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "blogs Listed!",
		"blogs": blogs,
	})
}

func (bc *BlogController) CreateBlog(ctx echo.Context) error {
	var blog model.Blog

	ctx.Bind(&blog)

	result := bc.model.InsertBlog(blog)

	if result == nil {
		ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something went wrong!"))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "blog Created!",
		"blog": result,
	})
}

func (bc *BlogController) GetBlog(ctx echo.Context) error {
	blog := bc.model.FindBlog(ctx.Param("id"))

	if blog["status"] != http.StatusOK {
		return ctx.JSON(blog["status"].(int), blog)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "blog Found!",
		"blog": blog["blog"],
	})
}

func (bc *BlogController) UpdateBlog(ctx echo.Context) error {
	blog := bc.model.FindBlog(ctx.Param("id"))

	if blog["status"] != http.StatusOK {
		return ctx.JSON(blog["status"].(int), blog)
	}

	var newBlog model.Blog

	ctx.Bind(&newBlog)

	update := bc.model.UpdateBlog(blog["id"].(int), newBlog)

	if !update {
		return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, helper.Response(200, "blog Updated!"))
}

func (bc *BlogController) DeleteBlog(ctx echo.Context) error {
	blog := bc.model.FindBlog(ctx.Param("id"))

	if blog["status"] != http.StatusOK {
		return ctx.JSON(blog["status"].(int), blog)
	}

	delete := bc.model.DeleteBlog(blog["id"].(int))

	if !delete {
		return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, helper.Response(200, "blog Deleted!"))
}