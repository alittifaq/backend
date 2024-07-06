package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Foto     string             `bson:"foto,omitempty" json:"foto,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Gallery struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Foto           string             `bson:"foto,omitempty" json:"foto,omitempty"`
	Judul_Kegiatan string             `bson:"judul_kegiatan,omitempty" json:"judul_kegiatan,omitempty"`
	Tahun          int                `bson:"tahun,omitempty" json:"tahun,omitempty"`
}

type AdminRegistration struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Fullname        string             `json:"fullname" bson:"fullname" validate:"required"`
	Email           string             `json:"email" bson:"email" validate:"required,email"`
	Password        string             `json:"password" bson:"password" validate:"required,min=8"`
	ConfirmPassword string             `json:"confirm_password" bson:"confirm_password" validate:"required,eqfield=Password"`
}

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Password string             `json:"password" bson:"password" validate:"required,min=8"`
}

type Feedback struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Rating  int                `json:"rating" bson:"rating" validate:"required,min=1,max=5"`
	Content string             `json:"content" bson:"content" validate:"required"`
	Sender  string             `json:"sender" bson:"sender" validate:"required"`
}
