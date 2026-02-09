package main

import (
	"fmt"
	rand2 "math/rand/v2"
	"sync"
	"time"
)

// WaitGroup'u pointer olarak geçmek zorundayız! (Kopya giderse sayaç çalışmaz)
func DownloadFile(dosyaAdi string, wg *sync.WaitGroup) {

	// Egerki defer kw ile fonksiyon kapanırken Done olarak isaretlemez isek deadlock hatası alırız önemli bir nokta
	// fatal error: all goroutines are asleep - deadlock!

	defer wg.Done() // bitince calıscak // e
	// wg.Done aslında , wg.Add(-1) ile aynı islevde

	//defer wg.Add(-1) //  yorum satirini kaldirip deneyebilirsin tekrardan

	rand := rand2.IntN(5) + 1
	fmt.Printf("%v Dosyasi indirilmeye baslanıldı \n", dosyaAdi)
	time.Sleep(time.Duration(rand) * time.Second)
	fmt.Printf("%v Dosyasi indirildi \n -------------\n", dosyaAdi)
}

func main() {
	var wg sync.WaitGroup
	dosyalar := []string{"oyun", "film", "goLand"}
	for _, v := range dosyalar {
		// 1. Sayacı artır (Her dosya için +1)
		// Bunu 'go' demeden ÖNCE yapmak zorundasın.
		wg.Add(1)
		go DownloadFile(v, &wg) // go yazmazsak işlem tamamen semkron geçeklesir
	}

	fmt.Println("⏳ Main: Tüm indirmeler bekleniyor...")

	// Burası olmazsa main direkt kapanır!
	// Sayaç 0 olana kadar (yani 3 Done gelene kadar) burada bekler.
	wg.Wait()

	fmt.Println(" Main: Bitti")

	// Senkron Output without (go kw) işlemler birbirini blokluyor
	//oyun Dosyasi indirilmeye baslanıldı
	//oyun Dosyasi indirildi
	//-------------
	//	film Dosyasi indirilmeye baslanıldı
	//film Dosyasi indirildi
	//-------------
	//	goLand Dosyasi indirilmeye baslanıldı
	//goLand Dosyasi indirildi
	//-------------

	// Asenkron Output with (go kw) işi biten cıkıyor
	//film Dosyasi indirilmeye baslanıldı
	//oyun Dosyasi indirilmeye baslanıldı
	//goLand Dosyasi indirilmeye baslanıldı
	//film Dosyasi indirildi
	//-------------
	//	goLand Dosyasi indirildi
	//-------------
	//	oyun Dosyasi indirildi
	//-------------
}
