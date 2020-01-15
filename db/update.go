package database

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	model "nc_user.com/v1/model"
	"nc_user.com/v1/utils"
)

func UpdateOneUser(req *model.UserUpdateReq) (interface{}, error) {
	user, _ := FindUserByID(req.ID)

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.PassWord != "" {
		user.PassWord = utils.Md5(req.PassWord)
	}

	fmt.Printf("user :%+v", user)

	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": user}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}

func FindUserByID(ID int) (*model.User, error) {
	var user model.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": ID} //map[string]interface{}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)

	return &user, err
}
