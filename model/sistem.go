package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Foto string             `bson:"foto,omitempty" json:"foto,omitempty"`
	Nama string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type Gallery struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Foto           string             `bson:"foto,omitempty" json:"foto,omitempty"`
	Judul_Kegiatan string             `bson:"judul_kegiatan,omitempty" json:"judul_kegiatan,omitempty"`
	Tahun          int                `bson:"tahun,omitempty" json:"tahun,omitempty"`
}

type AdminRegistration struct {
	Fullname        string `json:"fullname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
