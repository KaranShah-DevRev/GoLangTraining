package main

type Cash struct {
	Source string
}

type Card struct {
	CardNumber string
	CardHolder string
	ExpMonth   string
	ExpYear    string
	CVV        string
}

type BankTransfer struct {
	AccountHolder string
	AccountNumber string
}

type PaymentMethod struct {
	Card         Card
	BankTransfer BankTransfer
}


type PersonPaymentMethodSelect struct {
	PaymentMethod PaymentMethod
	Cash 		Cash
}
