package main

func main() {

	var card = Card{
		CardNumber: "1111 2222 3333 4444",
		CardHolder: "John Doe",
		ExpMonth:   "09",
		ExpYear:    "2021",
		CVV:        "123",
	}

	var bankTransfer = BankTransfer{
		AccountHolder: "John Doe",
		AccountNumber: "123456789",
	}

	paymentCard := PersonPaymentMethodSelect{
		PaymentMethod: PaymentMethod{Card: card},
	}
	paymentCard.Pay(100)

	paymentTransfer := PersonPaymentMethodSelect{
		PaymentMethod: PaymentMethod{BankTransfer: bankTransfer},
	}
	paymentTransfer.Pay(100)

	paymentCash := PersonPaymentMethodSelect{
		Cash: Cash{Source: "Wallet"},
	}
	paymentCash.Pay(100)
}
