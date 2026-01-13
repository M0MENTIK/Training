package main

import "fmt"

type PaymentMethode interface {
	Pay(amoint float64) string
}

func ProcessPayment(p PaymentMethode, amount float64) {
	fmt.Println(p.Pay(amount))
}

//

//

//

type CreditCard struct {
	cardNumber string
}

func (c CreditCard) Pay(amoint float64) string {
	return fmt.Sprintf("Payment with a credit card: %s. Sum operation: %.2f$", c.cardNumber, amoint)
}

//

//

//

type PayPal struct {
	email string
}

func (p PayPal) Pay(amoint float64) string {
	return fmt.Sprintf("Payment with a PayPal: %s. Sum operation:%.2f$", p.email, amoint)
}
