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
	status string
}

// SiparisVer: Producer (yazma-yönlü kanal)
func SiparisVer(siparisKanali chan<- Order, totalCount int) {
	coffeTypes := []string{"Latte", "Americano", "Filtre", "Mocha", "Espresso", "Macchiato", "marijuana", "Flat Dark"}
	for i := 1; i <= totalCount; i++ {
		ord := Order{
			ID:     i,
			Name:   coffeTypes[(i-1)%len(coffeTypes)],
			status: "Yeni",
		}
		fmt.Printf("Müşteri: Sipariş %d verildi. (%s)\n", ord.ID, ord.Name)
		siparisKanali <- ord
		// Müşteriler arası küçük gecikme
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("Müşteri: Tüm siparişler verildi. Sipariş kanalı kapatılıyor.")
	close(siparisKanali)
}

// Kasiyer: Middleman (okuma ve yazma yönlü kanallar) mutfakKanali Okuma kanali
func Kasiyer(siparisKanali <-chan Order, mutfakKanali chan<- Order) {
	for ord := range siparisKanali {
		fmt.Printf("Kasiyer: Sipariş %d alındı. Ödeme alınıyor...\n", ord.ID)
		time.Sleep(600 * time.Millisecond) // ödeme süresi
		ord.status = "Hazirlaniyor"
		// Kuyruk dolu mu? (yalnızca bilgilendirme amaçlı; len/cap racy olabilir ama simülasyon için kabul edilebilir)
		if len(mutfakKanali) == cap(mutfakKanali) {
			fmt.Println("Kasiyer: Mutfak kuyruğu dolu, yeni sipariş almak için bekliyorum...")
			mutfakKanali <- ord
			fmt.Printf("Kasiyer: Sipariş %d mutfağa iletildi. (Kuyruk: %d/%d)\n", ord.ID, len(mutfakKanali), cap(mutfakKanali))
		}

	}
	fmt.Println("Kasiyer: Müşteri kanalı kapandı. Mutfak kuyruğu kapatılıyor.")
	close(mutfakKanali)
}

func Barista(id int, mutfakChannel <-chan Order, siparisChannel chan<- Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for ord := range mutfakChannel {
		random := rand2.IntN(1500)
		time.Sleep(time.Duration(random) * time.Millisecond)
		ord.status = "Hazır"
		siparisChannel <- ord
		fmt.Printf("Barista %d: Sipariş %d bitti! ☕️ (Hazırlık süresi: %v)\n", id, ord.ID, random/100)
	}
	fmt.Printf("Barista %d: Mutfak kapandı, çıkıyorum.\n", id)
}

func main() {
	siparisKanal := make(chan Order)
	mutfakKuyruk := make(chan Order, 5)
	hazirKahveler := make(chan Order)
	var wg sync.WaitGroup

	SiparisVer(siparisKanal, 12)
	Kasiyer(siparisKanal, mutfakKuyruk)

	go Barista(1, mutfakKuyruk, hazirKahveler, &wg)
	go Barista(2, mutfakKuyruk, hazirKahveler, &wg)
	go Barista(3, mutfakKuyruk, hazirKahveler, &wg)
}
