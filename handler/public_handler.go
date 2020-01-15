package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	db "nc_user.com/v1/db"
	model "nc_user.com/v1/model"
)

// func GetAllStudentsWithKeyword(c echo.Context) error {
// 	var student model.User
// 	if err := c.Bind(&student); err != nil {
// 		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
// 	}
// 	students, err2 := db.FindNameStartsWith(&student)
// 	if err2 != nil {
// 		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
// 	}

// 	return c.JSON(http.StatusOK, students)
// }

func RegisterUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.AddOneUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

func LoginUser(c echo.Context) error {
	var req model.LoginReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}
