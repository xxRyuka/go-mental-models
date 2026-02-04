package main

import "fmt"

type EmailSender struct {
	Adress    string
	SendCount int
}

func (es EmailSender) GetProvider() string {
	return "emailProvider"
}

func (es *EmailSender) Send(msg string) bool {
	fmt.Printf("%v adresine %v mesajı başarıyla iletildi\n", es.Adress, msg)
	es.SendCount++
	return true
}
