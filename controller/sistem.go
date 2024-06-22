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
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProduk(respw http.ResponseWriter, req *http.Request) {
	produk, err := atdb.GetAllDoc[[]model.Product](config.Mongoconn, "product", bson.M{})
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, produk)
}

func PostProduk(respw http.ResponseWriter, req *http.Request) {
	var newProduk model.Product
	if err := json.NewDecoder(req.Body).Decode(&newProduk); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}
	newProduk.ID = primitive.NewObjectID()
	if _, err := atdb.InsertOneDoc(config.Mongoconn, "product", newProduk); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, newProduk)
}

func UpdateProduct(respw http.ResponseWriter, req *http.Request) {
	var product model.Product
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	// Definisikan filter untuk menemukan produk berdasarkan nama produk
	filter := bson.M{"nama": product.Nama}

	// Get data produk berdasarkan nama produk
	existingProduct, err := atdb.GetOneDoc[model.Product](config.Mongoconn, "product", filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	// Pastikan produk ditemukan sebelum melakukan update
	if existingProduct.ID == primitive.NilObjectID {
		helper.WriteJSON(respw, http.StatusNotFound, "Product not found")
		return
	}

	// Update data produk yang ditemukan
	if _, err := atdb.ReplaceOneDoc(config.Mongoconn, "product", filter, product); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	// Tulis respons sukses
	helper.WriteJSON(respw, http.StatusOK, product)
}

func DeleteProduk(respw http.ResponseWriter, req *http.Request) {
	var newProduk model.Product
	if err := json.NewDecoder(req.Body).Decode(&newProduk); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.M{"nama": newProduk.Nama}
	err := atdb.DeleteOneDoc(config.Mongoconn, "product", filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, newProduk)
}

func GetGallery(respw http.ResponseWriter, req *http.Request) {
	gallery, err := atdb.GetAllDoc[[]model.Gallery](config.Mongoconn, "gallery", bson.M{})
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, gallery)
}

func PostGallery(respw http.ResponseWriter, req *http.Request) {
	var newGallery model.Gallery
	if err := json.NewDecoder(req.Body).Decode(&newGallery); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}
	newGallery.ID = primitive.NewObjectID()
	if _, err := atdb.InsertOneDoc(config.Mongoconn, "gallery", newGallery); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, newGallery)
}

func PutGallery(respw http.ResponseWriter, req *http.Request) {
	var newGallery model.Gallery
	if err := json.NewDecoder(req.Body).Decode(&newGallery); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	filter := bson.M{"judul_kegiatan": newGallery.Judul_Kegiatan}
	update := bson.M{"$set": newGallery}
	if _, err := atdb.UpdateDoc(config.Mongoconn, "gallery", filter, update); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, newGallery)
}

func DeleteGallery(respw http.ResponseWriter, req *http.Request) {
	var newGallery model.Gallery
	if err := json.NewDecoder(req.Body).Decode(&newGallery); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.M{"judul_kegiatan": newGallery.Judul_Kegiatan}
	err := atdb.DeleteOneDoc(config.Mongoconn, "gallery", filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, "Gallery deleted")
}

func GetOneGallery(respw http.ResponseWriter, req *http.Request) {
	var requestBody struct {
		Judul_Kegiatan string `json:"judul_kegiatan"`
	}

	// Decode request body
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Invalid request body")
		return
	}

	if requestBody.Judul_Kegiatan == "" {
		helper.WriteJSON(respw, http.StatusBadRequest, "Missing gallery title")
		return
	}

	// Create filter to search for the document with the given activity title
	filter := bson.M{"judul_kegiatan": requestBody.Judul_Kegiatan}

	// Retrieve one gallery document
	gallery, err := atdb.GetOneDoc[model.Gallery](config.Mongoconn, "gallery", filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helper.WriteJSON(respw, http.StatusNotFound, "Gallery not found")
		} else {
			helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Return the gallery document in JSON format
	helper.WriteJSON(respw, http.StatusOK, gallery)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registrationData model.AdminRegistration

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&registrationData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Lakukan validasi dan pemrosesan data di sini
	if registrationData.Password != registrationData.ConfirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// Simpan data ke database atau lakukan tindakan lain yang diperlukan
	_, err = atdb.InsertOneDoc(config.Mongoconn, "user", registrationData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Registration successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUser(respw http.ResponseWriter, req *http.Request) {
	var loginDetails model.User
	if err := json.NewDecoder(req.Body).Decode(&loginDetails); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	var user model.User
	filter := bson.M{"email": loginDetails.Email, "password": loginDetails.Password}
	user, err := atdb.GetOneDoc[model.User](config.Mongoconn, "user", filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	helper.WriteJSON(respw, http.StatusOK, user)
}
