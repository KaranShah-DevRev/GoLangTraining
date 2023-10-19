package main

import (
	"backend/configs"
	"backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	paymentMethodCollection := configs.GetCollection(configs.ConnectDB())

	routes.PaymentMethodRoute(router, paymentMethodCollection)

	log.Fatal(http.ListenAndServe(":8080", router))
}
