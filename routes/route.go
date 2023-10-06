package routes

import (
	"github.com/labstack/echo/v4"

	controller "praktikum/controllers"
)

func UserRoutes(e *echo.Echo, uc controller.UserController) {
	users := e.Group("/users")
	users.GET("", uc.GetUsers())
	users.POST("", uc.CreateUser())

	users.GET("/:id", uc.GetUser())
	users.PUT("/:id", uc.EditUser())
	users.DELETE("/:id", uc.RemoveUser())
}

func BookRoutes(e *echo.Echo, bc controller.BookController) {
	books := e.Group("/books")
	books.GET("", bc.GetBooks)
	books.POST("", bc.CreateBook)

	books.GET("/:id", bc.GetBook)
	books.PUT("/:id", bc.EditBook)
	books.DELETE("/:id", bc.RemoveBook)
}

func BlogRoutes(e *echo.Echo, bc controller.BlogController) {
	blogs := e.Group("/blogs")
	blogs.GET("", bc.GetBlogs)
	blogs.POST("", bc.CreateBlog)

	blogs.GET("/:id", bc.GetBlog)
	blogs.PUT("/:id", bc.UpdateBlog)
	blogs.DELETE("/:id", bc.DeleteBlog)
}