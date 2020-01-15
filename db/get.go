package database

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	model "nc_user.com/v1/model"
	"nc_user.com/v1/utils"
)

func LoginUser(req model.LoginReq) (*model.LoginRes, error) {
	var user model.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"Email": req.Email, "PassWord": utils.Md5(req.PassWord)} //map[string]interface{}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)
	token, _ := utils.JwtGenerate(user.ID, user.Phone, user.Email)
	resp := model.LoginRes{&user, token}
	fmt.Println(err, resp)
	return &resp, err
}
