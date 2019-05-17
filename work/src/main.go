package main

import (
	"danmu_rest"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	router := danmu_rest.NewRouter()

	// These two lines are important in order to allow access from the front-end side to the methods
//	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
//	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})

//	log.Fatal(http.ListenAndServe(":8000", router))
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	//log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
