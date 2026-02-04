package main

import "fmt"

type Notifier interface {
	Send(msg string) bool
	GetProvider() string
}

type NotificationManager struct {
	// Slice yerine Map! Böylece isme göre bulabiliriz.
	//providers   []Notifier
	providerMap map[string]Notifier // key string : value Notifier
}

// NewManager Ctor for init map
// panic: assignment to entry in nil map hatasini engellemis olduk
func NewManager() *NotificationManager { // managerin bellek adresteki yerini donuyoruz :)
	return &NotificationManager{
		providerMap: make(map[string]Notifier),
	}
}

func (nm *NotificationManager) AddProvider(prv Notifier) {
	key := prv.GetProvider()
	nm.providerMap[key] = prv
	fmt.Printf("Sisteme eklendi: %v\n", key)
}

func (nm *NotificationManager) Broadcast(msg string) {
	for index, provider := range nm.providerMap {

		if provider.Send(msg) {
			fmt.Printf("%v. providera mesaj gonderildi \n", index)

		} else {
			fmt.Printf("%v nolu providera mesaj gonderilemedi\n", index)
		}
	}
}

func (nm *NotificationManager) SendByProviderId(id string, msg string) {
	if prv, ok := nm.providerMap[id]; ok {
		prv.Send(msg)
		fmt.Printf("Spesifik olarak %v provideri üzerinden %v mesajı iletilmiştir \n", prv, msg)
	} else {
		fmt.Println("provider mevcut değil")
	}

}
