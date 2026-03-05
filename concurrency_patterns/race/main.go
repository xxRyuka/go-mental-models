package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sayac := 0 // Tüm goroutine'lerin paylaştığı ortak bellek alanı (Shared State)

	// Aynı anda 1000 adet goroutine başlatıyoruz
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// TEHLİKE BURADA: 1000 goroutine aynı anda bu değişkene yazmaya çalışıyor!
			sayac++
		}()
	}

	wg.Wait()
	fmt.Printf("Beklenen Sonuç: 1000, Gerçekleşen Sonuç: %v\n", sayac)
}
