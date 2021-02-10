package main

import (
	"log"
	"net/http"

	"github.com/CezarGarrido/crypto-index/api/db"
	"github.com/CezarGarrido/crypto-index/api/entity"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// main : start program
func main() {

	database, err := db.NewWithMemoryDB("./api/storage/currencies.json")
	if err != nil {
		panic(err)
	}

	if err := database.MemoryDB.WriteFile(entity.Currency{
		BRL: "5.400",
		EUR: "0.920",
		CAD: "1.440",
	}); err != nil {
		log.Fatalln(err)
	}

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Cache-Control", "X-File-Name", "Origin", "X-Session-ID"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"})

	log.Println("Run on port:8089")
	r := mux.NewRouter()

	err = http.ListenAndServe(":8089", handlers.CORS(headersOk, methodsOk, originsOk)(r))
	if err != nil {
		log.Panic(err)
	}
}
