package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Produk struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Foto string             `bson:"foto,omitempty" json:"foto,omitempty"`
	Nama string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type Galeri struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Foto           string             `bson:"foto,omitempty" json:"foto,omitempty"`
	Judul_Kegiatan string             `bson:"judul_kegiatan,omitempty" json:"judul_kegiatan,omitempty"`
	Tahun          int                `bson:"tahun,omitempty" json:"tahun,omitempty"`
}
