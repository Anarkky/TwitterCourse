package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* User is a model of users for MongoDB data base*/
type User struct {
	//Objeto binario - slice of bits
	ID        primitive.ObjectID `bson:"_id, omitempty" json: "id`
	Name      string             `bson:"nombre" json: "name,omitempty`
	LastName  string             `bson:"lastname" json: lastname,omitempty`
	BirthDate time.Time          `bson:"birthDate" json: birthDate,omitempty`
	Email     string             `bson:"email" json: email`
	Password  string             `bson:"password" json: password,omitempty`
	Avatar    string             `bson:"avatar" json: avatar,omitempty`
	Banner    string             `bson:"banner" json: banner,omitempty`
	Biography string             `bson:"biography" json: biography,omitempty`
	Location  string             `bson:"location" json: location,omitempty`
	WebSite   string             `bson:"webSite" json: webSite,omitempty`
}
