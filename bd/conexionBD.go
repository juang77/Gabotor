package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is the conection with DataBase.
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://gabotorapp:Nicolas8032367@cluster0.fvcit.mongodb.net/Gabotor?retryWrites=true&w=majority")

// ConectarBD is The conection singleton.
func ConectarBD() *mongo.Client {
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

// ChequeoConnection is the function by check conection status.
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
