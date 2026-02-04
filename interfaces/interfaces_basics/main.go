package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64)
}

func ExecutePayment(pp PaymentProcessor, amount float64) {
	pp.Pay(amount)
}

type CreditCard struct{}

func (c CreditCard) Pay(a float64) {
	fmt.Printf("kredi kartiyla odendi tutar : %v\n", a)
}

type BitcoinWallet struct{}

func (bw BitcoinWallet) Pay(b float64) {
	fmt.Printf("bitcoin cuzdanıyla odendi tutar : %v\n", b)
}

func main() {
	c1 := CreditCard{}
	bw1 := BitcoinWallet{}

	ExecutePayment(c1, 50)
	ExecutePayment(bw1, 90)
	// output :
	//kredi kartiyla odendi tutar : 50
	//bitcoin cuzdanıyla odendi tutar : 90
}
