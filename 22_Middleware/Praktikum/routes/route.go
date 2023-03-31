package routes

import (
	"praktikum/constants"
	"praktikum/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func New() *echo.Echo {
	e := echo.New()
	jwtMiddleware := middleware.JWT([]byte(constants.SECRET_JWT))

	// users route
	e.GET("/users", controllers.GetUsersController, jwtMiddleware)
	e.GET("/users/:id", controllers.GetUserController, jwtMiddleware)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController, jwtMiddleware)
	e.PUT("/users/:id", controllers.UpdateUserController, jwtMiddleware)

	// books route
	e.GET("/books", controllers.GetBooksController, jwtMiddleware)
	e.GET("/books/:id", controllers.GetBookController, jwtMiddleware)
	e.POST("/books", controllers.CreateBookController, jwtMiddleware)
	e.DELETE("/books/:id", controllers.DeleteBookController, jwtMiddleware)
	e.PUT("/books/:id", controllers.UpdateBookController, jwtMiddleware)

	// blogs route
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)

	return e
}
