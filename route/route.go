package route

import (
	"log"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/handlers"
	"github.com/gocroot/helper"
	"github.com/gorilla/mux"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.ErrorMongoconn != nil {
		log.Println(config.ErrorMongoconn.Error())
	}

	var method, path string = r.Method, r.URL.Path
	switch {
	// case method == "GET" && path == "/":
	// 	controller.GetHome(w, r)
	// case method == "GET" && path == "/refresh/token":
	// 	controller.GetNewToken(w, r)
	case method == "POST" && helper.URLParam(path, "data/produk"):
		controller.PostProduk(w, r)
	case method == "DELETE" && helper.URLParam(path, "data/produk"):
		controller.DeleteProduk(w, r)
	}
}

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")

	return router
}
