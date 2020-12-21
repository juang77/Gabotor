package bd

import (
	"context"
	"log"
	"time"

	"github.com/juang77/Gabotor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoTweets is la funcion que devuelvo los tweets de un usuario
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("Gabotor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetSkip((pagina - 1) * 20)
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)

	}
	return resultados, true
}
