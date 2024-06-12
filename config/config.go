package config

import (
	"log"

	"github.com/gocroot/helper"
)

var IPPort, Net = helper.GetAddress()

func SetEnv() {
	if ErrorMongoconn != nil {
		log.Println(ErrorMongoconn.Error())
	}
	// 	profile, err := atdb.GetOneDoc[model.Profile](Mongoconn, "profile", primitive.M{})
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	PublicKeyWhatsAuth = profile.PublicKey
	// 	WAAPIToken = profile.Token
}
