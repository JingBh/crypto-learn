package app

import (
	"github.com/JingBh/crypto-learn/internal/app/controllers"
	"net/http"
)

var Mux = http.NewServeMux()

func init() {
	Mux.Handle("GET /{$}", controllers.ServeTemplate("index.html"))
	Mux.Handle("POST /cipher", http.HandlerFunc(controllers.PostCipher))
	Mux.Handle("GET /cipher/key", http.HandlerFunc(controllers.GetCipherKey))
	Mux.Handle("POST /perf", http.HandlerFunc(controllers.PostPerf))
}
