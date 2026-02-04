package main

import "fmt"

type Notifier interface {
	Send(msg string) bool
}

type NotificationManager struct {
	providers []Notifier
}

func (nm *NotificationManager) AddProvider(prv Notifier) {
	nm.providers = append(nm.providers, prv)
}

func (nm NotificationManager) Broadcast(msg string) {
	for index, provider := range nm.providers {

		if provider.Send(msg) {
			fmt.Printf("%v. providera mesaj gonderildi \n", index)

		} else {
			fmt.Printf("%v nolu providera mesaj gonderilemedi\n", index)
		}
	}
}
