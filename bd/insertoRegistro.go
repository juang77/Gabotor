package bd

import (
	"context"
	"time"

	"github.com/juang77/Gabotor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoRegistro is la funcion que permite insertar un nuevo registro.
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("Gabotor")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
