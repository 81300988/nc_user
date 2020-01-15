package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	db "nc_user.com/v1/db"
	model "nc_user.com/v1/model"
)

func UpdateUser(c echo.Context) error {
	var user model.UserUpdateReq
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.UpdateOneUser(&user)
	if err != nil {
		log.Printf("update error :%v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)

}
