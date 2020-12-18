package bd

import (
	"context"
	"time"

	"github.com/juang77/Gabotor/models"
)

//InsertoRelacion is la funcion que permite insertar relaciones
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("Gabotor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
