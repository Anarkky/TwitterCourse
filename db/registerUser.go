package db

import (
	"context"
	"time"

	"github.com/Anarkky/TwitterCourse/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(u models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCn.Database("twittercurse")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	//Get insert id
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
