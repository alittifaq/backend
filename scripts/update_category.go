package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Product struct
type Product struct {
    ID     primitive.ObjectID `bson:"_id,omitempty"`
    Foto   string             `bson:"foto"`
    Nama   string             `bson:"nama"`
    Kategori string           `bson:"kategori,omitempty"`
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://blkkalittifaq:blkkalittifaq1@cluster0.din9pla.mongodb.net/")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("blkkalittifaq").Collection("product")

	// Define the category based on product name
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var products []Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		var category string
		switch {
		case product.Nama == "Strawberry", product.Nama == "Buah Tin", product.Nama == "Jeruk Dekopon":
			category = "buah"
		case product.Nama == "Sawi Putih", product.Nama == "Daun Bawang", product.Nama == "Seledri":
			category = "sayur"
		default:
			category = "olahan"
		}

		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "kategori", Value: category},
			}},
		}

		_, err := collection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: product.ID}}, update)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("All documents updated with category field.")
}