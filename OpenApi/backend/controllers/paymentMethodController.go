package controllers

import (
	"backend/models"
	"backend/response"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

func CreatePaymentMethod(paymentMethodCollection *mongo.Collection) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var paymentMethod models.PaymentMethod
		defer cancel()
		if err := json.NewDecoder(req.Body).Decode(&paymentMethod); err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		if validationErr := validate.Struct(&paymentMethod); validationErr != nil {
			resp.WriteHeader(http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		newPaymentMethod := models.PaymentMethod{
			Id:                primitive.NewObjectID(),
			PaymentMethodType: paymentMethod.PaymentMethodType,
			CardNumber:        paymentMethod.CardNumber,
			CardHolder:        paymentMethod.CardHolder,
			ExpireDate:        paymentMethod.ExpireDate,
			CVV:               paymentMethod.CVV,
		}

		result, err := paymentMethodCollection.InsertOne(ctx, newPaymentMethod)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		resp.WriteHeader(http.StatusCreated)
		response := response.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(resp).Encode(response)
	}
}

func GetPaymentMethod(paymentMethodCollection *mongo.Collection) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(req)
		pmId := params["id"]
		var paymentMethod models.PaymentMethod
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(pmId)

		err := paymentMethodCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&paymentMethod)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		resp.WriteHeader(http.StatusOK)
		response := response.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": paymentMethod}}
		json.NewEncoder(resp).Encode(response)
	}
}

func UpdatePaymentMethod(paymentMethodCollection *mongo.Collection) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(req)
		pmId := params["id"]
		var paymentMethod models.PaymentMethod
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(pmId)

		if err := json.NewDecoder(req.Body).Decode(&paymentMethod); err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		if validationErr := validate.Struct(&paymentMethod); validationErr != nil {
			resp.WriteHeader(http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		update := bson.M{"expireDate": paymentMethod.ExpireDate}

		result, err := paymentMethodCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		var updatedPaymentMethod models.PaymentMethod
		if result.MatchedCount == 1 {
			err := paymentMethodCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&paymentMethod)

			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(resp).Encode(response)
				return
			}
		}

		resp.WriteHeader(http.StatusOK)
		response := response.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedPaymentMethod}}
		json.NewEncoder(resp).Encode(response)
	}
}

func DeletePaymentMethod(paymentMethodCollection *mongo.Collection) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(req)
		pmId := params["id"]
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(pmId)

		result, err := paymentMethodCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			resp.WriteHeader(http.StatusNotFound)
			response := response.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		resp.WriteHeader(http.StatusOK)
		response := response.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}}
		json.NewEncoder(resp).Encode(response)

	}
}

func GetAllPaymentMethods(paymentMethodCollection *mongo.Collection) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var paymentMethods []models.PaymentMethod
		defer cancel()

		cursor, err := paymentMethodCollection.Find(ctx, bson.M{})
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(resp).Encode(response)
			return
		}

		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var paymentMethod models.PaymentMethod
			cursor.Decode(&paymentMethod)
			paymentMethods = append(paymentMethods, paymentMethod)
		}

		resp.WriteHeader(http.StatusOK)
		response := response.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": paymentMethods}}
		json.NewEncoder(resp).Encode(response)
	}
}
