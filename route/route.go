package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nc_user.com/v1/config"
	"nc_user.com/v1/handler"
	model "nc_user.com/v1/model"
)

func All(e *echo.Echo) {
	Private(e)
	Staff(e)
	Public(e)
}

func Private(e *echo.Echo) {
	g := e.Group("api/user/v1/priavte")
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(config.Config.JWTSecret.JWTKey),
		Claims:     &model.UserClaims{},
	}
	g.Use(middleware.JWTWithConfig(JWTConfig))
	g.PUT("/user", handler.UpdateUser)
}

func Staff(e *echo.Echo) {
}

func Public(e *echo.Echo) {
	g := e.Group("api/user/v1/public")
	g.POST("/register", handler.RegisterUser)
	g.PATCH("/login", handler.LoginUser)
	// g.GET("/students", handler.GetAllStudents)
	// g.GET("/student", handler.GetOneStudent)
	// g.PATCH("/studentWithKeywords", handler.GetAllStudentsWithKeyword)
}
