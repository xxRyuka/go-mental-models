package main

import "fmt"

type SmsSender struct {
	PhoneNumber string
}

func (ss SmsSender) GetProvider() string {
	return "smsProvider"
}

func (ss SmsSender) Send(msg string) bool {

	fmt.Printf("%v numarasına %v mesaji iletildi\n", ss.PhoneNumber, msg)
	return true
}
