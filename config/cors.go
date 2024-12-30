package config

import (
	"log"
	"net/http"
)

var Origins = []string{
	"https://cdn.blkkalittifaq.id",
	"https://www.blkkalittifaq.id",
	"https://kabar.blkkalittifaq.id",
}

var Headers = []string{
	"Origin",
	"Content-Type",
	"Accept",
	"Authorization",
	"Access-Control-Request-Headers",
	"Token",
	"Login",
	"Access-Control-Allow-Origin",
	"Bearer",
	"X-Requested-With",
}

// Fungsi untuk memeriksa apakah origin diizinkan
func isAllowedOrigin(origin string) bool {
	for _, o := range Origins {
		if o == origin {
			return true
		}
	}
	return false
}

// Fungsi untuk mengatur header CORS
func SetAccessControlHeaders(w http.ResponseWriter, r *http.Request) bool {
	origin := r.Header.Get("Origin")

	if isAllowedOrigin(origin) {
		// Set header CORS untuk permintaan preflight (OPTIONS)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Max-Age", "3600")

		// Tangani preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			log.Println("Preflight request ditangani untuk origin:", origin)
			return true
		}

		// Tambahkan log untuk debugging
		log.Println("CORS header ditambahkan untuk origin:", origin)
		return false
	}

	// Log jika origin tidak diizinkan
	log.Println("Origin tidak diizinkan:", origin)
	return false
}