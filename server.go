package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	config "praktikum/config"
	controller "praktikum/controllers"
	model "praktikum/models"
	route "praktikum/routes"
)

var (
	conn config.Connection = config.LoadConfig()
	db *gorm.DB = config.InitDB(conn)

	userModel model.UserModel = model.UserModel{}
	bookModel model.BookModel = model.BookModel{}
	blogModel model.BlogModel = model.BlogModel{}
	
	userController controller.UserController = controller.UserController{}
	bookController controller.BookController = controller.BookController{}
	blogController controller.BlogController = controller.BlogController{}
)

func main() {
	userModel.Init(db)
	userController.InitUserController(userModel)
	
	bookModel.Init(db)
	bookController.InitBookController(bookModel)
	
	blogModel.Init(db)
	blogController.InitBlogController(blogModel)

    e := echo.New()

	route.UserRoutes(e, userController)
	route.BookRoutes(e, bookController)
	route.BlogRoutes(e, blogController)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%d", conn.SERVER_PORT)))
}