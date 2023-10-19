package models


type PaymentMethod struct {
	ID   string `json:"id"`
	PaymentMethodType string `json:"paymentMethodType"`
	CreditCardNumber string `json:"creditCardNumber"`
	CreditCardHolderName string `json:"creditCardHolderName"`
	ExpirationDate  string `json:"expirationDate"`
	CVV  string `json:"cvv"`
	
}