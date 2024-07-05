package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetProduk mengambil semua produk dari database dan mengembalikannya sebagai JSON.
func GetProduk(respw http.ResponseWriter, req *http.Request) {
	produk, err := atdb.GetAllDoc[[]model.Product](config.Mongoconn, "product", bson.M{})
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, produk)
}

func GetOneProduk(respw http.ResponseWriter, req *http.Request) {
	// Ambil parameter dari query string
	nama := req.URL.Query().Get("nama")

	if nama == "" {
		helper.WriteJSON(respw, http.StatusBadRequest, "Missing product title")
		return
	}

	// Buat filter untuk mencari dokumen dengan judul kegiatan yang diberikan
	filter := bson.M{"nama": nama}

	// Ambil satu dokumen galeri
	product, err := atdb.GetOneDoc[model.Product](config.Mongoconn, "product", filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helper.WriteJSON(respw, http.StatusNotFound, "Product not found")
		} else {
			helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Kembalikan dokumen galeri dalam format JSON
	helper.WriteJSON(respw, http.StatusOK, product)
}

// PostProduk menambahkan produk baru ke dalam database.
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

	// Validasi bahwa ID produk tidak boleh kosong
	if product.ID == primitive.NilObjectID {
		helper.WriteJSON(respw, http.StatusBadRequest, "ID produk tidak boleh kosong")
		return
	}

	// Definisikan filter untuk menemukan produk berdasarkan ID produk
	filter := bson.M{"_id": product.ID}

	// Definisikan update dengan set data baru
	update := bson.M{
		"$set": bson.M{
			"foto": product.Foto,
			"nama": product.Nama,
		},
	}

	// Update produk di MongoDB
	if _, err := atdb.UpdateDoc(config.Mongoconn, "product", filter, update); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	// Kirim respons sukses
	helper.WriteJSON(respw, http.StatusOK, product)
}

// DeleteProduk menghapus produk dari database berdasarkan namanya.
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
	helper.WriteJSON(respw, http.StatusOK, "Produk berhasil dihapus")
}

// GetGallery mengambil semua item galeri dari database dan mengembalikannya sebagai JSON.
func GetGallery(respw http.ResponseWriter, req *http.Request) {
	gallery, err := atdb.GetAllDoc[[]model.Gallery](config.Mongoconn, "gallery", bson.M{})
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, gallery)
}

func GetOneGallery(respw http.ResponseWriter, req *http.Request) {
	// Ambil parameter dari query string
	judulKegiatan := req.URL.Query().Get("judul_kegiatan")

	if judulKegiatan == "" {
		helper.WriteJSON(respw, http.StatusBadRequest, "Missing gallery title")
		return
	}

	// Buat filter untuk mencari dokumen dengan judul kegiatan yang diberikan
	filter := bson.M{"judul_kegiatan": judulKegiatan}

	// Ambil satu dokumen galeri
	gallery, err := atdb.GetOneDoc[model.Gallery](config.Mongoconn, "gallery", filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helper.WriteJSON(respw, http.StatusNotFound, "Gallery not found")
		} else {
			helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Kembalikan dokumen galeri dalam format JSON
	helper.WriteJSON(respw, http.StatusOK, gallery)
}

// PostGallery menambahkan item galeri baru ke dalam database.
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

// UpdateGallery memperbarui item galeri yang ada di dalam database.
func UpdateGallery(respw http.ResponseWriter, req *http.Request) {
	var gallery model.Gallery
	err := json.NewDecoder(req.Body).Decode(&gallery)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	// Validasi bahwa ID gallery tidak boleh kosong
	if gallery.ID == primitive.NilObjectID {
		helper.WriteJSON(respw, http.StatusBadRequest, "ID gallery tidak boleh kosong")
		return
	}

	// Definisikan filter untuk menemukan gallery berdasarkan ID gallery
	filter := bson.M{"_id": gallery.ID}

	// Definisikan update dengan set data baru
	update := bson.M{
		"$set": bson.M{
			"foto":           gallery.Foto,
			"judul_kegiatan": gallery.Judul_Kegiatan,
			"tahun":          gallery.Tahun,
		},
	}

	// Update gallery di MongoDB
	if _, err := atdb.UpdateDoc(config.Mongoconn, "gallery", filter, update); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	// Kirim respons sukses
	helper.WriteJSON(respw, http.StatusOK, gallery)
}

// DeleteGallery menghapus item galeri dari database berdasarkan judul kegiatan.
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
	helper.WriteJSON(respw, http.StatusOK, "Item galeri berhasil dihapus")
}

// RegisterHandler menghandle permintaan registrasi admin.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
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
		http.Error(w, "Password tidak sesuai", http.StatusBadRequest)
		return
	}

	// Simpan data ke database atau lakukan tindakan lain yang diperlukan
	_, err = atdb.InsertOneDoc(config.Mongoconn, "user", registrationData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Registrasi berhasil"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUser mengambil informasi user dari database berdasarkan email dan password.
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
		helper.WriteJSON(respw, http.StatusUnauthorized, "Email atau password salah")
		return
	}

	helper.WriteJSON(respw, http.StatusOK, user)
}

var validate = validator.New()

// PostFeedback menambahkan feedback baru.
func PostFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback model.Feedback

	// Decode request body ke struct Feedback
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validasi struct Feedback
	err = validate.Struct(feedback)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	feedback.ID = primitive.NewObjectID()

	// Simpan ke database
	_, err = atdb.InsertOneDoc(config.Mongoconn, "feedback", feedback)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	helper.WriteJSON(w, http.StatusOK, feedback)
}

// GetFeedback mengambil semua feedback dari database.
func GetFeedback(w http.ResponseWriter, r *http.Request) {
	var feedbacks []model.Feedback

	cursor, err := config.Mongoconn.Collection("feedback").Find(context.Background(), bson.D{})
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var feedback model.Feedback
		if err = cursor.Decode(&feedback); err != nil {
			helper.WriteJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		feedbacks = append(feedbacks, feedback)
	}

	if err := cursor.Err(); err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, feedbacks)
}

// UpdateFeedback mengupdate feedback berdasarkan ID.
func UpdateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback model.Feedback

	// Decode request body ke struct Feedback
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validasi struct Feedback
	err = validate.Struct(feedback)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Dapatkan koleksi "feedback" dari MongoDB
	collection := config.Mongoconn.Collection("feedback")
	// Tentukan filter berdasarkan ID feedback
	filter := bson.M{"_id": feedback.ID}

	// Tentukan operasi update dengan mengatur data baru
	update := bson.M{"$set": feedback}

	// Lakukan operasi UpdateOne untuk mengupdate dokumen di MongoDB
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, feedback)
}

// DeleteFeedback menghapus feedback berdasarkan ID.
func DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Dapatkan koleksi "feedback" dari MongoDB
	collection := config.Mongoconn.Collection("feedback")

	// Tentukan filter berdasarkan ID feedback
	filter := bson.M{"_id": id}

	// Lakukan operasi DeleteOne untuk menghapus dokumen di MongoDB
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


func init() {
	// Check if MongoDB connection error exists during initialization
	if config.ErrorMongoconn != nil {
		log.Fatal(config.ErrorMongoconn)
	}
}