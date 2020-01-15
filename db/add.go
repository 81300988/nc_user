package database

import (
	"context"
	"time"

	model "nc_user.com/v1/model"
	"nc_user.com/v1/utils"
)

func AddOneUser(user *model.User) (interface{}, error) {
	collection := Client.Database("homework2").Collection("user")
	user.PassWord = utils.Md5(user.PassWord)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, &user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
