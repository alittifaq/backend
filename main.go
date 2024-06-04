package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Route struct represents a document in MongoDB
type Route struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Rute           string             `json:"Rute,omitempty" bson:"Rute,omitempty"`
	JamOperasional string             `json:"Jam Operasional,omitempty" bson:"Jam Operasional,omitempty"`
	Tarif          string             `json:"Tarif,omitempty" bson:"Tarif,omitempty"`
}

var MongoClient *mongo.Client


func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGOSTRING")

	// Debugging line to print the connection string
	log.Println("MONGOSTRING:", mongoURI)

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	}

	log.Println("Connected to MongoDB!")
	MongoClient = client

	// Initialize Fiber
	app := fiber.New()
	app.Use(cors.New())

	// Define routes
	app.Get("/routes", getRoutes)
	app.Post("/routes", createRoute)
	app.Put("/routes/:id", editRoute)
	app.Delete("/routes/:id", deleteRoute) 

	// Start Fiber app
	log.Fatal(app.Listen(":3000"))
}

// getRoutes handles the GET /routes route
func getRoutes(c *fiber.Ctx) error {
	collection := MongoClient.Database("Angkutankotabdg").Collection("data json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query routes",
		})
	}
	defer cursor.Close(ctx)

	var routes []Route
	for cursor.Next(ctx) {
		var route Route
		if err := cursor.Decode(&route); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to decode route",
			})
		}
		routes = append(routes, route)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cursor error",
		})
	}

	return c.JSON(routes)
}

// createRoute handles the POST /routes route
func createRoute(c *fiber.Ctx) error {
	collection := MongoClient.Database("Angkutankotabdg").Collection("data json")

	var route Route
	if err := c.BodyParser(&route); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, route)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create route",
		})
	}

	// Jika berhasil tambahkan data, kembalikan respons berhasil bersama data yang ditambahkan
	return c.JSON(fiber.Map{
		"message": "Route created successfully",
		"data":    route,
	})
}

func editRoute(c *fiber.Ctx) error {
	collection := MongoClient.Database("Angkutankotabdg").Collection("data json")

	routeId := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(routeId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid route ID",
		})
	}

	var route Route
	if err := c.BodyParser(&route); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": route,
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update route",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Route updated successfully",
		"data":    route,
	})
}

// deleteRoute handles the DELETE /routes/:id route
func deleteRoute(c *fiber.Ctx) error {
	collection := MongoClient.Database("Angkutankotabdg").Collection("data json")

	routeID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(routeID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid route ID",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete route",
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Route not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Route deleted successfully",
	})
}
