package internal

import (
	"github.com/JingBh/crypto-learn/internal/app"
	"log"
	"net/http"
)

func StartServer(publish bool) {
	addr := "localhost:8206"
	if publish {
		addr = ":8206"
	}
	log.Fatal(http.ListenAndServe(addr, app.Mux))
}
