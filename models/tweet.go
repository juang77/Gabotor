package models

//Tweet is the model for Tweets on DB
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
