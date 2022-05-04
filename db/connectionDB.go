package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCn is the data base connections variable*/
var MongoCn = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://sphzer0:Guild$Wars2@cluster0.p3zaz.mongodb.net/test")

/* Connect db provides data base connection*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión Exitosa con la BD")
	return client
}

/* VerifyConnection makes a ping to data base*/
func VerifyConnection() int {
	err := MongoCn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
