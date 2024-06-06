package route

import (
	"log"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.ErrorMongoconn != nil {
		log.Println(config.ErrorMongoconn.Error())
	}

	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/":
		controller.GetHome(w, r)

	default:
		controller.NotFound(w, r)
	}
}
