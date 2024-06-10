package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProduk(respw http.ResponseWriter, req *http.Request) {
	produk, err := atdb.GetAllDoc[[]model.Produk](config.Mongoconn, "produk", bson.M{})
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, produk)
}

func PostProduk(respw http.ResponseWriter, req *http.Request) {
	var newProduk model.Produk
	if err := json.NewDecoder(req.Body).Decode(&newProduk); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}
	newProduk.ID = primitive.NewObjectID()
	if _, err := atdb.InsertOneDoc(config.Mongoconn, "produk", newProduk); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusCreated, newProduk)
}

func PutProduk(respw http.ResponseWriter, req *http.Request) {
	id := helper.GetParam(req)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Invalid ID")
		return
	}
	var updatedProduk model.Produk
	if err := json.NewDecoder(req.Body).Decode(&updatedProduk); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedProduk}
	if _, err := atdb.UpdateDoc(config.Mongoconn, "produk", filter, update); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, updatedProduk)
}

func DeleteProduk(respw http.ResponseWriter, req *http.Request) {
	id := helper.GetParam(req)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Invalid ID")
		return
	}
	filter := primitive.M{"_id": objectID}
	err = atdb.DeleteOneDoc(config.Mongoconn, "produk", filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, "Produk deleted")
}
