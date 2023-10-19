package routes

import (
	"backend/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func PaymentMethodRoute(router *mux.Router, paymentMethodCollection *mongo.Collection) {
	router.HandleFunc("/paymentMethod", controllers.CreatePaymentMethod(paymentMethodCollection)).Methods("POST")
	router.HandleFunc("/paymentMethod/{id}", controllers.GetPaymentMethod(paymentMethodCollection)).Methods("GET")
	router.HandleFunc("/paymentMethod/{id}", controllers.UpdatePaymentMethod(paymentMethodCollection)).Methods("PUT")
	router.HandleFunc("/paymentMethod/{id}", controllers.DeletePaymentMethod(paymentMethodCollection)).Methods("DELETE")
	router.HandleFunc("/paymentMethods", controllers.GetAllPaymentMethods(paymentMethodCollection)).Methods("GET")
}
