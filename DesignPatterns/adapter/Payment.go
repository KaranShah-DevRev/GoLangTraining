package main

import "fmt"

type Payment interface {
	Pay(amount float32)
}

func (p *PersonPaymentMethodSelect) Pay(amount float32) {
	if p.Cash != (Cash{}) {
		fmt.Println("Paying cash: ", amount)
		return
	} else if p.PaymentMethod != (PaymentMethod{}) {
		if p.PaymentMethod.Card != (Card{}) {
			fmt.Println("Paying with card: ", amount)
		} else {
			fmt.Println("Paying through bank: ", amount)
		}
	}
}
