package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentMethod struct {
	Id                primitive.ObjectID `json:"id,omitempty"`
	PaymentMethodType string             `json:"paymentMethodType" validate:"required"`
	CardNumber        string             `json:"cardNumber" validate:"required"`
	CardHolder        string             `json:"cardHolder" validate:"required"`
	ExpireDate        string             `json:"expireDate" validate:"required"`
	CVV               string             `json:"cvv,omitempty" validate:"required"`
}
