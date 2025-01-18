package main

import "fmt"

// Define an interface
type paymenter interface {
	pay (amount float32)
	quantity (amount int)
}

type payment struct {
	gateway paymenter
}

type stripe struct {}
func (s stripe) pay (amount float32) {
	fmt.Println("Paying using stripe:", amount)
	s.quantity(5)
}
func (s stripe) quantity (amount int) {
	fmt.Println("Quantity with stripe:", amount)
}	

type razorpay struct {}
func (r razorpay) pay (amount float32) {
	fmt.Println("Paying using razorpay:", amount)
	r.quantity(10)
}
func (r razorpay) quantity (amount int) {
	fmt.Println("Quantity with razorpay:", amount)
}

func main() {
	fmt.Println("Interfaces in Go")
	payment1 := payment{
		gateway: razorpay{},
	}
	payment2 := payment{
		gateway: stripe{},
	}
	// fmt.Println(payment1, payment2)
	payment1.gateway.pay(1000)
	payment2.gateway.pay(2000)
}