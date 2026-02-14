package main

import (
	"fmt"
	rand2 "math/rand/v2"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Name   string
	Status string
}

func SiparisVer(musteriKanali chan<- Order, siparisSayi int) {

	spesiyaller := []string{"safa", "samet", "fukan", "sena", "cakkal", "ahmak"}
	for i := 1; i < siparisSayi; i++ {
		randSayi := rand2.IntN(len(spesiyaller))
		ord := Order{
			ID:     i,
			Name:   spesiyaller[randSayi],
			Status: "Siparis Alindi",
		}
		fmt.Printf("Musteri Kanali : Siparis %v Alindi %+v\n", i, ord)
		musteriKanali <- ord // siparisler kanala eklendi
		//time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("Siparisler tamamlandı ve Musteri Kanali Kapatildi \n")
	close(musteriKanali) // kanali acan kapatmalı , altında yatan nedeni arastir
}

// Siparis kanalindan okucam mutfak kanalina yazcam
func Kasiyer(musteriKanali <-chan Order, mutfakKanal chan<- Order) {
	for siparis := range musteriKanali {
		fmt.Printf("%+v Siparisi hazirlaniliyor \n", siparis)
		time.Sleep(500 * time.Millisecond)
		siparis.Status = "Hazirlaniyor"
		fmt.Printf("Siparis %+v Mutfağa İletildi \n", siparis)
		mutfakKanal <- siparis
	}
	close(mutfakKanal)

}

func Barista(baristaNumarası int, mutfakKanali <-chan Order, hazirSiparisKanali chan<- Order, wg *sync.WaitGroup) {

	defer wg.Done()
	for order := range mutfakKanali {
		fmt.Printf("Siparis %v , %v Numarali Barista Tarafindan Hazirlanmaya baslandı\n", order.ID, baristaNumarası)
		time.Sleep((2) * time.Second)
		order.Status = "Hazir"
		fmt.Printf("%v numarali Barista %v numarali siparisi hazirladi \n", baristaNumarası, order.ID)
		hazirSiparisKanali <- order
	}
	//close(hazirSiparisKanali)
}

func Notification(hazirSiparisKanali <-chan Order) {
	for ord := range hazirSiparisKanali {
		fmt.Printf("%+v Hazir \n ", ord)
	}
}

func main() {

	chMusteri := make(chan Order)
	go SiparisVer(chMusteri, 10) // bunu go yapmazsak senkron oldugu için alıcısını bekleyecek ve deadlock olacak ,
	// go yaparsak birisi consume etmezse bu sefer çalısmayacak birisinin consume etmesi gerekiyor unbuffered oldugu için

	chMutfak := make(chan Order, 50)
	go Kasiyer(chMusteri, chMutfak)

	chHazirSiparisler := make(chan<- Order)
	var wg sync.WaitGroup
	wg.Add(2)
	go Barista(100, chMutfak, chHazirSiparisler, &wg)
	go Barista(200, chMutfak, chHazirSiparisler, &wg)

	go func() {
		wg.Wait()
		//close(chHazirSiparisler)
	}()
	//Notification(chHazirSiparisler)

}
