package route

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/helper"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}
	config.SetEnv()

	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/data/gallery":
		controller.GetGallery(w, r)
	case method == "GET" && path == "/data/product":
		controller.GetProduk(w, r)
	case method == "POST" && helper.URLParam(path, "/upload/:path"):
		controller.PostUploadGithub(w, r)
	case method == "POST" && path == "/data/product":
		controller.PostProduk(w, r)
	case method == "POST" && path == "/data/gallery":
		controller.PostGallery(w, r)
	case method == "DELETE" && path == "/data/produk":
		controller.DeleteProduk(w, r)
	case method == "POST" && path == "/data/adminregister":
		controller.RegisterHandler(w, r)
	}

}
