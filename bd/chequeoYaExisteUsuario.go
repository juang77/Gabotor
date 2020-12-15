package bd

import (
	"context"
	"time"

	"github.com/juang77/Gabotor/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoYaExisteUsuario is La funcion que permite revisar si un usuario ya existe.
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("Gabotor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
