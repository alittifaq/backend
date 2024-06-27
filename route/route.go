package route

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/helper"
)

// URL mengarahkan permintaan HTTP masuk ke fungsi controller yang tepat.
func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // Jika ini adalah permintaan preflight, kembalikan segera.
	}
	config.SetEnv()

	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/data/gallery":
		controller.GetGallery(w, r)
	case method == "GET" && path == "/data/product":
		controller.GetProduk(w, r)
	case method == "GET" && path == "/data/gallery/detail":
		controller.GetOneGallery(w, r)
	case method == "POST" && helper.URLParam(path, "/upload/:path"):
		controller.PostUploadGithub(w, r)
	case method == "POST" && path == "/data/product":
		controller.PostProduk(w, r)
	case method == "POST" && path == "/data/gallery":
		controller.PostGallery(w, r)
	case method == "PUT" && path == "/data/product":
		controller.UpdateProduct(w, r)
	case method == "PUT" && path == "/data/gallery":
		controller.UpdateGallery(w, r)
	case method == "DELETE" && path == "/data/product":
		controller.DeleteProduk(w, r)
	case method == "DELETE" && path == "/data/gallery":
		controller.DeleteGallery(w, r)
	case method == "POST" && path == "/data/adminregister":
		controller.RegisterHandler(w, r)
	case method == "POST" && path == "/data/user":
		controller.GetUser(w, r)
	}
}
