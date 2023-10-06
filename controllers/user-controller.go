package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	helper "praktikum/helpers"
	model "praktikum/models"
)

type UserController struct {
	model model.UserModel
}

func (uc *UserController) InitUserController(um model.UserModel) {
	uc.model = um
}

func (uc *UserController) GetUsers() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var users []model.User
	
		err := uc.model.SelectAllUser(&users)
	
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, err.Error()))
		}
	
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Users Listed!",
			"users": users,
		})
	}
}

func (uc *UserController) CreateUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var user model.User
	
		ctx.Bind(user)
	
		err := uc.model.InsertUser(&user)
	
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, err.Error()))
		}
	
		return ctx.JSON(http.StatusOK, map[string]interface{} {
			"message": "New User Created!",
			"user": user,
		})
	}
}

func (uc *UserController) GetUser() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		user := uc.model.FindUser(ctx.Param("id"))
	
		if user["status"] != http.StatusOK {
			return ctx.JSON(user["status"].(int), user)
		}
	
		return ctx.JSON(http.StatusOK, map[string]interface{} {
			"message": "User Found!",
			"user": user["user"],
		})
	}
}

func (uc *UserController) EditUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// check user if exist
		user := uc.model.FindUser(ctx.Param("id"))
	
		if user["status"] != http.StatusOK {
			return ctx.JSON(user["status"].(int), user)
		}
	
		var newUserData model.User
	
		ctx.Bind(&newUserData)
	
		err := uc.model.UpdateUser(user["id"].(int), newUserData)
	
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
		}
	
		return ctx.JSON(http.StatusOK, helper.Response(200, "User Updated!"))
	}
}

func (uc *UserController) RemoveUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// check user if exist
		user := uc.model.FindUser(ctx.Param("id"))
	
		if user["status"] != http.StatusOK {
			return ctx.JSON(user["status"].(int), user)
		}
	
		err := uc.model.DeleteUser(user["id"].(int))
	
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
		}
	
		return ctx.JSON(http.StatusOK, helper.Response(200, "User Deleted!"))
	}
}