package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RuteAngkot struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Rute           string             `bson:"Rute" json:"Rute"`
	JamOperasional string             `bson:"Jam Operasional" json:"Jam Operasional"`
	Tarif          string             `bson:"Tarif" json:"Tarif"`
}
