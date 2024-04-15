package app

import (
	"github.com/JingBh/crypto-learn/internal/app/controllers"
	"net/http"
)

var Mux = http.NewServeMux()

func init() {
	Mux.Handle("GET /{$}", controllers.ServeTemplate("index.html"))
	Mux.Handle("POST /des", http.HandlerFunc(controllers.PostDES))
	Mux.Handle("GET /des/key", http.HandlerFunc(controllers.GetDESKey))
}
