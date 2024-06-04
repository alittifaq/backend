package handler

import (
	"context"
	"gocroot/config"
	"gocroot/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetRoutes(c *fiber.Ctx) error {
	var routes []model.RuteAngkot
	collection := config.GetCollection("routes")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching routes"})
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var route model.RuteAngkot
		if err := cursor.Decode(&route); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error decoding route"})
		}
		routes = append(routes, route)
	}

	return c.JSON(routes)
}
func UpdateRoute(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid route ID"})
	}

	var route model.RuteAngkot
	if err := c.BodyParser(&route); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	collection := config.GetCollection("routes")
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": route})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error updating route"})
	}

	return c.JSON(route)
}
func DeleteRoute(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid route ID"})
	}

	collection := config.GetCollection("routes")
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting route"})
	}

	return c.SendStatus(http.StatusOK)
}
func GetPhoneNumber(c *fiber.Ctx) error {
	return c.SendString("Login")
}
